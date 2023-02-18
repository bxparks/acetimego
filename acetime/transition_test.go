package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestCompareTransitionToMatchFuzzy(t *testing.T) {
	match := matchingEra{
		startDt: DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		untilDt: DateTuple{2001, 1, 1, 0, zoneinfo.SuffixW},
	}

	tn := transition{
		match:          &match,
		transitionTime: DateTuple{1999, 11, 1, 0, zoneinfo.SuffixW},
	}
	status := compareTransitionToMatchFuzzy(&tn, &match)
	if !(status == compareStatusPrior) {
		t.Fatal("fatal")
	}

	tn = transition{
		match:          &match,
		transitionTime: DateTuple{1999, 12, 1, 0, zoneinfo.SuffixW},
	}
	status = compareTransitionToMatchFuzzy(&tn, &match)
	if !(status == compareStatusWithinMatch) {
		t.Fatal("fatal")
	}

	tn = transition{
		match:          &match,
		transitionTime: DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
	}
	status = compareTransitionToMatchFuzzy(&tn, &match)
	if !(status == compareStatusWithinMatch) {
		t.Fatal("fatal")
	}

	tn = transition{
		match:          &match,
		transitionTime: DateTuple{2001, 1, 1, 0, zoneinfo.SuffixW},
	}
	status = compareTransitionToMatchFuzzy(&tn, &match)
	if !(status == compareStatusWithinMatch) {
		t.Fatal("fatal")
	}

	tn = transition{
		match:          &match,
		transitionTime: DateTuple{2001, 3, 1, 0, zoneinfo.SuffixW},
	}
	status = compareTransitionToMatchFuzzy(&tn, &match)
	if !(status == compareStatusFarFuture) {
		t.Fatal("fatal")
	}
}

func TestCompareTransitionToMatch(t *testing.T) {
	// UNTIL = 2002-01-02T03:00
	era := zoneinfo.ZoneEra{
		OffsetSecondsCode:    0,
		DeltaMinutes:         0,
		UntilYear:            2,
		UntilMonth:           1,
		UntilDay:             2,
		UntilSecondsCode:     3 * 3600 / 15,
		UntilSecondsModifier: zoneinfo.SuffixW,
	}

	// matchingEra=[2000-01-01, 2001-01-01)
	match := matchingEra{
		startDt:           DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		untilDt:           DateTuple{2001, 1, 1, 0, zoneinfo.SuffixW},
		era:               &era,
		prevMatch:         nil,
		lastOffsetSeconds: 0,
		lastDeltaSeconds:  0,
	}

	transitions := []transition{
		// transitionTime = 1999-12-31
		transition{
			match:          &match,
			transitionTime: DateTuple{1999, 12, 31, 0, zoneinfo.SuffixW},
		},
		// transitionTime = 2000-01-01
		transition{
			match:          &match,
			transitionTime: DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW},
		},
		// transitionTime = 2000-01-02
		transition{
			match:          &match,
			transitionTime: DateTuple{2000, 1, 2, 0, zoneinfo.SuffixW},
		},
		// transitionTime = 2001-02-03
		transition{
			match:          &match,
			transitionTime: DateTuple{2001, 2, 3, 0, zoneinfo.SuffixW},
		},
	}
	transition0 := &transitions[0]
	transition1 := &transitions[1]
	transition2 := &transitions[2]
	transition3 := &transitions[3]

	// Populate the transitionTimeS and transitionTimeU fields.
	fixTransitionTimes(transitions)

	status := compareTransitionToMatch(transition0, &match)
	if !(status == compareStatusPrior) {
		t.Fatal("tt:", transition0.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition1, &match)
	if !(status == compareStatusExactMatch) {
		t.Fatal("tt:", transition1.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition2, &match)
	if !(status == compareStatusWithinMatch) {
		t.Fatal("tt:", transition2.transitionTime, "; status:", status)
	}

	status = compareTransitionToMatch(transition3, &match)
	if !(status == compareStatusFarFuture) {
		t.Fatal("tt:", transition3.transitionTime, "; status:", status)
	}
}
