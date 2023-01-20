package acetime

import (
	"fmt"
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
	s, h, m := minutesToHM(zdt.OffsetMinutes)
	var c byte
	if s < 0 {
		c = '-'
	} else {
		c = '+'
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%c%02d:%02d[%s]",
		zdt.Year, zdt.Month, zdt.Day, zdt.Hour, zdt.Minute, zdt.Second,
		c, h, m, zdt.Tz.String())
}
