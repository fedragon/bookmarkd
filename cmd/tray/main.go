package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/energye/systray"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sync/errgroup"

	"github.com/fedragon/bookmarkd/api"
	"github.com/fedragon/bookmarkd/internal"
)

//go:embed icon.svg
var icon []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	ctx, cancel := context.WithCancel(context.Background())
	group, gctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return run(gctx)
	})

	systray.SetIcon(icon)
	systray.SetTooltip("Store bookmarks in Obsidian")
	systray.SetOnClick(func(menu systray.IMenu) {
		if err := menu.ShowMenu(); err != nil {
			fmt.Println(err)
		}
	})

	systray.AddMenuItem("bookmarkd", "").Disable()

	mStatus := systray.AddMenuItem("Status: 🍏", "Status")
	mStatus.Disable()

	group.Go(func() error {
		checkStatus(gctx, mStatus)
		return nil
	})

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	mQuit.Click(func() {
		fmt.Println("quitting (1)")
		cancel()
		if err := group.Wait(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("quitting (3)")
		systray.Quit()
	})
}

func onExit() {}

func checkStatus(ctx context.Context, mStatus *systray.MenuItem) {
	client := &http.Client{Timeout: time.Second}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("quitting (2)")
			return
		case <-time.Tick(time.Second * 5):
			fmt.Println("updating status...")
			res, err := client.Get("http://localhost:3333/api/status")
			if err != nil {
				mStatus.SetTitle("Status: ❌")
				fmt.Println(err)
			} else if res.StatusCode != http.StatusOK {
				mStatus.SetTitle("Status: 🍎")
			} else {
				mStatus.SetTitle("Status: 🍏")
			}
		}
	}
}

func run(ctx context.Context) error {
	config := internal.Config{}
	if err := envconfig.Process("", &config); err != nil {
		return err
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/api/bookmarks", api.Handle)
	router.Get("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{Addr: config.HttpAddress, Handler: router}

	go func() {
		fmt.Println("starting server...")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
