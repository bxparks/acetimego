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
// captured by the OffsetDateTime instance.
type ZonedExtra struct {
	Zetype              uint8  // type of match (e.g. gap, overlap)
	StdOffsetSeconds    int32  // STD offset
	DstOffsetSeconds    int32  // DST offset
	ReqStdOffsetSeconds int32  // request STD offset
	ReqDstOffsetSeconds int32  // request DST offset
	Abbrev              string // abbreviation (e.g. PST, PDT)
}

func (extra *ZonedExtra) IsError() bool {
	return extra.Zetype == ZonedExtraErr
}
