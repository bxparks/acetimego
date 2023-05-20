package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
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

	// Convert the OffsetSeconds to +/-hh:mm, ignoring any remaining seconds. This
	// is valid for any time after Jan 7, 1972 when Africa/Monrovia became the
	// last zone to convert to a UTC Offset in whole minutes.
	strbuild.TimeOffset(b, odt.OffsetSeconds)
}
