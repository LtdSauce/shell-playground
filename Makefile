SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFALGS += --no-builtin-rules

# Suggestion to avoid tabs from https://tech.davis-hansson.com/p/make/
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

all: run
.PHONY: all

build-client:
> GOARCH=wasm GOOS=js go build -o web/app.wasm
.PHONY: build-client

build-server:
> go build -o build/server
.PHONY: build-server

run: build-client build-server
> ./build/server
.PHONY: run
