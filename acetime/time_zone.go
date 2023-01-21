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
	zoneProcessor *ZoneProcessor
}

func (tz *TimeZone) String() string {
	if tz.zoneProcessor == nil {
		return "UTC"
	} else {
		return tz.zoneProcessor.String()
	}
}

// NewTimeZoneUTC returns a TimeZone instance that represents the UTC timezone.
func NewTimeZoneUTC() TimeZone {
	return TimeZone{}
}

func (tz *TimeZone) IsUTC() bool {
	return tz.zoneProcessor == nil
}

func NewTimeZoneFromZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {
	var zoneProcessor ZoneProcessor
	zoneProcessor.InitForZoneInfo(zoneInfo)
	return TimeZone{&zoneProcessor}
}

// OffsetDateTimeFromEpochSeconds calculates the OffsetDateTime from the given
// epochSeconds.
//
// Adapted from atc_time_zone_offset_date_time_from_epoch_seconds() in the
// AceTimeC library and, TimeZone::getOffsetDateTime(epochSeconds) from the
// AceTime library.
func (tz *TimeZone) OffsetDateTimeFromEpochSeconds(
	epochSeconds ATime) OffsetDateTime {

	// UTC
	if tz.zoneProcessor == nil {
		return NewOffsetDateTimeFromEpochSeconds(epochSeconds, 0)
	}

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

	// UTC
	if tz.zoneProcessor == nil {
		return NewOffsetDateTimeFromLocalDateTime(ldt, 0)
	}

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewOffsetDateTimeError()
	}

	// Convert FindResult into OffsetDateTime using the request offset, and the
	// result fold.
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

func (tz *TimeZone) ZonedExtraFromEpochSeconds(epochSeconds ATime) ZonedExtra {
	if tz.zoneProcessor == nil {
		return ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetMinutes:    0,
			DstOffsetMinutes:    0,
			ReqStdOffsetMinutes: 0,
			ReqDstOffsetMinutes: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.zoneProcessor.FindByEpochSeconds(epochSeconds)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewZonedExtraError()
	}

	return ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetMinutes:    result.stdOffsetMinutes,
		DstOffsetMinutes:    result.dstOffsetMinutes,
		ReqStdOffsetMinutes: result.reqStdOffsetMinutes,
		ReqDstOffsetMinutes: result.reqDstOffsetMinutes,
		Abbrev:              result.abbrev,
	}
}

func (tz *TimeZone) ZonedExtraFromLocalDateTime(
	ldt *LocalDateTime) ZonedExtra {

	if tz.zoneProcessor == nil {
		return ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetMinutes:    0,
			DstOffsetMinutes:    0,
			ReqStdOffsetMinutes: 0,
			ReqDstOffsetMinutes: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == FindResultErr || result.frtype == FindResultNotFound {
		return NewZonedExtraError()
	}

	return ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetMinutes:    result.stdOffsetMinutes,
		DstOffsetMinutes:    result.dstOffsetMinutes,
		ReqStdOffsetMinutes: result.reqStdOffsetMinutes,
		ReqDstOffsetMinutes: result.reqDstOffsetMinutes,
		Abbrev:              result.abbrev,
	}
}
