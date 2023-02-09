TARGETS := zone_infos.go zone_policies.go zone_registry.go

TOOLS := $(abspath ../../AceTimeTools)
TZ_REPO := $(abspath ../../tz)
TZ_VERSION := 2022g
START_YEAR := 1970
UNTIL_YEAR := 10000

all:
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language go \
		--scope extended \
		--offset_granularity 1 \
		--delta_granularity 900 \
		--until_at_granularity 60 \
		--db_namespace zonedb \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

zonedb.json: Makefile
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language json \
		--scope extended \
		--granularity 1 \
		--generate_int16_years \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

test:

clean:
	rm -rf tzfiles zonedb.json
