BINARY := yask

VERSION := $(shell git describe --tags)
COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

.PHONY: build-dev
build-dev:
	go build -ldflags \
	"-X 'main.Commit=$(COMMIT)' -X 'main.BuildDate=$(DATE)'"\
	 -o $(BINARY) main.go

.PHONY: build-prod
build-prod:
	go build -ldflags \
	"-X 'main.Version=$(VERSION)' -X 'main.Commit=$(COMMIT)' -X 'main.BuildDate=$(DATE)'"\
	 -o $(BINARY) main.go

.PHONY: print
print:
	@echo "Version:    $(VERSION)"
	@echo "Commit:     $(COMMIT)"
	@echo "BuildDate:  $(DATE)"