package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/boltdb/bolt"
	"github.com/gocolly/colly"
	"github.com/microcosm-cc/bluemonday"
	"github.com/segmentio/ksuid"
)

var (
	bucketName          = []byte("bookmarks")
	errKeyAlreadyExists = errors.New("key already exists")
	frontmatter         = `
---
url: %s
fetched_at: %s
tags: [ %s ]
---
`
)

type Handler struct {
	DB *bolt.DB
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
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

	if err := h.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
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
		policy := bluemonday.UGCPolicy()
		converter := md.NewConverter(r.Request.URL.Path, true, nil)
		markdown, err := converter.ConvertBytes(policy.SanitizeBytes(r.Body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("docs/%s.md", ksuid.New()))
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

		if _, err = f.WriteString("\n"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err = f.Write(markdown); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := h.DB.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(bucketName)
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
}

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
		frontmatter,
		url,
		fetchedAt,
		formattedTags,
	)
}
