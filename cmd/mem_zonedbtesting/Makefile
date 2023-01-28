TARGETS := mem_zonedbtesting.out

TARGETS_TINY := mem_zonedbtesting.tiny.out

# Uncomment this to compile for an ESP32
#TINY_TARGET := -target=esp32-coreboard-v2

help:
	@echo 'Usage: make (build | tiny | clean | all)'

all: build tiny

build: $(TARGETS)

tiny: $(TARGETS_TINY)

mem_zonedbtesting.out: mem_zonedbtesting.go
	go build -o $@ $<

mem_zonedbtesting.tiny.out: mem_zonedbtesting.go
	tinygo build $(TINY_TARGET) \
		-size full \
		-print-allocs=acetime \
		-o $@ $< \
		> mem_zonedbtesting.size.txt \
		2> mem_zonedbtesting.allocs.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
