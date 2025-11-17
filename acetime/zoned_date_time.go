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

// Resolved disambiguation of the ZonedDateTime from its UnixSeconds or
// PlainDateTime.
type ResolvedType uint8

const (
	ResolvedError ResolvedType = iota
	ResolvedUnique
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
	Resolved ResolvedType
}

// Converts from Unix unixSeconds.
func ZonedDateTimeFromUnixSeconds(
	unixSeconds Time, tz *TimeZone) ZonedDateTime {

	odt, resolved := tz.findOffsetDateTimeForUnixSeconds(unixSeconds)
	return ZonedDateTime{
		OffsetDateTime: odt,
		Tz:             tz,
		Resolved:       resolved,
	}
}

func ZonedDateTimeFromPlainDateTime(
	pdt *PlainDateTime,
	tz *TimeZone,
	disambiguate uint8) ZonedDateTime {

	odt, resolved := tz.findOffsetDateTimeForPlainDateTime(pdt, disambiguate)
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
	unixSeconds := zdt.UnixSeconds()
	if unixSeconds == InvalidUnixSeconds {
		return ZonedDateTimeError
	}
	return ZonedDateTimeFromUnixSeconds(unixSeconds, tz)
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

	odt, resolved := zdt.Tz.findOffsetDateTimeForPlainDateTime(
		&zdt.OffsetDateTime.PlainDateTime, disambiguate)
	zdt.OffsetDateTime = odt
	zdt.Resolved = resolved
}

// String returns the given ZonedDateTime in ISO8601 format, in the form of
// "yyyy-mm-ddThh:mm:ss+/-hh:mm[timezone]".
func (zdt *ZonedDateTime) String() string {
	var b strings.Builder
	zdt.BuildString(&b)
	return b.String()
}

func (zdt *ZonedDateTime) BuildString(b *strings.Builder) {
	zdt.OffsetDateTime.PlainDateTime.BuildString(b)

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
