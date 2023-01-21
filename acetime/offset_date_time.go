package acetime

import (
	"fmt"
)

type OffsetDateTime struct {
	Year          int16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minute        uint8
	Second        uint8
	Fold          uint8
	OffsetMinutes int16
}

// NewOffsetDateTimeError returns an instance of OffsetDateTime that indicates
// an error condition such that IsError() returns true.
func NewOffsetDateTimeError() OffsetDateTime {
	return OffsetDateTime{Year: InvalidYear}
}

func NewOffsetDateTimeFromLocalDateTime(
	ldt *LocalDateTime, offsetMinutes int16) OffsetDateTime {

	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		ldt.Fold, offsetMinutes}
}

func (odt *OffsetDateTime) IsError() bool {
	return odt.Year == InvalidYear
}

func (odt *OffsetDateTime) ToEpochSeconds() ATime {
	if odt.IsError() {
		return InvalidEpochSeconds
	}

	epochSeconds := (&LocalDateTime{
		odt.Year, odt.Month, odt.Day,
		odt.Hour, odt.Minute, odt.Second, odt.Fold,
	}).ToEpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return epochSeconds
	}
	return epochSeconds - ATime(odt.OffsetMinutes)*60
}

func NewOffsetDateTimeFromEpochSeconds(
	epochSeconds ATime, offsetMinutes int16) OffsetDateTime {

	if epochSeconds == InvalidEpochSeconds {
		return NewOffsetDateTimeError()
	}

	epochSeconds += ATime(offsetMinutes) * 60
	ldt := NewLocalDateTimeFromEpochSeconds(epochSeconds)
	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		0 /*Fold*/, offsetMinutes}
}

func (odt *OffsetDateTime) String() string {
	s, h, m := minutesToHM(odt.OffsetMinutes)
	var c byte
	if s < 0 {
		c = '-'
	} else {
		c = '+'
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%c%02d:%02d",
		odt.Year, odt.Month, odt.Day, odt.Hour, odt.Minute, odt.Second,
		c, h, m)
}

func minutesToHM(minutes int16) (sign int8, h uint8, m uint8) {
	if minutes < 0 {
		sign = -1
		minutes = -minutes
	} else {
		sign = 1
	}
	h = uint8(minutes / 60)
	m = uint8(minutes % 60)
	return
}
