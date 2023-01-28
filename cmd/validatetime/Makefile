TARGETS := validatetime.out

TARGETS_TINY := validatetime.tiny.out

# Uncomment this to compile for an ESP32
#TINY_TARGET := -target=esp32-coreboard-v2

help:
	@echo 'Usage: make (build | tiny | all | clean )'

all: build tiny

build: $(TARGETS)

tiny: $(TARGETS_TINY)

validatetime.out: validatetime.go
	go build -o $@ $<

validatetime.tiny.out: validatetime.go
	tinygo build $(TINY_TARGET) \
		-size full \
		-print-allocs=acetime \
		-o $@ $< \
		> validatetime.size.txt \
		2> validatetime.allocs.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
