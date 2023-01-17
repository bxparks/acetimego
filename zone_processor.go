package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"strings"
)

//-----------------------------------------------------------------------------
// ZoneProcessor
//-----------------------------------------------------------------------------

type YearMonth struct {
	/** year [0,10000] */
	year int16
	/** month [1,12] */
	month uint8
}

type Err int8

const (
	ErrOk Err = iota
	ErrGeneric
)

const (
	maxMatches       = 4
	maxInteriorYears = 4
)

type ZoneProcessor struct {
	zoneInfo          *zoneinfo.ZoneInfo
	year              int16
	isFilled          bool
	numMatches        uint8
	matches           [maxMatches]MatchingEra
	transitionStorage TransitionStorage
}

func (zp *ZoneProcessor) isFilledForYear(year int16) bool {
	return zp.isFilled && (year == zp.year)
}

// InitForZoneInfo initializes the ZoneProcessor for the given zoneInfo.
func (zp *ZoneProcessor) InitForZoneInfo(zoneInfo *zoneinfo.ZoneInfo) {
	zp.zoneInfo = zoneInfo
	zp.isFilled = false
}

func (zp *ZoneProcessor) InitForYear(year int16) Err {
	if zp.isFilledForYear(year) {
		return ErrOk
	}

	zp.year = year
	zp.numMatches = 0
	zp.transitionStorage.Init()
	if year < zp.zoneInfo.StartYear-1 || zp.zoneInfo.UntilYear < year {
		return ErrGeneric
	}

	startYm := YearMonth{year - 1, 12}
	untilYm := YearMonth{year + 1, 2}

	// Step 1: Find matches.
	zp.numMatches = findMatches(zp.zoneInfo, startYm, untilYm, zp.matches[:])
	if zp.numMatches == 0 {
		return ErrGeneric
	}

	// Step 2: Create Transitions.
	createTransitions(&zp.transitionStorage, zp.matches[:zp.numMatches])

	// Step 3: Fix transition times.
	transitions := zp.transitionStorage.GetActives()
	fixTransitionTimes(transitions)

	// Step 4: Generate start and until times.
	generateStartUntilTimes(transitions)

	// Step 5: Calc abbreviations.
	calcAbbreviations(transitions)

	return ErrOk
}

func (zp *ZoneProcessor) InitForEpochSeconds(epochSeconds int32) Err {
	ldt := LocalDateTimeFromEpochSeconds(epochSeconds)
	if ldt.IsError() {
		return ErrGeneric
	}
	return zp.InitForYear(ldt.Year)
}

//---------------------------------------------------------------------------
// MonthDay
//-----------------------------------------------------------------------------

/** A tuple of month and day. */
type MonthDay struct {
	/** month [1,12] */
	month uint8

	/** day [1,31] */
	day uint8
}

// calcStartDayOfMonth Extracts the actual (month, day) pair from the expression
// used in the TZ data files of the form (OnDayOfWeek >= OnDayOfMonth) or
// (OnDayOfWeek <= OnDayOfMonth).
//
// There are 4 combinations:
//
// @verbatim
// OnDayOfWeek=0, OnDayOfMonth=(1-31): exact match
// OnDayOfWeek=1-7, OnDayOfMonth=1-31: dayOfWeek>=dayOfMonth
// OnDayOfWeek=1-7, OnDayOfMonth=0: last{dayOfWeek}
// OnDayOfWeek=1-7, OnDayOfMonth=-(1-31): dayOfWeek<=dayOfMonth
// @endverbatim
//
// Caveats: This function handles expressions which crosses month boundaries,
// but not year boundaries (e.g. Jan to Dec of the previous year, or Dec to
// Jan of the following year.)
func calcStartDayOfMonth(year int16, month uint8, onDayOfWeek uint8,
	onDayOfMonth int8) (md MonthDay) {

	if onDayOfWeek == 0 {
		md.month = month
		md.day = uint8(onDayOfMonth)
		return
	}

	if onDayOfMonth >= 0 {
		daysInMonth := int8(DaysInYearMonth(year, month))
		if onDayOfMonth == 0 {
			onDayOfMonth = daysInMonth - 6
		}
		dow := DayOfWeek(year, month, uint8(onDayOfMonth))
		dayOfWeekShift := (onDayOfWeek - dow + 7) % 7
		day := onDayOfMonth + int8(dayOfWeekShift)
		if day > daysInMonth {
			// TODO: Support shifting from Dec to Jan of following  year.
			day -= daysInMonth
			month++
		}
		md.month = month
		md.day = uint8(day)
	} else {
		onDayOfMonth = -onDayOfMonth
		dow := DayOfWeek(year, month, uint8(onDayOfMonth))
		dayOfWeekShift := (dow - onDayOfWeek + 7) % 7
		day := onDayOfMonth - int8(dayOfWeekShift)
		if day < 1 {
			// TODO: Support shifting from Jan to Dec of the previous year.
			month--
			daysInPrevMonth := DaysInYearMonth(year, month)
			day += int8(daysInPrevMonth)
		}
		md.month = month
		md.day = uint8(day)
	}
	return
}

