include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt


## install: Install missing dependencies. Runs `go mod tidy` internally.
install: go-get

## clean: Clean build files. Runs `go clean` internally.
clean:
	@(MAKEFILE) go-clean

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

go-compile: go-clean go-get go-build

go-build:
	@echo "  >  Building binary..."
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache..."
	@GOBIN=$(GOBIN) go clean

go-get:
	@echo "  >  Checking for missing dependencies..."
	@GOBIN=$(GOBIN) go mod tidy

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo