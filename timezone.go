package acetime

type DateTuple struct {
	/** [0,10000] */
	year int16

	/** [1-12] */
	month uint8

	/** [1-31] */
	day uint8

	/** negative values allowed */
	minutes int16

	/** suffixS, suffixW, suffixU */
	suffix uint8
}

// dateTupleCompare compare 2 DateTuple instances (a, b) and returns -1, 0, 1
// depending on whether a is less than, equal, or greater than b, respectively.
func dateTupleCompare(a *DateTuple, b *DateTuple) int8 {
	if a.year < b.year {
		return -1
	}
	if a.year > b.year {
		return 1
	}
	if a.month < b.month {
		return -1
	}
	if a.month > b.month {
		return 1
	}
	if a.day < b.day {
		return -1
	}
	if a.day > b.day {
		return 1
	}
	if a.minutes < b.minutes {
		return -1
	}
	if a.minutes > b.minutes {
		return 1
	}
	return 0
}

// dateTupleSubtract returns the number of seconds of (a - b).
func dateTupleSubtract(a *DateTuple, b *DateTuple) int32 {
	eda := LocalDateToEpochDays(a.year, a.month, a.day)
	esa := eda*86400 + int32(a.minutes)*60

	edb := LocalDateToEpochDays(b.year, b.month, b.day)
	esb := edb*86400 + int32(b.minutes)*60

	return esa - esb
}

func dateTupleNormalize(dt *DateTuple) {
	const oneDayAsMinutes = 60 * 24

	if dt.minutes <= -oneDayAsMinutes {
		dt.year, dt.month, dt.day = LocalDateDecrementOneDay(
			dt.year, dt.month, dt.day)
		dt.minutes += oneDayAsMinutes
	} else if oneDayAsMinutes <= dt.minutes {
		dt.year, dt.month, dt.day = LocalDateIncrementOneDay(
			dt.year, dt.month, dt.day)
		dt.minutes -= oneDayAsMinutes
	} else {
		// do nothing
	}
}

// dateTupleExpand converts the given 'tt', offsetMinutes, and deltaMinutes into
// the 'w', 's' and 'u' versions of the DateTuple. It is allowed for 'ttw' to
// be an alias of 'tt'.
func dateTupleExpand(
	tt *DateTuple,
	offsetMinutes int16,
	deltaMinutes int16,
	ttw *DateTuple,
	tts *DateTuple,
	ttu *DateTuple) {

	if tt.suffix == suffixS {
		*tts = *tt

		ttu.year = tt.year
		ttu.month = tt.month
		ttu.day = tt.day
		ttu.minutes = tt.minutes - offsetMinutes
		ttu.suffix = suffixU

		ttw.year = tt.year
		ttw.month = tt.month
		ttw.day = tt.day
		ttw.minutes = tt.minutes + deltaMinutes
		ttw.suffix = suffixW
	} else if tt.suffix == suffixU {
		*ttu = *tt

		tts.year = tt.year
		tts.month = tt.month
		tts.day = tt.day
		tts.minutes = tt.minutes + offsetMinutes
		tts.suffix = suffixS

		ttw.year = tt.year
		ttw.month = tt.month
		ttw.day = tt.day
		ttw.minutes = tt.minutes + (offsetMinutes + deltaMinutes)
		ttw.suffix = suffixW
	} else {
		// Explicit set the suffix to 'w' in case it was something else.
		*ttw = *tt
		ttw.suffix = suffixW

		tts.year = tt.year
		tts.month = tt.month
		tts.day = tt.day
		tts.minutes = tt.minutes - deltaMinutes
		tts.suffix = suffixS

		ttu.year = tt.year
		ttu.month = tt.month
		ttu.day = tt.day
		ttu.minutes = tt.minutes - (deltaMinutes + offsetMinutes)
		ttu.suffix = suffixU
	}

	dateTupleNormalize(ttw)
	dateTupleNormalize(tts)
	dateTupleNormalize(ttu)
}

const (
	matchStatusFarPast     = iota // 0
	matchStatusPrior              // 1
	matchStatusExactMatch         // 2
	matchStatusWithinMatch        // 3
	matchStatusFarFuture          // 4
)

