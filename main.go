package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/boltdb/bolt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gocolly/colly"
	"github.com/segmentio/ksuid"
)

var errKeyAlreadyExists = errors.New("key already exists")

func mkFrontMatter(url string, fetchedAt string, tags ...string) string {
	var formattedTags string
	for i, tag := range tags {
		if i == len(tags)-1 {
			formattedTags += fmt.Sprintf(`"%s"`, tag)
		} else {
			formattedTags += fmt.Sprintf(`"%s", `, tag)
		}
	}

	return fmt.Sprintf(
		`
---
url: %s
fetched_at: %s
tags: [ %s ]
---
`,
		url,
		fetchedAt,
		formattedTags,
	)
}

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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/bookmarks", func(w http.ResponseWriter, r *http.Request) {
		rawURL := r.URL.Query().Get("url")
		if len(rawURL) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		key := strings.Split(parsedURL.String(), "?")[0]
		if key[len(key)-1] == '/' {
			key = key[:len(key)-1]
		}

		if err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("bookmarks"))
			if bucket.Get([]byte(key)) != nil {
				return errKeyAlreadyExists
			}
			return nil
		}); err != nil {
			if errors.Is(err, errKeyAlreadyExists) {
				w.WriteHeader(http.StatusConflict)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tags := r.URL.Query()["tags"]

		c := colly.NewCollector()
		c.OnResponse(func(r *colly.Response) {
			converter := md.NewConverter(r.Request.URL.Path, true, nil)
			markdown, err := converter.ConvertBytes(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			f, err := os.Create(fmt.Sprintf("%s.md", ksuid.New()))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer f.Close()

			if _, err = f.WriteString(
				mkFrontMatter(
					r.Request.URL.String(),
					time.Now().Format(time.RFC3339),
					tags...,
				),
			); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if _, err = f.Write(markdown); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err := db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte("bookmarks"))
				return bucket.Put([]byte(key), []byte(time.Now().Format(time.RFC3339)))
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		})

		if err := c.Visit(rawURL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
