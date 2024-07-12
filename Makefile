.PHONY: build-server
build-server:
	go build -o bin/server cmd/server/main.go

.PHONY: build-bookmarklet
build-bookmarklet:
	go build -o bin/bookmarklet cmd/bookmarklet/main.go

.PHONY: build-tray
build-tray:
	go build -o bin/tray cmd/tray/main.go

.PHONY: macos
macos: bundle-macos build-macos

.PHONY: build-macos
build-macos: build-tray
	cp bin/tray macos/bookmarkd.app/Contents/MacOS/bookmarkd

.PHONY: bundle-macos
bundle-macos:
	cd macos && ./bundle.sh
