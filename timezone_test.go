package acetime

import (
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
	a := DateTuple{2000, 1, 1, 0, suffixW}
	b := DateTuple{2000, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &b) == 0) {
		t.Fatalf("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

	bb := DateTuple{2000, 1, 1, 0, suffixS}
	if !(dateTupleCompare(&a, &bb) == 0) {
		t.Fatalf("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

	c := DateTuple{2000, 1, 1, 1, suffixW}
	if !(dateTupleCompare(&a, &c) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

	d := DateTuple{2000, 1, 2, 0, suffixW}
	if !(dateTupleCompare(&a, &d) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

	e := DateTuple{2000, 2, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &e) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

	f := DateTuple{2001, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &f) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}

func TestDateTupleSubtract(t *testing.T) {
	var dta, dtb DateTuple
	var diff int32

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 1, 1, suffixW} // 2000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 2, 0, suffixW} // 2000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	dtb = DateTuple{2000, 3, 1, 0, suffixW} // 2000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 2000 is leap, 29 day
		t.Fatalf("fatal")
	}
}

func TestDateTupleNormalize(t *testing.T) {
	var dt DateTuple

	// 00:00
	dt = DateTuple{2000, 1, 1, 0, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// 23:45
	dt = DateTuple{2000, 1, 1, 15 * 95, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 15 * 95, suffixW}) {
		t.Fatalf("fatal")
	}

	// 24:00
	dt = DateTuple{2000, 1, 1, 15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// 24:15
	dt = DateTuple{2000, 1, 1, 15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 15, suffixW}) {
		t.Fatalf("fatal")
	}

	// -24:00
	dt = DateTuple{2000, 1, 1, -15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// -24:15
	dt = DateTuple{2000, 1, 1, -15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, -15, suffixW}) {
		t.Fatalf("fatal")
	}
}

func TestDateTupleExpand(t *testing.T) {
	var tt DateTuple
	var ttw DateTuple
	var tts DateTuple
	var ttu DateTuple

	const offsetMinutes = 2 * 60
	const deltaMinutes = 1 * 60

	tt = DateTuple{2000, 1, 30, 15 * 16, suffixW} // 04:00
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}

	tt = DateTuple{2000, 1, 30, 15 * 12, suffixS}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}

	tt = DateTuple{2000, 1, 30, 15 * 4, suffixU}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}
}

