BINARY := yask
PKG := github.com/av-ugolkov/yask/cmd

VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X '$(PKG).Version=$(VERSION)'" -o $(BINARY) main.go

.PHONY: build-mac
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -X '$(PKG).Version=$(VERSION)'" -o $(BINARY) main.go

.PHONY: build-win
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X '$(PKG).Version=$(VERSION)'" -o $(BINARY).exe main.go