//-----------------------------------------------------------------------------
// Step 1
//-----------------------------------------------------------------------------

func findMatches(
	zoneInfo *zoneinfo.ZoneInfo,
	startYm YearMonth,
	untilYm YearMonth,
	matches []MatchingEra) uint8 {

	var iMatch uint8 = 0
	var prevMatch *MatchingEra = nil
	var eras []zoneinfo.ZoneEra = zoneInfo.ErasActive()

	for iEra := range eras {
		era := &eras[iEra]
		var prevEra *zoneinfo.ZoneEra = nil
		if prevMatch != nil {
			prevEra = prevMatch.era
		}
		if eraOverlapsInterval(prevEra, era, startYm, untilYm) {
			if iMatch < uint8(len(matches)) {
				createMatchingEra(&matches[iMatch], prevMatch, era, startYm, untilYm)
				prevMatch = &matches[iMatch]
				iMatch++
			}
		}
	}
	return iMatch
}

/**
 * Determines if era overlaps the interval [startYm, untilYm). This does
 * not need to be exact since the startYm and untilYm are created to have
 * some slop of about one month at the low and high end, so we can ignore
 * the day, time and timeSuffix fields of the era. The start date of the
 * current era is represented by the UNTIL fields of the previous era, so
 * the interval of the current era is [era.start=prev.UNTIL,
 * era.until=era.UNTIL). Overlap happens if (era.start < untilYm) and
 * (era.until > startYm). If prev.isNull(), then interpret prev as the
 * earliest ZoneEra.
 */
func eraOverlapsInterval(
	prevEra *zoneinfo.ZoneEra,
	era *zoneinfo.ZoneEra,
	startYm YearMonth,
	untilYm YearMonth) bool {

	return (prevEra == nil ||
		compareEraToYearMonth(prevEra, untilYm.year, untilYm.month) < 0) &&
		compareEraToYearMonth(era, startYm.year, startYm.month) > 0
}

/** Return (1, 0, -1) depending on how era compares to (year, month). */
func compareEraToYearMonth(
	era *zoneinfo.ZoneEra, year int16, month uint8) int8 {

	if era.UntilYear < year {
		return -1
	}
	if era.UntilYear > year {
		return 1
	}
	if era.UntilMonth < month {
		return -1
	}
	if era.UntilMonth > month {
		return 1
	}
	if era.UntilDay > 1 {
		return 1
	}
	//if era.until_time_minutes < 0 { return -1; // never possible
	if era.UntilTimeCode > 0 {
		return 1
	}
	return 0
}

/**
 * Create a new MatchingEra object around the 'era' which intersects the
 * half-open [startYm, untilYm) interval. The interval is assumed to overlap
 * the ZoneEra using the eraOverlapsInterval() method. The 'prev' ZoneEra is
 * needed to define the startDateTime of the current era.
 */
