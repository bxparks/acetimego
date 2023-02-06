package acetime

import (
	"strings"
)

var (
	OffsetDateTimeError = OffsetDateTime{Year: InvalidYear}
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

func (odt *OffsetDateTime) EpochSeconds() ATime {
	if odt.IsError() {
		return InvalidEpochSeconds
	}

	epochSeconds := (&LocalDateTime{
		odt.Year, odt.Month, odt.Day,
		odt.Hour, odt.Minute, odt.Second, odt.Fold,
	}).EpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return epochSeconds
	}
	return epochSeconds - ATime(odt.OffsetMinutes)*60
}

func NewOffsetDateTimeFromEpochSeconds(
	epochSeconds ATime, offsetMinutes int16) OffsetDateTime {

	if epochSeconds == InvalidEpochSeconds {
		return OffsetDateTimeError
	}

	epochSeconds += ATime(offsetMinutes) * 60
	ldt := NewLocalDateTimeFromEpochSeconds(epochSeconds)
	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		0 /*Fold*/, offsetMinutes}
}

func (odt *OffsetDateTime) LocalDateTime() LocalDateTime {
	return LocalDateTime{
		Year:   odt.Year,
		Month:  odt.Month,
		Day:    odt.Day,
		Hour:   odt.Hour,
		Minute: odt.Minute,
		Second: odt.Second,
		Fold:   odt.Fold,
	}
}

func (odt *OffsetDateTime) String() string {
	var b strings.Builder
	odt.BuildString(&b)
	return b.String()
}

func (odt *OffsetDateTime) BuildString(b *strings.Builder) {
	ldt := odt.LocalDateTime()
	ldt.BuildString(b)
	BuildUTCOffset(b, odt.OffsetMinutes)
}

// Extract the UTC offset as +/-hh:mm
func BuildUTCOffset(b *strings.Builder, offsetMinutes int16) {
	s, h, m := minutesToHM(offsetMinutes)
	var c byte
	if s < 0 {
		c = '-'
	} else {
		c = '+'
	}

	b.WriteByte(c)
	WriteUint8Pad2(b, h, '0')
	b.WriteByte(':')
	WriteUint8Pad2(b, m, '0')
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
