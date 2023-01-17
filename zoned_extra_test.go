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

func TestZonedExtraFromEpochSeconds(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := TimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	ze := ZonedExtraFromEpochSeconds(InvalidEpochSeconds, &tz)
	if !ze.IsError() {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_FallBack(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := TimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back.
	odt := OffsetDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60}
	epochSeconds := odt.ToEpochSeconds()

	ze := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.IsError() {
		t.Fatal(ze)
	}
	if !(ze == ZonedExtra{-8 * 60, 1 * 60, "PDT"}) {
		t.Fatal(ze)
	}

	// Go forward an hour. Should be 01:29:00-08:00.
	epochSeconds += 3600
	ze = ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.IsError() {
		t.Fatal(ze)
	}
	if !(ze == ZonedExtra{-8 * 60, 0 * 60, "PST"}) {
		t.Fatal(ze)
	}
}

func TestZonedExtraFromEpochSeconds_SpringForward(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := TimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back.
	odt := OffsetDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60}
	epochSeconds := odt.ToEpochSeconds()

	ze := ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.IsError() {
		t.Fatal(ze)
	}
	if !(ze == ZonedExtra{-8 * 60, 0 * 60, "PST"}) {
		t.Fatal(ze)
	}

	// Go forward an hour. Should be 01:29:00-08:00.
	epochSeconds += 3600
	ze = ZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if ze.IsError() {
		t.Fatal(ze)
	}
	if !(ze == ZonedExtra{-8 * 60, 1 * 60, "PDT"}) {
		t.Fatal(ze)
	}
}