func createMatchingEra(
	newMatch *MatchingEra,
	prevMatch *MatchingEra,
	era *zoneinfo.ZoneEra,
	startYm YearMonth,
	untilYm YearMonth) {

	// If prevMatch is nil, set startDate to be earlier than all valid ZoneEra.
	var startDate DateTuple
	if prevMatch == nil {
		startDate.year = InvalidYear
		startDate.month = 1
		startDate.day = 1
		startDate.minutes = 0
		startDate.suffix = zoneinfo.SuffixW
	} else {
		startDate.year = prevMatch.era.UntilYear
		startDate.month = prevMatch.era.UntilMonth
		startDate.day = prevMatch.era.UntilDay
		startDate.minutes = prevMatch.era.UntilMinutes()
		startDate.suffix = prevMatch.era.UntilSuffix()
	}
	lowerBound := DateTuple{startYm.year, startYm.month, 1, 0, zoneinfo.SuffixW}
	if dateTupleCompare(&startDate, &lowerBound) < 0 {
		startDate = lowerBound
	}

	untilDate := DateTuple{
		era.UntilYear,
		era.UntilMonth,
		era.UntilDay,
		era.UntilMinutes(),
		era.UntilSuffix(),
	}
	upperBound := DateTuple{untilYm.year, untilYm.month, 1, 0, zoneinfo.SuffixW}
	if dateTupleCompare(&upperBound, &untilDate) < 0 {
		untilDate = upperBound
	}

	newMatch.startDt = startDate
	newMatch.untilDt = untilDate
	newMatch.era = era
	newMatch.prevMatch = prevMatch
	newMatch.lastOffsetMinutes = 0
	newMatch.lastDeltaMinutes = 0
}

//-----------------------------------------------------------------------------
// Step 2
//-----------------------------------------------------------------------------

func createTransitions(ts *TransitionStorage, matches []MatchingEra) {
	for i := range matches {
		createTransitionsForMatch(ts, &matches[i])
	}
}

func createTransitionsForMatch(ts *TransitionStorage, match *MatchingEra) {
	policy := match.era.ZonePolicy
	if policy == nil {
		// Step 2A
		createTransitionsFromSimpleMatch(ts, match)
	} else {
		// Step 2B
		createTransitionsFromNamedMatch(ts, match)
	}
}

//-----------------------------------------------------------------------------
// Step 2A
//-----------------------------------------------------------------------------

func createTransitionsFromSimpleMatch(
	ts *TransitionStorage, match *MatchingEra) {

	freeAgent := ts.GetFreeAgent()
	createTransitionForYear(freeAgent, 0, nil, match)
	freeAgent.matchStatus = matchStatusExactMatch
	match.lastOffsetMinutes = freeAgent.offsetMinutes
	match.lastDeltaMinutes = freeAgent.deltaMinutes
	ts.AddFreeAgentToActivePool()
}

func createTransitionForYear(
	t *Transition, year int16, rule *zoneinfo.ZoneRule, match *MatchingEra) {

	t.match = match
	t.rule = rule
	t.offsetMinutes = match.era.StdOffsetMinutes()
	t.letter = ""

	if rule != nil {
		t.transitionTime = getTransitionTime(year, rule)
		t.deltaMinutes = rule.DstOffsetMinutes()
		// If LETTER is a '-', treat it the same as an empty string.
		if rule.Letter != "-" {
			t.letter = rule.Letter
		}
	} else {
		// Create a Transition using the MatchingEra for the transitionTime.
		// Used for simple MatchingEra.
		t.transitionTime = match.startDt
		t.deltaMinutes = match.era.DstOffsetMinutes()
	}
}

func getTransitionTime(year int16, rule *zoneinfo.ZoneRule) DateTuple {
	md := calcStartDayOfMonth(
		year, rule.InMonth, rule.OnDayOfWeek, rule.OnDayOfMonth)
	return DateTuple{
		year:    year,
		month:   md.month,
		day:     md.day,
		minutes: rule.AtMinutes(),
		suffix:  rule.AtSuffix(),
	}
}

//-----------------------------------------------------------------------------
// Step 2B
//-----------------------------------------------------------------------------

