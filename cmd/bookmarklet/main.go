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

	"github.com/fedragon/bookmarkd/bookmarklet"
	"github.com/fedragon/bookmarkd/internal"
)

const html = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<a href="{{ .URL }}">Generate bookmarklet</a>
</body>
</html>
`

func main() {
	escaped := url.QueryEscape(strings.TrimSpace(bookmarklet.SourceFile))
	encoded := fmt.Sprintf(
		"javascript:%s",
		internal.EncodeURIComponent(
			fmt.Sprintf(
				"(function(){%s})();",
				escaped,
			),
		),
	)

	tpl, err := template.New("bookmarklet").Parse(html)
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	if err := tpl.Execute(buffer, struct{ URL string }{URL: encoded}); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("bookmarklet/index.html", buffer.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}
