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
// captured by the OffsetDateTime instance.
type ZonedExtra struct {
	FoldType            uint8  // type of fold (e.g. gap, overlap)
	StdOffsetSeconds    int32  // STD offset
	DstOffsetSeconds    int32  // DST offset
	ReqStdOffsetSeconds int32  // request STD offset
	ReqDstOffsetSeconds int32  // request DST offset
	Abbrev              string // abbreviation (e.g. PST, PDT)
}

func (extra *ZonedExtra) IsError() bool {
	return extra.FoldType == FoldTypeErr
}