func createTransitionsFromNamedMatch(
	ts *TransitionStorage, match *MatchingEra) {

	ts.ResetCandidatePool()

	// Pass 1: Find candidate transitions using whole years.
	findCandidateTransitions(ts, match)

	// Pass 2: Fix the transitions times, converting 's' and 'u' into 'w'
	// uniformly.
	transitions := ts.GetCandidates()
	fixTransitionTimes(transitions)

	// Pass 3: Select only those Transitions which overlap with the actual
	// start and until times of the MatchingEra.
	selectActiveTransitions(transitions)
	lastTransition := ts.AddActiveCandidatesToActivePool()
	match.lastOffsetMinutes = lastTransition.offsetMinutes
	match.lastDeltaMinutes = lastTransition.deltaMinutes
}

// Step 2B: Pass 1
func findCandidateTransitions(ts *TransitionStorage, match *MatchingEra) {
	policy := match.era.ZonePolicy
	startYear := match.startDt.year
	endYear := match.untilDt.year

	prior := ts.ReservePrior()
	prior.isValidPrior = false
	for ir := range policy.Rules {
		rule := &policy.Rules[ir]

		// Add transitions for interior years
		var interiorYears [maxInteriorYears]int16
		numYears := calcInteriorYears(
			interiorYears[:], rule.FromYear, rule.ToYear, startYear, endYear)
		for iy := uint8(0); iy < numYears; iy++ {
			year := interiorYears[iy]
			t := ts.GetFreeAgent()
			createTransitionForYear(t, year, rule, match)
			status := compareTransitionToMatchFuzzy(t, match)
			if status == matchStatusPrior {
				ts.SetFreeAgentAsPriorIfValid()
			} else if status == matchStatusWithinMatch {
				ts.AddFreeAgentToCandidatePool()
			} else {
				// Must be kFarFuture.
				// Do nothing, allowing the free agent to be reused.
			}
		}

		// Add Transition for prior year
		priorYear := getMostRecentPriorYear(rule.FromYear, rule.ToYear, startYear)
		if priorYear != InvalidYear {
			t := ts.GetFreeAgent()
			createTransitionForYear(t, priorYear, rule, match)
			ts.SetFreeAgentAsPriorIfValid()
		}
	}

	// Add the reserved prior into the Candidate pool only if 'isValidPrior' is
	// true.
	if prior.isValidPrior {
		ts.AddPriorToCandidatePool()
	}
}

// calcInteriorYears calculates the interior years that overlaps (FromYear,
// ToYear) and (startYear, endYear). The results are placed into the
// interiorYears slice, and the number of elements are returned.
func calcInteriorYears(
	interiorYears []int16,
	fromYear int16,
	toYear int16,
	startYear int16,
	endYear int16) uint8 {

	var i uint8 = 0
	for year := startYear; year <= endYear; year++ {
		if fromYear <= year && year <= toYear {
			interiorYears[i] = year
			i++
			if int(i) >= len(interiorYears) {
				break
			}
		}
	}
	return i
}

func getMostRecentPriorYear(
	fromYear int16, toYear int16, startYear int16) int16 {

	if fromYear < startYear {
		if toYear < startYear {
			return toYear
		} else {
			return startYear - 1
		}
	} else {
		return InvalidYear
	}
}

// Step 2B: Pass 3
func selectActiveTransitions(transitions []Transition) {
	var prior *Transition = nil
	for i := range transitions {
		transition := &transitions[i]
		prior = processTransitionMatchStatus(transition, prior)
	}

	// If the latest prior transition is found, shift it to start at the
	// startDateTime of the current match.
	if prior != nil {
		prior.transitionTime = prior.match.startDt
	}
}

func processTransitionMatchStatus(
	transition *Transition, prior *Transition) *Transition {

	status := compareTransitionToMatch(transition, transition.match)
	transition.matchStatus = status

	if status == matchStatusExactMatch {
		if prior != nil {
			prior.matchStatus = matchStatusFarPast
		}
		prior = transition
	} else if status == matchStatusPrior {
		if prior != nil {
			if dateTupleCompare(
				&prior.transitionTimeU, &transition.transitionTimeU) <= 0 {

				prior.matchStatus = matchStatusFarPast
				prior = transition
			} else {
				transition.matchStatus = matchStatusFarPast
			}
		} else {
			prior = transition
		}
	}

	return prior
}

