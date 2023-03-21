SHELL := /usr/bin/env bash

help:
	@echo 'Usage: make (build|tiny|test|all|zonedbs|clean)'

.PHONY: all build buildtiny test clean zonedbs

all: build tiny test

# Use nested for-loop because 'go build ./...' does not produce error messages
# compatible with vim quickfix (at least not a format that I can easily
# customize vim to handle).
build:
	set -e; \
	for i in acetime/Makefile zoneinfo/Makefile cmd/Makefile; do \
		$(MAKE) -C $$(dirname $$i) build; \
	done

buildtiny:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) buildtiny; \
	done

zonedbs:
	set -e; \
	for i in zonedb*/Makefile; do \
		$(MAKE) -C $$(dirname $$i); \
	done

# If we use 'go test ./...', the subdirectory is not recognized by vim so the
# direct navigation in quickfix mode does not work. Use a for-loop instead.
#
# TODO: Skip ds3231 until I get TinyGo working under GitHub actions.
test:
	set -e; \
	for i in */Makefile; do \
		if [[ $$i =~ ds3231 ]]; then \
			echo "Skipping $$i"; \
			continue; \
		fi; \
		$(MAKE) -C $$(dirname $$i) test; \
	done

clean:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) clean; \
	done
