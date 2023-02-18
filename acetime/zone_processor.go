package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"strings"
)

//-----------------------------------------------------------------------------
// ZoneProcessor
//-----------------------------------------------------------------------------

type yearMonth struct {
	// year [0,10000]
	year int16
	// month [1,12]
	month uint8
}

type Err uint8

const (
	ErrOk Err = iota
	ErrGeneric
)

const (
	maxMatches       = 4
	maxInteriorYears = 4
)

type ZoneProcessor struct {
	zoneInfo   *zoneinfo.ZoneInfo
	year       int16
	isFilled   bool
	numMatches uint8
	matches    [maxMatches]matchingEra
	tstorage   transitionStorage
}

func (zp *ZoneProcessor) isFilledForYear(year int16) bool {
	return zp.isFilled && (year == zp.year)
}

// InitForZoneInfo initializes the ZoneProcessor for the given zoneInfo.
func (zp *ZoneProcessor) InitForZoneInfo(zoneInfo *zoneinfo.ZoneInfo) {

	zp.zoneInfo = zoneInfo
	zp.isFilled = false
}

// Clear cache, used only for tests.
func (zp *ZoneProcessor) reset() {
	zp.isFilled = false
}

func (zp *ZoneProcessor) IsLink() bool {
	return zp.zoneInfo.IsLink()
}

func (zp *ZoneProcessor) InitForYear(year int16) Err {
	if zp.isFilledForYear(year) {
		return ErrOk
	}

	zp.year = year
	zp.isFilled = true
	zp.numMatches = 0
	zp.tstorage.init()
	if year < zp.zoneInfo.StartYear-1 || zp.zoneInfo.UntilYear < year {
		return ErrGeneric
	}

	startYm := yearMonth{year - 1, 12}
	untilYm := yearMonth{year + 1, 2}

	// Step 1: Find matches.
	zp.numMatches = findMatches(zp.zoneInfo, startYm, untilYm, zp.matches[:])
	if zp.numMatches == 0 {
		return ErrGeneric
	}

	// Step 2: Create Transitions.
	createTransitions(&zp.tstorage, zp.matches[:zp.numMatches])

	// Step 3: Fix transition times.
	transitions := zp.tstorage.getActives()
	fixTransitionTimes(transitions)

	// Step 4: Generate start and until times.
	generateStartUntilTimes(transitions)

	// Step 5: Calc abbreviations.
	calcAbbreviations(transitions)

	return ErrOk
}

func (zp *ZoneProcessor) InitForEpochSeconds(epochSeconds ATime) Err {
	ldt := NewLocalDateTimeFromEpochSeconds(epochSeconds)
	if ldt.IsError() {
		return ErrGeneric
	}
	return zp.InitForYear(ldt.Year)
}

func (zp *ZoneProcessor) Name() string {
	return zp.zoneInfo.Name
}

//---------------------------------------------------------------------------
// monthDay
//-----------------------------------------------------------------------------

// monthDay is a tuple of month and day.
type monthDay struct {
	month uint8 // [1,12]
	day   uint8 // [1,31]
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
	onDayOfMonth int8) (md monthDay) {

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
	startYm yearMonth,
	untilYm yearMonth,
	matches []matchingEra) uint8 {

	var iMatch uint8 = 0
	var prevMatch *matchingEra = nil
	var eras []zoneinfo.ZoneEra = zoneInfo.ErasActive()

	for iEra := range eras {
		era := &eras[iEra]
		var prevEra *zoneinfo.ZoneEra = nil
		if prevMatch != nil {
			prevEra = prevMatch.era
		}
		if eraOverlapsInterval(prevEra, era, startYm, untilYm) {
			if iMatch < uint8(len(matches)) {
				creatematchingEra(&matches[iMatch], prevMatch, era, startYm, untilYm)
				prevMatch = &matches[iMatch]
				iMatch++
			}
		}
	}
	return iMatch
}

// Determines if era overlaps the interval [startYm, untilYm). This does
// not need to be exact since the startYm and untilYm are created to have
// some slop of about one month at the low and high end, so we can ignore
// the day, time and timeSuffix fields of the era. The start date of the
// current era is represented by the UNTIL fields of the previous era, so
// the interval of the current era is [era.start=prev.UNTIL,
// era.until=era.UNTIL). Overlap happens if (era.start < untilYm) and
// (era.until > startYm). If prev.isNull(), then interpret prev as the
// earliest ZoneEra.
func eraOverlapsInterval(
	prevEra *zoneinfo.ZoneEra,
	era *zoneinfo.ZoneEra,
	startYm yearMonth,
	untilYm yearMonth) bool {

	return (prevEra == nil ||
		compareEraToYearMonth(prevEra, untilYm.year, untilYm.month) < 0) &&
		compareEraToYearMonth(era, startYm.year, startYm.month) > 0
}

