package acetime

import (
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
	a := DateTuple{2000, 1, 1, 0, suffixW}
	b := DateTuple{2000, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &b) == 0) {
		t.Fatal("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

	bb := DateTuple{2000, 1, 1, 0, suffixS}
	if !(dateTupleCompare(&a, &bb) == 0) {
		t.Fatal("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

	c := DateTuple{2000, 1, 1, 1, suffixW}
	if !(dateTupleCompare(&a, &c) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

	d := DateTuple{2000, 1, 2, 0, suffixW}
	if !(dateTupleCompare(&a, &d) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

	e := DateTuple{2000, 2, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &e) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

	f := DateTuple{2001, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &f) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}

func TestDateTupleSubtract(t *testing.T) {
	var dta, dtb DateTuple
	var diff int32

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 1, 1, suffixW} // 2000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 2, 0, suffixW} // 2000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	dtb = DateTuple{2000, 3, 1, 0, suffixW} // 2000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 2000 is leap, 29 day
		t.Fatal(diff)
	}
}

func TestDateTupleNormalize(t *testing.T) {
	var dt DateTuple

	// 00:00
	dt = DateTuple{2000, 1, 1, 0, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 0, suffixW}) {
		t.Fatal(dt)
	}

	// 23:45
	dt = DateTuple{2000, 1, 1, 15 * 95, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 15 * 95, suffixW}) {
		t.Fatal(dt)
	}

	// 24:00
	dt = DateTuple{2000, 1, 1, 15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 0, suffixW}) {
		t.Fatal(dt)
	}

	// 24:15
	dt = DateTuple{2000, 1, 1, 15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 15, suffixW}) {
		t.Fatal(dt)
	}

	// -24:00
	dt = DateTuple{2000, 1, 1, -15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, 0, suffixW}) {
		t.Fatal(dt)
	}

	// -24:15
	dt = DateTuple{2000, 1, 1, -15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, -15, suffixW}) {
		t.Fatal(dt)
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
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatal(ttu)
	}

	tt = DateTuple{2000, 1, 30, 15 * 12, suffixS}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatal(ttu)
	}

	tt = DateTuple{2000, 1, 30, 15 * 4, suffixU}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatal(ttu)
	}
}

