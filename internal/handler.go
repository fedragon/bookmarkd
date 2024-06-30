package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly"
	"github.com/microcosm-cc/bluemonday"
	"github.com/segmentio/ksuid"
)

var (
	frontmatter = "---\nurl: %s\nfetched_at: %s\n%s\n---\n"
)

type Handler struct{}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
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

	tags := r.URL.Query()["tags"]

	c := colly.NewCollector()
	c.OnResponse(func(cr *colly.Response) {
		body := bluemonday.UGCPolicy().SanitizeBytes(cr.Body)
		converter := md.NewConverter(cr.Request.URL.Path, true, nil)
		markdown, err := converter.ConvertBytes(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content := strings.Join(
			[]string{
				buildFrontmatter(cr.Request.URL.String(), time.Now().Format(time.RFC3339), tags...),
				string(markdown),
			},
			"\n",
		)

		link, err := buildObsidianLink("obsidian-plugin-dev", fmt.Sprintf("Clippings/%s", ksuid.New()), content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", link)
		http.Redirect(w, r, link, http.StatusFound)
		return
	})

	if err := c.Visit(rawURL); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func buildFrontmatter(url string, fetchedAt string, tags ...string) string {
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

func buildObsidianLink(vault string, path string, content string) (string, error) {
	// mimick Javascript's encodeURIComponent, which is looser than Go's url.QueryEscape
	encodeURIComponent := func(str string) string {
		result := strings.Replace(str, "+", "%20", -1)
		result = strings.Replace(result, "%21", "!", -1)
		result = strings.Replace(result, "%27", "'", -1)
		result = strings.Replace(result, "%28", "(", -1)
		result = strings.Replace(result, "%29", ")", -1)
		result = strings.Replace(result, "%2A", "*", -1)
		return result
	}

	baseURL, err := url.Parse("obsidian://new")
	if err != nil {
		return "", err
	}

	values := url.Values{}
	values.Add("vault", vault)
	values.Add("file", path)
	values.Add("content", content)
	baseURL.RawQuery = encodeURIComponent(values.Encode())

	return baseURL.String(), nil
}
