package acetime

import (
	"strings"
)

var (
	OffsetDateTimeError = OffsetDateTime{Year: InvalidYear}
)

// An OffsetDateTime represents a [LocalDateTime] with a fixed OffsetSeconds
// relative to UTC. This is mostly useful for the implementation of
// [ZonedDateTime], but it may be useful for end-user applications which need to
// represent a datetime with fixed offsets.
type OffsetDateTime struct {
	Year          int16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minute        uint8
	Second        uint8
	Fold          uint8
	OffsetSeconds int32
}

func NewOffsetDateTimeFromLocalDateTime(
	ldt *LocalDateTime, offsetSeconds int32) OffsetDateTime {

	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		ldt.Fold, offsetSeconds}
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
	return epochSeconds - ATime(odt.OffsetSeconds)
}

func NewOffsetDateTimeFromEpochSeconds(
	epochSeconds ATime, offsetSeconds int32) OffsetDateTime {

	if epochSeconds == InvalidEpochSeconds {
		return OffsetDateTimeError
	}

	epochSeconds += ATime(offsetSeconds)
	ldt := NewLocalDateTimeFromEpochSeconds(epochSeconds)
	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		0 /*Fold*/, offsetSeconds}
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
	BuildUTCOffset(b, odt.OffsetSeconds)
}

// Extract the UTC offset as +/-hh:mm. Ignore the seconds field for time zones
// before Jan 7, 1972 (Africa/Monrovia was the last one) whose UTC Offset is
// shifted in units of seconds instead of whole minutes.
func BuildUTCOffset(b *strings.Builder, offsetSeconds int32) {
	s, h, m, _ := secondsToHMS(offsetSeconds)
	var c byte
	if s < 0 {
		c = '-'
	} else {
		c = '+'
	}

	b.WriteByte(c)
	BuildUint8Pad2(b, h, '0')
	b.WriteByte(':')
	BuildUint8Pad2(b, m, '0')
}

func secondsToHMS(seconds int32) (sign int8, h uint8, m uint8, s uint8) {
	if seconds < 0 {
		sign = -1
		seconds = -seconds
	} else {
		sign = 1
	}
	s = uint8(seconds % 60)
	minutes := seconds / 60
	m = uint8(minutes % 60)
	hours := uint8(minutes / 60)
	h = hours

	return
}
