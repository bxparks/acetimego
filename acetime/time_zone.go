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

// findOffsetDateTimeForEpochSeconds returns the OffsetDateTime matching the
// given epochSeconds. The 'resolved' result is always set to ResolvedUnique
// because a given epochSeconds always resolves to a unique OffsetDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_epoch_seconds() in the
// acetimec library and, TimeZone::getOffsetDateTime(epochSeconds) from the
// AceTime library.
func (tz *TimeZone) findOffsetDateTimeForEpochSeconds(
	epochSeconds Time) (OffsetDateTime, uint8) {

	// UTC or Error
	if tz.processor == nil {
		return OffsetDateTimeFromEpochSeconds(epochSeconds, 0), ResolvedUnique
	}

	err := tz.processor.initForEpochSeconds(epochSeconds)
	if err != errOk {
		return OffsetDateTimeError, ResolvedUnique
	}

	result := tz.processor.findByEpochSeconds(epochSeconds)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return OffsetDateTimeError, ResolvedUnique
	}

	totalOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
	odt := OffsetDateTimeFromEpochSeconds(epochSeconds, totalOffsetSeconds)
	resolved := uint8(ResolvedUnique)
	return odt, resolved
}

// findZonedExtraForEpochSeconds returns the ZonedExtra matching the given
// epochSeconds.
//
// Adapted from atc_time_zone_zoned_extra_from_epoch_seconds() in the
// acetimec library and, TimeZone::getZonedExtra(epochSeconds) from the
// AceTime library.
func (tz *TimeZone) findZonedExtraForEpochSeconds(
	epochSeconds Time) ZonedExtra {

	// UTC or Error
	if tz.processor == nil {
		var foldType uint8
		if epochSeconds == InvalidEpochSeconds {
			foldType = FoldTypeErr
		} else {
			foldType = FoldTypeExact
		}
		return ZonedExtra{
			FoldType:            foldType,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	err := tz.processor.initForEpochSeconds(epochSeconds)
	if err != errOk {
		return ZonedExtraError
	}

	result := tz.processor.findByEpochSeconds(epochSeconds)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	return ZonedExtra{
		FoldType:            result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}

// findOffsetDateTimeForLocalDateTime returns the matching OffsetDateTime from
// the given LocalDateTime.
//
// Adapted from atc_time_zone_offset_date_time_from_local_date_time() from the
// acetimec library, and TimeZone::getOffsetDateTime(const LocalDateTime&) from
// the AceTime library.
func (tz *TimeZone) findOffsetDateTimeForLocalDateTime(
	ldt *LocalDateTime, disambiguate uint8) (OffsetDateTime, uint8) {

	// UTC or Error
	if tz.processor == nil {
		return OffsetDateTimeFromLocalDateTime(ldt, 0), ResolvedUnique
	}

	result := tz.processor.findByLocalDateTime(ldt, disambiguate)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return OffsetDateTimeError, ResolvedUnique
	}

	// Convert findResult into OffsetDateTime using the request offset.
	odt := OffsetDateTime{
		LocalDateTime: *ldt,
		OffsetSeconds: result.reqStdOffsetSeconds + result.reqDstOffsetSeconds,
	}

	// Special process for findResultGap: Convert to epochSeconds using the
	// reqStdOffsetSeconds and reqDstOffsetSeconds, then convert back to
	// OffsetDateTime using the target stdOffsetSeconds and
	// dstOffsetSeconds.
	if result.frtype == findResultGap {
		epochSeconds := odt.EpochSeconds()
		targetOffsetSeconds := result.stdOffsetSeconds + result.dstOffsetSeconds
		odt = OffsetDateTimeFromEpochSeconds(epochSeconds, targetOffsetSeconds)
	}

	resolved := resolveForResultTypeAndFold(result.frtype, result.fold)
	return odt, resolved
}

// Convert frtype and fold into a ZonedDateTime.Resolved field.
func resolveForResultTypeAndFold(frtype uint8, fold uint8) uint8 {
	if frtype == findResultOverlap {
		if fold == 0 {
			return ResolvedOverlapEarlier
		} else {
			return ResolvedOverlapLater
		}
	} else if frtype == findResultGap {
		if fold == 0 {
			return ResolvedGapLater
		} else {
			return ResolvedGapEarlier
		}
	} else {
		return ResolvedUnique
	}
}

// findZonedExtraForLocalDateTime returns the matching ZonedExtra from the
// given LocalDateTime.
//
// Adapted from atc_time_zone_zoned_extra_from_local_date_time() from the
// acetimec library, and TimeZone::getZonedExtra(const LocalDateTime&) from
// the AceTime library.
func (tz *TimeZone) findZonedExtraForLocalDateTime(
	ldt *LocalDateTime, disambiguate uint8) ZonedExtra {

	// UTC or Error
	if tz.processor == nil {
		var foldType uint8
		if ldt.IsError() {
			foldType = FoldTypeErr
		} else {
			foldType = FoldTypeExact
		}
		return ZonedExtra{
			FoldType:            foldType,
			StdOffsetSeconds:    0,
			DstOffsetSeconds:    0,
			ReqStdOffsetSeconds: 0,
			ReqDstOffsetSeconds: 0,
			Abbrev:              "UTC",
		}
	}

	result := tz.processor.findByLocalDateTime(ldt, disambiguate)
	if result.frtype == findResultErr || result.frtype == findResultNotFound {
		return ZonedExtraError
	}

	return ZonedExtra{
		FoldType:            result.frtype,
		StdOffsetSeconds:    result.stdOffsetSeconds,
		DstOffsetSeconds:    result.dstOffsetSeconds,
		ReqStdOffsetSeconds: result.reqStdOffsetSeconds,
		ReqDstOffsetSeconds: result.reqDstOffsetSeconds,
		Abbrev:              result.abbrev,
	}
}
