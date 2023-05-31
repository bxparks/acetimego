.PHONY: all
all: build buildtiny buildesp32

#------------------------------------------------------------------------------

.PHONY: build
build:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

test:

#------------------------------------------------------------------------------

.PHONY: buildtiny
buildtiny:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

testtiny:

#------------------------------------------------------------------------------

.PHONY: buildesp32
buildesp32:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

#------------------------------------------------------------------------------

.PHONY: clean
clean:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

