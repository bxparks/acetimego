package acetime

import (
	"testing"
)

//-----------------------------------------------------------------------------
// ZonedExtra.
// Extra meta information about a given instant in time, such as the
// the STD offset, the DST offset, and the abbreviation used.
//-----------------------------------------------------------------------------

// Test that ZonedExtraXxx constants are the same as findResultXxx constants.
func TestZonedExtraTypeConstantsMatch(t *testing.T) {
	if !(ZonedExtraErr == findResultErr) {
		t.Fatal("")
	}
	if !(ZonedExtraNotFound == findResultNotFound) {
		t.Fatal("")
	}
	if !(ZonedExtraExact == findResultExact) {
		t.Fatal("")
	}
	if !(ZonedExtraGap == findResultGap) {
		t.Fatal("")
	}
	if !(ZonedExtraOverlap == findResultOverlap) {
		t.Fatal("")
	}
}
