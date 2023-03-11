package acetime

import (
	"github.com/bxparks/AceTimeGo/strbuild"
	"strings"
)

const (
	InvalidEpochSeconds = ATime(-(1 << 63)) // math.MinInt64
)

var (
	LocalDateTimeError = LocalDateTime{Year: InvalidYear}
)

// A LocalDateTime represents a datetime without regards to the [TimeZone] or
// offset from UTC. Sometimes this is used to hold a localized datetime in the
// current implicitly-defined time zone. Sometimes this is used to represent a
// datetime in UTC.
type LocalDateTime struct {
	Year   int16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
	Fold   uint8
}

func (ldt *LocalDateTime) IsError() bool {
	return ldt.Year == InvalidYear
}

func (ldt *LocalDateTime) EpochSeconds() ATime {
	if ldt.IsError() {
		return InvalidEpochSeconds
	}

	days := LocalDateToEpochDays(ldt.Year, ldt.Month, ldt.Day)
	seconds := LocalTimeToSeconds(ldt.Hour, ldt.Minute, ldt.Second)
	return ATime(days)*86400 + ATime(seconds)
}

func NewLocalDateTimeFromEpochSeconds(epochSeconds ATime) LocalDateTime {
	if epochSeconds == InvalidEpochSeconds {
		return LocalDateTimeError
	}

	// Integer floor-division towards -infinity
	eps := int64(epochSeconds)
	var days int32
	if eps < 0 {
		days = int32((eps+1)/86400) - 1
	} else {
		days = int32(eps / 86400)
	}
	seconds := int32(eps - 86400*int64(days))

	year, month, day := LocalDateFromEpochDays(days)
	hour, minute, second := LocalTimeFromSeconds(seconds)

	return LocalDateTime{year, month, day, hour, minute, second, 0 /*Fold*/}
}

func (ldt *LocalDateTime) String() string {
	var b strings.Builder
	ldt.BuildString(&b)
	return b.String()
}

func (ldt *LocalDateTime) BuildString(b *strings.Builder) {
	strbuild.Uint16Pad4(b, uint16(ldt.Year), '0')
	b.WriteByte('-')
	strbuild.Uint8Pad2(b, ldt.Month, '0')
	b.WriteByte('-')
	strbuild.Uint8Pad2(b, ldt.Day, '0')
	b.WriteByte('T')
	strbuild.Uint8Pad2(b, ldt.Hour, '0')
	b.WriteByte(':')
	strbuild.Uint8Pad2(b, ldt.Minute, '0')
	b.WriteByte(':')
	strbuild.Uint8Pad2(b, ldt.Second, '0')
}
