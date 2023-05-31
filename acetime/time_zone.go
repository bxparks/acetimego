package acetime

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

const (
	TztypeError = iota
	TztypeUTC
	TztypeProcessor
)

// A TimeZone represents one of the IANA TZ time zones. It has reference
// semantics meaning that a copy of this will point to same underlying
// zoneProcessor and its cache. A TimeZone can be passed around by value or by
// pointer because it is a light-weight object.
type TimeZone struct {
	tztype    uint8
	processor *zoneProcessor
}

var (
	// TimeZoneUTC is a predefined instance that represents UTC time zone
	TimeZoneUTC = TimeZone{TztypeUTC, nil}

	// TimeZoneError is a predefined instance that represents an error
	TimeZoneError = TimeZone{TztypeError, nil}
)

func newTimeZoneFromZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {
	var processor zoneProcessor
	processor.initForZoneInfo(zoneInfo)
	return TimeZone{TztypeProcessor, &processor}
}

func (tz *TimeZone) IsError() bool {
	return tz.tztype == TztypeError
}

func (tz *TimeZone) IsUTC() bool {
	return tz.processor == nil
}

func (tz *TimeZone) IsLink() bool {
	return tz.processor.isLink()
}

func (tz *TimeZone) Name() string {
	if tz.tztype == TztypeError {
		return "Err" // Can be rendered on 7-segment LED
	} else if tz.tztype == TztypeUTC {
		return "UTC"
	} else {
		return tz.processor.name()
	}
}

func (tz *TimeZone) ZoneID() uint32 {
	if tz.processor == nil {
		return 0 // AceTimeTool guarantees that 0 is invalid
	} else {
		return tz.processor.zoneInfo.ZoneID
	}
}

// findForEpochSeconds calculates the OffsetDateTime and ZonedExtra from the
// given epochSeconds.
//
// Adapted from atc_time_zone_offset_date_time_from_epoch_seconds() in the
// acetimec library and, TimeZone::getOffsetDateTime(epochSeconds) from the
// AceTime library.
func (tz *TimeZone) findForEpochSeconds(
	epochSeconds Time) (OffsetDateTime, ZonedExtra) {

	// UTC (or Error)
	if tz.processor == nil {
		odt := NewOffsetDateTimeFromEpochSeconds(epochSeconds, 0)
		extra := ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
		return odt, extra
	}

	err := tz.processor.initForEpochSeconds(epochSeconds)
	if err != errOk {
		return OffsetDateTimeError, ZonedExtraError
	}

	result := tz.processor.findByEpochSeconds(epochSeconds)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return OffsetDateTimeError, ZonedExtraError
	}

	totalOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
	odt := NewOffsetDateTimeFromEpochSeconds(epochSeconds, totalOffsetSeconds)
	odt.Fold = result.fold

	extra := ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}

	return odt, extra
}

// findForLocalDateTime calculates the OffsetDateTime and ZonedExtra from the
// given LocalDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_local_date_time() from the
// acetimec library, and TimeZone::getOffsetDateTime(const LocalDateTime&) from
// the AceTime library.
func (tz *TimeZone) findForLocalDateTime(
	ldt *LocalDateTime) (OffsetDateTime, ZonedExtra) {

	// UTC (or Error)
	if tz.processor == nil {
		odt := NewOffsetDateTimeFromLocalDateTime(ldt, 0)
		extra := ZonedExtra{
			Zetype:              ZonedExtraExact,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}

		return odt, extra
	}

	result := tz.processor.findByLocalDateTime(ldt)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return OffsetDateTimeError, ZonedExtraError
	}

	// Convert findResult into OffsetDateTime using the request offset, and the
	// result fold.
	odt := OffsetDateTime{
		LocalDateTime: *ldt,
		OffsetSeconds: result.reqStdOffsetSeconds + result.reqDstOffsetSeconds,
	}
	odt.Fold = result.fold

	// Special processor for kAtcfindResultGap: Convert to epochSeconds using the
	// reqStdOffsetSeconds and reqDstOffsetSeconds, then convert back to
	// OffsetDateTime using the target stdOffsetSeconds and
	// dstOffsetSeconds.
	if result.frtype == findResultGap {
		epochSeconds := odt.EpochSeconds()
		targetOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
		odt = NewOffsetDateTimeFromEpochSeconds(epochSeconds, targetOffsetSeconds)
	}

	extra := ZonedExtra{
		Zetype:              result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}

	return odt, extra
}
