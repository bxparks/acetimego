package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

//-----------------------------------------------------------------------------
// ZonedExtra.
// Extra meta information about a given instant in time, such as the
// the STD offset, the DST offset, and the abbreviation used.
//-----------------------------------------------------------------------------

// Test that ZonedExtraXxx constants are the same as FindResultXxx constants.
func TestZonedExtraTypeConstantsMatch(t *testing.T) {
	if !(ZonedExtraErr == FindResultErr) {
		t.Fatal("")
	}
	if !(ZonedExtraNotFound == FindResultNotFound) {
		t.Fatal("")
	}
	if !(ZonedExtraExact == FindResultExact) {
		t.Fatal("")
	}
	if !(ZonedExtraGap == FindResultGap) {
		t.Fatal("")
	}
	if !(ZonedExtraOverlap == FindResultOverlap) {
		t.Fatal("")
	}
}

func TestZonedExtraFromEpochSeconds(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.NewTimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	ze := NewZonedExtraFromEpochSeconds(InvalidEpochSeconds, &tz)
	if !(ze.Zetype == ZonedExtraErr) {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_FallBack(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.NewTimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back, and occurs in the overlap.
	odt := OffsetDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60}
	epochSeconds := odt.ToEpochSeconds()

	ze := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected := ZonedExtra{
		Zetype:              ZonedExtraOverlap,
		StdOffsetMinutes:    -8 * 60,
		DstOffsetMinutes:    1 * 60,
		ReqStdOffsetMinutes: -8 * 60,
		ReqDstOffsetMinutes: 1 * 60,
		Abbrev:              "PDT",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}

	// Go forward an hour, should be 01:29:00-08:00, which is again in the
	// overlap.
	epochSeconds += 3600
	ze = NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected = ZonedExtra{
		Zetype:              ZonedExtraOverlap,
		StdOffsetMinutes:    -8 * 60,
		DstOffsetMinutes:    0 * 60,
		ReqStdOffsetMinutes: -8 * 60,
		ReqDstOffsetMinutes: 0 * 60,
		Abbrev:              "PST",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_SpringForward(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.NewTimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// spring forward.
	odt := OffsetDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60}
	epochSeconds := odt.ToEpochSeconds()

	ze := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected := ZonedExtra{
		Zetype:              ZonedExtraExact,
		StdOffsetMinutes:    -8 * 60,
		DstOffsetMinutes:    0 * 60,
		ReqStdOffsetMinutes: -8 * 60,
		ReqDstOffsetMinutes: 0 * 60,
		Abbrev:              "PST",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}

	// Go forward an hour. Should be 01:29:00-08:00.
	epochSeconds += 3600
	ze = NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected = ZonedExtra{
		Zetype:              ZonedExtraExact,
		StdOffsetMinutes:    -8 * 60,
		DstOffsetMinutes:    1 * 60,
		ReqStdOffsetMinutes: -8 * 60,
		ReqDstOffsetMinutes: 1 * 60,
		Abbrev:              "PDT",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}
}
