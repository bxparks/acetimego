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

type ZonedExtra struct {
	stdOffsetMinutes int16  // STD offset
	dstOffsetMinutes int16  // DST offset
	abbrev           string // abbreviation (e.g. PST, PDT)
}

func (ze ZonedExtra) IsError() bool {
	return ze.stdOffsetMinutes == InvalidOffsetMinutes
}

func NewZonedExtraError() ZonedExtra {
	return ZonedExtra{stdOffsetMinutes: InvalidOffsetMinutes}
}

func ZonedExtraFromEpochSeconds(epochSeconds int32, tz *TimeZone) ZonedExtra {
	if epochSeconds == InvalidEpochSeconds {
		return NewZonedExtraError()
	}
	return tz.ZonedExtraFromEpochSeconds(epochSeconds)
}
