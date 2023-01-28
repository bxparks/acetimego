TARGETS := mem_go_time.out

TARGETS_TINY := mem_go_time.tiny.out

# Uncomment this to compile for an ESP32
#TINY_TARGET := -target=esp32-coreboard-v2

help:
	@echo 'Usage: make (build | tiny | clean | all)'

all: build tiny

build: $(TARGETS)

tiny: $(TARGETS_TINY)

mem_go_time.out: mem_go_time.go
	go build -o $@ $<

mem_go_time.tiny.out: mem_go_time.go
	tinygo build $(TINY_TARGET) \
		-size full \
		-print-allocs=^time$ \
		-o $@ $< \
		> mem_go_time.size.txt \
		2> mem_go_time.allocs.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
