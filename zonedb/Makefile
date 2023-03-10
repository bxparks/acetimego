TARGETS := zone_infos.go zone_policies.go zone_registry.go

TOOLS := $(abspath ../../AceTimeTools)
TZ_REPO := $(abspath ../../tz)
TZ_VERSION := 2022g
START_YEAR := 1800  # smaller than min year in TZDB
UNTIL_YEAR := 10000


# The zonedb format for AceTimeGo now supports:
#	* one-second resolution for STDOFF, UNTIL, AT
#	* one-minute resolution for DSTOFF and SAVE.
#
# Here are a few zones in the 1930 whose DSTOFF or SAVE are at 00:20 minutes,
# instead of multiples of 00:15.
#
# * Asia/Brunei : Minute not equal:  1935-09-14 00:20:00 +0820 +0820
# 1935-09-14T00:15:00+08:15[Asia/Brunei]
# * Asia/Kuala_Lumpur : Minute not equal:  1933-01-01 00:20:00 +0720 +0720
# 1933-01-01T00:15:00+07:15[Asia/Kuala_Lumpur]
# * Asia/Kuching : Minute not equal:  1935-09-14 00:20:00 +0820 +0820 //
# 1935-09-14T00:15:00+08:15[Asia/Kuching]
# * Asia/Singapore : Minute not equal:  1933-01-01 00:20:00 +0720 +0720
# * 1933-01-01T00:15:00+07:15[Asia/Singapore]
# * Singapore: :  Singapore : Minute not equal:  1933-01-01 00:20:00 +0720
# +0720
all:
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language go \
		--scope extended \
		--offset_granularity 1 \
		--delta_granularity 60 \
		--until_at_granularity 1 \
		--db_namespace zonedb \
		--generate_int16_years \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

zonedb.json: Makefile
	$(TOOLS)/tzcompiler.sh \
		--tzrepo $(TZ_REPO) \
		--tag $(TZ_VERSION) \
		--action zonedb \
		--language json \
		--scope extended \
		--offset_granularity 1 \
		--delta_granularity 60 \
		--until_at_granularity 1 \
		--generate_int16_years \
		--start_year $(START_YEAR) \
		--until_year $(UNTIL_YEAR)

test:
	go test

clean:
	rm -rf tzfiles zonedb.json
