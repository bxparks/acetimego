package acetime

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

//-----------------------------------------------------------------------------
// matchingEra
//-----------------------------------------------------------------------------

type matchingEra struct {
	// The effective start time of the matching ZoneEra, which uses the
	// UTC offsets of the previous matching era.
	startDt dateTuple

	// The effective until time of the matching ZoneEra.
	untilDt dateTuple

	// The ZoneEra that matched the given year. NonNullable.
	era *zoneinfo.ZoneEra

	// The previous matchingEra, needed to interpret startDt.
	prevMatch *matchingEra

	// The STD offset of the last transition in this matchingEra.
	lastOffsetSeconds int32

	// The DST offset of the last transition in this matchingEra.
	lastDeltaSeconds int32

	// The format string from era.FormatIndex
	format string
}

//-----------------------------------------------------------------------------
// transition
//-----------------------------------------------------------------------------

type transition struct {
	// The matchingEra which generated this transition.
	match *matchingEra

	// The original transition time, usually 'w' but sometimes 's' or 'u'. After
	// expandDateTuple() is called, this field will definitely be a 'w'. We must
	// remember that the transitionTime* fields are expressed using the UTC
	// offset of the *previous* transition.
	transitionTime dateTuple

	//union {

	// Version of transitionTime in 's' mode, using the UTC offset of the
	// *previous* transition. Valid before
	// ExtendedZoneProcessor::generateStartUntilTimes() is called.
	transitionTimeS dateTuple

	// Start time expressed using the UTC offset of the current transition.
	// Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	startDt dateTuple

	//}

	//union {

	// Version of transitionTime in 'u' mode, using the UTC offset of the
	// *previous* transition. Valid before
	// ExtendedZoneProcessor::generateStartUntilTimes() is called.
	transitionTimeU dateTuple

	// Until time expressed using the UTC offset of the current transition.
	// Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	untilDt dateTuple

	//}

	// The calculated transition time of the given rule.
	startEpochSeconds Time

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

	// During processTransitionCompareStatus(), this flag indicates how the
	// transition falls within the time interval of the matchingEra.
	compareStatus uint8
}

