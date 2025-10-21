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
			PlainDateTime: PlainDateTime{2000, 1, 1, 1, 2, 3},
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
	pdt := PlainDateTime{2023, 1, 19, 17, 3, 23}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23UTC") {
		t.Fatal(s, zdt)
	}
}

func TestZonedDateTimeToString(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	pdt := PlainDateTime{2023, 1, 19, 17, 3, 23}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
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

	// Create a ZonedDateTime from a random unixSeconds.
	unixSeconds := Time(-32423234)
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}

	// Create the expected PlainDateTime.
	expected := PlainDateTimeFromUnixSeconds(unixSeconds)
	pdt := zdt.OffsetDateTime.PlainDateTime
	if !(expected == pdt) {
		t.Fatal(expected, zdt)
	}

	// String(). If TimeZone.IsUTC(), then the ISO8601 format is simplified.
	pdt = PlainDateTime{2023, 1, 19, 17, 3, 23}
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23UTC") {
		t.Fatal(s, zdt)
	}
}

//-----------------------------------------------------------------------------
// ZonedDateTimeFromUnixSeconds()
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromUnixSeconds(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	var unixSeconds Time = 946684800
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{1999, 12, 31, 16, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(unixSeconds == zdt.UnixSeconds()) {
		t.Fatal(zdt)
	}

	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromUnixSeconds_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	var unixSeconds Time = 2524608000
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2049, 12, 31, 16, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(unixSeconds == zdt.UnixSeconds()) {
		t.Fatal(zdt.UnixSeconds())
	}

	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromUnixSeconds_UnixMax(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var unixSeconds Time = (1 << 31) - 1
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2038, 1, 19, 3, 14, 7},
			OffsetSeconds: 0,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	if !(unixSeconds == zdt.UnixSeconds()) {
		t.Fatal(zdt.UnixSeconds())
	}
	expectedExtra := ZonedExtra{FoldTypeExact, 0, 0, 0, 0, "UTC"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromUnixSeconds_Invalid(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var unixSeconds Time = InvalidUnixSeconds
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if !zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(unixSeconds == zdt.UnixSeconds()) {
		t.Fatal(zdt)
	}
	extra := zdt.ZonedExtra()
	if !extra.IsError() {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromUnixSeconds_FallBack(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-07:00, 31 minutes before the overlap.
	odt := OffsetDateTime{
		PlainDateTime: PlainDateTime{2022, 11, 6, 1, 29, 0},
		OffsetSeconds: -7 * 3600,
	}
	unixSeconds := odt.UnixSeconds()
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2022, 11, 6, 1, 29, 0},
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
	unixSeconds += 3600
	zdt = ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2022, 11, 6, 1, 29, 0},
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
	unixSeconds += 3600
	zdt = ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2022, 11, 6, 2, 29, 0},
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

func TestZonedDateTimeFromUnixSeconds_SpringForward(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-08:00, 31 minutes before the DST gap.
	odt := OffsetDateTime{
		PlainDateTime: PlainDateTime{2022, 3, 13, 1, 29, 0},
		OffsetSeconds: -8 * 3600,
	}
	unixSeconds := odt.UnixSeconds()
	zdt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2022, 3, 13, 1, 29, 0},
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
	unixSeconds += 3600
	zdt = ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2022, 3, 13, 3, 29, 0},
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
// ZonedDateTimeFromPlainDateTime()
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromPlainDateTime(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	pdt := PlainDateTime{2000, 1, 1, 0, 0, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2000, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	unixSeconds := zdt.UnixSeconds()
	if !(unixSeconds == 946684800+8*60*60) {
		t.Fatal(unixSeconds)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, there is only one
	// match
	pdt = PlainDateTime{2000, 1, 1, 0, 0, 0}
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2000, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	unixSeconds = zdt.UnixSeconds()
	if !(unixSeconds == 946684800+8*60*60) {
		t.Fatal(unixSeconds)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromPlainDateTime_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	pdt := PlainDateTime{2050, 1, 1, 0, 0, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2050, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	unixSeconds := zdt.UnixSeconds()
	if !(unixSeconds == 2524608000+8*60*60) {
		t.Fatal(unixSeconds)
	}
	expectedExtra := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra := zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	pdt = PlainDateTime{2050, 1, 1, 0, 0, 0}
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2050, 1, 1, 0, 0, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedUnique,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	unixSeconds = zdt.UnixSeconds()
	if !(unixSeconds == 2524608000+8*60*60) {
		t.Fatal(unixSeconds)
	}
	expectedExtra = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	extra = zdt.ZonedExtra()
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromPlainDateTime_BeforeGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 01:59 should resolve to 01:59-08:00
	pdt := PlainDateTime{2018, 3, 11, 1, 59, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 1, 59, 0},
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
	pdt = PlainDateTime{2018, 3, 11, 1, 59, 0}
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 1, 59, 0},
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

func TestZonedDateTimeFromPlainDateTime_InGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 doesn't exist.
	pdt := PlainDateTime{2018, 3, 11, 2, 1, 0}

	// DisambiguateCompatible selects the later time, 03:01-07:00.
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra := ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	// Instead of calling 'zdt.ZonedExtra()', use the original PlainDateTime,
	// because the zdt has already been resolved to a real date time.
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateLater also selects the later time, 03:01-07:00.
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateLater)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 3, 1, 0},
			OffsetSeconds: -7 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapLater,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	// Instead of calling 'zdt.ZonedExtra()', use the original PlainDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateEarlier selects the earlier time, 01:01-08:00.
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateEarlier)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	// Instead of calling 'zdt.ZonedExtra()', use the original PlainDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra != expectedExtra {
		t.Fatal(extra)
	}

	// DisambiguateReversed also selects the earlier time, 01:01-08:00.
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 1, 1, 0},
			OffsetSeconds: -8 * 3600,
		},
		Tz:       &tz,
		Resolved: ResolvedGapEarlier,
	}
	if !(zdt == expected) {
		t.Fatal(zdt)
	}
	expectedExtra = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	// Instead of calling 'zdt.ZonedExtra()', use the original PlainDateTime,
	// because the zdt has already been resolved to a real date time.
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra != expectedExtra {
		t.Fatal(extra)
	}
}

