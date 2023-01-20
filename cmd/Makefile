TARGETS := printmemory printmemory.tiny

all: $(TARGETS)

printmemory:
	go build printmemory.go

printmemory.tiny:
	tinygo build -o printmemory.tiny printmemory.go

clean:
	rm -f $(TARGETS)
