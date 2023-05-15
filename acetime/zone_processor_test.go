package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

// Most of these tests were migrated from ExtendedZoneProcessorTest in the
// AceTime library.

//-----------------------------------------------------------------------------
// monthDay
//-----------------------------------------------------------------------------

func TestCalcStartDayOfMonth(t *testing.T) {
	// 2018-11, Sun>=1
	md := calcStartDayOfMonth(2018, 11, uint8(Sunday), 1)
	if !(md == monthDay{11, 4}) {
		t.Fatal("md:", md)
	}

	// 2018-11, lastSun
	md = calcStartDayOfMonth(2018, 11, uint8(Sunday), 0)
	if !(md == monthDay{11, 25}) {
		t.Fatal("md:", md)
	}

	// 2018-11, Sun>=30, should shift to 2018-12-2
	md = calcStartDayOfMonth(2018, 11, uint8(Sunday), 30)
	if !(md == monthDay{12, 2}) {
		t.Fatal("md:", md)
	}

	// 2018-11, Mon<=7
	md = calcStartDayOfMonth(2018, 11, uint8(Monday), -7)
	if !(md == monthDay{11, 5}) {
		t.Fatal("md:", md)
	}

	// 2018-11, Mon<=1, shifts back into October
	md = calcStartDayOfMonth(2018, 11, uint8(Monday), -1)
	if !(md == monthDay{10, 29}) {
		t.Fatal("md:", md)
	}

	// 2018-03, Thu>=9
	md = calcStartDayOfMonth(2018, 3, uint8(Thursday), 9)
	if !(md == monthDay{3, 15}) {
		t.Fatal("md:", md)
	}

	// 2018-03-30 exactly
	md = calcStartDayOfMonth(2018, 3, 0, 30)
	if !(md == monthDay{3, 30}) {
		t.Fatal("md:", md)
	}
}

//-----------------------------------------------------------------------------

func TestZoneProcessorToString(t *testing.T) {
	zoneManager := NewZoneManager(&zonedbtesting.DataContext)
	zoneInfo := zoneManager.store.ZoneInfoByID(
		zonedbtesting.ZoneIDAmerica_Los_Angeles)
	var zp zoneProcessor
	zp.initForZoneInfo(zoneInfo)
	if !(zp.name() == "America/Los_Angeles") {
		t.Fatal(zp.name(), zp)
	}
}

func TestZoneProcessorInitForYear(t *testing.T) {
	zoneManager := NewZoneManager(&zonedbtesting.DataContext)
	zoneInfo := zoneManager.store.ZoneInfoByID(
		zonedbtesting.ZoneIDAmerica_Los_Angeles)
	var zp zoneProcessor
	zp.initForZoneInfo(zoneInfo)
	if zp.year != InvalidYear {
		t.Fatal(zp)
	}
	zp.initForYear(2023)
	if zp.year == InvalidYear {
		t.Fatal(zp)
	}
}

//-----------------------------------------------------------------------------
// Step 1
//-----------------------------------------------------------------------------

