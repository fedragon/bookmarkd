# bookmd

Converts HTML pages to Markdown files and stores them in a local [Obsidian](https://obsidian.md) vault, in line with Steph Ango's [File over app](https://stephango.com/file-over-app) philosophy.

Whenever you'd like to bookmark a page, click on the provided bookmarklet, and it will be added to your Obsidian vault.

## Usage

Run the server with the following command (on your local machine, or anywhere you'd like to execute it from):

```
BOOKMD_OBSIDIAN_VAULT=<your_vault> BOOKMD_OBSIDIAN_FOLDER=<your_folder> go run cmd/main.go
```

where

- `BOOKMD_OBSIDIAN_VAULT` is the name of your Obsidian vault
- `BOOKMD_OBSIDIAN_FOLDER` is the name of the folder that will contain your bookmarks (defaulting to `Clippings`, if not set).

## Install bookmarklet

The bookmarklet assumes that your server is running on `localhost:3000`: if you'd like to run it remotely instead, you'll need to change the URL in `bookmarklet/src.js` and then run

```shell
cd bookmarklet
node maker.js
```

To install the bookmarklet in your browser, right-click on your bookmarks' bar, select "Add Bookmark" and paste the contents of [bookmarklet.js](bookmarklet/bookmarklet.js) in the "URL" field:

## Credits

Inspired by [downmark](https://github.com/alessandro-fazzi/downmark).
