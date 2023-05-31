package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

var (
	ZonedDateTimeError = ZonedDateTime{
		OffsetDateTime: OffsetDateTimeError,
		ZonedExtra:     ZonedExtraError,
	}
)

// ZonedDateTime represents a date/time with its associated TimeZone.
type ZonedDateTime struct {
	OffsetDateTime
	ZonedExtra
	Tz *TimeZone
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.OffsetDateTime.IsError() || zdt.ZonedExtra.IsError()
}

func NewZonedDateTimeFromEpochSeconds(
	epochSeconds Time, tz *TimeZone) ZonedDateTime {

	odt, extra := tz.findForEpochSeconds(epochSeconds)
	return ZonedDateTime{
		OffsetDateTime: odt,
		ZonedExtra:     extra,
		Tz:             tz,
	}
}

func NewZonedDateTimeFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedDateTime {

	odt, extra := tz.findForLocalDateTime(ldt)
	return ZonedDateTime{
		OffsetDateTime: odt,
		ZonedExtra:     extra,
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

	odt, extra := zdt.Tz.findForLocalDateTime(&zdt.OffsetDateTime.LocalDateTime)
	zdt.OffsetDateTime = odt
	zdt.ZonedExtra = extra
}

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
