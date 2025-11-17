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
	processor *zoneProcessor // nil for TztypeError and TztypeUTC
}

var (
	// TimeZoneUTC is a predefined instance that represents UTC time zone
	TimeZoneUTC = TimeZone{TztypeUTC, nil}

	// TimeZoneError is a predefined instance that represents an error
	TimeZoneError = TimeZone{TztypeError, nil}
)

func timeZoneFromZoneInfo(zoneInfo *zoneinfo.ZoneInfo) TimeZone {
	var processor zoneProcessor
	processor.initForZoneInfo(zoneInfo)
	return TimeZone{TztypeProcessor, &processor}
}

func (tz *TimeZone) IsError() bool {
	return tz.tztype == TztypeError
}

func (tz *TimeZone) IsUTC() bool {
	return tz.processor == nil // TimeZoneError acts like UTC
}

func (tz *TimeZone) IsLink() bool {
	if tz.processor == nil {
		return false
	}
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

// findOffsetDateTimeForUnixSeconds returns the OffsetDateTime matching the
// given unixSeconds. The 'resolved' result is always set to ResolvedUnique
// because a given unixSeconds always resolves to a unique OffsetDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_epoch_seconds() in the
// acetimec library and, TimeZone::getOffsetDateTime(unixSeconds) from the
// AceTime library.
func (tz *TimeZone) findOffsetDateTimeForUnixSeconds(
	unixSeconds Time) (OffsetDateTime, ResolvedType) {

	// UTC or Error
	if tz.processor == nil {
		return OffsetDateTimeFromUnixSeconds(unixSeconds, 0), ResolvedUnique
	}

	err := tz.processor.initForUnixSeconds(unixSeconds)
	if err != errOk {
		return OffsetDateTimeError, ResolvedUnique
	}

	result := tz.processor.findByUnixSeconds(unixSeconds)
	if result.frtype == findResultNotFound {
		return OffsetDateTimeError, ResolvedUnique
	}

	totalOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
	odt := OffsetDateTimeFromUnixSeconds(unixSeconds, totalOffsetSeconds)
	resolved := ResolvedUnique
	return odt, resolved
}

// findZonedExtraForUnixSeconds returns the ZonedExtra matching the given
// unixSeconds.
//
// Adapted from atc_time_zone_zoned_extra_from_epoch_seconds() in the
// acetimec library and, TimeZone::getZonedExtra(unixSeconds) from the
// AceTime library.
func (tz *TimeZone) findZonedExtraForUnixSeconds(
	unixSeconds Time) ZonedExtra {

	// UTC or Error
	if tz.processor == nil {
		var resolved ResolvedType
		if unixSeconds == InvalidUnixSeconds {
			resolved = ResolvedError
		} else {
			resolved = ResolvedUnique
		}
		return ZonedExtra{
			Resolved:            resolved,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	err := tz.processor.initForUnixSeconds(unixSeconds)
	if err != errOk {
		return ZonedExtraError
	}

	result := tz.processor.findByUnixSeconds(unixSeconds)
	if result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	return ZonedExtra{
		Resolved:            ResolvedUnique,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}

// findOffsetDateTimeForPlainDateTime returns the matching OffsetDateTime from
// the given PlainDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_local_date_time() from the
// acetimec library, and TimeZone::getOffsetDateTime(const PlainDateTime&) from
// the AceTime library.
func (tz *TimeZone) findOffsetDateTimeForPlainDateTime(
	pdt *PlainDateTime, disambiguate uint8) (OffsetDateTime, ResolvedType) {

	// UTC or Error
	if tz.processor == nil {
		return OffsetDateTimeFromPlainDateTime(pdt, 0), ResolvedUnique
	}

	result := tz.processor.findByPlainDateTime(pdt, disambiguate)
	if result.frtype == findResultNotFound {
		return OffsetDateTimeError, ResolvedUnique
	}

	// Convert findResult into OffsetDateTime using the request offset.
	odt := OffsetDateTime{
		PlainDateTime: *pdt,
		OffsetSeconds: result.reqStdOffsetSeconds + result.reqDstOffsetSeconds,
	}

	// Special process for findResultGap: Convert to unixSeconds using the
	// reqStdOffsetSeconds and reqDstOffsetSeconds, then convert back to
	// OffsetDateTime using the target stdOffsetSeconds and
	// dstOffsetSeconds.
	if result.frtype == findResultGap {
		unixSeconds := odt.UnixSeconds()
		targetOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
		odt = OffsetDateTimeFromUnixSeconds(unixSeconds, targetOffsetSeconds)
	}

	resolved := resolveForResultTypeAndFold(result.frtype, result.foldNumber)
	return odt, resolved
}

// Convert frtype and foldNumber into a ZonedDateTime.Resolved field.
func resolveForResultTypeAndFold(
	frtype findResultType, foldNumber uint8) ResolvedType {

	if frtype == findResultOverlap {
		if foldNumber == 0 {
			return ResolvedOverlapEarlier
		} else {
			return ResolvedOverlapLater
		}
	} else if frtype == findResultGap {
		if foldNumber == 0 {
			return ResolvedGapLater
		} else {
			return ResolvedGapEarlier
		}
	} else {
		return ResolvedUnique
	}
}

// findZonedExtraForPlainDateTime returns the matching ZonedExtra from the
// given PlainDateTime.
//
// Adapted from atc_time_zone_zoned_extra_from_local_date_time() from the
// acetimec library, and TimeZone::getZonedExtra(const PlainDateTime&) from
// the AceTime library.
func (tz *TimeZone) findZonedExtraForPlainDateTime(
	pdt *PlainDateTime, disambiguate uint8) ZonedExtra {

	// UTC or Error
	if tz.processor == nil {
		var resolved ResolvedType
		if pdt.IsError() {
			resolved = ResolvedError
		} else {
			resolved = ResolvedUnique
		}
		return ZonedExtra{
			Resolved:            resolved,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.processor.findByPlainDateTime(pdt, disambiguate)
	if result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	resolved := resolveForResultTypeAndFold(result.frtype, result.foldNumber)
	return ZonedExtra{
		Resolved:            resolved,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}
