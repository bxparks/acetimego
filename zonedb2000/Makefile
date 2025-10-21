TARGETS := zone_infos.go zone_policies.go zone_registry.go

TOOLS := $(abspath ../../../compiler)
TZ_REPO := $(abspath ../../../../tz)
TZ_VERSION := 2025b
START_YEAR := 2000
UNTIL_YEAR := 2200


# The zonedb format for acetimego now supports:
#	* one-second resolution for STDOFF, UNTIL, AT
#	* one-minute resolution for DSTOFF and SAVE.
zonedb:
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--actions zonedb \
		--language go \
		--scope complete \
		--db_namespace zonedb2000 \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

zonedb.json: Makefile
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--actions json \
		--language go \
		--scope complete \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

#------------------------------------------------------------------------------

build:
	go build

test:
	go test

#------------------------------------------------------------------------------

buildtiny:
	# tinygo build

testtiny:
	tinygo test

#------------------------------------------------------------------------------

clean:
	rm -rf tzfiles zonedb.json
