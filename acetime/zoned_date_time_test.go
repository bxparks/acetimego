package acetime

import (
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
	"unsafe"
)

//-----------------------------------------------------------------------------
// ZonedDateTime.
// Much of the following tests adapted from zoned_date_time_test.c from the
// acetimec library, which in turn, were adopted from
// ZonedDateTimeExtendedTest.ino in the AceTime library.
//-----------------------------------------------------------------------------

func TestZonedDateTimeSize(t *testing.T) {
	zdt := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2000, 1, 1, 1, 2, 3},
			OffsetSeconds: -8 * 3600,
		},
	}
	size := unsafe.Sizeof(zdt)
	if !(size == 32) { // assumes 64-bit alignment for *TimeZone pointer
		t.Fatal("Sizeof(ZonedDateTime): ", size)
	}
}

func TestZonedDateTimeUtcToString(t *testing.T) {
	tz := TimeZoneUTC
	ldt := LocalDateTime{2023, 1, 19, 17, 3, 23}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23UTC") {
		t.Fatal(s, zdt)
	}
}

func TestZonedDateTimeToString(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	ldt := LocalDateTime{2023, 1, 19, 17, 3, 23}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23-08:00[America/Los_Angeles]") {
		t.Fatal(s, zdt)
	}
}

//-----------------------------------------------------------------------------
// TimeZoneUTC()
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromUTC(t *testing.T) {
	// A UTC timezone
	tz := TimeZoneUTC
	if !(tz.Name() == "UTC") {
		t.Fatal(tz)
	}

	// Create a ZonedDateTime from a random epochSeconds.
	epochSeconds := Time(-32423234)
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}

	// Create the expected LocalDateTime.
	expected := LocalDateTimeFromEpochSeconds(epochSeconds)
	ldt := zdt.OffsetDateTime.LocalDateTime
	if !(expected == ldt) {
		t.Fatal(expected, zdt)
	}

	// String(). If TimeZone.IsUTC(), then the ISO8601 format is simplified.
	ldt = LocalDateTime{2023, 1, 19, 17, 3, 23}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23UTC") {
		t.Fatal(s, zdt)
	}
}

