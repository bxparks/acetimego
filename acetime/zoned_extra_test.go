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
// NewZonedExtraFromEpochSeconds(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestNewZonedExtraFromEpochSeconds(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	var epochSeconds Time = 946684800
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromEpochSeconds_2050(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	var epochSeconds Time = 2524608000
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromEpochSeconds_UnixMax(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = (1 << 31) - 1
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, 0, 0, 0, 0, "UTC"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromEpochSeconds_Invalid(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("Etc/UTC")
	var epochSeconds Time = InvalidEpochSeconds
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if !extra.IsError() {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromEpochSeconds_FallBack(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-07:00, which is 31 minutes before the DST
	// fall-back.
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 11, 6, 1, 29, 0, 0 /*Fold*/},
		OffsetSeconds: -7 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
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
	extra = NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
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
	extra = NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromEpochSeconds_SpringForward(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// Start our sampling at 01:29:00-08:00, which is 31 minutes before the DST
	// spring forward.
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2022, 3, 13, 1, 29, 0, 0 /*Fold*/},
		OffsetSeconds: -8 * 3600,
	}
	epochSeconds := odt.EpochSeconds()
	extra := NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// An hour later, we spring forward to 03:29:00-07:00.
	epochSeconds += 3600
	extra = NewZonedExtraFromEpochSeconds(epochSeconds, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

//-----------------------------------------------------------------------------
// NewZonedExtraFromLocalDateTime(). Same structure as zoned_date_time_test.go.
//-----------------------------------------------------------------------------

func TestNewZonedExtraFromLocalDateTime(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, there is only one match
	ldt = LocalDateTime{2000, 1, 1, 0, 0, 0, 1 /*Fold*/}
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_2050(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	ldt := LocalDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0, 1 /*Fold*/}
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_BeforeDst(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 01:59 should resolve to 01:59-08:00
	ldt := LocalDateTime{2018, 3, 11, 1, 59, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt = LocalDateTime{2018, 3, 11, 1, 59, 0, 1 /*Fold*/}
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_InGap(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 doesn't exist.
	// Setting (fold=0) causes the first transition to be selected, which has a
	// UTC offset of -08:00, so this is interpreted as 02:01-08:00 which gets
	// normalized to 03:01-07:00, which falls in the 2nd transition.
	ldt := LocalDateTime{2018, 3, 11, 2, 1, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	// fold == 0 to indicate only one match
	expected := ZonedExtra{FoldTypeGap, -8 * 3600, 3600, -8 * 3600, 0, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Setting (fold=1) causes the second transition to be selected, which has a
	// UTC offset of -07:00, so this is interpreted as 02:01-07:00 which gets
	// normalized to 01:01-08:00, which falls in the 1st transition.
	ldt = LocalDateTime{2018, 3, 11, 2, 1, 0, 1 /*Fold*/}
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	// fold == 0 to indicate the 1st transition
	expected = ZonedExtra{FoldTypeGap, -8 * 3600, 0, -8 * 3600, 3600, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_InDst(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 03:01 should resolve to 03:01-07:00.
	ldt := LocalDateTime{2018, 3, 11, 3, 1, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_BeforeSdt(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 00:59 is an hour before the DST->STD transition, so should return
	// 00:59-07:00.
	ldt := LocalDateTime{2018, 11, 4, 0, 59, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 3600, -8 * 3600, 3600, "PDT"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_InOverlap(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// There were two instances of 01:01
	// Setting (fold==0) selects the first instance, resolves to 01:01-07:00.
	ldt := LocalDateTime{2018, 11, 4, 1, 1, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{
		FoldTypeOverlap, -8 * 3600, 3600, -8 * 3600, 3600, "PDT",
	}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// Setting (fold==1) selects the second instance, resolves to 01:01-08:00.
	ldt.Fold = 1
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeOverlap, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}

func TestNewZonedExtraFromLocalDateTime_AfterOverlap(t *testing.T) {
	zm := NewZoneManager(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")

	// 02:01 should resolve to 02:01-08:00
	ldt := LocalDateTime{2018, 11, 4, 2, 1, 0, 0 /*Fold*/}
	extra := NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected := ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}

	// check that fold=1 gives identical results, since there is one match
	ldt.Fold = 1
	extra = NewZonedExtraFromLocalDateTime(&ldt, &tz)
	if extra.IsError() {
		t.Fatal(extra)
	}
	expected = ZonedExtra{FoldTypeExact, -8 * 3600, 0, -8 * 3600, 0, "PST"}
	if !(extra == expected) {
		t.Fatal(extra)
	}
}
