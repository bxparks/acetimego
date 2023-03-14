help:
	@echo 'Usage: make (build|tiny|tinyesp32|all|clean)'

all: build buildtiny buildesp32

# Dummy target to support recursive 'test' target in ../Makefile.
test:

build:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

buildtiny:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

buildesp32:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

clean:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done