// dateTupleCompareFuzzy compares the given 't' with the interval defined by
// [start, until). The comparison is fuzzy, with a slop of about one month so
// that we can ignore the day and minutes fields.
//
// The following values are returned:
//
//   - matchStatusPrior if 't' is less than 'start' by at least one month,
//   - matchStatusFarFuture if 't' is greater than 'until' by at least one
//     month,
//   - matchStatusWithinMatch if 't' is within [start, until) with a one
//     month slop,
//   - matchStatusExactMatch is never returned.
func dateTupleCompareFuzzy(
	t *DateTuple, start *DateTuple, until *DateTuple) uint8 {

	// Use int32 because a delta year of 2730 or greater will exceed
	// the range of an int16.
	var tMonths int32 = int32(t.year)*12 + int32(t.month)
	var startMonths int32 = int32(start.year)*12 + int32(start.month)
	if tMonths < startMonths-1 {
		return matchStatusPrior
	}
	var untilMonths int32 = int32(until.year)*12 + int32(until.month)
	if untilMonths+1 < tMonths {
		return matchStatusFarFuture
	}
	return matchStatusWithinMatch
}

//-----------------------------------------------------------------------------

type MatchingEra struct {
	/**
	 * The effective start time of the matching ZoneEra, which uses the
	 * UTC offsets of the previous matching era.
	 */
	startDt DateTuple

	/** The effective until time of the matching ZoneEra. */
	untilDt DateTuple

	/** The ZoneEra that matched the given year. NonNullable. */
	era *ZoneEra

	/** The previous MatchingEra, needed to interpret startDt.  */
	prevMatch *MatchingEra

	/** The STD offset of the last Transition in this MatchingEra. */
	lastOffsetMinutes int16

	/** The DST offset of the last Transition in this MatchingEra. */
	lastDeltaMinutes int16
}

//-----------------------------------------------------------------------------

type Transition struct {
	/** The MatchingEra which generated this Transition. */
	match *MatchingEra

	/**
	 * The Zone transition rule that matched for the the given year. Set to
	 * nullptr if the RULES column is '-', indicating that the MatchingEra was
	 * a "simple" ZoneEra.
	 */
	rule *ZoneRule

	/**
	 * The original transition time, usually 'w' but sometimes 's' or 'u'. After
	 * expandDateTuple() is called, this field will definitely be a 'w'. We must
	 * remember that the transitionTime* fields are expressed using the UTC
	 * offset of the *previous* Transition.
	 */
	transitionTime DateTuple

	//union {
	/**
	* Version of transitionTime in 's' mode, using the UTC offset of the
	* *previous* Transition. Valid before
	* ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	transitionTimeS DateTuple

	/**
	* Start time expressed using the UTC offset of the current Transition.
	* Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	startDt DateTuple
	//}

	//union {
	/**
	* Version of transitionTime in 'u' mode, using the UTC offset of the
	* *previous* transition. Valid before
	* ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	transitionTimeU DateTuple

	/**
	* Until time expressed using the UTC offset of the current Transition.
	* Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	untilDt DateTuple
	//}

	/** The calculated transition time of the given rule. */
	startEpochSeconds int32

	/**
	 * The base offset minutes, not the total effective UTC offset. Note that
	 * this is different than basic::Transition::offsetMinutes used by
	 * BasicZoneProcessor which is the total effective offsetMinutes. (It may be
	 * possible to make this into an effective offsetMinutes (i.e. offsetMinutes
	 * + deltaMinutes) but it does not seem worth making that change right now.)
	 */
	offsetMinutes int16

	/** The DST delta minutes. */
	deltaMinutes int16

	/** The calculated effective time zone abbreviation, e.g. "PST" or "PDT". */
	abbrev string

	/** Storage for the single letter 'letter' field if 'rule' is not null. */
	letter string

	/**
	 * During findCandidateTransitions(), this flag indicates whether the
	 * current transition is a valid "prior" transition that occurs before other
	 * transitions. It is set by setFreeAgentAsPriorIfValid() if the transition
	 * is a prior transition.
	 */
	isValidPrior bool

	/**
	 * During processTransitionMatchStatus(), this flag indicates how the
	 * transition falls within the time interval of the MatchingEra.
	 */
	matchStatus uint8
}

//-----------------------------------------------------------------------------

const (
	maxAbbrevSize         = 6
	transitionStorageSize = 8
	maxMatches            = 4
	maxInteriorYears      = 4
)

type TransitionStorage struct {
	/** Index of the most recent prior transition [0,kTransitionStorageSize) */
	indexPrior uint8
	/** Index of the candidate pool [0,kTransitionStorageSize) */
	indexCandidate uint8
	/** Index of the free agent transition [0, kTransitionStorageSize) */
	indexFree uint8
	/** Number of allocated transitions. */
	allocSize uint8

	/** Pool of Transition objects. */
	transitions [transitionStorageSize]Transition
}

