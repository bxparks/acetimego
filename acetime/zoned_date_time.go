package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

var (
	ZonedDateTimeError = ZonedDateTime{
		OffsetDateTime: OffsetDateTimeError,
	}
)

// Resolved disambiguation of the ZonedDateTime from its EpochSeconds or
// LocalDateTime.
const (
	ResolvedUnique = iota
	ResolvedOverlapEarlier
	ResolvedOverlapLater
	ResolvedGapEarlier
	ResolvedGapLater
)

// ZonedDateTime represents an OffsetDateTime associated with a TimeZone.
// The intention is that this object contains the most frequently used
// fields. Additional (and less often used) fields are placed into the
// ZonedExtra object.
//
// An alternative implementation is to insert the ZonedExtra object directly
// into the ZonedDateTime. But that causes the size of the object to grow from
// 24 bytes to 64 bytes, which causes the performance of this library to degrade
// noticeably. The validation/tools/compare_acetimego binary saw a performce
// decrease from 6.3 seconds to 8.1 seconds with the larger ZonedDateTime
// object.
type ZonedDateTime struct {
	OffsetDateTime
	Tz       *TimeZone
	Resolved uint8
}

func ZonedDateTimeFromEpochSeconds(
	epochSeconds Time, tz *TimeZone) ZonedDateTime {

	odt, resolved := tz.findOffsetDateTimeForEpochSeconds(epochSeconds)
	return ZonedDateTime{
		OffsetDateTime: odt,
		Tz:             tz,
		Resolved:       resolved,
	}
}

func ZonedDateTimeFromLocalDateTime(
	ldt *LocalDateTime,
	tz *TimeZone,
	disambiguate uint8) ZonedDateTime {

	odt, resolved := tz.findOffsetDateTimeForLocalDateTime(ldt, disambiguate)
	return ZonedDateTime{
		OffsetDateTime: odt,
		Tz:             tz,
		Resolved:       resolved,
	}
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.OffsetDateTime.IsError()
}

func (zdt *ZonedDateTime) ConvertToTimeZone(tz *TimeZone) ZonedDateTime {
	if zdt.IsError() {
		return ZonedDateTimeError
	}
	epochSeconds := zdt.EpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return ZonedDateTimeError
	}
	return ZonedDateTimeFromEpochSeconds(epochSeconds, tz)
}

// Normalize should be called if any of its date or time fields are changed
// manually. This corrects for invalid date or time fields, for example:
//  1. time fields which not exist (i.e. in a DST shift-forward gap),
//  2. dates which are inconsistent (e.g. Feb 29 in a non-leap year),
//  3. or changing the date from a DST date to a non-DST date, or vise versa.
func (zdt *ZonedDateTime) Normalize(disambiguate uint8) {
	if zdt.IsError() {
		return
	}

	odt, resolved := zdt.Tz.findOffsetDateTimeForLocalDateTime(
		&zdt.OffsetDateTime.LocalDateTime, disambiguate)
	zdt.OffsetDateTime = odt
	zdt.Resolved = resolved
}

// ZonedExtra returns the ZonedExtra object corresponding to the current
// ZonedDateTime. This is will be always identical to the value returned by
// ZonedExtraFromEpochSeconds().
//
// It will usually be identical to the value returned by
// ZonedExtraFromLocalDateTime() except when the LocalDateTime falls in a gap
// (FoldTypeGap). In that case, the ZonedDateTime has already been normalized
// into a real ZonedDateTime, and ZonedExtraFromLocalDateTime() should be
// called with the original LocalDateTime if the information about the
// non-existent LocalDateTime is required.
func (zdt *ZonedDateTime) ZonedExtra() ZonedExtra {
	return zdt.Tz.findZonedExtraForEpochSeconds(
		zdt.OffsetDateTime.EpochSeconds())
}

// String returns the given ZonedDateTime in ISO8601 format, in the form of
// "yyyy-mm-ddThh:mm:ss+/-hh:mm[timezone]".
func (zdt *ZonedDateTime) String() string {
	var b strings.Builder
	zdt.BuildString(&b)
	return b.String()
}

func (zdt *ZonedDateTime) BuildString(b *strings.Builder) {
	zdt.OffsetDateTime.LocalDateTime.BuildString(b)

	if zdt.Tz.IsUTC() {
		// Append just a "UTC" to simplify the ISO8601.
		b.WriteString("UTC")
	} else {
		// Append the "+/-hh:mm[tz]"
		strbuild.TimeOffset(b, zdt.OffsetSeconds)
		b.WriteByte('[')
		b.WriteString(zdt.Tz.Name())
		b.WriteByte(']')
	}
}