// Return (1, 0, -1) depending on how era compares to (year, month).
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
	//if era.UntilSeconds() < 0 return -1; // never possible
	if era.UntilSeconds() > 0 {
		return 1
	}
	return 0
}

// Create a new matchingEra object around the 'era' which intersects the
// half-open [startYm, untilYm) interval. The interval is assumed to overlap
// the ZoneEra using the eraOverlapsInterval() method. The 'prev' ZoneEra is
// needed to define the startDateTime of the current era.
func creatematchingEra(
	newMatch *matchingEra,
	prevMatch *matchingEra,
	era *zoneinfo.ZoneEra,
	startYm yearMonth,
	untilYm yearMonth) {

	// If prevMatch is nil, set startDate to be earlier than all valid ZoneEra.
	var startDate DateTuple
	if prevMatch == nil {
		startDate.year = InvalidYear
		startDate.month = 1
		startDate.day = 1
		startDate.seconds = 0
		startDate.suffix = zoneinfo.SuffixW
	} else {
		startDate.year = prevMatch.era.UntilYear
		startDate.month = prevMatch.era.UntilMonth
		startDate.day = prevMatch.era.UntilDay
		startDate.seconds = prevMatch.era.UntilSeconds()
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
		era.UntilSeconds(),
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
	newMatch.lastOffsetSeconds = 0
	newMatch.lastDeltaSeconds = 0
	newMatch.format = era.Format
}

//-----------------------------------------------------------------------------
// Step 2
//-----------------------------------------------------------------------------

func createTransitions(ts *transitionStorage, matches []matchingEra) {

	for i := range matches {
		createTransitionsForMatch(ts, &matches[i])
	}
}

func createTransitionsForMatch(ts *transitionStorage, match *matchingEra) {

	if match.era.HasPolicy() {
		// Step 2B
		createTransitionsFromNamedMatch(ts, match)
	} else {
		// Step 2A
		createTransitionsFromSimpleMatch(ts, match)
	}
}

//-----------------------------------------------------------------------------
// Step 2A
//-----------------------------------------------------------------------------

func createTransitionsFromSimpleMatch(
	ts *transitionStorage, match *matchingEra) {

	freeAgent := ts.getFreeAgent()
	createTransitionForYear(freeAgent, 0, nil, match)
	freeAgent.compareStatus = compareStatusExactMatch
	match.lastOffsetSeconds = freeAgent.offsetSeconds
	match.lastDeltaSeconds = freeAgent.deltaSeconds
	ts.addFreeAgentToActivePool()
}

func createTransitionForYear(
	t *transition, year int16, rule *zoneinfo.ZoneRule, match *matchingEra) {

	t.match = match
	t.offsetSeconds = match.era.StdOffsetSeconds()

	if rule != nil {
		t.transitionTime = getTransitionTime(year, rule)
		t.deltaSeconds = rule.DstOffsetSeconds()
		t.letter = rule.Letter
	} else {
		// Create a transition using the matchingEra for the transitionTime.
		// Used for simple matchingEra.
		t.transitionTime = match.startDt
		t.deltaSeconds = match.era.DstOffsetSeconds()
		t.letter = ""
	}
}

func getTransitionTime(year int16, rule *zoneinfo.ZoneRule) DateTuple {
	md := calcStartDayOfMonth(
		year, rule.InMonth, rule.OnDayOfWeek, rule.OnDayOfMonth)
	return DateTuple{
		year:    year,
		month:   md.month,
		day:     md.day,
		seconds: rule.AtSeconds(),
		suffix:  rule.AtSuffix(),
	}
}

//-----------------------------------------------------------------------------
// Step 2B
//-----------------------------------------------------------------------------

func createTransitionsFromNamedMatch(
	ts *transitionStorage, match *matchingEra) {

	ts.resetCandidatePool()

	// Pass 1: Find candidate transitions using whole years.
	findCandidateTransitions(ts, match)

	// Pass 2: Fix the transitions times, converting 's' and 'u' into 'w'
	// uniformly.
	transitions := ts.getCandidates()
	fixTransitionTimes(transitions)

	// Pass 3: Select only those Transitions which overlap with the actual
	// start and until times of the matchingEra.
	selectActiveTransitions(transitions)
	lastTransition := ts.addActiveCandidatesToActivePool()
	match.lastOffsetSeconds = lastTransition.offsetSeconds
	match.lastDeltaSeconds = lastTransition.deltaSeconds
}

// Step 2B: Pass 1
func findCandidateTransitions(ts *transitionStorage, match *matchingEra) {

	policy := match.era.Policy
	startYear := match.startDt.year
	endYear := match.untilDt.year

	prior := ts.reservePrior()
	prior.isValidPrior = false
	rules := policy.Rules
	for ir := range rules {
		rule := &rules[ir]

		// Add transitions for interior years
		var interiorYears [maxInteriorYears]int16
		numYears := calcInteriorYears(
			interiorYears[:], rule.FromYear, rule.ToYear, startYear, endYear)
		for iy := uint8(0); iy < numYears; iy++ {
			year := interiorYears[iy]
			t := ts.getFreeAgent()
			createTransitionForYear(t, year, rule, match)
			status := compareTransitionToMatchFuzzy(t, match)
			if status == compareStatusPrior {
				ts.setFreeAgentAsPriorIfValid()
			} else if status == compareStatusWithinMatch {
				ts.addFreeAgentToCandidatePool()
			} else {
				// Must be kFarFuture.
				// Do nothing, allowing the free agent to be reused.
			}
		}

		// Add transition for prior year
		priorYear := getMostRecentPriorYear(rule.FromYear, rule.ToYear, startYear)
		if priorYear != InvalidYear {
			t := ts.getFreeAgent()
			createTransitionForYear(t, priorYear, rule, match)
			ts.setFreeAgentAsPriorIfValid()
		}
	}

	// Add the reserved prior into the Candidate pool only if 'isValidPrior' is
	// true.
	if prior.isValidPrior {
		ts.addPriorToCandidatePool()
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
func selectActiveTransitions(transitions []transition) {
	var prior *transition = nil
	for i := range transitions {
		transition := &transitions[i]
		prior = processTransitionCompareStatus(transition, prior)
	}

	// If the latest prior transition is found, shift it to start at the
	// startDateTime of the current match.
	if prior != nil {
		prior.transitionTime = prior.match.startDt
	}
}

func processTransitionCompareStatus(
	transition *transition, prior *transition) *transition {

	status := compareTransitionToMatch(transition, transition.match)
	transition.compareStatus = status

	if status == compareStatusExactMatch {
		if prior != nil {
			prior.compareStatus = compareStatusFarPast
		}
		prior = transition
	} else if status == compareStatusPrior {
		if prior != nil {
			if dateTupleCompare(
				&prior.transitionTimeU, &transition.transitionTimeU) <= 0 {

				prior.compareStatus = compareStatusFarPast
				prior = transition
			} else {
				transition.compareStatus = compareStatusFarPast
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

func generateStartUntilTimes(transitions []transition) {
	prev := &transitions[0]
	isAfterFirst := false

	for i := range transitions {
		transition := &transitions[i]

		// 1) Update the untilDateTime of the previous transition
		tt := &transition.transitionTime
		if isAfterFirst {
			prev.untilDt = *tt
		}

		// 2) Calculate the current startDateTime by shifting the
		// transitionTime (represented in the UTC offset of the previous
		// transition) into the UTC offset of the *current* transition.
		seconds := tt.seconds +
			transition.offsetSeconds +
			transition.deltaSeconds -
			prev.offsetSeconds -
			prev.deltaSeconds
		transition.startDt.year = tt.year
		transition.startDt.month = tt.month
		transition.startDt.day = tt.day
		transition.startDt.seconds = seconds
		transition.startDt.suffix = tt.suffix
		dateTupleNormalize(&transition.startDt)

		// 3) The epochSecond of the 'transitionTime' is determined by the
		// UTC offset of the *previous* transition. However, the
		// transitionTime can be represented by an illegal time (e.g. 24:00).
		// So, it is better to use the properly normalized startDateTime
		// (calculated above) with the *current* UTC offset.
		//
		// NOTE: We should also be able to  calculate this directly from
		// 'transitionTimeU' which should still be a valid field, because it
		// hasn't been clobbered by 'untilDateTime' yet. Not sure if this saves
		// any CPU time though, since we still need to mutiply by 900.
		st := &transition.startDt
		offsetSeconds := ATime(st.seconds -
			(transition.offsetSeconds + transition.deltaSeconds))
		epochSeconds := 86400 * ATime(
			LocalDateToEpochDays(st.year, st.month, st.day))
		transition.startEpochSeconds = epochSeconds + offsetSeconds

		prev = transition
		isAfterFirst = true
	}

	// The last transition's until time is the until time of the matchingEra.
	var untilTimeW DateTuple
	var untilTimeS DateTuple
	var untilTimeU DateTuple
	dateTupleExpand(
		&prev.match.untilDt,
		prev.offsetSeconds,
		prev.deltaSeconds,
		&untilTimeW,
		&untilTimeS,
		&untilTimeU)
	prev.untilDt = untilTimeW
}

//-----------------------------------------------------------------------------
// Step 5
//-----------------------------------------------------------------------------

func calcAbbreviations(transitions []transition) {

	for i := range transitions {
		transition := &transitions[i]
		transition.abbrev = createAbbreviation(
			transition.match.era.Format,
			transition.deltaSeconds,
			transition.letter)
	}
}

func createAbbreviation(
	format string, deltaSeconds int32, letter string) string {

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
			if deltaSeconds == 0 {
				return format[:slashIndex]
			} else {
				return format[slashIndex+1:]
			}
		} else {
			// Just return FORMAT disregarding deltaSeconds and Letter.
			return format
		}
	}
}

//---------------------------------------------------------------------------
// FindByLocalDateTime() and FindByEpochSeconds()
//---------------------------------------------------------------------------

// Values of the findResult.type field.
const (
	findResultErr = iota
	findResultNotFound
	findResultExact
	findResultGap
	findResultOverlap
)

var (
	findResultError = findResult{frtype: findResultErr}
)

type findResult struct {
	frtype              uint8
	fold                uint8
	stdOffsetSeconds    int32  // STD offset
	dstOffsetSeconds    int32  // DST offset
	reqStdOffsetSeconds int32  // request STD offset
	reqDstOffsetSeconds int32  // request DST offset
	abbrev              string // abbreviation (e.g. PST, PDT)
}

// Find the AtcfindResult at the given epoch_seconds.
//
// Adapted from ExtendedZoneProcessor::findByEpochSeconds(epochSeconds)
// in the AceTime library and atc_processor_find_by_epoch_seconds() in the
// AceTimeC library.
func (zp *ZoneProcessor) FindByEpochSeconds(epochSeconds ATime) findResult {
	err := zp.InitForEpochSeconds(epochSeconds)
	if err != ErrOk {
		return findResultError
	}

	tfs := zp.tstorage.findTransitionForSeconds(epochSeconds)
	transition := tfs.curr
	if transition == nil {
		return findResultError
	}

	var frtype uint8
	if tfs.num == 2 {
		frtype = findResultOverlap
	} else {
		frtype = findResultExact
	}
	return findResult{
		frtype:              frtype,
		fold:                tfs.fold,
		stdOffsetSeconds:    transition.offsetSeconds,
		dstOffsetSeconds:    transition.deltaSeconds,
		reqStdOffsetSeconds: transition.offsetSeconds,
		reqDstOffsetSeconds: transition.deltaSeconds,
		abbrev:              transition.abbrev,
	}
}

// Return the findResult at the given LocalDateTime.
//
// Adapted from ExtendedZoneProcessor::findByLocalDateTime(const LocalDatetime&)
// in the AceTime library and atc_processor_find_by_local_date_time() in the
// AceTimeC library.
func (zp *ZoneProcessor) FindByLocalDateTime(ldt *LocalDateTime) findResult {

	err := zp.InitForYear(ldt.Year)
	if err != ErrOk {
		return findResultError
	}

	tfd := zp.tstorage.findTransitionForDateTime(ldt)

	// Extract the appropriate transition, depending on the requested 'fold'
	// and the 'tfd.searchStatus'.
	var transition *transition
	var result findResult
	if tfd.num == 1 {
		transition = tfd.curr
		result.frtype = findResultExact
		result.fold = 0
		result.reqStdOffsetSeconds = transition.offsetSeconds
		result.reqDstOffsetSeconds = transition.deltaSeconds
	} else { // num = 0 or 2
		if tfd.prev == nil || tfd.curr == nil {
			// ldt was far past or far future, and didn't match anything.
			transition = nil
			result.frtype = findResultNotFound
			result.fold = 0
		} else { // gap or overlap
			if tfd.num == 0 { // gap
				result.frtype = findResultGap
				result.fold = 0
				if ldt.Fold == 0 {
					// ldt wants to use the 'prev' transition to convert to
					// epochSeconds.
					result.reqStdOffsetSeconds = tfd.prev.offsetSeconds
					result.reqDstOffsetSeconds = tfd.prev.deltaSeconds
					// But after normalization, it will be shifted into the curr
					// transition, so select 'curr' as the target transition.
					transition = tfd.curr
				} else {
					// ldt wants to use the 'curr' transition to convert to
					// epochSeconds.
					result.reqStdOffsetSeconds = tfd.curr.offsetSeconds
					result.reqDstOffsetSeconds = tfd.curr.deltaSeconds
					// But after normalization, it will be shifted into the prev
					// transition, so select 'prev' as the target transition.
					transition = tfd.prev
				}
			} else {
				if ldt.Fold == 0 {
					transition = tfd.prev
				} else {
					transition = tfd.curr
				}
				result.frtype = findResultOverlap
				result.fold = ldt.Fold
				result.reqStdOffsetSeconds = transition.offsetSeconds
				result.reqDstOffsetSeconds = transition.deltaSeconds
			}
		}
	}

	if transition == nil {
		return findResultError
	}

	result.stdOffsetSeconds = transition.offsetSeconds
	result.dstOffsetSeconds = transition.deltaSeconds
	result.abbrev = transition.abbrev
	return result
}
