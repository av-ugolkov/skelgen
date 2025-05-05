BINARY := yask
PKG := github.com/av-ugolkov/yask/cmd

VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X '$(PKG).Version=$(VERSION)'" -o $(BINARY) main.go

.PHONY: build-mac
build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X '$(PKG).Version=$(VERSION)'" -o $(BINARY) main.go

.PHONY: build-win
build-win:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X '$(PKG).Version=$(VERSION)'" -o $(BINARY).exe main.go