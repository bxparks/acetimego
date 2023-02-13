package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

//-----------------------------------------------------------------------------
// MatchingEra
//-----------------------------------------------------------------------------

type MatchingEra struct {
	// The effective start time of the matching ZoneEra, which uses the
	// UTC offsets of the previous matching era.
	startDt DateTuple

	// The effective until time of the matching ZoneEra.
	untilDt DateTuple

	// The ZoneEra that matched the given year. NonNullable.
	era *zoneinfo.ZoneEra

	// The previous MatchingEra, needed to interpret startDt.
	prevMatch *MatchingEra

	// The STD offset of the last Transition in this MatchingEra.
	lastOffsetSeconds int32

	// The DST offset of the last Transition in this MatchingEra.
	lastDeltaSeconds int32

	// The format string from era.FormatIndex
	format string
}

//-----------------------------------------------------------------------------
// Transition
//-----------------------------------------------------------------------------

type Transition struct {
	// The MatchingEra which generated this Transition.
	match *MatchingEra

	// The original transition time, usually 'w' but sometimes 's' or 'u'. After
	// expandDateTuple() is called, this field will definitely be a 'w'. We must
	// remember that the transitionTime* fields are expressed using the UTC
	// offset of the *previous* Transition.
	transitionTime DateTuple

	//union {

	// Version of transitionTime in 's' mode, using the UTC offset of the
	// *previous* Transition. Valid before
	// ExtendedZoneProcessor::generateStartUntilTimes() is called.
	transitionTimeS DateTuple

	// Start time expressed using the UTC offset of the current Transition.
	// Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	startDt DateTuple

	//}

	//union {

	// Version of transitionTime in 'u' mode, using the UTC offset of the
	// *previous* transition. Valid before
	// ExtendedZoneProcessor::generateStartUntilTimes() is called.
	transitionTimeU DateTuple

	// Until time expressed using the UTC offset of the current Transition.
	// Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	untilDt DateTuple

	//}

	// The calculated transition time of the given rule.
	startEpochSeconds ATime

	// The base offset seconds, not the total effective UTC offset.
	offsetSeconds int32

	// The DST delta seconds.
	deltaSeconds int32

	// The calculated effective time zone abbreviation, e.g. "PST" or "PDT".
	abbrev string

	// letter field copied from matching rule if not null.
	letter string

	// During findCandidateTransitions(), this flag indicates whether the
	// current transition is a valid "prior" transition that occurs before other
	// transitions. It is set by setFreeAgentAsPriorIfValid() if the transition
	// is a prior transition.
	isValidPrior bool

	// During processTransitionMatchStatus(), this flag indicates how the
	// transition falls within the time interval of the MatchingEra.
	matchStatus uint8
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
			prev.offsetSeconds,
			prev.deltaSeconds,
			&curr.transitionTime,
			&curr.transitionTimeS,
			&curr.transitionTimeU)
		prev = curr
	}
}

//-----------------------------------------------------------------------------

const (
	transitionStorageSize = 8
)

// TransitionStorage holds 4 pools of Transitions indicated by the following
// half-open (inclusive to exclusive) index ranges:
//
//  1. Active pool: [0, indexPrior)
//  2. Prior pool: [indexPrior, indexCandidate), either 0 or 1 element
//  3. Candidate pool: [indexCandidate, indexFree)
//  4. Free agent pool: [indexFree, allocSize), 0 or 1 element
type TransitionStorage struct {
	// Index of the most recent prior transition [0,kTransitionStorageSize)
	indexPrior uint8
	// Index of the candidate pool [0,kTransitionStorageSize)
	indexCandidate uint8
	// Index of the free agent transition [0, kTransitionStorageSize)
	indexFree uint8
	// Number of allocated transitions.
	allocSize uint8
	// Pool of Transition objects.
	transitions [transitionStorageSize]Transition
}

func (ts *TransitionStorage) init() {
	ts.indexPrior = 0
	ts.indexCandidate = 0
	ts.indexFree = 0
	ts.allocSize = 0
}

// getActives returns the active transitions in the interval [0,indexPrior).
func (ts *TransitionStorage) getActives() []Transition {
	return ts.transitions[0:ts.indexPrior]
}