func (ts *TransitionStorage) Init() {
	ts.indexPrior = 0
	ts.indexCandidate = 0
	ts.indexFree = 0
	ts.allocSize = 0
}

func (ts *TransitionStorage) ResetCandidatePool() {
	ts.indexCandidate = ts.indexPrior
	ts.indexFree = ts.indexPrior
}

func (ts *TransitionStorage) GetFreeAgent() *Transition {
	if ts.indexFree < transitionStorageSize {
		if ts.indexFree >= ts.allocSize {
			ts.allocSize = ts.indexFree + 1
		}
		return &ts.transitions[ts.indexFree]
	} else {
		return &ts.transitions[transitionStorageSize-1]
	}
}

func (ts *TransitionStorage) AddFreeAgentToActivePool() {
	if ts.indexFree >= transitionStorageSize {
		return
	}
	ts.indexFree++
	ts.indexPrior = ts.indexFree
	ts.indexCandidate = ts.indexFree
}

func (ts *TransitionStorage) ReservePrior() *Transition {
	ts.GetFreeAgent()
	ts.indexCandidate++
	ts.indexFree++
	return &ts.transitions[ts.indexPrior]
}

func (ts *TransitionStorage) SetFreeAgentAsPriorIfValid() {
	ft := &ts.transitions[ts.indexFree]
	prior := &ts.transitions[ts.indexPrior]
	if (prior.isValidPrior && dateTupleCompare(
		&prior.transitionTime,
		&ft.transitionTime) < 0) || !prior.isValidPrior {

		ft.isValidPrior = true
		prior.isValidPrior = false

		// swap(prior, free)
		*ft, *prior = *prior, *ft
	}
}

func (ts *TransitionStorage) AddPriorToCandidatePool() {
	ts.indexCandidate--
}

func (ts *TransitionStorage) AddFreeAgentToCandidatePool() {
	if ts.indexFree >= transitionStorageSize {
		return
	}
	for i := ts.indexFree; i > ts.indexCandidate; i-- {
		curr := &ts.transitions[i]
		prev := &ts.transitions[i-1]
		if dateTupleCompare(&curr.transitionTime, &prev.transitionTime) >= 0 {
			break
		}
		*curr, *prev = *prev, *curr
	}
	ts.indexFree++
}

func isMatchStatusActive(status uint8) bool {
	return status == matchStatusExactMatch ||
		status == matchStatusWithinMatch ||
		status == matchStatusPrior
}

// AddActiveCandidatesToActivePool adds the candidate transitions to the active
// pool, and returns the last active transition added.
func (ts *TransitionStorage) AddActiveCandidatesToActivePool() *Transition {
	// Shift active candidates to the left into the Active pool.
	iActive := ts.indexPrior
	iCandidate := ts.indexCandidate
	for ; iCandidate < ts.indexFree; iCandidate++ {
		if isMatchStatusActive(ts.transitions[iCandidate].matchStatus) {
			if iActive != iCandidate {
				// Shift candidate into active slot
				ts.transitions[iActive] = ts.transitions[iCandidate]
			}
			iActive++
		}
	}

	ts.indexPrior = iActive
	ts.indexCandidate = iActive
	ts.indexFree = iActive

	return &ts.transitions[iActive-1]
}

func fixTransitionTimes(transitions []Transition) {
	if len(transitions) == 0 {
		return
	}

	prev := &transitions[0]
	for i := range transitions {
		curr := &transitions[i]
		dateTupleExpand(
			&curr.transitionTime,
			prev.offsetMinutes,
			prev.deltaMinutes,
			&curr.transitionTime,
			&curr.transitionTimeS,
			&curr.transitionTimeU)
		prev = curr
	}
}

//-----------------------------------------------------------------------------

