TARGETS := \
mem_go_time.out \
mem_zone_info.out \
mem_zone_registry.out \
validatetime.out

TARGETS_TINY := \
mem_go_time.tiny.out mem_go_time.tiny.txt \
mem_zone_info.tiny.out mem_zone_info.tiny.txt \
mem_zone_registry.tiny.out mem_zone_registry.tiny.txt \
validatetime.tiny.out validatetime.tiny.txt

# Uncomment this to compile for an ESP32
#BUILD_TARGET := -target=esp32-coreboard-v2

all: $(TARGETS)

alltiny: $(TARGETS_TINY)

mem_go_time.out: mem_go_time.go
	go build -o $@ $<

mem_go_time.tiny.out: mem_go_time.go
	tinygo build $(BUILD_TARGET) -size full -o $@ $<  > mem_go_time.tiny.txt

mem_zone_info.out: mem_zone_info.go
	go build -o $@ $<

mem_zone_info.tiny.out: mem_zone_info.go
	tinygo build $(BUILD_TARGET) -size full -o $@ $<  > mem_zone_info.tiny.txt

mem_zone_registry.out: mem_zone_registry.go
	go build -o $@ $<

mem_zone_registry.tiny.out: mem_zone_registry.go
	tinygo build $(BUILD_TARGET) -size full -o $@ $< > mem_zone_registry.tiny.txt

validatetime.out: validatetime.go
	go build -o $@ $<

validatetime.tiny.out: validatetime.go
	tinygo build $(BUILD_TARGET) -size full -o $@ $< > validatetime.tiny.txt

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
