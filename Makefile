VERSION ?= $(shell git describe --tags)
BUILD ?= $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard cmd/renum/*.go)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=github.com/gen0cide/renum.Version=$(VERSION) -X=github.com/gen0cide/renum.Build=$(BUILD)"

## lint: Run golangci-lint against the project.
lint:
	@-$(MAKE) -s go-lint

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install:

## compile: Compile the binary.
compile:
	@-$(MAKE) -s go-compile

## full: Clean's the build cache, gets dependencies, runs generators and linters, then compiles a fresh binary.
full:
	@-$(MAKE) -s go-full-compile

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOBIN=$(GOBIN) $(GOBIN)/$(PROJECTNAME) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-$(MAKE) go-clean

go-compile: go-template go-build

go-full-compile: go-lint go-template go-build go-install

go-lint:
	@echo "  >  Running linters..."
	@golangci-lint run --config=$(GOBASE)/.golangci.yml --fix

go-build:
	@echo "  >  Building binary..."
	@GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-template:
	@echo "  >  Bundling embedded assets..."
	@GOBIN=$(GOBIN) go generate ./generator/...

go-generate:
	@echo "  >  Generating dependency files..."
	@GOBIN=$(GOBIN) go generate ./...

go-install: go-build
	@echo "  > Installing binary..."
	@cp -f $(GOBIN)/$(PROJECTNAME) $(GOPATH)/bin/$(PROJECTNAME)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " ["$(PROJECTNAME)"] Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo






