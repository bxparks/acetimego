TARGETS := zone_infos.go zone_policies.go zone_registry.go

TOOLS := $(abspath ../../AceTimeTools)

TZ_VERSION := 2022g
START_YEAR := 2000
UNTIL_YEAR := 10000

all:
	$(TOOLS)/tzcompiler.sh \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language go \
		--scope extended \
		--db_namespace zonedb \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

zonedb.json: Makefile
	$(TOOLS)/tzcompiler.sh \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language json \
		--scope extended \
		--generate_int16_years \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

clean:
	rm -rf tzfiles zonedb.json
