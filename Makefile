BINARY := yask
PKG:= github.com/av-ugolkov/yask/cmd

VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: build
build:
	go build -ldflags "-X '$(PKG).Version=$(VERSION)'" -o $(BINARY) main.go