func TestDateTupleCompareFuzzy(t *testing.T) {
	if !(matchStatusPrior == dateTupleCompareFuzzy(
		&DateTuple{2000, 10, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2000, 11, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2002, 2, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2002, 3, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusFarFuture == dateTupleCompareFuzzy(
		&DateTuple{2002, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	// Verify dates whose delta months is greater than 32767. In
	// other words, delta years is greater than 2730.
	if !(matchStatusFarFuture == dateTupleCompareFuzzy(
		&DateTuple{5000, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}
	if !(matchStatusPrior == dateTupleCompareFuzzy(
		&DateTuple{1000, 4, 1, 1, 0},
		&DateTuple{4000, 12, 1, 1, 0},
		&DateTuple{4002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}
}

//-----------------------------------------------------------------------------

func TestCompareTransitionToMatchFuzzy(t *testing.T) {
	match := MatchingEra{
		startDt: DateTuple{2000, 1, 1, 0, suffixW},
		untilDt: DateTuple{2001, 1, 1, 0, suffixW},
	}

	transition := Transition{
		match:          &match,
		rule:           nil,
		transitionTime: DateTuple{1999, 11, 1, 0, suffixW},
	}
	status := compareTransitionToMatchFuzzy(&transition, &match)
	if !(status == matchStatusPrior) {
		t.Fatal("fatal")
	}

	transition = Transition{
		match:          &match,
		rule:           nil,
		transitionTime: DateTuple{1999, 12, 1, 0, suffixW},
	}
	status = compareTransitionToMatchFuzzy(&transition, &match)
	if !(status == matchStatusWithinMatch) {
		t.Fatal("fatal")
	}

	transition = Transition{
		match:          &match,
		rule:           nil,
		transitionTime: DateTuple{2000, 1, 1, 0, suffixW},
	}
	status = compareTransitionToMatchFuzzy(&transition, &match)
	if !(status == matchStatusWithinMatch) {
		t.Fatal("fatal")
	}

	transition = Transition{
		match:          &match,
		rule:           nil,
		transitionTime: DateTuple{2001, 1, 1, 0, suffixW},
	}
	status = compareTransitionToMatchFuzzy(&transition, &match)
	if !(status == matchStatusWithinMatch) {
		t.Fatal("fatal")
	}

	transition = Transition{
		match:          &match,
		rule:           nil,
		transitionTime: DateTuple{2001, 3, 1, 0, suffixW},
	}
	status = compareTransitionToMatchFuzzy(&transition, &match)
	if !(status == matchStatusFarFuture) {
		t.Fatal("fatal")
	}
}

func TestCompareTransitionToMatch(t *testing.T) {
	// UNTIL = 2002-01-02T03:00
	era := ZoneEra{
		zonePolicy:        nil,
		format:            "",
		offsetCode:        0,
		deltaCode:         0,
		untilYear:         2,
		untilMonth:        1,
		untilDay:          2,
		untilTimeCode:     12,
		untilTimeModifier: suffixW,
	}

	// MatchingEra=[2000-01-01, 2001-01-01)
	match := MatchingEra{
		startDt:           DateTuple{2000, 1, 1, 0, suffixW},
		untilDt:           DateTuple{2001, 1, 1, 0, suffixW},
		era:               &era,
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}

	transitions := []Transition{
		// transitionTime = 1999-12-31
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{1999, 12, 31, 0, suffixW},
		},
		// transitionTime = 2000-01-01
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 1, 0, suffixW},
		},
		// transitionTime = 2000-01-02
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 2, 0, suffixW},
		},
		// transitionTime = 2001-02-03
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2001, 2, 3, 0, suffixW},
		},
	}
	transition0 := &transitions[0]
	transition1 := &transitions[1]
	transition2 := &transitions[2]
	transition3 := &transitions[3]

	// Populate the transitionTimeS and transitionTimeU fields.
	fixTransitionTimes(transitions)

	status := compareTransitionToMatch(transition0, &match)
	if !(status == matchStatusPrior) {
		t.Fatal("tt:", transition0.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition1, &match)
	if !(status == matchStatusExactMatch) {
		t.Fatal("tt:", transition1.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition2, &match)
	if !(status == matchStatusWithinMatch) {
		t.Fatal("tt:", transition2.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition3, &match)
	if !(status == matchStatusFarFuture) {
		t.Fatal("tt:", transition3.transitionTime, "; status:", status)
	}
}

//-----------------------------------------------------------------------------

func TestCalcStartDayOfMonth(t *testing.T) {
	// 2018-11, Sun>=1
	monthDay := calcStartDayOfMonth(2018, 11, IsoWeekdaySunday, 1)
	if !(monthDay == MonthDay{11, 4}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-11, lastSun
	monthDay = calcStartDayOfMonth(2018, 11, IsoWeekdaySunday, 0)
	if !(monthDay == MonthDay{11, 25}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-11, Sun>=30, should shift to 2018-12-2
	monthDay = calcStartDayOfMonth(2018, 11, IsoWeekdaySunday, 30)
	if !(monthDay == MonthDay{12, 2}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-11, Mon<=7
	monthDay = calcStartDayOfMonth(2018, 11, IsoWeekdayMonday, -7)
	if !(monthDay == MonthDay{11, 5}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-11, Mon<=1, shifts back into October
	monthDay = calcStartDayOfMonth(2018, 11, IsoWeekdayMonday, -1)
	if !(monthDay == MonthDay{10, 29}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-03, Thu>=9
	monthDay = calcStartDayOfMonth(2018, 3, IsoWeekdayThursday, 9)
	if !(monthDay == MonthDay{3, 15}) {
		t.Fatal("monthDay:", monthDay)
	}

	// 2018-03-30 exactly
	monthDay = calcStartDayOfMonth(2018, 3, 0, 30)
	if !(monthDay == MonthDay{3, 30}) {
		t.Fatal("monthDay:", monthDay)
	}
}

//-----------------------------------------------------------------------------
// Step 1
//-----------------------------------------------------------------------------

func TestCompareEraToYearMonth(t *testing.T) {
	era := ZoneEra{
		untilYear:         2000,
		untilMonth:        1,
		untilDay:          2,
		untilTimeCode:     12,
		untilTimeModifier: suffixW,
	}

	if !(1 == compareEraToYearMonth(&era, 2000, 1)) {
		t.Fatal("fatal")
	}
	if !(1 == compareEraToYearMonth(&era, 2000, 1)) {
		t.Fatal("fatal")
	}
	if !(-1 == compareEraToYearMonth(&era, 2000, 2)) {
		t.Fatal("fatal")
	}
	if !(-1 == compareEraToYearMonth(&era, 2000, 3)) {
		t.Fatal("fatal")
	}

	era2 := ZoneEra{
		untilYear:         2000,
		untilMonth:        1,
		untilDay:          0,
		untilTimeCode:     0,
		untilTimeModifier: suffixW,
	}
	if !(0 == compareEraToYearMonth(&era2, 2000, 1)) {
		t.Fatal("fatal")
	}
}

func TestCreateMatchingEra(t *testing.T) {
	// 14-month interval, from 2000-12 until 2002-02
	startYm := YearMonth{2000, 12}
	untilYm := YearMonth{2002, 2}

	// UNTIL = 2000-12-02 3:00
	era1 := ZoneEra{
		untilYear:         2000,
		untilMonth:        12,
		untilDay:          2,
		untilTimeCode:     3 * (60 / 15),
		untilTimeModifier: suffixW,
	}

	// UNTIL = 2001-02-03 4:00
	era2 := ZoneEra{
		untilYear:         2001,
		untilMonth:        2,
		untilDay:          3,
		untilTimeCode:     4 * (60 / 15),
		untilTimeModifier: suffixW,
	}

	// UNTIL = 2002-10-11 4:00
	era3 := ZoneEra{
		untilYear:         2002,
		untilMonth:        10,
		untilDay:          11,
		untilTimeCode:     4 * (60 / 15),
		untilTimeModifier: suffixW,
	}

	// No previous matching era, so startDt is set to startYm.
	var match1 MatchingEra
	createMatchingEra(&match1, nil, &era1, startYm, untilYm)
	if !(match1.startDt == DateTuple{2000, 12, 1, 60 * 0, suffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.untilDt == DateTuple{2000, 12, 2, 60 * 3, suffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.era == &era1) {
		t.Fatal("match1.startDt:", match1.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is < untilYm, so is retained.
	var match2 MatchingEra
	createMatchingEra(&match2, &match1, &era2, startYm, untilYm)
	if !(match2.startDt == DateTuple{2000, 12, 2, 60 * 3, suffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.untilDt == DateTuple{2001, 2, 3, 60 * 4, suffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.era == &era2) {
		t.Fatal("match2.startDt:", match2.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is > untilYm so truncated to untilYm.
	var match3 MatchingEra
	createMatchingEra(&match3, &match2, &era3, startYm, untilYm)
	if !(match3.startDt == DateTuple{2001, 2, 3, 60 * 4, suffixW}) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
	if !(match3.untilDt == DateTuple{2002, 2, 1, 60 * 0, suffixW}) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
	if !(match3.era == &era3) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
}