func TestCompareEraToYearMonth(t *testing.T) {
	// 2000-01-02 03:00
	era := zoneinfo.ZoneEra{
		UntilYear:            2000,
		UntilMonth:           1,
		UntilDay:             2,
		UntilSecondsCode:     3 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
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

	// 2000-01-01 00:00
	era2 := zoneinfo.ZoneEra{
		UntilYear:            2000,
		UntilMonth:           1,
		UntilDay:             1,
		UntilSecondsCode:     0,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}
	if !(0 == compareEraToYearMonth(&era2, 2000, 1)) {
		t.Fatal("fatal")
	}
}

func TestCreatematchingEra(t *testing.T) {
	// 14-month interval, from 2000-12 until 2002-02
	startYm := yearMonth{2000, 12}
	untilYm := yearMonth{2002, 2}

	// UNTIL = 2000-12-02 3:00
	era1 := zoneinfo.ZoneEra{
		UntilYear:            2000,
		UntilMonth:           12,
		UntilDay:             2,
		UntilSecondsCode:     3 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}

	// UNTIL = 2001-02-03 4:00
	era2 := zoneinfo.ZoneEra{
		UntilYear:            2001,
		UntilMonth:           2,
		UntilDay:             3,
		UntilSecondsCode:     4 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}

	// UNTIL = 2002-10-11 4:00
	era3 := zoneinfo.ZoneEra{
		UntilYear:            2002,
		UntilMonth:           10,
		UntilDay:             11,
		UntilSecondsCode:     4 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}

	// No previous matching era, so startDt is set to startYm.
	var match1 matchingEra
	creatematchingEra(&match1, nil, &era1, startYm, untilYm)
	if !(match1.startDt == dateTuple{2000, 12, 1, 3600 * 0, zoneinfo.SuffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.untilDt == dateTuple{2000, 12, 2, 3600 * 3, zoneinfo.SuffixW}) {
		t.Fatal("match1.startDt:", match1.startDt)
	}
	if !(match1.era == &era1) {
		t.Fatal("match1.startDt:", match1.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is < untilYm, so is retained.
	var match2 matchingEra
	creatematchingEra(&match2, &match1, &era2, startYm, untilYm)
	if !(match2.startDt == dateTuple{2000, 12, 2, 3600 * 3, zoneinfo.SuffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.untilDt == dateTuple{2001, 2, 3, 3600 * 4, zoneinfo.SuffixW}) {
		t.Fatal("match2.startDt:", match2.startDt)
	}
	if !(match2.era == &era2) {
		t.Fatal("match2.startDt:", match2.startDt)
	}

	// startDt is set to the prevMatch.untilDt.
	// untilDt is > untilYm so truncated to untilYm.
	var match3 matchingEra
	creatematchingEra(&match3, &match2, &era3, startYm, untilYm)
	if !(match3.startDt == dateTuple{2001, 2, 3, 3600 * 4, zoneinfo.SuffixW}) {
		t.Fatal("match3.startDt: ", match3.startDt)
	}
	if !(match3.untilDt == dateTuple{2002, 2, 1, 3600 * 0, zoneinfo.SuffixW}) {
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
	manager := NewZoneManager(&zonedbtesting.DataContext)
	info := manager.store.ZoneInfoByID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	era := &info.Eras[0]
	policy := era.Policy

	// Rule 6, [2007,9999]
	// Rule    US    2007    max    -    Nov    Sun>=1    2:00    0    S
	rule := &policy.Rules[6]

	// Nov 4 2018
	dt := getTransitionTime(2018, rule)
	if !(dt == dateTuple{2018, 11, 4, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// Nov 3 2019
	dt = getTransitionTime(2019, rule)
	if !(dt == dateTuple{2019, 11, 3, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}
}

func TestCreateTransitionForYear(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	info := manager.store.ZoneInfoByID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	era := &info.Eras[0]
	policy := era.Policy

	match := matchingEra{
		startDt:           dateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           dateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               era,
		prevMatch:         nil,
		lastOffsetSeconds: 0,
		lastDeltaSeconds:  0,
	}
	rule := &policy.Rules[6]

	// Nov Sun>=1
	var transition transition
	createTransitionForYear(&transition, 2019, rule, &match)
	if !(transition.offsetSeconds == -3600*8) {
		t.Fatal(transition.offsetSeconds)
	}
	if !(transition.deltaSeconds == 0) {
		t.Fatal(transition.deltaSeconds)
	}
	tt := &transition.transitionTime
	if !(*tt == dateTuple{2019, 11, 3, 3600 * 2, zoneinfo.SuffixW}) {
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
	manager := NewZoneManager(&zonedbtesting.DataContext)
	info := manager.store.ZoneInfoByID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	era := &info.Eras[0]

	match := matchingEra{
		startDt:           dateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           dateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               era,
		prevMatch:         nil,
		lastOffsetSeconds: 0,
		lastDeltaSeconds:  0,
	}

	// Reserve storage for the Transitions
	var ts transitionStorage

	// Verify compareTransitionToMatchFuzzy() elminates various transitions
	// to get down to 5:
	//    * 2018 Mar Sun>=8 (Mar 11)
	//    * 2019 Nov Sun>=1 (Nov 4)
	//    * 2019 Mar Sun>=8 (Mar 10)
	//    * 2019 Nov Sun>=1 (Nov 3)
	//    * 2020 Mar Sun>=8 (Mar 8)
	ts.resetCandidatePool()
	findCandidateTransitions(&ts, &match)
	candidates := ts.getCandidates()
	if !(5 == len(candidates)) {
		t.Fatal()
	}

	tt := &candidates[0].transitionTime
	if !(*tt == dateTuple{2018, 3, 11, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[1].transitionTime
	if !(*tt == dateTuple{2018, 11, 4, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[2].transitionTime
	if !(*tt == dateTuple{2019, 3, 10, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[3].transitionTime
	if !(*tt == dateTuple{2019, 11, 3, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &candidates[4].transitionTime
	if !(*tt == dateTuple{2020, 3, 8, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 2B Pass 3
//-----------------------------------------------------------------------------

func TestProcessTransitionCompareStatus(t *testing.T) {
	// UNTIL = 2002-01-02T03:00
	era := zoneinfo.ZoneEra{
		OffsetSecondsCode:    0,
		DeltaMinutes:         0,
		UntilYear:            2002,
		UntilMonth:           1,
		UntilDay:             2,
		UntilSecondsCode:     3 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}

	// [2000-01-01, 2001-01-01)
	match := matchingEra{
		startDt:           dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		untilDt:           dateTuple{2001, 1, 1, 0, zoneinfo.SuffixW},
		era:               &era,
		prevMatch:         nil,
		lastOffsetSeconds: 0,
		lastDeltaSeconds:  0,
	}

	// This transition occurs before the match, so prior should be filled.
	// transitionTime = 1999-12-31
	transitions := []transition{
		transition{
			match:          &match,
			transitionTime: dateTuple{1999, 12, 31, 0, zoneinfo.SuffixW},
		},
		// This occurs at exactly match.startDateTime, so should replace the prior.
		// transitionTime = 2000-01-01
		transition{
			match:          &match,
			transitionTime: dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		},
		// An interior transition. Prior should not change.
		// transitionTime = 2000-01-02
		transition{
			match:          &match,
			transitionTime: dateTuple{2000, 1, 2, 0, zoneinfo.SuffixW},
		},
		// Occurs after match.untilDateTime, so should be rejected.
		// transitionTime = 2001-01-02
		transition{
			match:          &match,
			transitionTime: dateTuple{2001, 1, 2, 0, zoneinfo.SuffixW},
		},
	}
	transition0 := &transitions[0]
	transition1 := &transitions[1]
	transition2 := &transitions[2]
	transition3 := &transitions[3]

	// Populate the transitionTimeS and transitionTimeU fields.
	var prior *transition = nil
	fixTransitionTimes(transitions)

	prior = processTransitionCompareStatus(transition0, prior)
	if !(compareStatusPrior == transition0.compareStatus) {
		t.Fatal(transition0.compareStatus)
	}
	if !(prior == transition0) {
		t.Fatal(transition0)
	}

	prior = processTransitionCompareStatus(transition1, prior)
	if !(compareStatusExactMatch == transition1.compareStatus) {
		t.Fatal(transition1.compareStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}

	prior = processTransitionCompareStatus(transition2, prior)
	if !(compareStatusWithinMatch == transition2.compareStatus) {
		t.Fatal(transition2.compareStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}

	prior = processTransitionCompareStatus(transition3, prior)
	if !(compareStatusFarFuture == transition3.compareStatus) {
		t.Fatal(transition3.compareStatus)
	}
	if !(prior == transition1) {
		t.Fatal(transition1)
	}
}

//-----------------------------------------------------------------------------
// Step 2B
//-----------------------------------------------------------------------------

func TestCreateTransitionsFromNamedMatch(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	info := manager.store.ZoneInfoByID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	era := &info.Eras[0]

	match := matchingEra{
		startDt:           dateTuple{2018, 12, 1, 0, zoneinfo.SuffixW},
		untilDt:           dateTuple{2020, 2, 1, 0, zoneinfo.SuffixW},
		era:               era,
		prevMatch:         nil,
		lastOffsetSeconds: 0,
		lastDeltaSeconds:  0,
	}

	// Reserve storage for the Transitions
	var ts transitionStorage

	createTransitionsFromNamedMatch(&ts, &match)
	if !(3 == ts.indexPrior) {
		t.Fatal(ts.indexPrior)
	}

	tt := &ts.transitions[0].transitionTime
	if !(*tt == dateTuple{2018, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[1].transitionTime
	if !(*tt == dateTuple{2019, 3, 10, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
	tt = &ts.transitions[2].transitionTime
	if !(*tt == dateTuple{2019, 11, 3, 3600 * 2, zoneinfo.SuffixW}) {
		t.Fatal(tt)
	}
}

//-----------------------------------------------------------------------------
// Step 3, Step 4. Use America/Los_Angeles to calculate the transitions
// beause I am familiar with it.
//-----------------------------------------------------------------------------

func TestFixTransitionTimesGenerateStartUntilTimes(t *testing.T) {
	manager := NewZoneManager(&zonedbtesting.DataContext)
	info := manager.store.ZoneInfoByID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	// Step 1: America/Los_Angeles matches one era, which points to US policy.
	var startYm = yearMonth{2017, 12}
	var untilYm = yearMonth{2019, 2}
	var matches [maxMatches]matchingEra

	numMatches := findMatches(info, startYm, untilYm, matches[:])
	if !(1 == numMatches) {
		t.Fatal(numMatches)
	}

	// Step 2: Create transitions.
	// Create a custom template instantiation to use a different SIZE than the
	// pre-defined typedef in ExtendedZoneProcess::transitionStorage.
	var storage transitionStorage
	createTransitions(&storage, matches[:numMatches])
	transitions := storage.getActives()
	if !(len(transitions) == 3) {
		t.Fatal(len(transitions))
	}

	// Step 3: Chain the transitions by fixing the transition times.
	fixTransitionTimes(transitions)

	// Step 3: Verification. The first transition is extended to -infinity.
	transition0 := &transitions[0]
	tt := &transition0.transitionTime
	if !(*tt == dateTuple{2017, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts := &transition0.transitionTimeS
	if !(*tts == dateTuple{2017, 12, 1, 0, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu := &transition0.transitionTimeU
	if !(*ttu == dateTuple{2017, 12, 1, 8 * 3600, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 3: Verification: Second transition springs forward at 2018-03-11
	// 02:00.
	transition1 := &transitions[1]
	tt = &transition1.transitionTime
	if !(*tt == dateTuple{2018, 3, 11, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition1.transitionTimeS
	if !(*tts == dateTuple{2018, 3, 11, 2 * 3600, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition1.transitionTimeU
	if !(*ttu == dateTuple{2018, 3, 11, 10 * 3600, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 3: Verification: Third transition falls back at 2018-11-04 02:00.
	transition2 := &transitions[2]
	tt = &transition2.transitionTime
	if !(*tt == dateTuple{2018, 11, 4, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("tt:", tt)
	}
	tts = &transition2.transitionTimeS
	if !(*tts == dateTuple{2018, 11, 4, 1 * 3600, zoneinfo.SuffixS}) {
		t.Fatal("tts:", tts)
	}
	ttu = &transition2.transitionTimeU
	if !(*ttu == dateTuple{2018, 11, 4, 9 * 3600, zoneinfo.SuffixU}) {
		t.Fatal("ttu:", ttu)
	}

	// Step 4: Generate the startDateTime and untilDateTime of the transitions.
	generateStartUntilTimes(transitions)

	// Step 4: Verification: The first transition startTime should be the same as
	// its transitionTime.
	sdt := &transition0.startDt
	if !(*sdt == dateTuple{2017, 12, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt := &transition0.untilDt
	if !(*udt == dateTuple{2018, 3, 11, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt := OffsetDateTime{
		2017, 12, 1, 0, 0, 0, 0 /*Fold*/, -8 * 3600 /*offsetSeconds*/}
	eps := odt.EpochSeconds()
	if !(eps == transition0.startEpochSeconds) {
		t.Fatal(eps, transition0.startEpochSeconds)
	}

	// Step 4: Verification: Second transition startTime is shifted forward one
	// hour into PDT.
	sdt = &transition1.startDt
	if !(*sdt == dateTuple{2018, 3, 11, 3 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition1.untilDt
	if !(*udt == dateTuple{2018, 11, 4, 2 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2018, 3, 11, 3, 0, 0, 0 /*Fold*/, -7 * 3600 /*offsetSeconds*/}
	eps = odt.EpochSeconds()
	if !(eps == transition1.startEpochSeconds) {
		t.Fatal(transition1.startEpochSeconds)
	}

	// Step 4: Verification: Third transition startTime is shifted back one hour
	// into PST.
	sdt = &transition2.startDt
	if !(*sdt == dateTuple{2018, 11, 4, 1 * 3600, zoneinfo.SuffixW}) {
		t.Fatal("sdt:", sdt)
	}
	udt = &transition2.untilDt
	if !(*udt == dateTuple{2019, 2, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal("udt:", udt)
	}
	odt = OffsetDateTime{
		2018, 11, 4, 1, 0, 0, 0 /*Fold*/, -8 * 3600 /*offsetSeconds*/}
	eps = odt.EpochSeconds()
	if !(eps == transition2.startEpochSeconds) {
		t.Fatal(transition2.startEpochSeconds)
	}
}

//-----------------------------------------------------------------------------
// Step 5
//-----------------------------------------------------------------------------

func TestCreateAbbreviation(t *testing.T) {
	// If no '%', deltaSeconds and Letter should not matter
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

	// If '/', then deltaSeconds selects the first or second component.
	abbrev = createAbbreviation("GMT/BST", 0, "")
	if !("GMT" == abbrev) {
		t.Fatal(abbrev)
	}

	abbrev = createAbbreviation("GMT/BST", 60, "")
	if !("BST" == abbrev) {
		t.Fatal(abbrev)
	}
}
