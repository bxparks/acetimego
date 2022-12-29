package acetime

import (
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
	a := DateTuple{2000, 1, 1, 0, suffixW}
	b := DateTuple{2000, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &b) == 0) {
		t.Fatalf("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

	bb := DateTuple{2000, 1, 1, 0, suffixS}
	if dateTupleCompare(&a, &bb) != 0 {
		t.Fatalf("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

	c := DateTuple{2000, 1, 1, 1, suffixW}
	if !(dateTupleCompare(&a, &c) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

	d := DateTuple{2000, 1, 2, 0, suffixW}
	if !(dateTupleCompare(&a, &d) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

	e := DateTuple{2000, 2, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &e) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

	f := DateTuple{2001, 1, 1, 0, suffixW}
	if !(dateTupleCompare(&a, &f) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}

func TestDateTupleSubtract(t *testing.T) {
	var dta, dtb DateTuple
	var diff int32

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 1, 1, suffixW} // 2000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 2, 0, suffixW} // 2000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 1, 1, 0, suffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatalf("fatal")
	}

	dta = DateTuple{2000, 2, 1, 0, suffixW} // 2000-02-01 00:00
	dtb = DateTuple{2000, 3, 1, 0, suffixW} // 2000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 2000 is leap, 29 day
		t.Fatalf("fatal")
	}
}

func TestDateTupleNormalize(t *testing.T) {
	var dt DateTuple

	// 00:00
	dt = DateTuple{2000, 1, 1, 0, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// 23:45
	dt = DateTuple{2000, 1, 1, 15 * 95, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 15 * 95, suffixW}) {
		t.Fatalf("fatal")
	}

	// 24:00
	dt = DateTuple{2000, 1, 1, 15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// 24:15
	dt = DateTuple{2000, 1, 1, 15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 15, suffixW}) {
		t.Fatalf("fatal")
	}

	// -24:00
	dt = DateTuple{2000, 1, 1, -15 * 96, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, 0, suffixW}) {
		t.Fatalf("fatal")
	}

	// -24:15
	dt = DateTuple{2000, 1, 1, -15 * 97, suffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, -15, suffixW}) {
		t.Fatalf("fatal")
	}
}

func TestDateTupleExpand(t *testing.T) {
	var tt DateTuple
	var ttw DateTuple
	var tts DateTuple
	var ttu DateTuple

	const offsetMinutes = 2 * 60
	const deltaMinutes = 1 * 60

	tt = DateTuple{2000, 1, 30, 15 * 16, suffixW} // 04:00
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}

	tt = DateTuple{2000, 1, 30, 15 * 12, suffixS}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}

	tt = DateTuple{2000, 1, 30, 15 * 4, suffixU}
	dateTupleExpand(&tt, offsetMinutes, deltaMinutes, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 15 * 16, suffixW}) {
		t.Fatalf("fatal")
	}
	if !(tts == DateTuple{2000, 1, 30, 15 * 12, suffixS}) {
		t.Fatalf("fatal")
	}
	if !(ttu == DateTuple{2000, 1, 30, 15 * 4, suffixU}) {
		t.Fatalf("fatal")
	}
}

func TestDateTupleCompareFuzzy(t *testing.T) {
	if !(matchStatusPrior == dateTupleCompareFuzzy(
		&DateTuple{2000, 10, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2000, 11, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2002, 2, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusWithinMatch == dateTupleCompareFuzzy(
		&DateTuple{2002, 3, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	if !(matchStatusFarFuture == dateTupleCompareFuzzy(
		&DateTuple{2002, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}

	// Verify dates whose delta months is greater than 32767. In
	// other words, delta years is greater than 2730.
	if !(matchStatusFarFuture == dateTupleCompareFuzzy(
		&DateTuple{5000, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}
	if !(matchStatusPrior == dateTupleCompareFuzzy(
		&DateTuple{1000, 4, 1, 1, 0},
		&DateTuple{4000, 12, 1, 1, 0},
		&DateTuple{4002, 2, 1, 1, 0})) {
		t.Fatalf("fatal")
	}
}
