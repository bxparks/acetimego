package acetime

import (
	"strings"
)

//-----------------------------------------------------------------------------
// ZonedDateTime represents a date/time stamp with its associated TimeZone.
//-----------------------------------------------------------------------------

var (
	ZonedDateTimeError = ZonedDateTime{Year: InvalidYear}
)

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

	odt := tz.OffsetDateTimeFromEpochSeconds(epochSeconds)
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

	odt := tz.OffsetDateTimeFromLocalDateTime(ldt)
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
		BuildUTCOffset(b, zdt.OffsetSeconds)
		b.WriteByte('[')
		b.WriteString(zdt.Tz.Name())
		b.WriteByte(']')
	}
}
