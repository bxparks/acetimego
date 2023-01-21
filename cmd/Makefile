TARGETS := \
mem_zone_info \
mem_zone_registry \
validatetime

TARGETS_TINY := \
mem_zone_info.tiny \
mem_zone_registry.tiny \
validatetime.tiny

all: $(TARGETS)

alltiny: $(TARGETS_TINY)

mem_zone_info: mem_zone_info.go
	go build $<

mem_zone_info.tiny: mem_zone_info.go
	tinygo build -o mem_zone_info.tiny $<

mem_zone_registry: mem_zone_registry.go
	go build $<

mem_zone_registry.tiny: mem_zone_registry.go
	tinygo build -o mem_zone_registry.tiny $<

validatetime: validatetime.go
	go build $<

validatetime.tiny: validatetime.go
	tinygo build -o validatetime.tiny $<

clean:
	rm -f $(TARGETS) $(TARGETS_TINY)
