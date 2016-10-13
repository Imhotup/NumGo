#
#  Makefile for Go
#

# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
BUILDTAGS=

.PHONY: clean all fmt build test
.DEFAULT: default

all: clean build fmt test

build:
	@echo "+ $@"
	@go build -tags "$(BUILDTAGS) cgo" .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

test: fmt
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

clean:
	@echo "+ $@"
	@rm -rf NumGo

# vim: set noexpandtab shiftwidth=8 softtabstop=0:
