package acetime

// FoldType
const (
	FoldTypeErr = iota
	FoldTypeNotFound
	FoldTypeExact
	FoldTypeGap
	FoldTypeOverlap
)

var (
	ZonedExtraError = ZonedExtra{FoldType: FoldTypeErr}
)

// ZonedExtra contains information about a specific instant in time (either at a
// specific epochSeconds or a specific LocalDateTime) which are not fully
// captured by the OffsetDateTime. These include the STD offset, the DST offset,
// and the abbreviation.
type ZonedExtra struct {
	FoldType            uint8  // type of fold (e.g. gap, overlap)
	StdOffsetSeconds    int32  // STD offset
	DstOffsetSeconds    int32  // DST offset
	ReqStdOffsetSeconds int32  // request STD offset
	ReqDstOffsetSeconds int32  // request DST offset
	Abbrev              string // abbreviation (e.g. PST, PDT)
}

func NewZonedExtraFromEpochSeconds(
	epochSeconds Time, tz *TimeZone) ZonedExtra {

	return tz.findZonedExtraForEpochSeconds(epochSeconds)
}

func NewZonedExtraFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedExtra {

	return tz.findZonedExtraForLocalDateTime(ldt)
}

func (extra *ZonedExtra) IsError() bool {
	return extra.FoldType == FoldTypeErr
}

// OffsetSeconds returns the total offset from UTC in seconds (StdOffsetSeconds
// + DstOffsetSeconds). This is a convenience function because it is needed
// frequently.
func (extra *ZonedExtra) OffsetSeconds() int32 {
	return extra.StdOffsetSeconds + extra.DstOffsetSeconds
}