func TestDateTupleCompareFuzzy(t *testing.T) {
	status := dateTupleCompareFuzzy(
		&DateTuple{2000, 10, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusPrior == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2000, 11, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 2, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 3, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusFarFuture == status) {
		t.Fatal(status)
	}

	// Verify dates whose delta months is greater than 32767. In
	// other words, delta years is greater than 2730.
	status = dateTupleCompareFuzzy(
		&DateTuple{5000, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(matchStatusFarFuture == status) {
		t.Fatal(status)
	}
	status = dateTupleCompareFuzzy(
		&DateTuple{1000, 4, 1, 1, 0},
		&DateTuple{4000, 12, 1, 1, 0},
		&DateTuple{4002, 2, 1, 1, 0})
	if !(matchStatusPrior == status) {
		t.Fatal(status)
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

//-----------------------------------------------------------------------------
// Step 2A
//-----------------------------------------------------------------------------

var ZoneRulesTestUS = []ZoneRule{
	// Rule    US    1967    2006    -    Oct    lastSun    2:00    0    S
	{
		1967,    /*fromYear*/
		2006,    /*toYear*/
		10,      /*inMonth*/
		7,       /*onDayOfWeek*/
		0,       /*onDayOfMonth*/
		8,       /*atTimeCode*/
		suffixW, /*atTimeModifier*/
		0 + 4,   /*deltaCode*/
		"S",     /*letter*/
	},
	// Rule    US    1976    1986    -    Apr    lastSun    2:00    1:00    D
	{
		1976,    /*fromYear*/
		1986,    /*toYear*/
		4,       /*inMonth*/
		7,       /*onDayOfWeek*/
		0,       /*onDayOfMonth*/
		8,       /*atTimeCode*/
		suffixW, /*atTimeModifier*/
		4 + 4,   /*deltaCode*/
		"D",     /*letter*/
	},
	// Rule    US    1987    2006    -    Apr    Sun>=1    2:00    1:00    D
	{
		1987,    /*fromYear*/
		2006,    /*toYear*/
		4,       /*inMonth*/
		7,       /*onDayOfWeek*/
		1,       /*onDayOfMonth*/
		8,       /*atTimeCode*/
		suffixW, /*atTimeModifier*/
		4 + 4,   /*deltaCode*/
		"D",     /*letter*/
	},
	// Rule    US    2007    max    -    Mar    Sun>=8    2:00    1:00    D
	{
		2007,    /*fromYear*/
		9999,    /*toYear*/
		3,       /*inMonth*/
		7,       /*onDayOfWeek*/
		8,       /*onDayOfMonth*/
		8,       /*atTimeCode*/
		suffixW, /*atTimeModifier*/
		4 + 4,   /*deltaCode*/
		"D",     /*letter*/
	},
	// Rule    US    2007    max    -    Nov    Sun>=1    2:00    0    S
	{
		2007,    /*fromYear*/
		9999,    /*toYear*/
		11,      /*inMonth*/
		7,       /*onDayOfWeek*/
		1,       /*onDayOfMonth*/
		8,       /*atTimeCode*/
		suffixW, /*atTimeModifier*/
		0 + 4,   /*deltaCode*/
		"S",     /*letter*/
	},
}

var ZonePolicyTestUS = ZonePolicy{
	ZoneRulesTestUS, /*rules*/
	nil,             /* letters */
}

var ZoneEraTestLos_Angeles = []ZoneEra{
	//             -8:00    US    P%sT
	{
		&ZonePolicyTestUS, /*zonePolicy*/
		"P%T",             /*format*/
		-32,               /*offsetCode*/
		0 + 4,             /*deltaCode*/
		10000,             /*untilYear*/
		1,                 /*untilMonth*/
		1,                 /*untilDay*/
		0,                 /*untilTimeCode*/
		suffixW,           /*untilTimeModifier*/
	},
}

var ZoneTestLosAngeles = ZoneInfo{
	"America/Los_Angeles",  /*name*/
	0xb7f7e8f2,             /*zoneId*/
	2000,                   /*startYear*/
	10000,                  /*untilYear*/
	ZoneEraTestLos_Angeles, /*eras*/
	nil,                    /*targetInfo*/
}

func TestGetTransitionTime(t *testing.T) {
	// Nov Sun>=1
	rule := &ZoneRulesTestUS[4]

	// Nov 4 2018
	dt := getTransitionTime(2018, rule)
	if !(dt == DateTuple{2018, 11, 4, 15 * 8, suffixW}) {
		t.Fatal(dt)
	}

	// Nov 3 2019
	dt = getTransitionTime(2019, rule)
	if !(dt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal(dt)
	}
}

func TestCreateTransitionForYear(t *testing.T) {
	match := MatchingEra{
		startDt:           DateTuple{2018, 12, 1, 0, suffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, suffixW},
		era:               &ZoneEraTestLos_Angeles[0],
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}
	rule := &ZoneRulesTestUS[4]

	// Nov Sun>=1
	var transition Transition
	createTransitionForYear(&transition, 2019, rule, &match)
	if !(transition.offsetMinutes == -15*32) {
		t.Fatal(transition.offsetMinutes)
	}
	if !(transition.deltaMinutes == 0) {
		t.Fatal(transition.deltaMinutes)
	}
	tt := &transition.transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 2B Pass 1
//-----------------------------------------------------------------------------

func TestCalcInteriorYears(t *testing.T) {
	var interiorYears [maxInteriorYears]int16

	num := calcInteriorYears(interiorYears[:], 1998, 1999, 2000, 2002)
	if !(0 == num) {
		t.Fatal(num)
	}

	num = calcInteriorYears(interiorYears[:], 2003, 2005, 2000, 2002)
	if !(0 == num) {
		t.Fatal(num)
	}

	num = calcInteriorYears(interiorYears[:], 1998, 2000, 2000, 2002)
	if !(1 == num) {
		t.Fatal(num)
	}
	if !(2000 == interiorYears[0]) {
		t.Fatal(interiorYears[0])
	}

	num = calcInteriorYears(interiorYears[:], 2002, 2004, 2000, 2002)
	if !(1 == num) {
		t.Fatal(num)
	}
	if !(2002 == interiorYears[0]) {
		t.Fatal(interiorYears[0])
	}

	num = calcInteriorYears(interiorYears[:], 2001, 2002, 2000, 2002)
	if !(2 == num) {
		t.Fatal(num)
	}
	if !(2001 == interiorYears[0]) {
		t.Fatal(interiorYears[0])
	}
	if !(2002 == interiorYears[1]) {
		t.Fatal(interiorYears[1])
	}

	num = calcInteriorYears(interiorYears[:], 1999, 2003, 2000, 2002)
	if !(3 == num) {
		t.Fatal(num)
	}
	if !(2000 == interiorYears[0]) {
		t.Fatal(interiorYears[0])
	}
	if !(2001 == interiorYears[1]) {
		t.Fatal(interiorYears[1])
	}
	if !(2002 == interiorYears[2]) {
		t.Fatal(interiorYears[2])
	}
}

func TestGetMostRecentPriorYear(t *testing.T) {
	year := getMostRecentPriorYear(1998, 1999, 2000)
	if !(1999 == year) {
		t.Fatal(year)
	}

	year = getMostRecentPriorYear(2003, 2005, 2000)
	if !(InvalidYear == year) {
		t.Fatal(year)
	}

	year = getMostRecentPriorYear(1998, 2000, 2000)
	if !(1999 == year) {
		t.Fatal(year)
	}

	year = getMostRecentPriorYear(2002, 2004, 2000)
	if !(InvalidYear == year) {
		t.Fatal(year)
	}

	year = getMostRecentPriorYear(2001, 2002, 2000)
	if !(InvalidYear == year) {
		t.Fatal(year)
	}

	year = getMostRecentPriorYear(1999, 2003, 2000)
	if !(1999 == year) {
		t.Fatal(year)
	}
}

func TestFindCandidateTransitions(t *testing.T) {
	match := MatchingEra{
		startDt:           DateTuple{2018, 12, 1, 0, suffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, suffixW},
		era:               &ZoneEraTestLos_Angeles[0],
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}

	// Reserve storage for the Transitions
	var ts TransitionStorage
	ts.Init()

	// Verify compareTransitionToMatchFuzzy() elminates various transitions
	// to get down to 5:
	//    * 2018 Mar Sun>=8 (11)
	//    * 2019 Nov Sun>=1 (4)
	//    * 2019 Mar Sun>=8 (10)
	//    * 2019 Nov Sun>=1 (3)
	//    * 2020 Mar Sun>=8 (8)
	ts.ResetCandidatePool()
	findCandidateTransitions(&ts, &match)
	candidates := ts.GetCandidates()
	if !(5 == len(candidates)) {
		t.Fatal()
	}

	tt := &candidates[0].transitionTime
	if !(*tt == DateTuple{2018, 3, 11, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	//
	tt = &candidates[1].transitionTime
	if !(*tt == DateTuple{2018, 11, 4, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	//
	tt = &candidates[2].transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	//
	tt = &candidates[3].transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	//
	tt = &candidates[4].transitionTime
	if !(*tt == DateTuple{2020, 3, 8, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 2B Pass 3
//-----------------------------------------------------------------------------

func TestProcessTransitionMatchStatus(t *testing.T) {
	// UNTIL = 2002-01-02T03:00
	era := ZoneEra{
		zonePolicy:        nil,
		format:            "",
		offsetCode:        0,
		deltaCode:         0,
		untilYear:         2002,
		untilMonth:        1,
		untilDay:          2,
		untilTimeCode:     12,
		untilTimeModifier: suffixW,
	}

	// [2000-01-01, 2001-01-01)
	match := MatchingEra{
		startDt:           DateTuple{2000, 1, 1, 0, suffixW},
		untilDt:           DateTuple{2001, 1, 1, 0, suffixW},
		era:               &era,
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}

	// This transition occurs before the match, so prior should be filled.
	// transitionTime = 1999-12-31
	transitions := []Transition{
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{1999, 12, 31, 0, suffixW},
		},
		// This occurs at exactly match.startDateTime, so should replace the prior.
		// transitionTime = 2000-01-01
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 1, 0, suffixW},
		},
		// An interior transition. Prior should not change.
		// transitionTime = 2000-01-02
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 2, 0, suffixW},
		},
		// Occurs after match.untilDateTime, so should be rejected.
		// transitionTime = 2001-01-02
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2001, 1, 2, 0, suffixW},
		},
	}
	transition0 := &transitions[0]
	transition1 := &transitions[1]
	transition2 := &transitions[2]
	transition3 := &transitions[3]

	// Populate the transitionTimeS and transitionTimeU fields.
	var prior *Transition = nil
	fixTransitionTimes(transitions[:])

	prior = processTransitionMatchStatus(transition0, prior)
	if !(matchStatusPrior == transition0.matchStatus) {
		t.Fatal(transition0.matchStatus)
	}
	if !(prior == transition0) {
		t.Fatal(transition0)
	}

	prior = processTransitionMatchStatus(transition1, prior)
	if !(matchStatusExactMatch == transition1.matchStatus) {
		t.Fatal(transition1.matchStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}

	prior = processTransitionMatchStatus(transition2, prior)
	if !(matchStatusWithinMatch == transition2.matchStatus) {
		t.Fatal(transition2.matchStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}

	prior = processTransitionMatchStatus(transition3, prior)
	if !(matchStatusFarFuture == transition3.matchStatus) {
		t.Fatal(transition3.matchStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}
}

//-----------------------------------------------------------------------------
// Step 2B
//-----------------------------------------------------------------------------

func TestCreateTransitionsFromNamedMatch(t *testing.T) {
	match := MatchingEra{
		startDt:           DateTuple{2018, 12, 1, 0, suffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, suffixW},
		era:               &ZoneEraTestLos_Angeles[0],
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}

	// Reserve storage for the Transitions
	var ts TransitionStorage
	ts.Init()

	createTransitionsFromNamedMatch(&ts, &match)
	if !(3 == ts.indexPrior) {
		t.Fatal(ts.indexPrior)
	}

	//
	tt := &ts.transitions[0].transitionTime
	if !(*tt == DateTuple{2018, 12, 1, 0, suffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[1].transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[2].transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
}
