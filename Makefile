.PHONY: build trust

NAME     := baize
PACKAGE  := github.com/btwiuse/$(NAME)
IMG_NAME := btwiuse/baize
IMAGE    := ${IMG_NAME}
SHELL    := bash
BAZEL    := ./bazelw

default: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'

gazelle:             ## auto generate BUILD.bazel files from go.mod
	@ go mod tidy
	@ $(BAZEL) run //:gazelle -- update-repos --from_file=go.mod
	@ $(BAZEL) run //:gazelle

