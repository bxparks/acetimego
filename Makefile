help:
	@echo 'Usage: make (build | tiny | test | all | clean)'

all: build tiny test

build:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) build; \
	done

tiny:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) tiny; \
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
