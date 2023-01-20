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

func TestZonedDateTimeToString(t *testing.T) {
	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	ldt := LocalDateTime{2023, 1, 19, 17, 3, 23, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	s := zdt.String()
	if !(s == "2023-01-19T17:03:23-08:00[America/Los_Angeles]") {
		t.Fatal(s, zdt)
	}
}

func TestNewZonedDateTimeFromEpochSeconds(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	var epochSeconds int32 = 0
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
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

func TestNewZonedDateTimeFromEpochSeconds_2050(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	var epochSeconds int32 = 0
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
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

func TestNewZonedDateTimeFromEpochSeconds_UnixMax(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneEtc_UTC)
	var epochSeconds int32 = 1200798847
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
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

func TestNewZonedDateTimeFromEpochSeconds_Invalid(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneEtc_UTC)
	var epochSeconds int32 = InvalidEpochSeconds
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if !zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(epochSeconds == zdt.ToEpochSeconds()) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromEpochSeconds_FallBack(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back.
	odt := OffsetDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60}
	epochSeconds := odt.ToEpochSeconds()
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Go forward an hour. Should return 01:29:00-08:00, the second time this
	// was seen, so fold should be 1.
	epochSeconds += 3600
	zdt = NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 1, 29, 0, 1 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Go forward another hour. Should return 02:29:00-08:00, which occurs only
	// once, so fold should be 0.
	epochSeconds += 3600
	zdt = NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 11, 6, 2, 29, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromEpochSeconds_SpringForward(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// Start our sampling at 01:29:00-08:00, which is 31 minutes before the DST
	// spring forward.
	odt := OffsetDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60}
	epochSeconds := odt.ToEpochSeconds()
	zdt := NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	epochSeconds += 3600
	zdt = NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2022, 3, 13, 3, 29, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

//-----------------------------------------------------------------------------

func TestNewZonedDateTimeFromLocalDateTime(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2000)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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

func TestNewZonedDateTimeFromLocalDateTime_2050(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	ldt := LocalDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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

func TestNewZonedDateTimeFromLocalDateTime_BeforeDst(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 01:59 should resolve to 01:59-08:00
	ldt := LocalDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt = LocalDateTime{2018, 3, 11, 1, 59, 0, 1 /*Fold*/}
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromLocalDateTime_InGap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 02:01 doesn't exist.
	// Setting (fold=0) causes the first transition to be selected, which has a
	// UTC offset of -08:00, so this is interpreted as 02:01-08:00 which gets
	// normalized to 03:01-07:00, which falls in the 2nd transition.
	ldt := LocalDateTime{2018, 3, 11, 2, 1, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	// fold == 0 to indicate the 1st transition
	if !(zdt == ZonedDateTime{2018, 3, 11, 1, 1, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromLocalDateTime_InDst(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 03:01 should resolve to 03:01-07:00.
	ldt := LocalDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromLocalDateTime_BeforeSdt(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	ldt := LocalDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromLocalDateTime_InOverlap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// There were two instances of 01:01
	// Setting (fold==0) selects the first instance, resolves to 01:01-07:00.
	ldt := LocalDateTime{2018, 11, 4, 1, 1, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 1, 1, 0, 0 /*Fold*/, -7 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// Setting (fold==1) selects the second instance, resolves to 01:01-08:00.
	ldt.Fold = 1
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 1, 1, 0, 1 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}
}

func TestNewZonedDateTimeFromLocalDateTime_AfterOverlap(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

	// 02:01 should resolve to 02:01-08:00
	ldt := LocalDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/}
	zdt := NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	if zdt.IsError() {
		t.Fatal(zdt)
	}
	if !(zdt == ZonedDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/, -8 * 60, &tz}) {
		t.Fatal(zdt)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
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

	tzLosAngeles := NewTimeZoneFromZoneInfo(
		&zonedbtesting.ZoneAmerica_Los_Angeles)
	tzNewYork := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_New_York)

	// 2022-08-30 20:00-07:00 in LA
	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0, 0 /*Fold*/}
	ladt := NewZonedDateTimeFromLocalDateTime(&ldt, &tzLosAngeles)
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

	tzLosAngeles := NewTimeZoneFromZoneInfo(
		&zonedbtesting.ZoneAmerica_Los_Angeles)
	tzPacific := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneUS_Pacific)

	if !zonedbtesting.ZoneUS_Pacific.IsLink() {
		t.Fatal("US/Pacific should be a Link")
	}

	ldt := LocalDateTime{2022, 8, 30, 20, 0, 0, 0 /*Fold*/}
	ladt := NewZonedDateTimeFromLocalDateTime(&ldt, &tzLosAngeles)
	padt := NewZonedDateTimeFromLocalDateTime(&ldt, &tzPacific)

	if !(ladt.ToEpochSeconds() == padt.ToEpochSeconds()) {
		t.Fatal("epochSeconds not equal")
	}
}

//-----------------------------------------------------------------------------
// UnixSeconds64
//-----------------------------------------------------------------------------

func TestZonedDateTimeFromUnixSeconds64(t *testing.T) {
	savedEpochYear := GetCurrentEpochYear()
	SetCurrentEpochYear(2050)
	defer SetCurrentEpochYear(savedEpochYear)

	// TODO: Change to NewTimeZoneUTC() when it becomes available
	tz := NewTimeZoneFromZoneInfo(
		&zonedbtesting.ZoneAmerica_Los_Angeles)
	zdt := NewZonedDateTimeFromUnixSeconds64(InvalidUnixSeconds64, &tz)
	if !zdt.IsError() {
		t.Fatal(zdt)
	}

	// Test FromUnixSeconds64().
	// Unix seconds from 'date +%s -d 2050-01-02T03:04:05-08:00'.
	unixSeconds64 := int64(2524734245)
	zdt = NewZonedDateTimeFromUnixSeconds64(unixSeconds64, &tz)
	ldt := LocalDateTime{2050, 1, 2, 3, 4, 5, 0 /*Fold*/}
	if !(ldt == zdt.ToLocalDateTime()) {
		t.Fatal(zdt)
	}

	// Test ToUnixSeconds64(). Use +1 day after the previous ldt.
	ldt.Day++
	zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	unixSeconds64 = zdt.ToUnixSeconds64()
	expected := int64(2524734245 + 24*60*60)
	if !(expected == unixSeconds64) {
		t.Fatal(unixSeconds64)
	}
}

//-----------------------------------------------------------------------------
// Benchmarks
// $ go test -run=NOMATCH -bench=.
//-----------------------------------------------------------------------------

var epochSeconds int32
var zdt ZonedDateTime
var ldt = LocalDateTime{2023, 1, 19, 22, 11, 0, 0 /*Fold*/}
var tz = NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)

func BenchmarkZonedDateTimeFromEpochSeconds_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = NewZonedDateTimeFromEpochSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromEpochSeconds_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.zoneProcessor.Reset()
		zdt = NewZonedDateTimeFromEpochSeconds(3423423, &tz)
	}
}

func BenchmarkZonedDateTimeFromLocalDateTime_Cache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	}
}

func BenchmarkZonedDateTimeFromLocalDateTime_NoCache(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tz.zoneProcessor.Reset()
		zdt = NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	}
}

func BenchmarkZonedDateTimeToEpochSeconds(b *testing.B) {
	zdt = NewZonedDateTimeFromEpochSeconds(3423423, &tz)
	for n := 0; n < b.N; n++ {
		epochSeconds = zdt.ToEpochSeconds()
	}
}