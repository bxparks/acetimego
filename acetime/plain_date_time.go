package acetime

import (
	"github.com/bxparks/acetimego/internal/strbuild"
	"strings"
)

const (
	InvalidUnixSeconds = Time(-(1 << 63)) // math.MinInt64
)

var (
	PlainDateTimeError = PlainDateTime{Year: InvalidYear}
)

// A PlainDateTime represents a datetime without regards to the [TimeZone] or
// offset from UTC. Sometimes this is used to hold a localized datetime in the
// current implicitly-defined time zone. Sometimes this is used to represent a
// datetime in UTC.
type PlainDateTime struct {
	Year   int16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// Convert from Unix unixSeconds.
func PlainDateTimeFromUnixSeconds(unixSeconds Time) PlainDateTime {
	if unixSeconds == InvalidUnixSeconds {
		return PlainDateTimeError
	}

	// Integer floor-division towards -infinity
	eps := int64(unixSeconds)
	var days int32
	if eps < 0 {
		days = int32((eps+1)/86400) - 1
	} else {
		days = int32(eps / 86400)
	}
	seconds := int32(eps - 86400*int64(days))

	year, month, day := PlainDateFromUnixDays(days)
	hour, minute, second := PlainTimeFromSeconds(seconds)

	return PlainDateTime{year, month, day, hour, minute, second}
}

func (pdt *PlainDateTime) IsError() bool {
	return pdt.Year == InvalidYear
}

// Convert to Unix unixSeconds.
func (pdt *PlainDateTime) UnixSeconds() Time {
	if pdt.IsError() {
		return InvalidUnixSeconds
	}

	days := PlainDateToUnixDays(pdt.Year, pdt.Month, pdt.Day)
	seconds := PlainTimeToSeconds(pdt.Hour, pdt.Minute, pdt.Second)
	return Time(days)*86400 + Time(seconds)
}

func (pdt *PlainDateTime) String() string {
	var b strings.Builder
	pdt.BuildString(&b)
	return b.String()
}

func (pdt *PlainDateTime) BuildString(b *strings.Builder) {
	strbuild.Uint16Pad4(b, uint16(pdt.Year), '0')
	b.WriteByte('-')
	strbuild.Uint8Pad2(b, pdt.Month, '0')
	b.WriteByte('-')
	strbuild.Uint8Pad2(b, pdt.Day, '0')
	b.WriteByte('T')
	strbuild.Uint8Pad2(b, pdt.Hour, '0')
	b.WriteByte(':')
	strbuild.Uint8Pad2(b, pdt.Minute, '0')
	b.WriteByte(':')
	strbuild.Uint8Pad2(b, pdt.Second, '0')
}
