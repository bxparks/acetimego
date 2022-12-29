package acetime

type OffsetDateTime struct {
	Year  int16
	Month uint8
	Day   uint8

	Hour   uint8
	Minute uint8
	Second uint8

	Fold uint8

	OffsetMinutes int16
}

func (odt *OffsetDateTime) IsError() bool {
	return odt.Year == InvalidYear
}

func (odt *OffsetDateTime) ToEpochSeconds() int32 {
	if odt.IsError() {
		return InvalidEpochSeconds
	}

	epochSeconds := (&LocalDateTime{odt.Year, odt.Month, odt.Day,
		odt.Hour, odt.Minute, odt.Second}).ToEpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return epochSeconds
	}
	return epochSeconds - int32(odt.OffsetMinutes)*60
}

func OffsetDateTimeFromEpochSeconds(
	epochSeconds int32, offsetMinutes int16) OffsetDateTime {

	if epochSeconds == InvalidEpochSeconds {
		return OffsetDateTime{InvalidYear, 1, 1, 0, 0, 0, 0, 0}
	}

	epochSeconds += int32(offsetMinutes) * 60
	ldt := LocalDateTimeFromEpochSeconds(epochSeconds)
	return OffsetDateTime{
		ldt.Year, ldt.Month, ldt.Day,
		ldt.Hour, ldt.Minute, ldt.Second,
		0, offsetMinutes}
}
