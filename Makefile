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

test:
	go test ./...

clean:
	set -e; \
	for i in cmd/*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) clean; \
	done
