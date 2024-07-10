# bookmark'd

Converts HTML pages to Markdown files and stores them in a local [Obsidian](https://obsidian.md) vault, in line with Steph Ango's [File over app](https://stephango.com/file-over-app) philosophy.

Whenever you'd like to bookmark a page, click on the provided bookmarklet, and you will be prompted to add it to your Obsidian vault.

## Usage

The code can run either locally or as a [Vercel Function](https://vercel.com/docs/functions/runtimes/go). In either case, after the deployment you have to change the URL in `bookmarklet/src.js` to point to your deployed code and then follow the instructions in [Install bookmarklet](README.md#install-bookmarklet).

### Run locally

#### Option 1: as macOS app

Running

```bash
make bundle-macos
```

will create a macOS app in `macos/bookmarkd.app`. Open the app, and it will keep the server running, showing its status in the system tray.

#### Option 2: Headless

Build the binary with

```bash
make build-server
```

and then run it (on your local machine, or anywhere you'd like):

```
./bin/server
```

The default server address is `http://localhost:11235`, and can be configured via the `BOOKMD_HTTP_ADDRESS` environment variable.

The endpoint will be available at `<your_url>/api/bookmarks`.

### Run as Vercel Function

Deploy it to your Vercel account. The endpoint will be available at `<vercel_url>/api/bookmarks`. No code changes are required.

## Generate and install bookmarklet

Run the following to update `bookmarklet.js` as required by your setup:

```shell
make build-bookmarklet  # build the binary
./bin/bookmarklet       # update `bookmarklet.js` (you will be prompted for details)
```

To install the bookmarklet in your browser:

- right-click on your bookmarks' bar,
- select "Add Bookmark", then
- paste the contents of [bookmarklet/bookmarklet.js](bookmarklet/bookmarklet.js) in the "URL" field

## Bonus: Import Pocket saves

The repository contains an optional importer for [Pocket](https://getpocket.com/) saves in the `pocket-importer` directory. See its own README for details.

## Credits

Initial design inspired by [downmark](https://github.com/alessandro-fazzi/downmark).

Packaging script adapted from [xeoncross/macappshell](https://github.com/xeoncross/macappshell).

Tray icon courtesy of [ionicons](https://ionic.io/ionicons/usage#bookmarks-outline).