//-----------------------------------------------------------------------------
// Step 4
//-----------------------------------------------------------------------------

func generateStartUntilTimes(transitions []Transition) {
	prev := &transitions[0]
	isAfterFirst := false

	for i := range transitions {
		transition := &transitions[i]

		// 1) Update the untilDateTime of the previous Transition
		tt := &transition.transitionTime
		if isAfterFirst {
			prev.untilDt = *tt
		}

		// 2) Calculate the current startDateTime by shifting the
		// transitionTime (represented in the UTC offset of the previous
		// transition) into the UTC offset of the *current* transition.
		var minutes int16 = tt.minutes +
			transition.offsetMinutes +
			transition.deltaMinutes -
			prev.offsetMinutes -
			prev.deltaMinutes
		transition.startDt.year = tt.year
		transition.startDt.month = tt.month
		transition.startDt.day = tt.day
		transition.startDt.minutes = minutes
		transition.startDt.suffix = tt.suffix
		dateTupleNormalize(&transition.startDt)

		// 3) The epochSecond of the 'transitionTime' is determined by the
		// UTC offset of the *previous* Transition. However, the
		// transitionTime can be represented by an illegal time (e.g. 24:00).
		// So, it is better to use the properly normalized startDateTime
		// (calculated above) with the *current* UTC offset.
		//
		// NOTE: We should also be able to  calculate this directly from
		// 'transitionTimeU' which should still be a valid field, because it
		// hasn't been clobbered by 'untilDateTime' yet. Not sure if this saves
		// any CPU time though, since we still need to mutiply by 900.
		st := &transition.startDt
		offsetSeconds := 60 * int32(st.minutes-
			(transition.offsetMinutes+transition.deltaMinutes))
		epochSeconds := 86400 * LocalDateToEpochDays(st.year, st.month, st.day)
		transition.startEpochSeconds = epochSeconds + offsetSeconds

		prev = transition
		isAfterFirst = true
	}

	// The last Transition's until time is the until time of the MatchingEra.
	var untilTimeW DateTuple
	var untilTimeS DateTuple
	var untilTimeU DateTuple
	dateTupleExpand(
		&prev.match.untilDt,
		prev.offsetMinutes,
		prev.deltaMinutes,
		&untilTimeW,
		&untilTimeS,
		&untilTimeU)
	prev.untilDt = untilTimeW
}

//-----------------------------------------------------------------------------
// Step 5
//-----------------------------------------------------------------------------

func calcAbbreviations(transitions []Transition) {
	for i := range transitions {
		transition := &transitions[i]
		transition.abbrev = createAbbreviation(
			transition.match.era.Format,
			transition.deltaMinutes,
			transition.getLetter())
	}
}

func createAbbreviation(
	format string, deltaMinutes int16, letter string) string {

	// Check if FORMAT contains a '%'.
	if strings.IndexByte(format, '%') >= 0 {
		// If RULES column empty, then letter == "" because Go lang does not allow
		// strings to be set to nil. So we cannot distinguish between "" and not
		// existing. In Go lang then, always replace "%" with "".
		return strings.ReplaceAll(format, "%", letter)
	} else {
		// Check if FORMAT contains a '/'.
		slashIndex := strings.IndexByte(format, '/')
		if slashIndex != -1 {
			if deltaMinutes == 0 {
				return format[:slashIndex]
			} else {
				return format[slashIndex+1:]
			}
		} else {
			// Just return FORMAT disregarding deltaMinutes and Letter.
			return format
		}
	}
}

//---------------------------------------------------------------------------
// FindByLocalDateTime() and FindByEpochSeconds()
//---------------------------------------------------------------------------

// Values of the FindResult.type field.
const (
	FindResultNotFound = iota
	FindResultExact
	FindResultGap
	FindResultOverlap
)

type FindResult struct {
	frtype              uint8
	fold                uint8
	stdOffsetMinutes    int16  // STD offset
	dstOffsetMinutes    int16  // DST offset
	reqStdOffsetMinutes int16  // request STD offset
	reqDstOffsetMinutes int16  // request DST offset
	abbrev              string // abbreviation (e.g. PST, PDT)
}

