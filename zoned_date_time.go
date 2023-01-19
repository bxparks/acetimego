package acetime

//-----------------------------------------------------------------------------
// ZonedDateTime represents a date/time stamp with its associated TimeZone.
//-----------------------------------------------------------------------------

type ZonedDateTime struct {
	Year          int16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minute        uint8
	Second        uint8
	Fold          uint8
	OffsetMinutes int16
	Tz            *TimeZone
}

// NewZonedDateTimeError returns an instance of ZonedDateTime that indicates
// an error condition such that IsError() returns true.
func NewZonedDateTimeError() ZonedDateTime {
	return ZonedDateTime{Year: InvalidYear}
}

func (zdt *ZonedDateTime) IsError() bool {
	return zdt.Year == InvalidYear
}

func (zdt *ZonedDateTime) ToEpochSeconds() int32 {
	if zdt.IsError() {
		return InvalidEpochSeconds
	}
	return (&OffsetDateTime{
		Year:          zdt.Year,
		Month:         zdt.Month,
		Day:           zdt.Day,
		Hour:          zdt.Hour,
		Minute:        zdt.Minute,
		Second:        zdt.Second,
		Fold:          zdt.Fold,
		OffsetMinutes: zdt.OffsetMinutes,
	}).ToEpochSeconds()
}

func NewZonedDateTimeFromEpochSeconds(
	epochSeconds int32, tz *TimeZone) ZonedDateTime {

	odt := tz.OffsetDateTimeFromEpochSeconds(epochSeconds)
	return ZonedDateTime{
		Year:          odt.Year,
		Month:         odt.Month,
		Day:           odt.Day,
		Hour:          odt.Hour,
		Minute:        odt.Minute,
		Second:        odt.Second,
		Fold:          odt.Fold,
		OffsetMinutes: odt.OffsetMinutes,
		Tz:            tz,
	}
}

func NewZonedDateTimeFromLocalDateTime(
	ldt *LocalDateTime, tz *TimeZone) ZonedDateTime {

	odt := tz.OffsetDateTimeFromLocalDateTime(ldt)
	return ZonedDateTime{
		Year:          odt.Year,
		Month:         odt.Month,
		Day:           odt.Day,
		Hour:          odt.Hour,
		Minute:        odt.Minute,
		Second:        odt.Second,
		Fold:          odt.Fold,
		OffsetMinutes: odt.OffsetMinutes,
		Tz:            tz,
	}
}

func (zdt *ZonedDateTime) ConvertToTimeZone(tz *TimeZone) ZonedDateTime {
	if zdt.IsError() {
		return NewZonedDateTimeError()
	}
	epochSeconds := zdt.ToEpochSeconds()
	if epochSeconds == InvalidEpochSeconds {
		return NewZonedDateTimeError()
	}
	return NewZonedDateTimeFromEpochSeconds(epochSeconds, tz)
}
