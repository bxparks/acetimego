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

// ZonedDateTime represents an OffsetDateTime associated with a TimeZone.
// The intention is that this object contains the most frequently used
// fields. Additional (and less often used) fields are placed into the
// ZonedExtra object.
//
// An alternative implementation is to insert the ZonedExtra object directly
// into the ZonedDateTime. But that causes the size of the object to grow from
// 24 bytes to 64 bytes, which causes the performance of this library to degrade
// noticeably. The compare_acetimego binary in the AceTimeValidation project
// saw a performce decrease from 6.3 seconds to 8.1 seconds with the larger
// ZonedDateTime object.
type ZonedDateTime struct {
	OffsetDateTime
	Tz *TimeZone
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.OffsetDateTime.IsError()
}

func NewZonedDateTimeFromEpochSeconds(
	epochSeconds Time, tz *TimeZone) ZonedDateTime {

	odt := tz.findOffsetDateTimeForEpochSeconds(epochSeconds)
	return ZonedDateTime{
		OffsetDateTime: odt,
		Tz:             tz,
	}
}

func NewZonedDateTimeFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedDateTime {

	odt := tz.findOffsetDateTimeForLocalDateTime(ldt)
	return ZonedDateTime{
		OffsetDateTime: odt,
		Tz:             tz,
	}
}

func (zdt *ZonedDateTime) ConvertToTimeZone(tz *TimeZone) ZonedDateTime {
	if zdt.IsError() {
		return ZonedDateTimeError
	}
	epochSeconds := zdt.EpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return ZonedDateTimeError
	}
	return NewZonedDateTimeFromEpochSeconds(epochSeconds, tz)
}

// Normalize should be called if any of its date or time fields are changed
// manually. This corrects for invalid date or time fields, for example:
//  1. time fields which not exist (i.e. in a DST shift-forward gap),
//  2. dates which are inconsistent (e.g. Feb 29 in a non-leap year),
//  3. or changing the date from a DST date to a non-DST date, or vise versa.
func (zdt *ZonedDateTime) Normalize() {
	if zdt.IsError() {
		return
	}

	odt := zdt.Tz.findOffsetDateTimeForLocalDateTime(
		&zdt.OffsetDateTime.LocalDateTime)
	zdt.OffsetDateTime = odt
}

// ZonedExtra returns the ZonedExtra object corresponding to the current
// ZonedDateTime. This is will be always identical to the value returned by
// NewZonedExtraFromEpochSeconds().
//
// It will usually be identical to the value returned by
// NewZonedExtraFromLocalDateTime() except when the LocalDateTime falls in a gap
// (FoldTypeGap). In that case, the ZonedDateTime has already been normalized
// into a real ZonedDateTime, and NewZonedExtraFromLocalDateTime() should be
// called with the original LocalDateTime if the information about the
// non-existent LocalDateTime is required.
func (zdt *ZonedDateTime) ZonedExtra() ZonedExtra {
	return zdt.Tz.findZonedExtraForLocalDateTime(
		&zdt.OffsetDateTime.LocalDateTime)
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