// TODO: Merge error condition into frtype field
func NewFindResultError() FindResult {
	return FindResult{stdOffsetMinutes: InvalidOffsetMinutes}
}

func (ti *FindResult) IsError() bool {
	return ti.stdOffsetMinutes == InvalidOffsetMinutes
}

// Find the AtcFindResult at the given epoch_seconds.
//
// Adapted from ExtendedZoneProcessor::findByEpochSeconds(epochSeconds)
// in the AceTime library and atc_processor_find_by_epoch_seconds() in the
// AceTimeC library.
func (zp *ZoneProcessor) FindByEpochSeconds(epochSeconds int32) FindResult {
	err := zp.InitForEpochSeconds(epochSeconds)
	if err != ErrOk {
		return NewFindResultError()
	}

	tfs := zp.transitionStorage.findTransitionForSeconds(epochSeconds)
	transition := tfs.curr
	if transition == nil {
		return NewFindResultError()
	}

	var frtype uint8
	if tfs.num == 2 {
		frtype = FindResultOverlap
	} else {
		frtype = FindResultExact
	}
	return FindResult{
		frtype:              frtype,
		fold:                tfs.fold,
		stdOffsetMinutes:    transition.offsetMinutes,
		dstOffsetMinutes:    transition.deltaMinutes,
		reqStdOffsetMinutes: transition.offsetMinutes,
		reqDstOffsetMinutes: transition.deltaMinutes,
		abbrev:              transition.abbrev,
	}
}

// Return the FindResult at the given LocalDateTime.
//
// Adapted from ExtendedZoneProcessor::findByLocalDateTime(const LocalDatetime&)
// in the AceTime library and atc_processor_find_by_local_date_time() in the
// AceTimeC library.
func (zp *ZoneProcessor) FindByLocalDateTime(
	ldt *LocalDateTime, fold uint8) FindResult {

	err := zp.InitForYear(ldt.Year)
	if err != ErrOk {
		return NewFindResultError()
	}

	tfd := zp.transitionStorage.findTransitionForDateTime(ldt)

	// Extract the appropriate Transition, depending on the requested 'fold'
	// and the 'tfd.searchStatus'.
	var transition *Transition
	var result FindResult
	if tfd.num == 1 {
		transition = tfd.curr
		result.frtype = FindResultExact
		result.fold = 0
		result.reqStdOffsetMinutes = transition.offsetMinutes
		result.reqDstOffsetMinutes = transition.deltaMinutes
	} else { // num = 0 or 2
		if tfd.prev == nil || tfd.curr == nil {
			// ldt was far past or far future, and didn't match anything.
			transition = nil
			result.frtype = FindResultNotFound
			result.fold = 0
		} else { // gap or overlap
			if tfd.num == 0 { // gap
				result.frtype = FindResultGap
				result.fold = 0
				if fold == 0 {
					// ldt wants to use the 'prev' transition to convert to
					// epochSeconds.
					result.reqStdOffsetMinutes = tfd.prev.offsetMinutes
					result.reqDstOffsetMinutes = tfd.prev.deltaMinutes
					// But after normalization, it will be shifted into the curr
					// transition, so select 'curr' as the target transition.
					transition = tfd.curr
				} else {
					// ldt wants to use the 'curr' transition to convert to
					// epochSeconds.
					result.reqStdOffsetMinutes = tfd.curr.offsetMinutes
					result.reqDstOffsetMinutes = tfd.curr.deltaMinutes
					// But after normalization, it will be shifted into the prev
					// transition, so select 'prev' as the target transition.
					transition = tfd.prev
				}
			} else {
				if fold == 0 {
					transition = tfd.prev
				} else {
					transition = tfd.curr
				}
				result.frtype = FindResultOverlap
				result.fold = fold
				result.reqStdOffsetMinutes = transition.offsetMinutes
				result.reqDstOffsetMinutes = transition.deltaMinutes
			}
		}
	}

	if transition == nil {
		return NewFindResultError()
	}

	result.stdOffsetMinutes = transition.offsetMinutes
	result.dstOffsetMinutes = transition.deltaMinutes
	result.abbrev = transition.abbrev
	return result
}
