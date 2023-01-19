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

func NewTimeZoneFromZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {
	var tz TimeZone
	tz.zoneProcessor.InitForZoneInfo(zoneInfo)
	return tz
}

// OffsetDateTimeFromEpochSeconds calculates the OffsetDateTime from the given
// epochSeconds.
//
// Adapted from atc_time_zone_offset_date_time_from_epoch_seconds() in the
// AceTimeC library and, TimeZone::getOffsetDateTime(epochSeconds) from the
// AceTime library.
func (tz *TimeZone) OffsetDateTimeFromEpochSeconds(
	epochSeconds int32) OffsetDateTime {

	err := tz.zoneProcessor.InitForEpochSeconds(epochSeconds)
	if err != ErrOk {
		return NewOffsetDateTimeError()
	}

	result := tz.zoneProcessor.FindByEpochSeconds(epochSeconds)
	if result.frtype == FindResultNotFound {
		return NewOffsetDateTimeError()
	}

	totalOffsetMinutes := result.stdOffsetMinutes + result.dstOffsetMinutes
	odt := NewOffsetDateTimeFromEpochSeconds(epochSeconds, totalOffsetMinutes)
	if !odt.IsError() {
		odt.Fold = result.fold
	}
	return odt
}

// OffsetDateTimeFromLocalDateTime calculates the OffsetDateTime from the given
// LocalDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_local_date_time() from the
// AceTimeC library, and TimeZone::getOffsetDateTime(const LocalDatetime&) from
// the AceTime library.
func (tz *TimeZone) OffsetDateTimeFromLocalDateTime(
	ldt *LocalDateTime) OffsetDateTime {

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewOffsetDateTimeError()
	}

	// Convert FindResult into OffsetDateTime using the requested offset.
	odt := OffsetDateTime{
		Year:          ldt.Year,
		Month:         ldt.Month,
		Day:           ldt.Day,
		Hour:          ldt.Hour,
		Minute:        ldt.Minute,
		Second:        ldt.Second,
		OffsetMinutes: result.reqStdOffsetMinutes + result.reqDstOffsetMinutes,
		Fold:          result.fold,
	}

	// Special processor for kAtcFindResultGap: Convert to epochSeconds using the
	// reqStdOffsetMinutes and reqDstOffsetMinutes, then convert back to
	// OffsetDateTime using the target stdOffsetMinutes and
	// dstOffsetMinutes.
	if result.frtype == FindResultGap {
		epochSeconds := odt.ToEpochSeconds()
		targetOffset := result.stdOffsetMinutes + result.dstOffsetMinutes
		odt = NewOffsetDateTimeFromEpochSeconds(epochSeconds, targetOffset)
	}

	return odt
}

func (tz *TimeZone) ZonedExtraFromEpochSeconds(epochSeconds int32) ZonedExtra {
	result := tz.zoneProcessor.FindByEpochSeconds(epochSeconds)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewZonedExtraError()
	}

	return ZonedExtra{
		zetype:              result.frtype,
		stdOffsetMinutes:    result.stdOffsetMinutes,
		dstOffsetMinutes:    result.dstOffsetMinutes,
		reqStdOffsetMinutes: result.reqStdOffsetMinutes,
		reqDstOffsetMinutes: result.reqDstOffsetMinutes,
		abbrev:              result.abbrev,
	}
}

func (tz *TimeZone) ZonedExtraFromLocalDateTime(
	ldt *LocalDateTime) ZonedExtra {

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewZonedExtraError()
	}

	return ZonedExtra{
		zetype:              result.frtype,
		stdOffsetMinutes:    result.stdOffsetMinutes,
		dstOffsetMinutes:    result.dstOffsetMinutes,
		reqStdOffsetMinutes: result.reqStdOffsetMinutes,
		reqDstOffsetMinutes: result.reqDstOffsetMinutes,
		abbrev:              result.abbrev,
	}
}
