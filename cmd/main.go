package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"bookdown/internal"
)

func main() {
	handler := &internal.Handler{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/bookmarks", handler.Handle)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
