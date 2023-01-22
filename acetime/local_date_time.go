package acetime

import (
	"math"
	"strings"
)

const (
	InvalidEpochSeconds  ATime = math.MinInt32
	InvalidUnixSeconds64 int64 = math.MinInt64
)

type LocalDateTime struct {
	Year   int16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
	Fold   uint8
}

func NewLocalDateTimeError() LocalDateTime {
	return LocalDateTime{Year: InvalidYear}
}

func (ldt *LocalDateTime) IsError() bool {
	return ldt.Year == InvalidYear
}

func (ldt *LocalDateTime) ToEpochSeconds() ATime {
	if ldt.IsError() {
		return InvalidEpochSeconds
	}

	days := LocalDateToEpochDays(ldt.Year, ldt.Month, ldt.Day)
	seconds := LocalTimeToSeconds(ldt.Hour, ldt.Minute, ldt.Second)
	return ATime(days*86400 + seconds)
}

func NewLocalDateTimeFromEpochSeconds(epochSeconds ATime) LocalDateTime {
	if epochSeconds == InvalidEpochSeconds {
		return NewLocalDateTimeError()
	}

	// Integer floor-division towards -infinity
	eps := int32(epochSeconds)
	var days int32
	if eps < 0 {
		days = (eps+1)/86400 - 1
	} else {
		days = eps / 86400
	}
	seconds := eps - 86400*days

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
	WriteUint16Pad4(b, uint16(ldt.Year), '0')
	b.WriteByte('-')
	WriteUint8Pad2(b, ldt.Month, '0')
	b.WriteByte('-')
	WriteUint8Pad2(b, ldt.Day, '0')
	b.WriteByte('T')
	WriteUint8Pad2(b, ldt.Hour, '0')
	b.WriteByte(':')
	WriteUint8Pad2(b, ldt.Minute, '0')
	b.WriteByte(':')
	WriteUint8Pad2(b, ldt.Second, '0')
}
