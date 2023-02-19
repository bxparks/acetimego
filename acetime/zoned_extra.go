package acetime

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

// ZonedExtra contains information about a specific instant in time (either at a
// specific epochSeconds or a specific LocalDateTime) which are not fully
// captured by the ZonedDateTime instance.
type ZonedExtra struct {
	Zetype              uint8  // type of match (e.g. gap, overlap)
	StdOffsetSeconds    int32  // STD offset
	DstOffsetSeconds    int32  // DST offset
	ReqStdOffsetSeconds int32  // request STD offset
	ReqDstOffsetSeconds int32  // request DST offset
	Abbrev              string // abbreviation (e.g. PST, PDT)
}

func NewZonedExtraFromEpochSeconds(
	epochSeconds ATime, tz *TimeZone) ZonedExtra {

	if epochSeconds == InvalidEpochSeconds {
		return ZonedExtraError
	}
	return tz.zonedExtraFromEpochSeconds(epochSeconds)
}

func NewZonedExtraFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedExtra {

	if ldt.IsError() {
		return ZonedExtraError
	}
	return tz.zonedExtraFromLocalDateTime(ldt)
}
