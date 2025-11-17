package acetime

import (
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
)

//-----------------------------------------------------------------------------
// ZonedExtra.
//-----------------------------------------------------------------------------

func TestOffsetSeconds(t *testing.T) {
	extra := ZonedExtra{ResolvedUnique, 1, 2, 3, 4, "ABC"}
	if extra.OffsetSeconds() != 1+2 {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// ZonedExtraFromUnixSeconds(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestZonedExtraFromUnixSeconds(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	var unixSeconds Time = 946684800
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromUnixSeconds_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	var unixSeconds Time = 2524608000
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromUnixSeconds_UnixMax(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var unixSeconds Time = (1 << 31) - 1
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, 0, 0, 0, 0, "UTC"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromUnixSeconds_Invalid(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var unixSeconds Time = InvalidUnixSeconds
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if !extra.IsError() {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromUnixSeconds_FallBack(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start sampling at 01:29:00-07:00, 31 minutes before overlap
	odt := OffsetDateTime{
		PlainDateTime: PlainDateTime{2022, 11, 6, 1, 29, 0},
		OffsetSeconds: -7 * 3600,
	}
	unixSeconds := odt.UnixSeconds()
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Go forward an hour. Should return 01:29:00-08:00, the second time this
	// was seen.
	unixSeconds += 3600
	extra = ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Go forward another hour. Should return 02:29:00-08:00, which occurs only
	// once, so fold should be 0.
	unixSeconds += 3600
	extra = ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromUnixSeconds_SpringForward(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start sampling at 01:29:00-08:00, 31 minutes before gap
	odt := OffsetDateTime{
		PlainDateTime: PlainDateTime{2022, 3, 13, 1, 29, 0},
		OffsetSeconds: -8 * 3600,
	}
	unixSeconds := odt.UnixSeconds()
	extra := ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	unixSeconds += 3600
	extra = ZonedExtraFromUnixSeconds(unixSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// ZonedExtraFromPlainDateTime(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestZonedExtraFromPlainDateTime(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	pdt := PlainDateTime{2000, 1, 1, 0, 0, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, there is only one
	// match
	pdt = PlainDateTime{2000, 1, 1, 0, 0, 0}
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	pdt := PlainDateTime{2050, 1, 1, 0, 0, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	pdt = PlainDateTime{2050, 1, 1, 0, 0, 0}
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_BeforeGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 01:59 should resolve to 01:59-08:00
	pdt := PlainDateTime{2018, 3, 11, 1, 59, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	pdt = PlainDateTime{2018, 3, 11, 1, 59, 0}
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_InGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 doesn't exist.
	// DisambiguateCompatible selects the later time in the gap.
	pdt := PlainDateTime{2018, 3, 11, 2, 1, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedGapLater, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// DisambiguateReversed selects the earlier time in the gap.
	pdt = PlainDateTime{2018, 3, 11, 2, 1, 0}
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	// fold == 0 to indicate the 1st transition
	expected = ZonedExtra{
		ResolvedGapEarlier, -8 * 3600, 0, -8 * 3600, 3600, "PST",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_AfterGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 03:01 should resolve to 03:01-07:00.
	pdt := PlainDateTime{2018, 3, 11, 3, 1, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_BeforeOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	pdt := PlainDateTime{2018, 11, 4, 0, 59, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_InOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// There were two instances of 01:01
	// DisambiguateCompatible selects the first instance, resolves to 01:01-07:00.
	pdt := PlainDateTime{2018, 11, 4, 1, 1, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		ResolvedOverlapEarlier, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// DisambiguateReversed selects the second instance, resolves to 01:01-08:00.
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedOverlapLater, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromPlainDateTime_AfterOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 should resolve to 02:01-08:00
	pdt := PlainDateTime{2018, 11, 4, 2, 1, 0}
	extra := ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromPlainDateTime(&pdt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{ResolvedUnique, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}
