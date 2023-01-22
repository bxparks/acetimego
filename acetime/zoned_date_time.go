package acetime

import (
	"strings"
)

//-----------------------------------------------------------------------------
// ZonedDateTime represents a date/time stamp with its associated TimeZone.
//-----------------------------------------------------------------------------

type ZonedDateTime struct {
	Year          int16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minute        uint8
	Second        uint8
	Fold          uint8
	OffsetMinutes int16
	Tz            *TimeZone
}

// NewZonedDateTimeError returns an instance of ZonedDateTime that indicates
// an error condition such that IsError() returns true.
func NewZonedDateTimeError() ZonedDateTime {
	return ZonedDateTime{Year: InvalidYear}
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.Year == InvalidYear
}

func (zdt *ZonedDateTime) ToLocalDateTime() LocalDateTime {
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

func (zdt *ZonedDateTime) ToEpochSeconds() ATime {
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
		OffsetMinutes: zdt.OffsetMinutes,
	}).ToEpochSeconds()
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
		OffsetMinutes: odt.OffsetMinutes,
		Tz:            tz,
	}
}

func NewZonedDateTimeFromUnixSeconds64(
	unixSeconds64 int64, tz *TimeZone) ZonedDateTime {

	if unixSeconds64 == InvalidUnixSeconds64 {
		return NewZonedDateTimeError()
	}

	epochSeconds := ATime(unixSeconds64 -
		GetSecondsToCurrentEpochFromUnixEpoch64())

	odt := tz.OffsetDateTimeFromEpochSeconds(epochSeconds)
	return ZonedDateTime{
		Year:          odt.Year,
		Month:         odt.Month,
		Day:           odt.Day,
		Hour:          odt.Hour,
		Minute:        odt.Minute,
		Second:        odt.Second,
		Fold:          odt.Fold,
		OffsetMinutes: odt.OffsetMinutes,
		Tz:            tz,
	}
}

func (zdt *ZonedDateTime) ToUnixSeconds64() int64 {
	epochSeconds := zdt.ToEpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return InvalidUnixSeconds64
	}
	return int64(epochSeconds) + GetSecondsToCurrentEpochFromUnixEpoch64()
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
		OffsetMinutes: odt.OffsetMinutes,
		Tz:            tz,
	}
}

func (zdt *ZonedDateTime) ConvertToTimeZone(tz *TimeZone) ZonedDateTime {
	if zdt.IsError() {
		return NewZonedDateTimeError()
	}
	epochSeconds := zdt.ToEpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return NewZonedDateTimeError()
	}
	return NewZonedDateTimeFromEpochSeconds(epochSeconds, tz)
}

func (zdt *ZonedDateTime) String() string {
	var b strings.Builder
	zdt.BuildString(&b)
	return b.String()
}

func (zdt *ZonedDateTime) BuildString(b *strings.Builder) {
	ldt := zdt.ToLocalDateTime()
	ldt.BuildString(b)

	if zdt.Tz.IsUTC() {
		// Append just a "UTC" to simplify the ISO8601.
		b.WriteString(" UTC")
	} else {
		// Append the "+/-hh:mm[tz]"
		BuildUTCOffset(b, zdt.OffsetMinutes)
		b.WriteByte('[')
		b.WriteString(zdt.Tz.String())
		b.WriteByte(']')
	}
}
