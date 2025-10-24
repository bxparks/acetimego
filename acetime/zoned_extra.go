package acetime

// FoldKind, the type of the FoldType. Must be identical to findResult enums.
type FoldType uint8

const (
	FoldTypeNotFound FoldType = iota
	FoldTypeExact
	FoldTypeGap
	FoldTypeOverlap
)

var (
	ZonedExtraError = ZonedExtra{ResolvedFold: FoldTypeNotFound}
)

// ZonedExtra contains information about a specific instant in time (either at a
// specific unixSeconds or a specific PlainDateTime) which are not fully
// captured by the OffsetDateTime. These include the STD offset, the DST offset,
// and the abbreviation.
type ZonedExtra struct {
	ResolvedFold        FoldType // type of fold (e.g. gap, overlap)
	StdOffsetSeconds    int32    // STD offset
	DstOffsetSeconds    int32    // DST offset
	ReqStdOffsetSeconds int32    // request STD offset
	ReqDstOffsetSeconds int32    // request DST offset
	Abbrev              string   // abbreviation (e.g. PST, PDT)
}

func ZonedExtraFromUnixSeconds(
	unixSeconds Time, tz *TimeZone) ZonedExtra {

	return tz.findZonedExtraForUnixSeconds(unixSeconds)
}

func ZonedExtraFromPlainDateTime(
	pdt *PlainDateTime, tz *TimeZone, disambiguate uint8) ZonedExtra {

	return tz.findZonedExtraForPlainDateTime(pdt, disambiguate)
}

func (extra *ZonedExtra) IsError() bool {
	return extra.ResolvedFold == FoldTypeNotFound
}

// OffsetSeconds returns the total offset from UTC in seconds (StdOffsetSeconds
// + DstOffsetSeconds). This is a convenience function because it is needed
// frequently.
func (extra *ZonedExtra) OffsetSeconds() int32 {
	return extra.StdOffsetSeconds + extra.DstOffsetSeconds
}
