package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"

	"github.com/fedragon/bookmd/internal"
)

func main() {
	config := internal.Config{}
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	handler := &internal.Handler{
		Vault:  config.ObsidianVault,
		Folder: config.ObsidianFolder,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/bookmarks", handler.Handle)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HttpPort), r); err != nil {
		log.Fatal(err)
	}
}
