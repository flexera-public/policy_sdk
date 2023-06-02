#!/usr/bin/make
#
# Makefile for Flexera Policy Tool
#
# Features:
# - runs tests recursively, computes code coverage report
# - build for linux/amd64, darwin/amd64, windows/amd64
# - produces .tgz/.zip build output
# - produces version.go for each build with string in global variable VV, please
#   print this using a --version option in the executable
#
# Top-level targets:
# default: compile the program, you can thus use make && ./NAME -options ...
# build: builds binaries for the current platform
# test: runs unit tests recursively and produces code coverage stats and shows them
# depend: installs executable dependencies
# clean: removes build stuff

NAME=fpt
EXE:=$(NAME)$(shell go env GOEXE)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# === below this line ideally remains unchanged, add new targets at the end ===

TRAVIS_BRANCH?=dev
DATE=$(shell date '+%F %T')
TRAVIS_COMMIT?=$(shell git rev-parse HEAD)
GIT_BRANCH:=$(shell git symbolic-ref --short -q HEAD || echo 'master')
SHELL:=$(shell which bash)

# if we are building a tag overwrite TRAVIS_BRANCH with TRAVIS_TAG
ifneq ($(TRAVIS_TAG),)
TRAVIS_BRANCH:=$(TRAVIS_TAG)
endif

ifeq ($(GOOS),windows)
export CC:=x86_64-w64-mingw32-gcc
export CXX:=x86_64-w64-mingw32-g++
endif

# This works around an issue between dep and Cygwin git by using Windows git instead.
ifeq ($(shell go env GOHOSTOS),windows)
  ifeq ($(shell git version | grep windows),)
    export PATH:=$(shell cygpath 'C:\Program Files\Git\cmd'):$(PATH)
  endif
endif

ifeq ($(GOOS),windows)
UPLOADS?=build/$(NAME)-$(GOOS)-$(GOARCH).zip
else
UPLOADS?=build/$(NAME)-$(GOOS)-$(GOARCH).tgz
endif

GO_SOURCE:=$(shell find . -name "*.go")

default: $(EXE)
$(EXE): $(GO_SOURCE) version
	go build -tags $(NAME)_make -o $(EXE) ./cmd/$(NAME)

install: $(EXE)
	go install

build: $(EXE) $(UPLOADS)

# create a tgz with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos
build/$(NAME)-%.tgz: $(GO_SOURCE) version
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt="$*"; GOOS="$${tgt%-*}" GOARCH="$${tgt#*-}" go build -tags $(NAME)_make -o build/$(NAME)/$(NAME)`GOOS="$${tgt%-*}" GOARCH="$${tgt#*-}" go env GOEXE` ./cmd/$(NAME)
	chmod +x build/$(NAME)/$(NAME)*
	tar -zcf $@ -C build $(NAME)
	rm -r build/$(NAME)

# create a zip with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.zip: $(GO_SOURCE) version
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt="$*"; GOOS="$${tgt%-*}" GOARCH="$${tgt#*-}" go build -tags $(NAME)_make -o build/$(NAME)/$(NAME)`GOOS="$${tgt%-*}" GOARCH="$${tgt#*-}" go env GOEXE` ./cmd/$(NAME)
	chmod +x build/$(NAME)/$(NAME)*
	cd build; 7z a -r $(notdir $@) $(NAME)
	rm -r build/$(NAME)

# produce a version string that is embedded into the binary that captures the branch, the date
# and the commit we're building
version:
	@echo -e "// +build $(NAME)_make\n\npackage main\n\nconst VV = \"$(NAME) $(TRAVIS_BRANCH) - $(DATE) - $(TRAVIS_COMMIT)\"" \
	  >./cmd/$(NAME)/version.go
	@echo "version.go: `tail -1 ./cmd/$(NAME)/version.go`"

# install vendored dependencies, as needed
vendor: go.mod go.sum
	go mod vendor

clean:
	rm -rf build $(EXE)
	rm -f version.go

lint:
	@if gofmt -l $(shell git ls-files | grep .go | grep -v sdk) 2>&1 | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	go vet -composites=false ./...

test: lint
	go test -cover -race ./...

sdk:
	rm -rf sdk
	cp -af $(GOPATH)/src/github.com/rightscale/governance/front_service/gen sdk
	rm -rf sdk/*/convert.go sdk/http/*/server
	find sdk -type f -name '*.go' -exec sed -i -e 's#github.com/rightscale/governance/front_service/gen#github.com/flexera-public/policy_sdk/sdk#' {} \;

# ===== SPECIAL TARGETS FOR fpt =====

.PHONY: fpt test sdk
