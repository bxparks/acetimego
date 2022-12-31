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

	// Use int32_t because a delta year of 2730 or greater will exceed
	// the range of an int16_t.
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

func fixTransitionTimes(ts []Transition) {
	if len(ts) == 0 {
		return
	}

	prev := &ts[0]
	for i := range ts {
		curr := &ts[i]
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
type MonthDay struct{
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

type Err int8

const (
	ErrOk = iota
  ErrGeneric
)

type ZoneProcessor struct {
	zoneInfo          *ZoneInfo
	year              int16
	isFilled          bool
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

func (zp *ZoneProcessor) InitForYear(
  zoneInfo *ZoneInfo, year int16) Err {

	return ErrOk
}

func (zp *ZoneProcessor) InitForEpochSeconds(
  zoneInfo *ZoneInfo, epochSeconds int32) Err {

	return ErrOk
}
