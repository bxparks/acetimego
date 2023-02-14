package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
	a := DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	b := DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &b) == 0) {
		t.Fatal("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

	bb := DateTuple{2000, 1, 1, 0, zoneinfo.SuffixS}
	if !(dateTupleCompare(&a, &bb) == 0) {
		t.Fatal("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

	c := DateTuple{2000, 1, 1, 1, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &c) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

	d := DateTuple{2000, 1, 2, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &d) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

	e := DateTuple{2000, 2, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &e) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

	f := DateTuple{2001, 1, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &f) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}

func TestDateTupleSubtract(t *testing.T) {
	var dta, dtb DateTuple
	var diff ATime

	dta = DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}      // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 1, 1 * 60, zoneinfo.SuffixW} // 2000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 1, 2, 0, zoneinfo.SuffixW} // 2000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW} // 2000-01-01 00:00
	dtb = DateTuple{2000, 2, 1, 0, zoneinfo.SuffixW} // 2000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatal(diff)
	}

	dta = DateTuple{2000, 2, 1, 0, zoneinfo.SuffixW} // 2000-02-01 00:00
	dtb = DateTuple{2000, 3, 1, 0, zoneinfo.SuffixW} // 2000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 2000 is leap, 29 day
		t.Fatal(diff)
	}
}

// Test that there is no overflow for year 6000, which is far beyond the
// Epoch.currentEpochYear.
func TestDateTupleSubtractNoOverflow(t *testing.T) {
	var dta, dtb DateTuple
	var diff ATime

	dta = DateTuple{6000, 1, 1, 0, zoneinfo.SuffixW}      // 6000-01-01 00:00
	dtb = DateTuple{6000, 1, 1, 1 * 60, zoneinfo.SuffixW} // 6000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{6000, 1, 1, 0, zoneinfo.SuffixW} // 6000-01-01 00:00
	dtb = DateTuple{6000, 1, 2, 0, zoneinfo.SuffixW} // 6000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatal(diff)
	}

	dta = DateTuple{6000, 1, 1, 0, zoneinfo.SuffixW} // 6000-01-01 00:00
	dtb = DateTuple{6000, 2, 1, 0, zoneinfo.SuffixW} // 6000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatal(diff)
	}

	dta = DateTuple{6000, 2, 1, 0, zoneinfo.SuffixW} // 6000-02-01 00:00
	dtb = DateTuple{6000, 3, 1, 0, zoneinfo.SuffixW} // 6000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 6000 is leap, 29 day
		t.Fatal(diff)
	}
}

func TestDateTupleNormalize(t *testing.T) {
	var dt DateTuple

	// 00:00
	dt = DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 23:45
	dt = DateTuple{2000, 1, 1, 23*3600 + 45*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 1, 23*3600 + 45*60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 24:00
	dt = DateTuple{2000, 1, 1, 24 * 3600, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 24:15
	dt = DateTuple{2000, 1, 1, 24*3600 + 15*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{2000, 1, 2, 15 * 60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// -24:00
	dt = DateTuple{2000, 1, 1, -24 * 3600, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// -24:15
	dt = DateTuple{2000, 1, 1, -24*3600 - 15*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == DateTuple{1999, 12, 31, -15 * 60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}
}

func TestDateTupleExpand(t *testing.T) {
	var tt DateTuple
	var ttw DateTuple
	var tts DateTuple
	var ttu DateTuple

	const offsetSeconds = 2 * 60 * 60
	const deltaSeconds = 1 * 60 * 60

	tt = DateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW} // 04:00
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}

	tt = DateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}

	tt = DateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == DateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == DateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == DateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}
}

func TestDateTupleCompareFuzzy(t *testing.T) {
	status := dateTupleCompareFuzzy(
		&DateTuple{2000, 10, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusPrior == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2000, 11, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 2, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 3, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&DateTuple{2002, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusFarFuture == status) {
		t.Fatal(status)
	}

	// Verify dates whose delta months is greater than 32767. In
	// other words, delta years is greater than 2730.
	status = dateTupleCompareFuzzy(
		&DateTuple{5000, 4, 1, 1, 0},
		&DateTuple{2000, 12, 1, 1, 0},
		&DateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusFarFuture == status) {
		t.Fatal(status)
	}
	status = dateTupleCompareFuzzy(
		&DateTuple{1000, 4, 1, 1, 0},
		&DateTuple{4000, 12, 1, 1, 0},
		&DateTuple{4002, 2, 1, 1, 0})
	if !(compareStatusPrior == status) {
		t.Fatal(status)
	}
}