// getCandidates returns the candidate transitions in the interval
// [indexCandidate,indexFree).
func (ts *TransitionStorage) getCandidates() []Transition {
	return ts.transitions[ts.indexCandidate:ts.indexFree]
}

// resetCandidatePool deletes the candidate pool by collapsing it into the prior
// pool.
func (ts *TransitionStorage) resetCandidatePool() {
	ts.indexCandidate = ts.indexPrior
	ts.indexFree = ts.indexPrior
}

func (ts *TransitionStorage) getFreeAgent() *Transition {
	if ts.indexFree < transitionStorageSize {
		if ts.indexFree >= ts.allocSize {
			ts.allocSize = ts.indexFree + 1
		}
		return &ts.transitions[ts.indexFree]
	} else {
		return &ts.transitions[transitionStorageSize-1]
	}
}

func (ts *TransitionStorage) addFreeAgentToActivePool() {
	if ts.indexFree >= transitionStorageSize {
		return
	}
	ts.indexFree++
	ts.indexPrior = ts.indexFree
	ts.indexCandidate = ts.indexFree
}

func (ts *TransitionStorage) reservePrior() *Transition {
	ts.getFreeAgent()
	ts.indexCandidate++
	ts.indexFree++
	return &ts.transitions[ts.indexPrior]
}

