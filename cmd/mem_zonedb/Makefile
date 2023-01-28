TARGETS := mem_zonedb.out

TARGETS_TINY := mem_zonedb.tiny.out

# Uncomment this to compile for an ESP32
#TINY_TARGET := -target=esp32-coreboard-v2

help:
	@echo 'Usage: make (build | tiny | clean | all)'

all: build tiny

build: $(TARGETS)

tiny: $(TARGETS_TINY)

mem_zonedb.out: mem_zonedb.go
	go build -o $@ $<

mem_zonedb.tiny.out: mem_zonedb.go
	tinygo build $(TINY_TARGET) -size full -o $@ $< > mem_zonedb.tiny.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