//-----------------------------------------------------------------------------
// ZonedDateTimeFromEpochSeconds()
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromEpochSeconds(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	var epochSeconds Time = 946684800
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{1999, 12, 31, 16, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.EpochSeconds()) {
		t.Fatal(zdt)
	}

	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromEpochSeconds_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	var epochSeconds Time = 2524608000
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2049, 12, 31, 16, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.EpochSeconds()) {
		t.Fatal(zdt.EpochSeconds())
	}

	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromEpochSeconds_UnixMax(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = (1 << 31) - 1
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2038, 1, 19, 3, 14, 7},
			OffsetSeconds: 0,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.EpochSeconds()) {
		t.Fatal(zdt.EpochSeconds())
	}
	expectedExtra := ZonedExtra{FoldTypeExact, 0, 0, 0, 0, "UTC"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromEpochSeconds_Invalid(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = InvalidEpochSeconds
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if !zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.EpochSeconds()) {
		t.Fatal(zdt)
	}
	extra := zdt.ZonedExtra()
	if !extra.IsError() {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromEpochSeconds_FallBack(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-07:00, 31 minutes before the overlap.
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 11, 6, 1, 29, 0},
		OffsetSeconds: -7 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 11, 6, 1, 29, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// Go forward an hour. Should return 01:29:00-08:00.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 11, 6, 1, 29, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST",
	}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// Go forward another hour. Should return 02:29:00-08:00.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 11, 6, 2, 29, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromEpochSeconds_SpringForward(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-08:00, 31 minutes before the DST gap.
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 3, 13, 1, 29, 0},
		OffsetSeconds: -8 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 3, 13, 1, 29, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 3, 13, 3, 29, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// ZonedDateTimeFromLocalDateTime()
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromLocalDateTime(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2000, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	epochSeconds := zdt.EpochSeconds()
	if !(epochSeconds == 946684800+8*60*60) {
		t.Fatal(epochSeconds)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, there is only one
	// match
	ldt = LocalDateTime{2000, 1, 1, 0, 0, 0}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2000, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	epochSeconds = zdt.EpochSeconds()
	if !(epochSeconds == 946684800+8*60*60) {
		t.Fatal(epochSeconds)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2050, 1, 1, 0, 0, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2050, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	epochSeconds := zdt.EpochSeconds()
	if !(epochSeconds == 2524608000+8*60*60) {
		t.Fatal(epochSeconds)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2050, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	epochSeconds = zdt.EpochSeconds()
	if !(epochSeconds == 2524608000+8*60*60) {
		t.Fatal(epochSeconds)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_BeforeGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 01:59 should resolve to 01:59-08:00
	ldt := LocalDateTime{2018, 3, 11, 1, 59, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 1, 59, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	ldt = LocalDateTime{2018, 3, 11, 1, 59, 0}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 1, 59, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_InGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 doesn't exist.
	ldt := LocalDateTime{2018, 3, 11, 2, 1, 0}

	// DisambiguateCompatible selects the later time, 03:01-07:00.
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	// Instead of calling 'zdt.ZonedExtra()', use the original LocalDateTime,
	// because the zdt has already been resolved to a real date time.
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateLater also selects the later time, 03:01-07:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateLater)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	// Instead of calling 'zdt.ZonedExtra()', use the original LocalDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateEarlier selects the earlier time, 01:01-08:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateEarlier)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	// Instead of calling 'zdt.ZonedExtra()', use the original LocalDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateReversed also selects the earlier time, 01:01-08:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	// Instead of calling 'zdt.ZonedExtra()', use the original LocalDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_AfterGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 03:01 should resolve to 03:01-07:00.
	ldt := LocalDateTime{2018, 3, 11, 3, 1, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_BeforeOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	ldt := LocalDateTime{2018, 11, 4, 0, 59, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 0, 59, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 0, 59, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_InOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// There were two instances of 01:01.
	ldt := LocalDateTime{2018, 11, 4, 1, 1, 0}

	// DisambiguateCompatible selects the earlier instance, 01:01-07:00.
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 1, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedOverlapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateEarlier selects the earlier instance, 01:01-07:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateEarlier)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 1, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedOverlapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateLater selects the later instance, 01:01-08:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateLater)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedOverlapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateReversed also selects the later instance, 01:01-08:00.
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedOverlapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromLocalDateTime_AfterOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 should resolve to 02:01-08:00
	ldt := LocalDateTime{2018, 11, 4, 2, 1, 0}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 2, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2018, 11, 4, 2, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------

func TestZonedDateTimeConvertToTimeZone(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tzLosAngeles := zm.TimeZoneFromName("America/Los_Angeles")
	tzNewYork := zm.TimeZoneFromName("America/New_York")

	// 2022-08-30 20:00-07:00 in LA
	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0}
	ladt := ZonedDateTimeFromLocalDateTime(
		&ldt, &tzLosAngeles, DisambiguateCompatible)
	if ladt.IsError() {
		t.Fatal(ladt)
	}

	// 2022-08-30 23:00-04:00 in NYC
	nydt := ladt.ConvertToTimeZone(&tzNewYork)
	if nydt.IsError() {
		t.Fatal(nydt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2022, 8, 30, 23, 0, 0},
			OffsetSeconds: -4 * 3600,
		},
		Tz:       &tzNewYork,
		Resolved: ResolvedUnique,
	}
	if !(nydt == expected) {
		t.Fatal(nydt)
	}
	expectedExtra := ZonedExtra{
		FoldTypeExact, -5 * 3600, 3600, -5 * 3600, 3600, "EDT",
	}
	extra := nydt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------

func TestZonedDateTimeToZonedExtra(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tzLosAngeles := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0}
	zdt := ZonedDateTimeFromLocalDateTime(
		&ldt, &tzLosAngeles, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}

	expected := ZonedExtra{
		FoldType:            FoldTypeExact,
		StdOffsetSeconds:    -8 * 3600,
		DstOffsetSeconds:    1 * 3600,
		ReqStdOffsetSeconds: -8 * 3600,
		ReqDstOffsetSeconds: 1 * 3600,
		Abbrev:              "PDT",
	}
	extra := zdt.ZonedExtra()
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// Test US/Pacific which is a Link to America/Los_Angeles
//-----------------------------------------------------------------------------

func TestZonedDateTimeForLink(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tzLosAngeles := zm.TimeZoneFromName("America/Los_Angeles")
	tzPacific := zm.TimeZoneFromName("US/Pacific")

	if !tzPacific.IsLink() {
		t.Fatal("US/Pacific should be a Link")
	}

	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0}
	ladt := ZonedDateTimeFromLocalDateTime(
		&ldt, &tzLosAngeles, DisambiguateCompatible)
	padt := ZonedDateTimeFromLocalDateTime(
		&ldt, &tzPacific, DisambiguateCompatible)

	if !(ladt.EpochSeconds() == padt.EpochSeconds()) {
		t.Fatal("epochSeconds not equal")
	}
}

//-----------------------------------------------------------------------------
// Test Normalize()
//-----------------------------------------------------------------------------

// Test ported from test(ZonedDateTimExtendedTest, normalize) in AceTime
// library.
func TestZonedDateTimeNormalize(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start with epochSeconds = 946684800. Should translate to
	// 1999-12-31T16:00:00-08:00. Note: epochSeconds = 0 does not work because
	// zonedbtesting database is not valid before 1980.
	zdt := ZonedDateTimeFromEpochSeconds(946684800, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{1999, 12, 31, 16, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if zdt != expected {
		t.Fatal(zdt)
	}

	// Set date-time to 2021-04-20T09:00:00, which happens to be in DST.
	zdt.Year = 2021
	zdt.Month = 4
	zdt.Day = 20
	zdt.Hour = 9
	zdt.Minute = 0
	zdt.Second = 0

	// If we blindly use the resulting epochSeconds to convert to a new
	// ZonedDateTime, we will be off by one hour, because the previous
	// OffsetDateTime had an offset of (-08:00), which does not match offset of
	// the explicitly specified date (-07:00).
	epochSeconds := zdt.EpochSeconds()
	newDt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2021, 4, 20, 10, 0, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if newDt != expected {
		t.Fatal(newDt)
	}

	// We must Normalize() after mutation.
	zdt.Normalize(DisambiguateCompatible)
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			LocalDateTime: LocalDateTime{2021, 4, 20, 9, 0, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if zdt != expected {
		t.Fatal(zdt.Year)
	}
}

//-----------------------------------------------------------------------------
// Benchmarks
// $ go test -run=NOMATCH -bench=.
//-----------------------------------------------------------------------------

var epochSeconds Time
var zdt ZonedDateTime
var ldt = LocalDateTime{2023, 1, 19, 22, 11, 0}
var zoneManager = ZoneManagerFromDataContext(&zonedbtesting.DataContext)
var tz = zoneManager.TimeZoneFromZoneID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

func BenchmarkZonedDateTimeFromEpochSeconds_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = ZonedDateTimeFromEpochSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromEpochSeconds_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.processor.reset()
		zdt = ZonedDateTimeFromEpochSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromLocalDateTime_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	}
}

func BenchmarkZonedDateTimeFromLocalDateTime_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.processor.reset()
		zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	}
}

func BenchmarkZonedDateTimeEpochSeconds(b *testing.B) {
	zdt = ZonedDateTimeFromEpochSeconds(3423423, &tz)
	for n := 0; n < b.N; n++ {
		epochSeconds = zdt.EpochSeconds()
	}
}
