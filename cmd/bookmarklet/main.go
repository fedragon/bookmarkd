package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"text/template"

	"github.com/charmbracelet/huh"

	"github.com/fedragon/bookmarkd/bookmarklet"
	"github.com/fedragon/bookmarkd/internal"
)

func main() {
	address := "http://localhost:11235"
	vault := "Vault"
	folder := "Clippings"

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("HTTP server address").
				Placeholder("http://localhost:11235").
				Value(&address),
			huh.NewInput().
				Title("Vault").
				Placeholder("Vault").
				Value(&vault),
			huh.NewInput().
				Title("Folder").
				Placeholder("Clippings").
				Value(&folder),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New("bookmarklet").Parse(bookmarklet.SourceFile)
	if err != nil {
		log.Fatal(err)
	}

	writer := bytes.NewBuffer([]byte{})
	if err := tpl.Execute(
		writer,
		struct {
			Address string
			Vault   string
			Folder  string
		}{
			Address: address,
			Vault:   vault,
			Folder:  folder,
		},
	); err != nil {
		log.Fatal(err)
	}

	encoded := url.QueryEscape(strings.TrimSpace(writer.String()))
	content := fmt.Sprintf(
		"javascript:%s",
		internal.EncodeURIComponent(
			fmt.Sprintf(
				"(function(){%s})();",
				encoded,
			),
		),
	)
	if err := os.WriteFile("bookmarklet/bookmarklet.js", []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
}
