package acetime

import (
	"math"
)

const InvalidEpochSeconds int32 = math.MinInt32

type LocalDateTime struct {
	Year   int16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

func NewLocalDateTimeError() LocalDateTime {
	return LocalDateTime{Year: InvalidYear}
}

func (ldt *LocalDateTime) IsError() bool {
	return ldt.Year == InvalidYear
}

func (ldt *LocalDateTime) ToEpochSeconds() int32 {
	if ldt.IsError() {
		return InvalidEpochSeconds
	}

	days := LocalDateToEpochDays(ldt.Year, ldt.Month, ldt.Day)
	seconds := LocalTimeToSeconds(ldt.Hour, ldt.Minute, ldt.Second)
	return days*86400 + seconds
}

func LocalDateTimeFromEpochSeconds(epochSeconds int32) LocalDateTime {
	if epochSeconds == InvalidEpochSeconds {
		return NewLocalDateTimeError()
	}

	// Integer floor-division towards -infinity
	var days int32
	if epochSeconds < 0 {
		days = (epochSeconds+1)/86400 - 1
	} else {
		days = epochSeconds / 86400
	}
	seconds := epochSeconds - 86400*days

	year, month, day := LocalDateFromEpochDays(days)
	hour, minute, second := LocalTimeFromSeconds(seconds)

	return LocalDateTime{year, month, day, hour, minute, second}
}
