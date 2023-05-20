package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

var (
	ZonedDateTimeError = ZonedDateTime{Year: InvalidYear}
)

// ZonedDateTime represents a date/time with its associated TimeZone.
type ZonedDateTime struct {
	Year          int16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minute        uint8
	Second        uint8
	Fold          uint8
	OffsetSeconds int32
	Tz            *TimeZone
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.Year == InvalidYear
}

func (zdt *ZonedDateTime) LocalDateTime() LocalDateTime {
	return LocalDateTime{
		Year:   zdt.Year,
		Month:  zdt.Month,
		Day:    zdt.Day,
		Hour:   zdt.Hour,
		Minute: zdt.Minute,
		Second: zdt.Second,
		Fold:   zdt.Fold,
	}
}

func (zdt *ZonedDateTime) EpochSeconds() ATime {
	if zdt.IsError() {
		return InvalidEpochSeconds
	}
	return (&OffsetDateTime{
		Year:          zdt.Year,
		Month:         zdt.Month,
		Day:           zdt.Day,
		Hour:          zdt.Hour,
		Minute:        zdt.Minute,
		Second:        zdt.Second,
		Fold:          zdt.Fold,
		OffsetSeconds: zdt.OffsetSeconds,
	}).EpochSeconds()
}

func NewZonedDateTimeFromEpochSeconds(
	epochSeconds ATime, tz *TimeZone) ZonedDateTime {

	odt := tz.offsetDateTimeFromEpochSeconds(epochSeconds)
	return ZonedDateTime{
		Year:          odt.Year,
		Month:         odt.Month,
		Day:           odt.Day,
		Hour:          odt.Hour,
		Minute:        odt.Minute,
		Second:        odt.Second,
		Fold:          odt.Fold,
		OffsetSeconds: odt.OffsetSeconds,
		Tz:            tz,
	}
}

func NewZonedDateTimeFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedDateTime {

	odt := tz.offsetDateTimeFromLocalDateTime(ldt)
	return ZonedDateTime{
		Year:          odt.Year,
		Month:         odt.Month,
		Day:           odt.Day,
		Hour:          odt.Hour,
		Minute:        odt.Minute,
		Second:        odt.Second,
		Fold:          odt.Fold,
		OffsetSeconds: odt.OffsetSeconds,
		Tz:            tz,
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

	ldt := zdt.LocalDateTime()
	odt := zdt.Tz.offsetDateTimeFromLocalDateTime(&ldt)
	zdt.Year = odt.Year
	zdt.Month = odt.Month
	zdt.Day = odt.Day
	zdt.Hour = odt.Hour
	zdt.Minute = odt.Minute
	zdt.Second = odt.Second
	zdt.Fold = odt.Fold
	zdt.OffsetSeconds = odt.OffsetSeconds
}

// Return additional information about the current date time in the ZonedExtra
// object.
func (zdt *ZonedDateTime) ZonedExtra() ZonedExtra {
	ldt := zdt.LocalDateTime()
	return NewZonedExtraFromLocalDateTime(&ldt, zdt.Tz)
}

func (zdt *ZonedDateTime) String() string {
	var b strings.Builder
	zdt.BuildString(&b)
	return b.String()
}

func (zdt *ZonedDateTime) BuildString(b *strings.Builder) {
	ldt := zdt.LocalDateTime()
	ldt.BuildString(b)

	if zdt.Tz.IsUTC() {
		// Append just a "UTC" to simplify the ISO8601.
		b.WriteString(" UTC")
	} else {
		// Append the "+/-hh:mm[tz]"
		strbuild.TimeOffset(b, zdt.OffsetSeconds)
		b.WriteByte('[')
		b.WriteString(zdt.Tz.Name())
		b.WriteByte(']')
	}
}
