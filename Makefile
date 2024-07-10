.PHONY: build-server
build-server:
	go build -o bin/server cmd/server/main.go

.PHONY: build-bookmarklet
build-bookmarklet:
	go build -o bin/bookmarklet cmd/bookmarklet/main.go

.PHONY: build-tray
build-tray:
	go build -o bin/tray cmd/tray/main.go

.PHONY: bundle-macos
bundle-macos:
	cd macos && ./bundle.sh
