package acetime

import (
	"math"
)

//-----------------------------------------------------------------------------
// ZonedExtra contains additional information about a specific instant
// in time (either at a specific epochSeconds or a specific LocalDateTime{}.
//-----------------------------------------------------------------------------

const (
	InvalidOffsetMinutes = math.MinInt16
)

const (
	ZonedExtraErr = iota
	ZonedExtraNotFound
	ZonedExtraExact
	ZonedExtraGap
	ZonedExtraOverlap
)

type ZonedExtra struct {
	zetype              uint8
	stdOffsetMinutes    int16  // STD offset
	dstOffsetMinutes    int16  // DST offset
	reqStdOffsetMinutes int16  // request STD offset
	reqDstOffsetMinutes int16  // request DST offset
	abbrev              string // abbreviation (e.g. PST, PDT)
}

func NewZonedExtraError() ZonedExtra {
	return ZonedExtra{zetype: ZonedExtraErr}
}

func NewZonedExtraFromEpochSeconds(
	epochSeconds int32, tz *TimeZone) ZonedExtra {

	if epochSeconds == InvalidEpochSeconds {
		return NewZonedExtraError()
	}
	return tz.ZonedExtraFromEpochSeconds(epochSeconds)
}

func NewZonedExtraFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedExtra {

	if ldt.IsError() {
		return NewZonedExtraError()
	}
	return tz.ZonedExtraFromLocalDateTime(ldt)
}
