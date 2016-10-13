#
#  Makefile for Go
#
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_FMT=$(GO_CMD) fmt

# Packages
TOP_PACKAGE_DIR := github.com/Ch3ck
PACKAGE_LIST :=  NumGo

.PHONY: all build clean test fmt

all: build

build:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Build $$p."; \
		$(GO_BUILD) $(TOP_PACKAGE_DIR)/$$p. || exit 1; \
	done
test:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Unit Testing $$p."; \
		$(GO_TEST) $(TOP_PACKAGE_DIR)/$$p. || exit 1; \
	done

clean:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Clean $$p."; \
		$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$$p.; \
	done

fmt:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Formatting $$p."; \
		$(GO_FMT) $(TOP_PACKAGE_DIR)/$$p. || exit 1; \
	done

# vim: set noexpandtab shiftwidth=8 softtabstop=0:
