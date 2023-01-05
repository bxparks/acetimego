package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

//-----------------------------------------------------------------------------
// TimeZone represents one of the IANA TZ time zones. This is a reference type,
// and meant to be passed around as a pointer and garbage collected when it is
// no longer used.
//-----------------------------------------------------------------------------

type TimeZone struct {
	zoneProcessor ZoneProcessor
}

func TimeZoneForZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {
	var tz TimeZone
	tz.zoneProcessor.InitForZoneInfo(zoneInfo)
	return tz
}

func (tz *TimeZone) OffsetDateDateTimeFromLocalDateTime(
	ldt *LocalDateTime, fold uint8) OffsetDateTime {

	return tz.zoneProcessor.OffsetDateTimeFromLocalDateTime(ldt, fold)
}

func (tz *TimeZone) OffsetDateDateTimeFromEpochSeconds(
	epochSeconds int32) OffsetDateTime {

	return tz.zoneProcessor.OffsetDateTimeFromEpochSeconds(epochSeconds)
}

func (tz *TimeZone) ZonedExtraFromEpochSeconds(epochSeconds int32) ZonedExtra {
	ti := tz.zoneProcessor.TransitionInfoFromEpochSeconds(epochSeconds)
	return ZonedExtra{
		stdOffsetMinutes: ti.stdOffsetMinutes,
		dstOffsetMinutes: ti.dstOffsetMinutes,
		abbrev:           ti.abbrev,
	}
}
