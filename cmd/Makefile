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

# NOTE: Only cmd/mem* modules support TinyGo
.PHONY: buildtiny
buildtiny:
	set -e; \
	for i in mem*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

testtiny:

#------------------------------------------------------------------------------

# NOTE: Only cmd/mem* modules support TinyGo
.PHONY: buildesp32
buildesp32:
	set -e; \
	for i in mem*/Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

#------------------------------------------------------------------------------

.PHONY: clean
clean:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done