func TestZonedDateTimeFromPlainDateTime_AfterGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 03:01 should resolve to 03:01-07:00.
	pdt := PlainDateTime{2018, 3, 11, 3, 1, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 3, 1, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 3, 11, 3, 1, 0},
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

func TestZonedDateTimeFromPlainDateTime_BeforeOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	pdt := PlainDateTime{2018, 11, 4, 0, 59, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 0, 59, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 0, 59, 0},
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

func TestZonedDateTimeFromPlainDateTime_InOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// There were two instances of 01:01.
	pdt := PlainDateTime{2018, 11, 4, 1, 1, 0}

	// DisambiguateCompatible selects the earlier instance, 01:01-07:00.
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 1, 1, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateEarlier)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 1, 1, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateLater)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 1, 1, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 1, 1, 0},
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

func TestZonedDateTimeFromPlainDateTime_AfterOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 should resolve to 02:01-08:00
	pdt := PlainDateTime{2018, 11, 4, 2, 1, 0}
	zdt := ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 2, 1, 0},
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
	zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2018, 11, 4, 2, 1, 0},
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
	pdt := PlainDateTime{2022, 8, 30, 20, 0, 0}
	ladt := ZonedDateTimeFromPlainDateTime(
		&pdt, &tzLosAngeles, DisambiguateCompatible)
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
			PlainDateTime: PlainDateTime{2022, 8, 30, 23, 0, 0},
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

	pdt := PlainDateTime{2022, 8, 30, 20, 0, 0}
	zdt := ZonedDateTimeFromPlainDateTime(
		&pdt, &tzLosAngeles, DisambiguateCompatible)
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

	pdt := PlainDateTime{2022, 8, 30, 20, 0, 0}
	ladt := ZonedDateTimeFromPlainDateTime(
		&pdt, &tzLosAngeles, DisambiguateCompatible)
	padt := ZonedDateTimeFromPlainDateTime(
		&pdt, &tzPacific, DisambiguateCompatible)

	if !(ladt.UnixSeconds() == padt.UnixSeconds()) {
		t.Fatal("unixSeconds not equal")
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

	// Start with unixSeconds = 946684800. Should translate to
	// 1999-12-31T16:00:00-08:00. Note: unixSeconds = 0 does not work because
	// zonedbtesting database is not valid before 1980.
	zdt := ZonedDateTimeFromUnixSeconds(946684800, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	expected := ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{1999, 12, 31, 16, 0, 0},
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

	// If we blindly use the resulting unixSeconds to convert to a new
	// ZonedDateTime, we will be off by one hour, because the previous
	// OffsetDateTime had an offset of (-08:00), which does not match offset of
	// the explicitly specified date (-07:00).
	unixSeconds := zdt.UnixSeconds()
	newDt := ZonedDateTimeFromUnixSeconds(unixSeconds, &tz)
	expected = ZonedDateTime{
		OffsetDateTime: OffsetDateTime{
			PlainDateTime: PlainDateTime{2021, 4, 20, 10, 0, 0},
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
			PlainDateTime: PlainDateTime{2021, 4, 20, 9, 0, 0},
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

var unixSeconds Time
var zdt ZonedDateTime
var pdt = PlainDateTime{2023, 1, 19, 22, 11, 0}
var zoneManager = ZoneManagerFromDataContext(&zonedbtesting.DataContext)
var tz = zoneManager.TimeZoneFromZoneID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

func BenchmarkZonedDateTimeFromUnixSeconds_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = ZonedDateTimeFromUnixSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromUnixSeconds_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.processor.reset()
		zdt = ZonedDateTimeFromUnixSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromPlainDateTime_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	}
}

func BenchmarkZonedDateTimeFromPlainDateTime_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.processor.reset()
		zdt = ZonedDateTimeFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	}
}

func BenchmarkZonedDateTimeUnixSeconds(b *testing.B) {
	zdt = ZonedDateTimeFromUnixSeconds(3423423, &tz)
	for n := 0; n < b.N; n++ {
		unixSeconds = zdt.UnixSeconds()
	}
}
