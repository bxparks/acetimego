package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

var (
	OffsetDateTimeError = OffsetDateTime{
		PlainDateTime: PlainDateTimeError,
	}
)

// An OffsetDateTime represents a [PlainDateTime] with a fixed OffsetSeconds
// relative to UTC. This is mostly useful for the implementation of
// [ZonedDateTime], but it may be useful for end-user applications which need to
// represent a datetime with fixed offsets.
type OffsetDateTime struct {
	PlainDateTime
	OffsetSeconds int32
}

func OffsetDateTimeFromPlainDateTime(
	pdt *PlainDateTime, offsetSeconds int32) OffsetDateTime {

	return OffsetDateTime{
		PlainDateTime: *pdt,
		OffsetSeconds: offsetSeconds,
	}
}

// Converts from Unix unixSeconds.
func OffsetDateTimeFromUnixSeconds(
	unixSeconds Time, offsetSeconds int32) OffsetDateTime {

	if unixSeconds == InvalidUnixSeconds {
		return OffsetDateTimeError
	}

	unixSeconds += Time(offsetSeconds)
	pdt := PlainDateTimeFromUnixSeconds(unixSeconds)
	return OffsetDateTime{
		PlainDateTime: pdt,
		OffsetSeconds: offsetSeconds,
	}
}

func (odt *OffsetDateTime) IsError() bool {
	return odt.PlainDateTime.IsError()
}

// Converts to Unix unixSeconds.
func (odt *OffsetDateTime) UnixSeconds() Time {
	if odt.IsError() {
		return InvalidUnixSeconds
	}

	unixSeconds := odt.PlainDateTime.UnixSeconds()
	if unixSeconds == InvalidUnixSeconds {
		return unixSeconds
	}
	return unixSeconds - Time(odt.OffsetSeconds)
}

func (odt *OffsetDateTime) String() string {
	var b strings.Builder
	odt.BuildString(&b)
	return b.String()
}

func (odt *OffsetDateTime) BuildString(b *strings.Builder) {
	odt.PlainDateTime.BuildString(b)

	// Convert the OffsetSeconds to +/-hh:mm, ignoring any remaining seconds. This
	// is valid for any time after Jan 7, 1972 when Africa/Monrovia became the
	// last zone to convert to a UTC Offset in whole minutes.
	strbuild.TimeOffset(b, odt.OffsetSeconds)
}
