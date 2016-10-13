#
#  Makefile for Go
#
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean

# Packages
TOP_PACKAGE_DIR := github.com/Ch3ck
PACKAGE_LIST := NumGo

.PHONY: all build clean fmt

all: build

build:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Build $$p ..."; \
		$(GO_BUILD) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

clean:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Clean $$p ..."; \
		$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$$p; \
	done

# vim: set noexpandtab shiftwidth=8 softtabstop=0:
