TARGETS := \
mem_go_time.out \
mem_zonedb.out \
mem_zonedbtesting.out \
validatetime.out

TARGETS_TINY := \
mem_go_time.tiny.out mem_go_time.tiny.txt \
mem_zonedb.tiny.out mem_zonedb.tiny.txt \
mem_zonedbtesting.tiny.out mem_zonedbtesting.tiny.txt \
validatetime.tiny.out validatetime.tiny.txt

# Uncomment this to compile for an ESP32
#TINY_TARGET := -target=esp32-coreboard-v2

all: $(TARGETS)

alltiny: $(TARGETS_TINY)

mem_go_time.out: mem_go_time.go
	go build -o $@ $<

mem_go_time.tiny.out: mem_go_time.go
	tinygo build $(TINY_TARGET) -size full -o $@ $<  > mem_go_time.tiny.txt

mem_zonedb.out: mem_zonedb.go
	go build -o $@ $<

mem_zonedb.tiny.out: mem_zonedb.go
	tinygo build $(TINY_TARGET) -size full -o $@ $< > mem_zonedb.tiny.txt

mem_zonedbtesting.out: mem_zonedbtesting.go
	go build -o $@ $<

mem_zonedbtesting.tiny.out: mem_zonedbtesting.go
	tinygo build $(TINY_TARGET) -size full -o $@ $< > mem_zonedbtesting.tiny.txt

validatetime.out: validatetime.go
	go build -o $@ $<

validatetime.tiny.out: validatetime.go
	tinygo build $(TINY_TARGET) -size full -o $@ $< > validatetime.tiny.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
