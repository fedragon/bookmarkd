package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"bookdown/internal"
)

func main() {
	db, err := bolt.Open("bookdown.db", 0600, &bolt.Options{Timeout: time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("bookmarks"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	handler := &internal.Handler{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/bookmarks", handler.HandlePost)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
