package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

// Most of these tests were migrated from ExtendedZoneProcessorTest in the
// AceTime library.

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

func TestZoneProcessorToString(t *testing.T) {
	var zp ZoneProcessor
	zp.InitForZoneInfo(
		&zonedbtesting.Context, &zonedbtesting.ZoneAmerica_Los_Angeles)
	if !(zp.String() == "America/Los_Angeles") {
		t.Fatal(zp.String(), zp)
	}
}

func TestZoneProcessorInitForYear(t *testing.T) {
	var zp ZoneProcessor
	zp.InitForZoneInfo(
		&zonedbtesting.Context, &zonedbtesting.ZoneAmerica_Los_Angeles)
	if zp.isFilled {
		t.Fatal(zp)
	}
	zp.InitForYear(2023)
	if !zp.isFilled {
		t.Fatal(zp)
	}
}

//-----------------------------------------------------------------------------
// Step 1
//-----------------------------------------------------------------------------

func TestCompareEraToYearMonth(t *testing.T) {
	era := zoneinfo.ZoneEra{
		UntilYear:         2000,
		UntilMonth:        1,
		UntilDay:          2,
		UntilTimeCode:     12,
		UntilTimeModifier: zoneinfo.SuffixW,
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

	era2 := zoneinfo.ZoneEra{
		UntilYear:         2000,
		UntilMonth:        1,
		UntilDay:          0,
		UntilTimeCode:     0,
		UntilTimeModifier: zoneinfo.SuffixW,
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
	era1 := zoneinfo.ZoneEra{
		UntilYear:         2000,
		UntilMonth:        12,
		UntilDay:          2,
		UntilTimeCode:     3 * (60 / 15),
		UntilTimeModifier: zoneinfo.SuffixW,
	}

	// UNTIL = 2001-02-03 4:00
	era2 := zoneinfo.ZoneEra{
		UntilYear:         2001,
		UntilMonth:        2,
		UntilDay:          3,
		UntilTimeCode:     4 * (60 / 15),
		UntilTimeModifier: zoneinfo.SuffixW,
	}

	// UNTIL = 2002-10-11 4:00
	era3 := zoneinfo.ZoneEra{
		UntilYear:         2002,
		UntilMonth:        10,
		UntilDay:          11,
		UntilTimeCode:     4 * (60 / 15),
		UntilTimeModifier: zoneinfo.SuffixW,
	}

	// Fake format offsets, with only a single empty string, and terminating "~"
	formatsOffset := []uint16{0, 0}
	formatsBuffer := "~"

	// No previous matching era, so startDt is set to startYm.
	var match1 MatchingEra
	createMatchingEra(formatsOffset, formatsBuffer,
		&match1, nil, &era1, startYm, untilYm)
	if !(match1.startDt == DateTuple{2000, 12, 1, 60 * 0, zoneinfo.SuffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.untilDt == DateTuple{2000, 12, 2, 60 * 3, zoneinfo.SuffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.era == &era1) {
		t.Fatal("match1.startDt:", match1.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is < untilYm, so is retained.
	var match2 MatchingEra
	createMatchingEra(formatsOffset, formatsBuffer,
		&match2, &match1, &era2, startYm, untilYm)
	if !(match2.startDt == DateTuple{2000, 12, 2, 60 * 3, zoneinfo.SuffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.untilDt == DateTuple{2001, 2, 3, 60 * 4, zoneinfo.SuffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.era == &era2) {
		t.Fatal("match2.startDt:", match2.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is > untilYm so truncated to untilYm.
	var match3 MatchingEra
	createMatchingEra(formatsOffset, formatsBuffer,
		&match3, &match2, &era3, startYm, untilYm)
	if !(match3.startDt == DateTuple{2001, 2, 3, 60 * 4, zoneinfo.SuffixW}) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
	if !(match3.untilDt == DateTuple{2002, 2, 1, 60 * 0, zoneinfo.SuffixW}) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
	if !(match3.era == &era3) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
}

//-----------------------------------------------------------------------------
// Step 2A
//-----------------------------------------------------------------------------

func TestGetTransitionTime(t *testing.T) {
	// Rule 5, [2007,9999]
	// Rule    US    2007    max    -    Nov    Sun>=1    2:00    0    S
	rule := &zonedbtesting.ZoneRulesUS[5]

	// Nov 4 2018
	dt := getTransitionTime(2018, rule)
	if !(dt == DateTuple{2018, 11, 4, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// Nov 3 2019
	dt = getTransitionTime(2019, rule)
	if !(dt == DateTuple{2019, 11, 3, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}
}

func TestCreateTransitionForYear(t *testing.T) {
	match := MatchingEra{
		startDt:           DateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               &zonedbtesting.ZoneEraAmerica_Los_Angeles[0],
		prevMatch:         nil,
		lastOffsetMinutes: 0,
		lastDeltaMinutes:  0,
	}
	rule := &zonedbtesting.ZoneRulesUS[5]

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
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, zoneinfo.SuffixW}) {
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
		startDt:           DateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               &zonedbtesting.ZoneEraAmerica_Los_Angeles[0],
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
	if !(*tt == DateTuple{2018, 3, 11, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[1].transitionTime
	if !(*tt == DateTuple{2018, 11, 4, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[2].transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[3].transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[4].transitionTime
	if !(*tt == DateTuple{2020, 3, 8, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 2B Pass 3
//-----------------------------------------------------------------------------

func TestProcessTransitionMatchStatus(t *testing.T) {
	// UNTIL = 2002-01-02T03:00
	era := zoneinfo.ZoneEra{
		ZonePolicy:        nil,
		FormatIndex:       0,
		OffsetCode:        0,
		DeltaCode:         0,
		UntilYear:         2002,
		UntilMonth:        1,
		UntilDay:          2,
		UntilTimeCode:     12,
		UntilTimeModifier: zoneinfo.SuffixW,
	}

	// [2000-01-01, 2001-01-01)
	match := MatchingEra{
		startDt:           DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		untilDt:           DateTuple{2001, 1, 1, 0, zoneinfo.SuffixW},
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
			transitionTime: DateTuple{1999, 12, 31, 0, zoneinfo.SuffixW},
		},
		// This occurs at exactly match.startDateTime, so should replace the prior.
		// transitionTime = 2000-01-01
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		},
		// An interior transition. Prior should not change.
		// transitionTime = 2000-01-02
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2000, 1, 2, 0, zoneinfo.SuffixW},
		},
		// Occurs after match.untilDateTime, so should be rejected.
		// transitionTime = 2001-01-02
		Transition{
			match:          &match,
			rule:           nil,
			transitionTime: DateTuple{2001, 1, 2, 0, zoneinfo.SuffixW},
		},
	}
	transition0 := &transitions[0]
	transition1 := &transitions[1]
	transition2 := &transitions[2]
	transition3 := &transitions[3]

	// Populate the transitionTimeS and transitionTimeU fields.
	var prior *Transition = nil
	fixTransitionTimes(transitions)

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
		startDt:           DateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           DateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               &zonedbtesting.ZoneEraAmerica_Los_Angeles[0],
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
	if !(*tt == DateTuple{2018, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[1].transitionTime
	if !(*tt == DateTuple{2019, 3, 10, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[2].transitionTime
	if !(*tt == DateTuple{2019, 11, 3, 15 * 8, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 3, Step 4. Use America/Los_Angeles to calculate the transitions
// beause I am familiar with it.
//-----------------------------------------------------------------------------

func TestFixTransitionTimesGenerateStartUntilTimes(t *testing.T) {
	// Step 1: America/Los_Angeles matches one era, which points to US policy.
	var startYm = YearMonth{2017, 12}
	var untilYm = YearMonth{2019, 2}
	var matches [maxMatches]MatchingEra

	numMatches := findMatches(
		zonedbtesting.Context.FormatOffsets,
		zonedbtesting.Context.FormatBuffer,
		&zonedbtesting.ZoneAmerica_Los_Angeles,
		startYm, untilYm, matches[:])
	if !(1 == numMatches) {
		t.Fatal(numMatches)
	}

	// Step 2: Create transitions.
	// Create a custom template instantiation to use a different SIZE than the
	// pre-defined typedef in ExtendedZoneProcess::TransitionStorage.
	var storage TransitionStorage
	createTransitions(&storage, matches[:numMatches])
	transitions := storage.GetActives()
	if !(len(transitions) == 3) {
		t.Fatal(len(transitions))
	}

	// Step 3: Chain the transitions by fixing the transition times.
	fixTransitionTimes(transitions)

	// Step 3: Verification. The first Transition is extended to -infinity.
	transition0 := &transitions[0]
	tt := &transition0.transitionTime
	if !(*tt == DateTuple{2017, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts := &transition0.transitionTimeS
	if !(*tts == DateTuple{2017, 12, 1, 0, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu := &transition0.transitionTimeU
	if !(*ttu == DateTuple{2017, 12, 1, 8 * 60, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 3: Verification: Second transition springs forward at 2018-03-11
	// 02:00.
	transition1 := &transitions[1]
	tt = &transition1.transitionTime
	if !(*tt == DateTuple{2018, 3, 11, 2 * 60, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition1.transitionTimeS
	if !(*tts == DateTuple{2018, 3, 11, 2 * 60, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition1.transitionTimeU
	if !(*ttu == DateTuple{2018, 3, 11, 10 * 60, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 3: Verification: Third transition falls back at 2018-11-04 02:00.
	transition2 := &transitions[2]
	tt = &transition2.transitionTime
	if !(*tt == DateTuple{2018, 11, 4, 2 * 60, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition2.transitionTimeS
	if !(*tts == DateTuple{2018, 11, 4, 1 * 60, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition2.transitionTimeU
	if !(*ttu == DateTuple{2018, 11, 4, 9 * 60, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 4: Generate the startDateTime and untilDateTime of the transitions.
	generateStartUntilTimes(transitions)

	// Step 4: Verification: The first transition startTime should be the same as
	// its transitionTime.
	sdt := &transition0.startDt
	if !(*sdt == DateTuple{2017, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt := &transition0.untilDt
	if !(*udt == DateTuple{2018, 3, 11, 2 * 60, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt := OffsetDateTime{
		2017, 12, 1, 0, 0, 0, 0 /*Fold*/, -8 * 60 /*offsetMinutes*/}
	eps := odt.ToEpochSeconds()
	if !(eps == transition0.startEpochSeconds) {
		t.Fatal(transition0.startEpochSeconds)
	}

	// Step 4: Verification: Second transition startTime is shifted forward one
	// hour into PDT.
	sdt = &transition1.startDt
	if !(*sdt == DateTuple{2018, 3, 11, 3 * 60, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition1.untilDt
	if !(*udt == DateTuple{2018, 11, 4, 2 * 60, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2018, 3, 11, 3, 0, 0, 0 /*Fold*/, -7 * 60 /*offsetMinutes*/}
	eps = odt.ToEpochSeconds()
	if !(eps == transition1.startEpochSeconds) {
		t.Fatal(transition1.startEpochSeconds)
	}

	// Step 4: Verification: Third transition startTime is shifted back one hour
	// into PST.
	sdt = &transition2.startDt
	if !(*sdt == DateTuple{2018, 11, 4, 1 * 60, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition2.untilDt
	if !(*udt == DateTuple{2019, 2, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2018, 11, 4, 1, 0, 0, 0 /*Fold*/, -8 * 60 /*offsetMinutes*/}
	eps = odt.ToEpochSeconds()
	if !(eps == transition2.startEpochSeconds) {
		t.Fatal(transition2.startEpochSeconds)
	}
}

//-----------------------------------------------------------------------------
// Step 5
//-----------------------------------------------------------------------------

func TestCreateAbbreviation(t *testing.T) {
	// If no '%', deltaMinutes and Letter should not matter
	abbrev := createAbbreviation("SAST", 0, "")
	if !("SAST" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("SAST", 60, "A")
	if !("SAST" == abbrev) {
		t.Fatal(abbrev)
	}

	// If '%', and Letter is "", remove the "%" (unlike AceTimeC where Letter is
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
