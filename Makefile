SHELL := /usr/bin/env bash

.PHONY: help
help:
	@echo 'Usage: make (build|tiny|test|all|zonedbs|clean)'

.PHONY: all
all: build buildtiny

.PHONY: zonedbs
zonedbs:
	set -e; \
	for i in zonedb*/Makefile; do \
		$(MAKE) -C $$(dirname $$i); \
	done

#------------------------------------------------------------------------------

# Use nested for-loop because 'go build ./...' does not produce error messages
# compatible with vim quickfix (at least not a format that I can easily
# customize vim to handle).
.PHONY: build
build:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

# If we use 'go test ./...', the subdirectory is not recognized by vim so the
# direct navigation in quickfix mode does not work. Use a for-loop instead.
.PHONY: test
test:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

#------------------------------------------------------------------------------

.PHONY: buildtiny
buildtiny:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

.PHONY: testtiny
testtiny:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

#------------------------------------------------------------------------------

.PHONY: clean
clean:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done
