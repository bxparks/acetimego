CMD_NAME := mem_zonedbtesting
TARGETS_UNIX := $(CMD_NAME).out
TARGETS_TINY := $(CMD_NAME).tiny.out
TARGETS_TINY_ESP32 := $(CMD_NAME).esp32.out

.PHONY := build tiny tinyesp32 flashesp32 all clean help

all: build

help:
	@echo 'Usage: make [all | build | tiny | tinyesp32 | flashesp32 | clean]'

build: $(TARGETS_UNIX)

tiny: $(TARGETS_TINY)

tinyesp32: $(TARGETS_TINY_ESP32)

flashesp32:
	tinygo flash -x -target=esp32-coreboard-v2

clean:
	rm -f $(TARGETS_UNIX) $(TARGETS_TINY) $(TARGETS_TINY_ESP32)

#------------------------------------------------------------------------------

# Use normal Go to compile to native host (Linux).
$(CMD_NAME).out: $(CMD_NAME).go
	go build -o $@ $<

# Use TinyGo to compile to native host (Linux).
$(CMD_NAME).tiny.out: $(CMD_NAME).go
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-o $@ $< \
		> $(CMD_NAME).tiny.size.txt \
		2> $(CMD_NAME).tiny.allocs.txt

# Use TinyGo to compile to ESP32.
$(CMD_NAME).esp32.out: $(CMD_NAME).go
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-target=esp32-coreboard-v2 \
		-o $@ $< \
		> $(CMD_NAME).esp32.size.txt \
		2> $(CMD_NAME).esp32.allocs.txt
