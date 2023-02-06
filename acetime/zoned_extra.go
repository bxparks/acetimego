package acetime

//-----------------------------------------------------------------------------
// ZonedExtra contains additional information about a specific instant
// in time (either at a specific epochSeconds or a specific LocalDateTime{}.
//-----------------------------------------------------------------------------

const (
	ZonedExtraErr = iota
	ZonedExtraNotFound
	ZonedExtraExact
	ZonedExtraGap
	ZonedExtraOverlap
)

var (
	ZonedExtraError = ZonedExtra{Zetype: ZonedExtraErr}
)

type ZonedExtra struct {
	Zetype              uint8  // type of match (e.g. gap, overlap)
	StdOffsetMinutes    int16  // STD offset
	DstOffsetMinutes    int16  // DST offset
	ReqStdOffsetMinutes int16  // request STD offset
	ReqDstOffsetMinutes int16  // request DST offset
	Abbrev              string // abbreviation (e.g. PST, PDT)
}

func NewZonedExtraFromEpochSeconds(
	epochSeconds ATime, tz *TimeZone) ZonedExtra {

	if epochSeconds == InvalidEpochSeconds {
		return ZonedExtraError
	}
	return tz.ZonedExtraFromEpochSeconds(epochSeconds)
}

func NewZonedExtraFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedExtra {

	if ldt.IsError() {
		return ZonedExtraError
	}
	return tz.ZonedExtraFromLocalDateTime(ldt)
}
