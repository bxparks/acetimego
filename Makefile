help:
	@echo 'Usage: make (build|tiny|test|all|zonedb|clean)'

.PHONY: all build tiny test clean zonedb

all: build tiny test

# Use nested for-loop because 'go build ./...' does not produce error messages
# compatible with vim quickfix (at least not a format that I can easily
# customize vim to handle).
build:
	set -e; \
	for i in acetime/Makefile zoneinfo/Makefile cmd/Makefile; do \
		$(MAKE) -C $$(dirname $$i) build; \
	done

tiny:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) tiny; \
	done

zonedb:
	set -e; \
	for i in zonedb*/Makefile; do \
		$(MAKE) -C $$(dirname $$i); \
	done

# If we use 'go test ./...', the subdirectory is not recognized by vim so the
# direct navigation in quickfix mode does not work. Use a for-loop instead.
test:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) test; \
	done

clean:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) clean; \
	done
