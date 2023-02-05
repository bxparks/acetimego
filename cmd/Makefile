help:
	@echo 'Usage: make (build | tiny | tinyesp32 | clean)'

all: build

build:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

tiny:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

tinyesp32:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done

clean:
	set -e; \
	for i in */Makefile; do \
		$(MAKE) -C $$(dirname $$i) $@; \
	done