func fixTransitionTimes(transitions []transition) {
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

// transitionStorage holds 4 pools of Transitions indicated by the following
// half-open (inclusive to exclusive) index ranges:
//
//  1. Active pool: [0, indexPrior)
//  2. Prior pool: [indexPrior, indexCandidate), either 0 or 1 element
//  3. Candidate pool: [indexCandidate, indexFree)
//  4. Free agent pool: [indexFree, allocSize), 0 or 1 element
type transitionStorage struct {
	// Index of the most recent prior transition [0,transitionStorageSize)
	indexPrior uint8
	// Index of the candidate pool [0,transitionStorageSize)
	indexCandidate uint8
	// Index of the free agent transition [0, transitionStorageSize)
	indexFree uint8
	// Number of allocated transitions.
	allocSize uint8
	// Pool of transition objects.
	transitions [transitionStorageSize]transition
}

func (ts *transitionStorage) init() {
	ts.indexPrior = 0
	ts.indexCandidate = 0
	ts.indexFree = 0
	ts.allocSize = 0
}

// getActives returns the active transitions in the interval [0,indexPrior).
func (ts *transitionStorage) getActives() []transition {
	return ts.transitions[0:ts.indexPrior]
}

// getCandidates returns the candidate transitions in the interval
// [indexCandidate,indexFree).
func (ts *transitionStorage) getCandidates() []transition {
	return ts.transitions[ts.indexCandidate:ts.indexFree]
}

// resetCandidatePool deletes the candidate pool by collapsing it into the prior
// pool.
func (ts *transitionStorage) resetCandidatePool() {
	ts.indexCandidate = ts.indexPrior
	ts.indexFree = ts.indexPrior
}

func (ts *transitionStorage) getFreeAgent() *transition {
	if ts.indexFree < transitionStorageSize {
		if ts.indexFree >= ts.allocSize {
			ts.allocSize = ts.indexFree + 1
		}
		return &ts.transitions[ts.indexFree]
	} else {
		return &ts.transitions[transitionStorageSize-1]
	}
}

func (ts *transitionStorage) addFreeAgentToActivePool() {
	if ts.indexFree >= transitionStorageSize {
		return
	}
	ts.indexFree++
	ts.indexPrior = ts.indexFree
	ts.indexCandidate = ts.indexFree
}

func (ts *transitionStorage) reservePrior() *transition {
	ts.getFreeAgent()
	ts.indexCandidate++
	ts.indexFree++
	return &ts.transitions[ts.indexPrior]
}

func (ts *transitionStorage) setFreeAgentAsPriorIfValid() {
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

func (ts *transitionStorage) addPriorToCandidatePool() {
	// This simple decrement works because there is only one prior, and it is
	// allocated just before the candidate pool.
	ts.indexCandidate--
}

func (ts *transitionStorage) addFreeAgentToCandidatePool() {
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

func isCompareStatusActive(status uint8) bool {
	return status == compareStatusExactMatch ||
		status == compareStatusWithinMatch ||
		status == compareStatusPrior
}

// Useful for debugging, commented out instead of deleting.
//
//func (ts *transitionStorage) printPoolSizes() {
//	fmt.Printf("indexPrior=%d; indexCandidate=%d; indexFree=%d; allocSize=%d\n",
//		ts.indexPrior, ts.indexCandidate, ts.indexFree, ts.allocSize)
//}

// addActiveCandidatesToActivePool adds the candidate transitions to the active
// pool, and returns the last active transition added.
func (ts *transitionStorage) addActiveCandidatesToActivePool() *transition {
	// Shift active candidates to the left into the Active pool.
	iActive := ts.indexPrior
	iCandidate := ts.indexCandidate
	for ; iCandidate < ts.indexFree; iCandidate++ {
		if isCompareStatusActive(ts.transitions[iCandidate].compareStatus) {
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

type transitionForSeconds struct {
	// The matching transition, nil if not found.
	curr *transition

	// 0 for the first or exact transition; 1 for the second transition
	fold uint8

	// Number of occurrences of the resulting LocalDateTime: 0, 1, or 2. This is
	// needed because a fold=0 can mean that the LocalDateTime occurs exactly
	// once, or that the first of two occurrences of LocalDateTime was selected by
	// the epochSeconds.
	num uint8
}

func (ts *transitionStorage) findTransitionForSeconds(
	epochSeconds Time) transitionForSeconds {

	var prev *transition = nil
	var curr *transition = nil
	var next *transition = nil

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
	return transitionForSeconds{curr, fold, num}
}

// calculateFold determines the 'fold' parameter at the given epochSeconds. This
// will become the output parameter of the corresponding LocalDateTime. A 0
// indicates that the LocalDateTime was the first ocurrence. A 1 indicates a
// LocalDateTime that occurred a second time.
func calculateFoldAndOverlap(
	epochSeconds Time,
	prev *transition,
	curr *transition,
	next *transition) (fold uint8, num uint8) {

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
// Adapted from transitionForDateTime in transition.h of the AceTime library,
// and transition.h from the acetimec library.
type transitionForDateTime struct {
	// The previous transition, or null if the first transition matches.
	prev *transition

	// The matching transition or null if not found.
	curr *transition

	// Number of matches for given LocalDateTime: 0, 1, or 2.
	num uint8
}

func (ts *transitionStorage) findTransitionForDateTime(
	ldt *LocalDateTime) transitionForDateTime {

	// Convert LocalDateTime to dateTuple.
	localDt := dateTuple{
		ldt.Year,
		ldt.Month,
		ldt.Day,
		int32(ldt.Hour)*60*60 + int32(ldt.Minute)*60,
		zoneinfo.SuffixW,
	}

	// Examine adjacent pairs of Transitions, looking for an exact match, gap,
	// or overlap.
	var prev *transition = nil
	var curr *transition = nil
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

	return transitionForDateTime{prev, curr, num}
}

//-----------------------------------------------------------------------------

func compareTransitionToMatch(t *transition, match *matchingEra) uint8 {
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
	var stw dateTuple
	var sts dateTuple
	var stu dateTuple
	dateTupleExpand(
		&match.startDt,
		prevMatchOffsetSeconds,
		prevMatchDeltaSeconds,
		&stw,
		&sts,
		&stu)

	// transition times.
	ttw := &t.transitionTime
	tts := &t.transitionTimeS
	ttu := &t.transitionTimeU

	// Compare transition to Match, where equality is assumed if *any* of the
	// 'w', 's', or 'u' versions of the dateTuple are equal. This prevents
	// duplicate transition instances from being created in a few cases.
	if dateTupleCompare(ttw, &stw) == 0 ||
		dateTupleCompare(tts, &sts) == 0 ||
		dateTupleCompare(ttu, &stu) == 0 {
		return compareStatusExactMatch
	}

	if dateTupleCompare(ttu, &stu) < 0 {
		return compareStatusPrior
	}

	// Now check if the transition occurs after the given match. The
	// untilDateTime of the current match uses the same UTC offsets as the
	// transitionTime of the current transition, so no complicated adjustments
	// are needed. We just make sure we compare 'w' with 'w', 's' with 's',
	// and 'u' with 'u'.
	matchUntil := &match.untilDt
	var transitionTime *dateTuple
	if matchUntil.suffix == zoneinfo.SuffixS {
		transitionTime = tts
	} else if matchUntil.suffix == zoneinfo.SuffixU {
		transitionTime = ttu
	} else { // assume 'w'
		transitionTime = ttw
	}
	if dateTupleCompare(transitionTime, matchUntil) < 0 {
		return compareStatusWithinMatch
	}

	return compareStatusFarFuture
}

func compareTransitionToMatchFuzzy(t *transition, m *matchingEra) uint8 {
	return dateTupleCompareFuzzy(&t.transitionTime, &m.startDt, &m.untilDt)
}