func (ts *TransitionStorage) setFreeAgentAsPriorIfValid() {
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

func (ts *TransitionStorage) addPriorToCandidatePool() {
	// This simple decrement works because there is only one prior, and it is
	// allocated just before the candidate pool.
	ts.indexCandidate--
}

func (ts *TransitionStorage) addFreeAgentToCandidatePool() {
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

// Useful for debugging, commented out instead of deleting.
//
//func (ts *TransitionStorage) printPoolSizes() {
//	fmt.Printf("indexPrior=%d; indexCandidate=%d; indexFree=%d; allocSize=%d\n",
//		ts.indexPrior, ts.indexCandidate, ts.indexFree, ts.allocSize)
//}

// addActiveCandidatesToActivePool adds the candidate transitions to the active
// pool, and returns the last active transition added.
func (ts *TransitionStorage) addActiveCandidatesToActivePool() *Transition {
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

//-----------------------------------------------------------------------------

type TransitionForSeconds struct {
	// The matching Transition, nil if not found.
	curr *Transition

	// 0 for the first or exact transition; 1 for the second transition
	fold uint8

	// Number of occurrences of the resulting LocalDateTime: 0, 1, or 2. This is
	// needed because a fold=0 can mean that the LocalDateTime occurs exactly
	// once, or that the first of two occurrences of LocalDateTime was selected by
	// the epochSeconds.
	num uint8
}

func (ts *TransitionStorage) findTransitionForSeconds(
	epochSeconds ATime) TransitionForSeconds {

	var prev *Transition = nil
	var curr *Transition = nil
	var next *Transition = nil

	transitions := ts.getActives()
	for i := range transitions {
		next = &transitions[i] // do not use := here (bitten twice by this bug)
		if next.startEpochSeconds > epochSeconds {
			break
		}
		prev = curr
		curr = next
		next = nil // clear 'next' in case we roll off the array
	}

	fold, num := calculateFoldAndOverlap(epochSeconds, prev, curr, next)
	return TransitionForSeconds{curr, fold, num}
}

// calculateFold determines the 'fold' parameter at the given epochSeconds. This
// will become the output parameter of the corresponding LocalDateTime. A 0
// indicates that the LocalDateTime was the first ocurrence. A 1 indicates a
// LocalDateTime that occurred a second time.
func calculateFoldAndOverlap(
	epochSeconds ATime,
	prev *Transition,
	curr *Transition,
	next *Transition) (fold uint8, num uint8) {

	if curr == nil {
		fold = 0
		num = 0
		return
	}

	// Check if within forward overlap shadow from prev
	var isOverlap bool
	if prev == nil {
		isOverlap = false
	} else {
		shiftSeconds := dateTupleSubtract(&curr.startDt, &prev.untilDt)
		if shiftSeconds >= 0 {
			// sprint forward, or unchanged
			isOverlap = false
		} else {
			isOverlap = epochSeconds-curr.startEpochSeconds < -shiftSeconds
		}
	}
	if isOverlap {
		// epochSeconds selects the second match
		fold = 1
		num = 2
		return
	}

	// Check if within backward overlap shawdow from next
	if next == nil {
		isOverlap = false
	} else {
		// Extract the shift to next transition. Can be 0 in some cases where
		// the zone changed from DST of one zone to the STD into another zone,
		// causing the overall UTC offset to remain unchanged.
		shiftSeconds := dateTupleSubtract(&next.startDt, &curr.untilDt)
		if shiftSeconds >= 0 {
			// spring forward, or unchanged
			isOverlap = false
		} else {
			// Check if within the backward overlap shadow from next
			delta := next.startEpochSeconds - epochSeconds
			isOverlap = delta <= -shiftSeconds
		}
	}
	if isOverlap {
		// epochSeconds selects the first match
		fold = 0
		num = 2
		return
	}

	// Normal single match, no overlap
	fold = 0
	num = 1
	return
}

//-----------------------------------------------------------------------------

// The result returned by findTransitionForDateTime(). There are 5
// possibilities:
//
//   - num=0, prev==NULL, curr=curr: datetime is far past
//   - num=1, prev==prev, curr=prev: exact match to datetime
//   - num=2, prev==prev, curr=curr: datetime in overlap
//   - num=0, prev==prev, curr=curr: datetime in gap
//   - num=0, prev==prev, curr=NULL: datetime is far future
//
// Adapted from TransitionForDateTime in Transition.h of the AceTime library,
// and transition.h from the AceTimeC library.
type TransitionForDateTime struct {
	// The previous transition, or null if the first transition matches.
	prev *Transition

	// The matching transition or null if not found.
	curr *Transition

	// Number of matches for given LocalDateTime: 0, 1, or 2.
	num uint8
}

func (ts *TransitionStorage) findTransitionForDateTime(
	ldt *LocalDateTime) TransitionForDateTime {

	// Convert LocalDateTime to DateTuple.
	localDt := DateTuple{
		ldt.Year,
		ldt.Month,
		ldt.Day,
		int32(ldt.Hour)*60*60 + int32(ldt.Minute)*60,
		zoneinfo.SuffixW,
	}

	// Examine adjacent pairs of Transitions, looking for an exact match, gap,
	// or overlap.
	var prev *Transition = nil
	var curr *Transition = nil
	var num uint8 = 0
	transitions := ts.getActives()
	for i := range transitions {
		curr = &ts.transitions[i]

		startDt := &curr.startDt
		untilDt := &curr.untilDt
		isExactMatch := dateTupleCompare(startDt, &localDt) <= 0 &&
			dateTupleCompare(&localDt, untilDt) < 0

		if isExactMatch {
			// Check for a previous exact match to detect an overlap.
			if num == 1 {
				num++
				break
			}

			// Loop again to detect an overlap.
			num = 1
		} else if dateTupleCompare(startDt, &localDt) > 0 {
			// Exit loop since no more candidate transition.
			break
		}

		prev = curr

		// Set nil so that if the loop runs off the end of the list of Transitions,
		// curr is marked as nullptr.
		curr = nil
	}

	// Check if the prev was an exact match, and clear the current to
	// avoid confusion.
	if num == 1 {
		curr = prev
	}

	return TransitionForDateTime{prev, curr, num}
}

//-----------------------------------------------------------------------------

func compareTransitionToMatch(t *Transition, match *MatchingEra) uint8 {
	// Find the previous Match offsets.
	var prevMatchOffsetSeconds int32
	var prevMatchDeltaSeconds int32
	if match.prevMatch != nil {
		prevMatchOffsetSeconds = match.prevMatch.lastOffsetSeconds
		prevMatchDeltaSeconds = match.prevMatch.lastDeltaSeconds
	} else {
		prevMatchOffsetSeconds = match.era.StdOffsetSeconds()
		prevMatchDeltaSeconds = 0
	}

	// Expand start times.
	var stw DateTuple
	var sts DateTuple
	var stu DateTuple
	dateTupleExpand(
		&match.startDt,
		prevMatchOffsetSeconds,
		prevMatchDeltaSeconds,
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
	if matchUntil.suffix == zoneinfo.SuffixS {
		transitionTime = tts
	} else if matchUntil.suffix == zoneinfo.SuffixU {
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
