SRCS := main.go

#------------------------------------------------------------------------------

.PHONY: build
build:
	go build -o main.out

#------------------------------------------------------------------------------

.PHONY: buildtiny
buildtiny:
	tinygo build -o tiny.out

#------------------------------------------------------------------------------

.PHONY: buildesp32
buildesp32: $(SRCS)
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-target=esp32-coreboard-v2 \
		-o esp32.out \
		> esp32.size.txt

.PHONY: flashesp32
flashesp32:
	tinygo flash -x -target=esp32-coreboard-v2

#------------------------------------------------------------------------------

.PHONY: clean
clean:
	rm -f *.out
