package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
	"unsafe"
)

//-----------------------------------------------------------------------------
// ZonedDateTime.
// Much of the following tests adapted from zoned_date_time_test.c from the
// AceTimeC library, which in turn, were adopted from
// ZonedDateTimeExtendedTest.ino in the AceTime library.
//-----------------------------------------------------------------------------

func TestZonedDateTimeSize(t *testing.T) {
	zdt := ZonedDateTime{2000, 1, 1, 1, 2, 3, 0 /*Fold*/, -8 * 60, nil}
	size := unsafe.Sizeof(zdt)
	if !(size == 24) { // assumes 64-bit alignment for *TimeZone pointer
		t.Fatal("Sizeof(ZonedDateTime): ", size)
	}
}

func TestZonedDateTimeFromEpochSeconds(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	var epochSeconds int32 = 0
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{1999, 12, 31, 16, 0, 0, 0, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.ToEpochSeconds()) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromEpochSeconds_2050(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	var epochSeconds int32 = 0
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2049, 12, 31, 16, 0, 0, 0, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.ToEpochSeconds()) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromEpochSeconds_UnixMax(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneEtc_UTC)
	var epochSeconds int32 = 1200798847
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2038, 1, 19, 3, 14, 7, 0, 0, &tz}) {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.ToEpochSeconds()) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromEpochSeconds_Invalid(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneEtc_UTC)
	var epochSeconds int32 = InvalidEpochSeconds
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if !zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.ToEpochSeconds()) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromEpochSeconds_FallBack(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back.
	odt := OffsetDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60}
	epochSeconds := odt.ToEpochSeconds()
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Go forward an hour. Should return 01:29:00-08:00, the second time this
	// was seen, so fold should be 1.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 1, 29, 0, 1 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Go forward another hour. Should return 02:29:00-08:00, which occurs only
	// once, so fold should be 0.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 2, 29, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromEpochSeconds_SpringForward(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-08:00, which is 31 minutes before the DST
	// spring forward.
	odt := OffsetDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60}
	epochSeconds := odt.ToEpochSeconds()
	zdt := ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	epochSeconds += 3600
	zdt = ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 3, 13, 3, 29, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

//-----------------------------------------------------------------------------

func TestZonedDateTimeFromLocalDateTime(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	epochSeconds := zdt.ToEpochSeconds()
	if !(epochSeconds == 8*60*60) {
		t.Fatal(epochSeconds)
	}

	// check that fold=1 gives identical results, there is only one match
	ldt = LocalDateTime{2000, 1, 1, 0, 0, 0, 1 /*Fold*/}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	epochSeconds = zdt.ToEpochSeconds()
	if !(epochSeconds == 8*60*60) {
		t.Fatal(epochSeconds)
	}
}

func TestZonedDateTimeFromLocalDateTime_2050(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	ldt := LocalDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	epochSeconds := zdt.ToEpochSeconds()
	if !(epochSeconds == 8*60*60) {
		t.Fatal(epochSeconds)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0, 1 /*Fold*/}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
	epochSeconds = zdt.ToEpochSeconds()
	if !(epochSeconds == 8*60*60) {
		t.Fatal(epochSeconds)
	}
}

func TestZonedDateTimeFromLocalDateTime_BeforeDst(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 01:59 should resolve to 01:59-08:00
	ldt := LocalDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt = LocalDateTime{2018, 3, 11, 1, 59, 0, 1 /*Fold*/}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromLocalDateTime_InGap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 02:01 doesn't exist.
	// Setting (fold=0) causes the first transition to be selected, which has a
	// UTC offset of -08:00, so this is interpreted as 02:01-08:00 which gets
	// normalized to 03:01-07:00, which falls in the 2nd transition.
	ldt := LocalDateTime{2018, 3, 11, 2, 1, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	// fold == 0 to indicate only one match
	if !(zdt == ZonedDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Setting (fold=1) causes the second transition to be selected, which has a
	// UTC offset of -07:00, so this is interpreted as 02:01-07:00 which gets
	// normalized to 01:01-08:00, which falls in the 1st transition.
	ldt = LocalDateTime{2018, 3, 11, 2, 1, 0, 1 /*Fold*/}
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	// fold == 0 to indicate the 1st transition
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 1, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromLocalDateTime_InDst(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 03:01 should resolve to 03:01-07:00.
	ldt := LocalDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromLocalDateTime_BeforeSdt(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	ldt := LocalDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromLocalDateTime_InOverlap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// There were two instances of 01:01
	// Setting (fold==0) selects the first instance, resolves to 01:01-07:00.
	ldt := LocalDateTime{2018, 11, 4, 1, 1, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 1, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Setting (fold==1) selects the second instance, resolves to 01:01-08:00.
	ldt.Fold = 1
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 1, 1, 0, 1 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestZonedDateTimeFromLocalDateTime_AfterOverlap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 02:01 should resolve to 02:01-08:00
	ldt := LocalDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/}
	zdt := ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = ZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

//-----------------------------------------------------------------------------

func TestZonedDateTimeConvertToTimeZone(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tzLosAngeles := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	tzNewYork := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_New_York)

	// 2022-08-30 20:00-07:00 in LA
	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0, 0 /*Fold*/}
	ladt := ZonedDateTimeFromLocalDateTime(&ldt, &tzLosAngeles)
	if ladt.IsError() {
		t.Fatal(ladt)
	}

	// 2022-08-30 23:00-04:00 in NYC
	nydt := ladt.ConvertToTimeZone(&tzNewYork)
	if nydt.IsError() {
		t.Fatal(nydt)
	}
	if !(nydt == ZonedDateTime{
		2022, 8, 30, 23, 0, 0, 0 /*Fold*/, -4 * 60, &tzNewYork}) {
		t.Fatal(nydt)
	}
}

//-----------------------------------------------------------------------------
// Test US/Pacific which is a Link to America/Los_Angeles
//-----------------------------------------------------------------------------

func TestZonedDateTimeForLink(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tzLosAngeles := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	tzPacific := NewTimeZoneForZoneInfo(&zonedbtesting.ZoneUS_Pacific)

	if !zonedbtesting.ZoneUS_Pacific.IsLink() {
		t.Fatal("US/Pacific should be a Link")
	}

	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0, 0 /*Fold*/}
	ladt := ZonedDateTimeFromLocalDateTime(&ldt, &tzLosAngeles)
	padt := ZonedDateTimeFromLocalDateTime(&ldt, &tzPacific)

	if !(ladt.ToEpochSeconds() == padt.ToEpochSeconds()) {
		t.Fatal("epochSeconds not equal")
	}
}
