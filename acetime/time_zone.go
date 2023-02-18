package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

const (
	TztypeError = iota
	TztypeUTC
	TztypeProcessor
)

// A TimeZone represents one of the IANA TZ time zones. It has reference
// semantics meaning that a copy of this will point to same underlying
// ZoneProcessor ands its cache. A TimeZone can be passed around by value or by
// pointer because it is a light-weight object.
type TimeZone struct {
	tztype        uint8
	zoneProcessor *ZoneProcessor
}

var (
	// TimeZoneUTC is a predefined instance that represents UTC time zone
	TimeZoneUTC = TimeZone{TztypeUTC, nil}

	// TimeZoneError is a predefined instance that represents an error
	TimeZoneError = TimeZone{TztypeError, nil}
)

func NewTimeZoneFromZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {

	var zoneProcessor ZoneProcessor
	zoneProcessor.InitForZoneInfo(zoneInfo)
	return TimeZone{TztypeProcessor, &zoneProcessor}
}

func (tz *TimeZone) IsError() bool {
	return tz.tztype == TztypeError
}

func (tz *TimeZone) IsUTC() bool {
	return tz.zoneProcessor == nil
}

func (tz *TimeZone) IsLink() bool {
	return tz.zoneProcessor.IsLink()
}

func (tz *TimeZone) Name() string {
	if tz.tztype == TztypeError {
		return "<Error>"
	} else if tz.tztype == TztypeUTC {
		return "UTC"
	} else {
		return tz.zoneProcessor.Name()
	}
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
		return OffsetDateTimeError
	}

	result := tz.zoneProcessor.FindByEpochSeconds(epochSeconds)
	if result.frtype == findResultNotFound {
		return OffsetDateTimeError
	}

	totalOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
	odt := NewOffsetDateTimeFromEpochSeconds(epochSeconds, totalOffsetSeconds)
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

	// UTC (or Error)
	if tz.zoneProcessor == nil {
		return NewOffsetDateTimeFromLocalDateTime(ldt, 0)
	}

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return OffsetDateTimeError
	}

	// Convert findResult into OffsetDateTime using the request offset, and the
	// result fold.
	odt := OffsetDateTime{
		Year:          ldt.Year,
		Month:         ldt.Month,
		Day:           ldt.Day,
		Hour:          ldt.Hour,
		Minute:        ldt.Minute,
		Second:        ldt.Second,
		OffsetSeconds: result.reqStdOffsetSeconds + result.reqDstOffsetSeconds,
		Fold:          result.fold,
	}

	// Special processor for kAtcfindResultGap: Convert to epochSeconds using the
	// reqStdOffsetSeconds and reqDstOffsetSeconds, then convert back to
	// OffsetDateTime using the target stdOffsetSeconds and
	// dstOffsetSeconds.
	if result.frtype == findResultGap {
		epochSeconds := odt.EpochSeconds()
		targetOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
		odt = NewOffsetDateTimeFromEpochSeconds(epochSeconds, targetOffsetSeconds)
	}

	return odt
}

func (tz *TimeZone) ZonedExtraFromEpochSeconds(epochSeconds ATime) ZonedExtra {
	if tz.zoneProcessor == nil {
		return ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.zoneProcessor.FindByEpochSeconds(epochSeconds)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	return ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}

func (tz *TimeZone) ZonedExtraFromLocalDateTime(
	ldt *LocalDateTime) ZonedExtra {

	if tz.zoneProcessor == nil {
		return ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.zoneProcessor.FindByLocalDateTime(ldt)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	return ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}
