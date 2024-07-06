package api

import (
	"fmt"
	"log"
	"net/http"
	nurl "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/go-shiori/go-readability"
	"github.com/microcosm-cc/bluemonday"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	vault := r.URL.Query().Get("vault")
	if len(vault) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	folder := r.URL.Query().Get("folder")
	if len(folder) == 0 {
		folder = "Clippings"
	}

	url := r.URL.Query().Get("url")
	if len(url) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tags := r.URL.Query()["tag"]

	silent := strings.ToLower(r.URL.Query().Get("silent")) == "true"

	fetchedAt := time.Now()
	if rawEpoch := r.URL.Query().Get("epoch"); len(rawEpoch) > 0 {
		epoch, err := strconv.Atoi(rawEpoch)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if epoch > 0 {
			fetchedAt = time.Unix(int64(epoch), 0)
		}
	}

	if _, err := nurl.Parse(url); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article, err := readability.FromURL(url, 10*time.Second)
	if err != nil {
		log.Printf("failed to parse %s: %v\n", url, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename := time.Now().Format("20060102_150405")
	if len(article.Title) > 0 {
		filename = regexp.MustCompile(`[^a-zA-Z0-9 _\-]+`).ReplaceAllString(article.Title, "")
	}

	converter := md.NewConverter(url, true, nil)
	mdContent, err := converter.ConvertString(bluemonday.UGCPolicy().Sanitize(article.Content))
	if err != nil {
		log.Printf("failed to convert body to Markdown: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parts := []string{buildFrontmatter(url, fetchedAt.Format(time.RFC3339), tags...)}
	if len(article.Title) > 0 {
		parts = append(parts, fmt.Sprintf("# %s\n", article.Title))
	}

	content := strings.Join(append(parts, mdContent), "\n")
	link, err := buildObsidianLink(vault, fmt.Sprintf("%s/%s", folder, filename), content, silent)
	if err != nil {
		log.Printf("failed to build obsidian link: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", link)
	http.Redirect(w, r, link, http.StatusFound)
}

func buildFrontmatter(url string, fetchedAt string, tags ...string) string {
	var formattedTags string
	if len(tags) > 0 {
		formattedTags = "tags: ["
	}
	for i, tag := range tags {
		if i == len(tags)-1 {
			formattedTags += fmt.Sprintf(`"%s"`, tag)
		} else {
			formattedTags += fmt.Sprintf(`"%s", `, tag)
		}
	}
	if len(tags) > 0 {
		formattedTags += "]"
	}

	return fmt.Sprintf(
		"---\nurl: %s\nfetched_at: %s\n%s\n---\n",
		url,
		fetchedAt,
		formattedTags,
	)
}

func buildObsidianLink(vault string, path string, content string, silent bool) (string, error) {
	// mimic Javascript's encodeURIComponent, which is looser than Go's url.QueryEscape
	encodeURIComponent := func(str string) string {
		result := strings.Replace(str, "+", "%20", -1)
		result = strings.Replace(result, "%21", "!", -1)
		result = strings.Replace(result, "%27", "'", -1)
		result = strings.Replace(result, "%28", "(", -1)
		result = strings.Replace(result, "%29", ")", -1)
		result = strings.Replace(result, "%2A", "*", -1)
		return result
	}

	baseURL, err := nurl.Parse("obsidian://new")
	if err != nil {
		return "", err
	}

	values := nurl.Values{}
	values.Add("vault", vault)
	values.Add("file", path)
	values.Add("content", content)
	values.Add("overwrite", "true")
	if silent {
		values.Add("silent", "true")
	}

	baseURL.RawQuery = encodeURIComponent(values.Encode())

	return baseURL.String(), nil
}
