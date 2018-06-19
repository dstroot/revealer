#
# Variables
#

# NOTE: Simply expanded variables are defined by lines using `:='
# When a simply expanded variable is referenced, its value is substituted
# verbatim. There is another assignment operator for variables, `?='.
# This is called a conditional variable assignment operator, because it
# only has an effect if the variable is not yet defined.

# NOTE: I am not a fan of versions - people frequently forget to increment them.
# The commit ID and the buid time are more precise and automatic. However
# versions can be useful for humans so I still keep a `VERSION` file in the
# root so that anyone can clearly check the VERSION of `master`.

OWNER := dstroot
REPO := github.com
NAME := $(shell basename $(CURDIR))

PROJECT := ${REPO}/${OWNER}/${NAME}
DOCKER_NAME := ${OWNER}/${NAME}

BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COMMIT_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo nosha)
VERSION := $(shell cat ./VERSION)


#
# Help
#


.PHONY: help
default help:
	@echo "Usage: make <command>\n"
	@echo "The commands are:"
	@echo "   gettools    Download and install Go-based build toolchain (uses go-get)."
	@echo "   test        Execute all development tests."
	@echo "   cover       Examine code test coverage."
	@echo "   lint        Run gometalinter against the source."
	@echo "   release     Build production release(s). Runs dependent rules."
	@echo "   todo        Display all TODO's in the source."
	@echo "   docs        Display the application documentation.\n"


#
# Development
#
#

# NOTE: Add @ to the beginning of a command to tell make not to print
# the command being executed.'

.PHONY: gettools
gettools:
	@go get -u github.com/alecthomas/gometalinter
	@go get -u golang.org/x/tools/cmd/cover
	@gometalinter --install


#
# Code Hygiene
#


# Go test cover with multiple packages support
.PHONY: test
test:
	@echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

# Get code coverage report
.PHONY: cover
cover: test
	@go tool cover -html=coverage.txt

# Lint all the things
.PHONY: lint
lint:
	@gometalinter --vendor ./...


#
# Release
#


.PHONY: release
release:
	@echo "Releasing: $(VERSION)"
	git add -A
	git commit -m "Releasing $(VERSION)"
	git push origin
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)


#
# Misc Stuff
#


# Show any to-do items per file.
.PHONY: todo
todo:
	@grep \
	--exclude-dir=vendor \
	--exclude-dir=node_modules \
	--exclude=Makefile \
	--text \
	--color \
	-nRo -E ' TODO.*|SkipNow|nolint:.*' .

# Show documentation
.PHONY: docs
docs:
	@godoc $(shell PWD)
