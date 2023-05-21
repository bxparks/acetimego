SRCS := main.go

.PHONY := buildesp32 flashesp32 clean

#------------------------------------------------------------------------------

buildesp32: $(SRCS)
	tinygo build \
		-size full \
		-print-allocs=acetime \
		-target=esp32-coreboard-v2 \
		-o esp32.out \
		> esp32.size.txt

flashesp32:
	tinygo flash -x -target=esp32-coreboard-v2

#------------------------------------------------------------------------------

clean:
	rm -f *.out
