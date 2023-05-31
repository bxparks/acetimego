SRCS := main.go transitions.go

#------------------------------------------------------------------------------
# Use normal Go to compile to native host (Linux).

.PHONY: build
build: $(SRCS) Makefile
	go build -o main.out

buildtiny:

#------------------------------------------------------------------------------

.PHONY: clean
clean:
	rm -f *.out
