# bookmd

Converts HTML pages to Markdown files and stores them in a local [Obsidian](https://obsidian.md) vault, in line with Steph Ango's [File over app](https://stephango.com/file-over-app) philosophy.

Whenever you'd like to bookmark a page, click on the provided bookmarklet, and it will be added to your Obsidian vault.

## Usage

Run the server with the following command (on your local machine, or anywhere you'd like to execute it from):

```
go run cmd/main.go
```

## Install bookmarklet

The bookmarklet assumes that your server is running on `localhost:3333`: if you'd like to run it remotely instead, you'll need to change the URL in `bookmarklet/src.js` and then run

```shell
cd bookmarklet
node maker.js
```

To install the bookmarklet in your browser, right-click on your bookmarks' bar, select "Add Bookmark" and paste the contents of [bookmarklet.js](bookmarklet/bookmarklet.js) in the "URL" field:

## Credits

Inspired by [downmark](https://github.com/alessandro-fazzi/downmark).
