package acetime

import (
	"testing"
)

//-----------------------------------------------------------------------------
// MonthDay
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
		zonePolicy:        &ZonePolicyTestUS,
		format:            "P%T",
		offsetCode:        -32,
		deltaCode:         0 + 4,
		untilYear:         10000,
		untilMonth:        1,
		untilDay:          1,
		untilTimeCode:     0,
		untilTimeModifier: suffixW,
	},
}

var ZoneTestLosAngeles = ZoneInfo{
	name:      "America/Los_Angeles",
	zoneID:    0xb7f7e8f2,
	startYear: 2000,
	untilYear: 10000,
	eras:      ZoneEraTestLos_Angeles,
	target:    nil,
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
	tt = &candidates[1].transitionTime
	if !(*tt == DateTuple{2018, 11, 4, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[2].transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[3].transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal(tt)
	}
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

	createTransitionsFromNamedMatch(&ts, &match)
	if !(3 == ts.indexPrior) {
		t.Fatal(ts.indexPrior)
	}

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

//-----------------------------------------------------------------------------
// Step 3, 4
//-----------------------------------------------------------------------------

func TestFixTransitionTimesGenerateStartUntilTimes(t *testing.T) {
	// Create 3 matches for the AlmostLosAngeles test zone.
	var startYm = YearMonth{2018, 12}
	var untilYm = YearMonth{2020, 2}
	var matches [maxMatches]MatchingEra

	numMatches := findMatches(&ZoneAlmostLosAngeles, startYm, untilYm, matches[:])
	if !(3 == numMatches) {
		t.Fatal(numMatches)
	}

	// Create a custom template instantiation to use a different SIZE than the
	// pre-defined typedef in ExtendedZoneProcess::TransitionStorage.
	var storage TransitionStorage

	// Create 3 Transitions corresponding to the matches.
	// Implements ExtendedZoneProcessor::createTransitionsFromSimpleMatch().
	transition1 := storage.GetFreeAgent()
	createTransitionForYear(transition1, 0, nil, &matches[0])
	transition1.matchStatus = matchStatusExactMatch // synthetic example
	storage.AddFreeAgentToCandidatePool()

	transition2 := storage.GetFreeAgent()
	createTransitionForYear(transition2, 0, nil, &matches[1])
	transition2.matchStatus = matchStatusWithinMatch // synthetic example
	storage.AddFreeAgentToCandidatePool()

	transition3 := storage.GetFreeAgent()
	createTransitionForYear(transition3, 0, nil, &matches[2])
	transition3.matchStatus = matchStatusWithinMatch // synthetic example
	storage.AddFreeAgentToCandidatePool()

	// Move actives to Active pool.
	storage.AddActiveCandidatesToActivePool()
	transitions := storage.GetActives()
	if !(3 == len(transitions)) {
		t.Fatal(len(transitions))
	}
	if !(&transitions[0] == transition1) {
		t.Fatal(transitions[0])
	}
	if !(&transitions[1] == transition2) {
		t.Fatal(transitions[1])
	}
	if !(&transitions[2] == transition3) {
		t.Fatal(transitions[2])
	}

	// Chain the transitions.
	fixTransitionTimes(transitions)

	// Verify. The first Transition is extended to -infinity.
	tt := &transition1.transitionTime
	if !(*tt == DateTuple{2018, 12, 1, 0, suffixW}) {
		t.Fatal("tt:", tt)
	}
	tts := &transition1.transitionTimeS
	if !(*tts == DateTuple{2018, 12, 1, 0, suffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu := &transition1.transitionTimeU
	if !(*ttu == DateTuple{2018, 12, 1, 15 * 32, suffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Second transition uses the UTC offset of the first.
	tt = &transition2.transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, suffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition2.transitionTimeS
	if !(*tts == DateTuple{2019, 3, 10, 15 * 8, suffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition2.transitionTimeU
	if !(*ttu == DateTuple{2019, 3, 10, 15 * 40, suffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Third transition uses the UTC offset of the second.
	tt = &transition3.transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition3.transitionTimeS
	if !(*tts == DateTuple{2019, 11, 3, 15 * 4, suffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition3.transitionTimeU
	if !(*ttu == DateTuple{2019, 11, 3, 15 * 36, suffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Generate the startDateTime and untilDateTime of the transitions.
	generateStartUntilTimes(transitions)

	// Verify. The first transition startTime should be the same as its
	// transitionTime.
	sdt := &transition1.startDt
	if !(*sdt == DateTuple{2018, 12, 1, 0, suffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt := &transition1.untilDt
	if !(*udt == DateTuple{2019, 3, 10, 15 * 8, suffixW}) {
		t.Fatal("udt:", udt)
	}
	odt := OffsetDateTime{
		2018, 12, 1, 0, 0, 0, 0 /*fold*/, -8 * 60 /*offsetMinutes*/}
	eps := odt.ToEpochSeconds()
	if !(eps == transition1.startEpochSeconds) {
		t.Fatal(transition1.startEpochSeconds)
	}

	// Second transition startTime is shifted forward one hour into PDT.
	sdt = &transition2.startDt
	if !(*sdt == DateTuple{2019, 3, 10, 15 * 12, suffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition2.untilDt
	if !(*udt == DateTuple{2019, 11, 3, 15 * 8, suffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2019, 3, 10, 3, 0, 0, 0 /*fold*/, -7 * 60 /*offsetMinutes*/}
	eps = odt.ToEpochSeconds()
	if !(eps == transition2.startEpochSeconds) {
		t.Fatal(transition2.startEpochSeconds)
	}

	// Third transition startTime is shifted back one hour into PST.
	sdt = &transition3.startDt
	if !(*sdt == DateTuple{2019, 11, 3, 15 * 4, suffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition3.untilDt
	if !(*udt == DateTuple{2020, 2, 1, 0, suffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2019, 11, 3, 1, 0, 0, 0 /*fold*/, -8 * 60 /*offsetMinutes*/}
	eps = odt.ToEpochSeconds()
	if !(eps == transition3.startEpochSeconds) {
		t.Fatal(transition3.startEpochSeconds)
	}
}

//---------------------------------------------------------------------------
// A simplified version of America/Los_Angeles, using only simple ZoneEras
// (i.e. no references to a ZonePolicy). Valid only for 2018.
//---------------------------------------------------------------------------

// Create simplified ZoneEras which approximate America/Los_Angeles
var ZoneEraAlmostLosAngeles = []ZoneEra{
	{
		zonePolicy:        nil,
		format:            "PST",
		offsetCode:        -32,
		deltaCode:         0 + 4,
		untilYear:         2019,
		untilMonth:        3,
		untilDay:          10,
		untilTimeCode:     2 * 4,
		untilTimeModifier: suffixW,
	},
	{
		zonePolicy:        nil,
		format:            "PDT",
		offsetCode:        -32,
		deltaCode:         4 + 4,
		untilYear:         2019,
		untilMonth:        11,
		untilDay:          3,
		untilTimeCode:     2 * 4,
		untilTimeModifier: suffixW,
	},
	{
		zonePolicy:        nil,
		format:            "PST",
		offsetCode:        -32,
		deltaCode:         0 + 4,
		untilYear:         2020,
		untilMonth:        3,
		untilDay:          8,
		untilTimeCode:     2 * 4,
		untilTimeModifier: suffixW,
	},
}

var ZoneAlmostLosAngeles = ZoneInfo{
	name:      "America/Almost_Los_Angeles",
	zoneID:    0x70166020,
	startYear: 2000,
	untilYear: 10000,
	eras:      ZoneEraAlmostLosAngeles,
	target:    nil,
}

//-----------------------------------------------------------------------------
// Step 5
//-----------------------------------------------------------------------------

func TestCreateAbbreviation(t *testing.T) {
	// If no '%', deltaMinutes and letter should not matter
	abbrev := createAbbreviation("SAST", 0, "")
	if !("SAST" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("SAST", 60, "A")
	if !("SAST" == abbrev) {
		t.Fatal(abbrev)
	}

	// If '%', and letter is "", remove the "%" (unlike AceTimeC where letter is
	// NULL.
	abbrev = createAbbreviation("SA%ST", 0, "")
	if !("SAST" == abbrev) {
		t.Fatal(abbrev)
	}

	// If '%', then replaced with (non-null) letterString.
	abbrev = createAbbreviation("P%T", 60, "D")
	if !("PDT" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("P%T", 0, "S")
	if !("PST" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("P%T", 0, "")
	if !("PT" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("%", 60, "CAT")
	if !("CAT" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("%", 0, "WAT")
	if !("WAT" == abbrev) {
		t.Fatal(abbrev)
	}

	// If '/', then deltaMinutes selects the first or second component.
	abbrev = createAbbreviation("GMT/BST", 0, "")
	if !("GMT" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("GMT/BST", 60, "")
	if !("BST" == abbrev) {
		t.Fatal(abbrev)
	}
}