func compareTransitionToMatch(t *Transition, match *MatchingEra) uint8 {
	// Find the previous Match offsets.
	var prevMatchOffsetMinutes int16
	var prevMatchDeltaMinutes int16
	if match.prevMatch != nil {
		prevMatchOffsetMinutes = match.prevMatch.lastOffsetMinutes
		prevMatchDeltaMinutes = match.prevMatch.lastDeltaMinutes
	} else {
		prevMatchOffsetMinutes = match.era.StdOffsetMinutes()
		prevMatchDeltaMinutes = 0
	}

	// Expand start times.
	var stw DateTuple
	var sts DateTuple
	var stu DateTuple
	dateTupleExpand(
		&match.startDt,
		prevMatchOffsetMinutes,
		prevMatchDeltaMinutes,
		&stw,
		&sts,
		&stu)

	// Transition times.
	ttw := &t.transitionTime
	tts := &t.transitionTimeS
	ttu := &t.transitionTimeU

	// Compare Transition to Match, where equality is assumed if *any* of the
	// 'w', 's', or 'u' versions of the DateTuple are equal. This prevents
	// duplicate Transition instances from being created in a few cases.
	if dateTupleCompare(ttw, &stw) == 0 ||
		dateTupleCompare(tts, &sts) == 0 ||
		dateTupleCompare(ttu, &stu) == 0 {
		return matchStatusExactMatch
	}

	if dateTupleCompare(ttu, &stu) < 0 {
		return matchStatusPrior
	}

	// Now check if the transition occurs after the given match. The
	// untilDateTime of the current match uses the same UTC offsets as the
	// transitionTime of the current transition, so no complicated adjustments
	// are needed. We just make sure we compare 'w' with 'w', 's' with 's',
	// and 'u' with 'u'.
	matchUntil := &match.untilDt
	var transitionTime *DateTuple
	if matchUntil.suffix == suffixS {
		transitionTime = tts
	} else if matchUntil.suffix == suffixU {
		transitionTime = ttu
	} else { // assume 'w'
		transitionTime = ttw
	}
	if dateTupleCompare(transitionTime, matchUntil) < 0 {
		return matchStatusWithinMatch
	}

	return matchStatusFarFuture
}

func compareTransitionToMatchFuzzy(t *Transition, m *MatchingEra) uint8 {
	return dateTupleCompareFuzzy(&t.transitionTime, &m.startDt, &m.untilDt)
}

//---------------------------------------------------------------------------

/** A tuple of month and day. */
type MonthDay struct {
	/** month [1,12] */
	month uint8

	/** day [1,31] */
	day uint8
}

// calcStartDayOfMonth Extracts the actual (month, day) pair from the expression
// used in the TZ data files of the form (onDayOfWeek >= onDayOfMonth) or
// (onDayOfWeek <= onDayOfMonth).
//
// There are 4 combinations:
//
// @verbatim
// onDayOfWeek=0, onDayOfMonth=(1-31): exact match
// onDayOfWeek=1-7, onDayOfMonth=1-31: dayOfWeek>=dayOfMonth
// onDayOfWeek=1-7, onDayOfMonth=0: last{dayOfWeek}
// onDayOfWeek=1-7, onDayOfMonth=-(1-31): dayOfWeek<=dayOfMonth
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
	ErrOk = iota
	ErrGeneric
)

type ZoneProcessor struct {
	zoneInfo          *ZoneInfo
	year              int16
	isFilled          bool
	numMatches        uint8
	numTransitions    uint8
	matches           [maxMatches]MatchingEra
	transitionStorage TransitionStorage
}

// ZoneProcessorFromZoneInfo creates a new ZoneProcessor from the given ZoneInfo
// instance.
func ZoneProcessorFromZoneInfo(zoneInfo *ZoneInfo) *ZoneProcessor {
	return &ZoneProcessor{zoneInfo: zoneInfo}
}

func (zp *ZoneProcessor) isFilledForYear(year int16) bool {
	return zp.isFilled && (year == zp.year)
}

func (zp *ZoneProcessor) InitForYear(zoneInfo *ZoneInfo, year int16) Err {
	if zp.isFilledForYear(year) {
		return ErrOk
	}

	zp.year = year
	zp.numMatches = 0
	zp.transitionStorage.Init()
	if year < zoneInfo.startYear-1 || zoneInfo.untilYear < year {
		return ErrGeneric
	}

	startYm := YearMonth{year - 1, 12}
	untilYm := YearMonth{year + 1, 2}

	// Step 1: Find matches.
	zp.numMatches = findMatches(zp.zoneInfo, startYm, untilYm, zp.matches[:])

	// Step 2: Create Transitions.
	createTransitions(&zp.transitionStorage, zp.matches[:zp.numMatches])

	// Step 3: Fix transition times.
	ts := &zp.transitionStorage
	transitions := ts.transitions[0:ts.indexFree]
	fixTransitionTimes(transitions)

	// Step 4: Generate start and until times.
	generateStartUntilTimes(transitions)

	// Step 5: Calc abbreviations.
	calcAbbreviations(transitions)

	return ErrOk
}

func (zp *ZoneProcessor) InitForEpochSeconds(
	zoneInfo *ZoneInfo, epochSeconds int32) Err {

	return ErrOk
}

//-----------------------------------------------------------------------------
// Step 1
//-----------------------------------------------------------------------------

