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

// Test that ZonedExtraXxx constants are the same as findResultXxx constants.
func TestZonedExtraTypeConstantsMatch(t *testing.T) {
	if !(ZonedExtraErr == findResultErr) {
		t.Fatal("")
	}
	if !(ZonedExtraNotFound == findResultNotFound) {
		t.Fatal("")
	}
	if !(ZonedExtraExact == findResultExact) {
		t.Fatal("")
	}
	if !(ZonedExtraGap == findResultGap) {
		t.Fatal("")
	}
	if !(ZonedExtraOverlap == findResultOverlap) {
		t.Fatal("")
	}
}

func TestZonedExtraFromEpochSeconds(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.TimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	ze := NewZonedExtraFromEpochSeconds(InvalidEpochSeconds, &tz)
	if !(ze.Zetype == ZonedExtraErr) {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_FallBack(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.TimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back, and occurs in the overlap.
	odt := OffsetDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 3600}
	epochSeconds := odt.EpochSeconds()

	ze := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected := ZonedExtra{
		Zetype:              ZonedExtraOverlap,
		StdOffsetSeconds:    -8 * 3600,
		DstOffsetSeconds:    1 * 3600,
		ReqStdOffsetSeconds: -8 * 3600,
		ReqDstOffsetSeconds: 1 * 3600,
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
		StdOffsetSeconds:    -8 * 3600,
		DstOffsetSeconds:    0 * 3600,
		ReqStdOffsetSeconds: -8 * 3600,
		ReqDstOffsetSeconds: 0 * 3600,
		Abbrev:              "PST",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_SpringForward(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	tz := manager.TimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// spring forward.
	odt := OffsetDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 3600}
	epochSeconds := odt.EpochSeconds()

	ze := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.Zetype == ZonedExtraErr {
		t.Fatal(ze)
	}
	expected := ZonedExtra{
		Zetype:              ZonedExtraExact,
		StdOffsetSeconds:    -8 * 3600,
		DstOffsetSeconds:    0 * 3600,
		ReqStdOffsetSeconds: -8 * 3600,
		ReqDstOffsetSeconds: 0 * 3600,
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
		StdOffsetSeconds:    -8 * 3600,
		DstOffsetSeconds:    1 * 3600,
		ReqStdOffsetSeconds: -8 * 3600,
		ReqDstOffsetSeconds: 1 * 3600,
		Abbrev:              "PDT",
	}
	if !(ze == expected) {
		t.Fatal(ze)
	}
}
