package acetime

import (
	"testing"
)

//-----------------------------------------------------------------------------
// ZonedExtra.
// Extra meta information about a given instant in time, such as the
// the STD offset, the DST offset, and the abbreviation used.
//-----------------------------------------------------------------------------

// Test that FoldTypeXxx constants are the same as findResultXxx constants.
func TestFoldTypeTypeConstantsMatch(t *testing.T) {
	if !(FoldTypeErr == findResultErr) {
		t.Fatal("")
	}
	if !(FoldTypeNotFound == findResultNotFound) {
		t.Fatal("")
	}
	if !(FoldTypeExact == findResultExact) {
		t.Fatal("")
	}
	if !(FoldTypeGap == findResultGap) {
		t.Fatal("")
	}
	if !(FoldTypeOverlap == findResultOverlap) {
		t.Fatal("")
	}
}
