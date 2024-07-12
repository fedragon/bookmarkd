# bookmark'd

Converts HTML pages to Markdown files and stores them in a local [Obsidian](https://obsidian.md) vault, in line with Steph Ango's [File over app](https://stephango.com/file-over-app) philosophy.

Whenever you'd like to bookmark a page, click on the provided bookmarklet, and you will be prompted to add it to your Obsidian vault.

## Usage

Below are a few ways to run the bookmark manager. In any case, you will need to install the bookmarklet in your browser (see instructions in [Install bookmarklet](README.md#install-bookmarklet)).

### Run locally

#### Option 1: as macOS app

Running

```bash
make build-macos
```

will create a macOS app in `macos/bookmarkd.app`. Open the app, and it will keep the server running, showing its status in the system tray.

#### Option 2: standalone HTTP server

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

### Run as [Vercel Function](https://vercel.com/docs/functions/runtimes/go)

Deploy it to your Vercel account. The endpoint will be available at `<vercel_url>/api/bookmarks`. No code changes are required.

### Install bookmarklet

Open [this page](bookmarklet/index.html) in your browser and click on the link in it. You will be prompted for details about your setup and finally see an alert dialog with javascript code: that's your bookmarklet!

Copy the whole content of the final prompt and install the bookmarklet in your browser by:

- right-clicking on your bookmarks' bar,
- selecting "Add Bookmark", and finally
- pasting the copied contents in the "URL" field (the name is up to you).

## Bonus: Import Pocket saves

The repository contains an optional importer for [Pocket](https://getpocket.com/) saves in the `pocket-importer` directory. See its own README for details.

## Credits

Initial design inspired by [downmark](https://github.com/alessandro-fazzi/downmark).

Packaging script adapted from [xeoncross/macappshell](https://github.com/xeoncross/macappshell).

Tray icon created by [SVG Repo](https://www.svgrepo.com/svg/118921/bookmark-sketched-symbol-outline) and licensed under [CC0 License](https://creativecommons.org/publicdomain/zero/1.0/).
