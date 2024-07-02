# bookmd

Converts HTML pages to Markdown files and stores them in a local [Obsidian](https://obsidian.md) vault, in line with Steph Ango's [File over app](https://stephango.com/file-over-app) philosophy.

Whenever you'd like to bookmark a page, click on the provided bookmarklet, and you will be prompted to add it to your Obsidian vault.

## Usage

The code can run either as a standalone HTTP server or as a [Vercel Function](https://vercel.com/docs/functions/runtimes/go). In either case, after the deployment you have to change the URL in `bookmarklet/src.js` to point to your deployed code and then follow the instructions in [Install bookmarklet](README.md#install-bookmarklet).

### Run as HTTP server

Build the binary with

```bash
go build -o bin/server cmd/server/main.go
```

and then run it (on your local machine, or anywhere you'd like):

```
./bin/server
```

The default server address is `http://localhost:3333`, and can be configured via the `BOOKMD_HTTP_ADDRESS` environment variable.

The endpoint will be available at `<your_url>/api/bookmarks`.

### Run as Vercel Function

Deploy it to your Vercel account. The endpoint will be available at `<vercel_url>/api/bookmarks`. No code changes are required.

## Install bookmarklet

Update the `addr`, `vault`, and `folder` variables in [bookmarklet/src.js](bookmarklet/src.js) as applicable to your setup, then run

```shell
cd bookmarklet
node maker.js
```

To install the bookmarklet in your browser:

- right-click on your bookmarks' bar,
- select "Add Bookmark", then
- paste the contents of [bookmarklet/bookmarklet.js](bookmarklet/bookmarklet.js) in the "URL" field

## Credits

Inspired by [downmark](https://github.com/alessandro-fazzi/downmark).
