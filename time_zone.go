package acetime

//-----------------------------------------------------------------------------
// TimeZone represents one of the IANA TZ time zones. This is a reference type,
// and meant to be passed around as a pointer and garbage collected when it is
// no longer used.
//-----------------------------------------------------------------------------

type TimeZone struct {
	zoneProcessor ZoneProcessor
}

func NewTimeZoneForZoneInfo(zoneInfo *ZoneInfo) TimeZone {
	var tz TimeZone
	tz.zoneProcessor.InitForZoneInfo(zoneInfo)
	return tz
}

func (tz *TimeZone) OffsetDateDateTimeFromLocalDateTime(
	ldt *LocalDateTime, fold uint8) OffsetDateTime {

	return tz.zoneProcessor.OffsetDateTimeFromLocalDateTime(ldt, fold)
}

func (tz *TimeZone) OffsetDateDateTimeFromEpochSeconds(
	epochSeconds int32) OffsetDateTime {

	return tz.zoneProcessor.OffsetDateTimeFromEpochSeconds(epochSeconds)
}
