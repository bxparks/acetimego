TARGETS := printmemory validatetime
TINY_TARGETS := printmemory.tiny validatetime.tiny

all: $(TARGETS)

tiny: $(TINY_TARGETS)

printmemory: printmemory.go
	go build $<

printmemory.tiny: printmemory.go
	tinygo build -o printmemory.tiny $<

validatetime: validatetime.go
	go build $<

validatetime.tiny: validatetime.go
	tinygo build -o validatetime.tiny $<

clean:
	rm -f $(TARGETS) $(TINY_TARGETS)
