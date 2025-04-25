package acetime

import (
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
)

//-----------------------------------------------------------------------------
// ZonedExtra.
//-----------------------------------------------------------------------------

// Test that FoldTypeXxx constants are the same as findResultXxx constants.
func TestFoldTypeTypeConstantsMatch(t *testing.T) {
	if !(FoldTypeErr == findResultErr) {
		t.Fatal("FoldTypeErr")
	}
	if !(FoldTypeNotFound == findResultNotFound) {
		t.Fatal("FoldTypeNotFound")
	}
	if !(FoldTypeExact == findResultExact) {
		t.Fatal("FoldTypeExact")
	}
	if !(FoldTypeGap == findResultGap) {
		t.Fatal("FoldTypeGap")
	}
	if !(FoldTypeOverlap == findResultOverlap) {
		t.Fatal("FoldTypeOverlap")
	}
}

func TestOffsetSeconds(t *testing.T) {
	extra := ZonedExtra{FoldTypeExact, 1, 2, 3, 4, "ABC"}
	if extra.OffsetSeconds() != 1+2 {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// ZonedExtraFromEpochSeconds(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestZonedExtraFromEpochSeconds(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	var epochSeconds Time = 946684800
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromEpochSeconds_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	var epochSeconds Time = 2524608000
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromEpochSeconds_UnixMax(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = (1 << 31) - 1
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, 0, 0, 0, 0, "UTC"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromEpochSeconds_Invalid(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = InvalidEpochSeconds
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if !extra.IsError() {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromEpochSeconds_FallBack(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start sampling at 01:29:00-07:00, 31 minutes before overlap
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 11, 6, 1, 29, 0},
		OffsetSeconds: -7 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Go forward an hour. Should return 01:29:00-08:00, the second time this
	// was seen, so fold should be 1.
	epochSeconds += 3600
	extra = ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Go forward another hour. Should return 02:29:00-08:00, which occurs only
	// once, so fold should be 0.
	epochSeconds += 3600
	extra = ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromEpochSeconds_SpringForward(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start sampling at 01:29:00-08:00, 31 minutes before gap
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 3, 13, 1, 29, 0},
		OffsetSeconds: -8 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	extra := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	epochSeconds += 3600
	extra = ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// ZonedExtraFromLocalDateTime(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestZonedExtraFromLocalDateTime(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, there is only one
	// match
	ldt = LocalDateTime{2000, 1, 1, 0, 0, 0}
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_2050(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2050, 1, 1, 0, 0, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0}
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_BeforeGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 01:59 should resolve to 01:59-08:00
	ldt := LocalDateTime{2018, 3, 11, 1, 59, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	ldt = LocalDateTime{2018, 3, 11, 1, 59, 0}
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_InGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 doesn't exist.
	// DisambiguateCompatible selects the later time in the gap.
	ldt := LocalDateTime{2018, 3, 11, 2, 1, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// DisambiguateReversed selects the earlier time in the gap.
	ldt = LocalDateTime{2018, 3, 11, 2, 1, 0}
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	// fold == 0 to indicate the 1st transition
	expected = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_AfterGap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 03:01 should resolve to 03:01-07:00.
	ldt := LocalDateTime{2018, 3, 11, 3, 1, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_BeforeOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	ldt := LocalDateTime{2018, 11, 4, 0, 59, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_InOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// There were two instances of 01:01
	// DisambiguateCompatible selects the first instance, resolves to 01:01-07:00.
	ldt := LocalDateTime{2018, 11, 4, 1, 1, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// DisambiguateReversed selects the second instance, resolves to 01:01-08:00.
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestZonedExtraFromLocalDateTime_AfterOverlap(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 should resolve to 02:01-08:00
	ldt := LocalDateTime{2018, 11, 4, 2, 1, 0}
	extra := ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateCompatible)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that DisambiguateReversed gives identical results, since there is one
	// match
	extra = ZonedExtraFromLocalDateTime(&ldt, &tz, DisambiguateReversed)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}