func findMatches(
	zoneInfo *ZoneInfo,
	startYm YearMonth,
	untilYm YearMonth,
	matches []MatchingEra) uint8 {

	var iMatch uint8 = 0
	var prevMatch *MatchingEra = nil
	var numEras uint8 = zoneInfo.numEras()

	var iEra uint8
	for iEra = 0; iEra < numEras; iEra++ {
		era := &zoneInfo.eras[iEra]
		var prevEra *ZoneEra = nil
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
	prevEra *ZoneEra,
	era *ZoneEra,
	startYm YearMonth,
	untilYm YearMonth) bool {

	return (prevEra == nil ||
		compareEraToYearMonth(prevEra, untilYm.year, untilYm.month) < 0) &&
		compareEraToYearMonth(era, startYm.year, startYm.month) > 0
}

/** Return (1, 0, -1) depending on how era compares to (year, month). */
func compareEraToYearMonth(era *ZoneEra, year int16, month uint8) int8 {
	if era.untilYear < year {
		return -1
	}
	if era.untilYear > year {
		return 1
	}
	if era.untilMonth < month {
		return -1
	}
	if era.untilMonth > month {
		return 1
	}
	if era.untilDay > 1 {
		return 1
	}
	//if era.until_time_minutes < 0 { return -1; // never possible
	if era.untilTimeCode > 0 {
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
	era *ZoneEra,
	startYm YearMonth,
	untilYm YearMonth) {

	// If prevMatch is nil, set startDate to be earlier than all valid ZoneEra.
	var startDate DateTuple
	if prevMatch == nil {
		startDate.year = InvalidYear
		startDate.month = 1
		startDate.day = 1
		startDate.minutes = 0
		startDate.suffix = suffixW
	} else {
		startDate.year = prevMatch.era.untilYear
		startDate.month = prevMatch.era.untilMonth
		startDate.day = prevMatch.era.untilDay
		startDate.minutes = prevMatch.era.UntilMinutes()
		startDate.suffix = prevMatch.era.UntilSuffix()
	}
	lowerBound := DateTuple{startYm.year, startYm.month, 1, 0, suffixW}
	if dateTupleCompare(&startDate, &lowerBound) < 0 {
		startDate = lowerBound
	}

	untilDate := DateTuple{
		era.untilYear,
		era.untilMonth,
		era.untilDay,
		era.UntilMinutes(),
		era.UntilSuffix(),
	}
	upperBound := DateTuple{untilYm.year, untilYm.month, 1, 0, suffixW}
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
	policy := match.era.zonePolicy
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
	t *Transition, year int16, rule *ZoneRule, match *MatchingEra) {

	t.match = match
	t.rule = rule
	t.offsetMinutes = match.era.StdOffsetMinutes()
	t.letter = ""

	if rule != nil {
		t.transitionTime = getTransitionTime(year, rule)
		t.deltaMinutes = rule.DstOffsetMinutes()
		// If LETTER is a '-', treat it the same as an empty string.
		if rule.letter != "-" {
			t.letter = rule.letter
		}
	} else {
		// Create a Transition using the MatchingEra for the transitionTime.
		// Used for simple MatchingEra.
		t.transitionTime = match.startDt
		t.deltaMinutes = match.era.DstOffsetMinutes()
	}
}

func getTransitionTime(year int16, rule *ZoneRule) DateTuple {
	md := calcStartDayOfMonth(
		year, rule.inMonth, rule.onDayOfWeek, rule.onDayOfMonth)
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
	transitions := ts.transitions[ts.indexCandidate:ts.indexFree]
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
	policy := match.era.zonePolicy
	startYear := match.startDt.year
	endYear := match.untilDt.year

	prior := ts.ReservePrior()
	prior.isValidPrior = false
	for ir := range policy.rules {
		rule := &policy.rules[ir]

		// Add transitions for interior years
		var interiorYears [maxInteriorYears]int16
		numYears := calcInteriorYears(
			interiorYears[:], rule.fromYear, rule.toYear, startYear, endYear)
		var iy uint8
		for iy = 0; iy < numYears; iy++ {
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
		priorYear := getMostRecentPriorYear(rule.fromYear, rule.toYear, startYear)
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

// calcInteriorYears calculates the interior years that overlaps (fromYear,
// toYear) and (startYear, endYear). The results are placed into the
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
func generateStartUntilTimes(transitions []Transition) {
}

//-----------------------------------------------------------------------------

// Step 5
func calcAbbreviations(transitions []Transition) {
}
