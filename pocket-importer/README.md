# Pocket importer

Imports Pocket saves into Bookmd.

**Disclaimer:** Use at your own risk, no guarantees whatsoever. I've used it to migrate my own Pocket saves and it worked fine except for a few websites.

## Usage

This currently only works with Chrome. Before running, make sure you have Chrome installed and that it is in your PATH.

### Export from Pocket

Open the [Pocket Export](https://getpocket.com/export) page and click on "Export HTML file" link. Copy the file to the `pocket-importer` directory as `pocket-export.html`.

### Run server

From the root of the `bookmd` project, run

```shell
make
bin/server
```

### Run Chrome in debug mode

```shell
/path/to/chrome --remote-debugging-port=1243
```

Playwright [cannot currently deal with 'open in app' dialogs](https://github.com/microsoft/playwright/issues/11014), so we have to ensure they are dismissed: to do so, type `obsidian://` in the address bar, tick the checkbox of the dialog so that it won't show up again, and select `Open in Obsidian`.

### Run the importer

From the `pocket-importer` directory, run

```shell
npm install
npx tsx ./importer.ts
```
