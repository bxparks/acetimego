SRCS := main.go

.PHONY := build buildtiny buildesp32 flashesp32 clean help

help:
	@echo 'Usage: make [build|buildtiny|buildesp32|flashesp32|clean]'

#------------------------------------------------------------------------------
# Use normal Go to compile to native host (Linux).

build: $(SRCS) Makefile
	go build -o main.out

#------------------------------------------------------------------------------
# Use TinyGo to compile to native host (Linux).

buildtiny: $(SRCS) Makefile
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-o tiny.out \
		> tiny.size.txt

#------------------------------------------------------------------------------
# Use TinyGo to compile to ESP32.

buildesp32: $(SRCS) Makefile
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-target=esp32-coreboard-v2 \
		-o esp32.out \
		> esp32.size.txt

flashesp32: $(SRCS) Makefile
	tinygo flash -x -target=esp32-coreboard-v2

#------------------------------------------------------------------------------

clean:
	rm -f *.out
