help:
	@echo 'Usage: make (build|tiny|tinyesp32|all|clean)'

all: build buildtiny buildesp32

build:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

test:

#------------------------------------------------------------------------------

buildtiny:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

testtiny:

#------------------------------------------------------------------------------

buildesp32:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

#------------------------------------------------------------------------------

clean:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done
