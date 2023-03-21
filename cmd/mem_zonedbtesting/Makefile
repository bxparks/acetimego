SRCS := mem_zonedtesting.go

.PHONY := build buildtiny buildesp32 flashesp32 all clean help

help:
	@echo 'Usage: make [build|buildtiny|buildesp32|flashesp32|clean]'

#------------------------------------------------------------------------------
# Use normal Go to compile to native host (Linux).

build: $(TARGETS_UNIX)
	go build -o mem_zonedtesting.out

#------------------------------------------------------------------------------
# Use TinyGo to compile to native host (Linux).

buildtiny: $(TARGETS_TINY)
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-o tiny.out \
		> tiny.size.txt

#------------------------------------------------------------------------------
# Use TinyGo to compile to ESP32.

buildesp32: $(TARGETS_TINY_ESP32)
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-target=esp32-coreboard-v2 \
		-o esp32.out \
		> esp32.size.txt

flashesp32:
	tinygo flash -x -target=esp32-coreboard-v2

#------------------------------------------------------------------------------

clean:
	rm -f *.out
