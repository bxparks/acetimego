package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

var (
	OffsetDateTimeError = OffsetDateTime{
		LocalDateTime: LocalDateTimeError,
	}
)

// An OffsetDateTime represents a [LocalDateTime] with a fixed OffsetSeconds
// relative to UTC. This is mostly useful for the implementation of
// [ZonedDateTime], but it may be useful for end-user applications which need to
// represent a datetime with fixed offsets.
type OffsetDateTime struct {
	LocalDateTime
	OffsetSeconds int32
}

func OffsetDateTimeFromLocalDateTime(
	ldt *LocalDateTime, offsetSeconds int32) OffsetDateTime {

	return OffsetDateTime{
		LocalDateTime: *ldt,
		OffsetSeconds: offsetSeconds,
	}
}

func OffsetDateTimeFromEpochSeconds(
	epochSeconds Time, offsetSeconds int32) OffsetDateTime {

	if epochSeconds == InvalidEpochSeconds {
		return OffsetDateTimeError
	}

	epochSeconds += Time(offsetSeconds)
	ldt := LocalDateTimeFromEpochSeconds(epochSeconds)
	return OffsetDateTime{
		LocalDateTime: ldt,
		OffsetSeconds: offsetSeconds,
	}
}

func (odt *OffsetDateTime) IsError() bool {
	return odt.LocalDateTime.IsError()
}

func (odt *OffsetDateTime) EpochSeconds() Time {
	if odt.IsError() {
		return InvalidEpochSeconds
	}

	epochSeconds := odt.LocalDateTime.EpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return epochSeconds
	}
	return epochSeconds - Time(odt.OffsetSeconds)
}

func (odt *OffsetDateTime) String() string {
	var b strings.Builder
	odt.BuildString(&b)
	return b.String()
}

func (odt *OffsetDateTime) BuildString(b *strings.Builder) {
	odt.LocalDateTime.BuildString(b)

	// Convert the OffsetSeconds to +/-hh:mm, ignoring any remaining seconds. This
	// is valid for any time after Jan 7, 1972 when Africa/Monrovia became the
	// last zone to convert to a UTC Offset in whole minutes.
	strbuild.TimeOffset(b, odt.OffsetSeconds)
}
