package acetime

import (
	"github.com/bxparks/acetimego/zoneinfo"
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
	a := dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	b := dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &b) == 0) {
		t.Fatal("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

	bb := dateTuple{2000, 1, 1, 0, zoneinfo.SuffixS}
	if !(dateTupleCompare(&a, &bb) == 0) {
		t.Fatal("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

	c := dateTuple{2000, 1, 1, 1, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &c) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

	d := dateTuple{2000, 1, 2, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &d) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

	e := dateTuple{2000, 2, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &e) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

	f := dateTuple{2001, 1, 1, 0, zoneinfo.SuffixW}
	if !(dateTupleCompare(&a, &f) < 0) {
		t.Fatal("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}

func TestDateTupleSubtract(t *testing.T) {
	var dta, dtb dateTuple
	var diff Time

	dta = dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}      // 2000-01-01 00:00
	dtb = dateTuple{2000, 1, 1, 1 * 60, zoneinfo.SuffixW} // 2000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatal(diff)
	}

	dta = dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW} // 2000-01-01 00:00
	dtb = dateTuple{2000, 1, 2, 0, zoneinfo.SuffixW} // 2000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatal(diff)
	}

	dta = dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW} // 2000-01-01 00:00
	dtb = dateTuple{2000, 2, 1, 0, zoneinfo.SuffixW} // 2000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatal(diff)
	}

	dta = dateTuple{2000, 2, 1, 0, zoneinfo.SuffixW} // 2000-02-01 00:00
	dtb = dateTuple{2000, 3, 1, 0, zoneinfo.SuffixW} // 2000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 2000 is leap, 29 day
		t.Fatal(diff)
	}
}

// Test that there is no overflow for year 6000, which is far beyond the
// Epoch.currentEpochYear.
func TestDateTupleSubtractNoOverflow(t *testing.T) {
	var dta, dtb dateTuple
	var diff Time

	dta = dateTuple{6000, 1, 1, 0, zoneinfo.SuffixW}      // 6000-01-01 00:00
	dtb = dateTuple{6000, 1, 1, 1 * 60, zoneinfo.SuffixW} // 6000-01-01 00:01
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-60 == diff) {
		t.Fatal(diff)
	}

	dta = dateTuple{6000, 1, 1, 0, zoneinfo.SuffixW} // 6000-01-01 00:00
	dtb = dateTuple{6000, 1, 2, 0, zoneinfo.SuffixW} // 6000-01-02 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400 == diff) {
		t.Fatal(diff)
	}

	dta = dateTuple{6000, 1, 1, 0, zoneinfo.SuffixW} // 6000-01-01 00:00
	dtb = dateTuple{6000, 2, 1, 0, zoneinfo.SuffixW} // 6000-02-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*31 == diff) { // January has 31 day
		t.Fatal(diff)
	}

	dta = dateTuple{6000, 2, 1, 0, zoneinfo.SuffixW} // 6000-02-01 00:00
	dtb = dateTuple{6000, 3, 1, 0, zoneinfo.SuffixW} // 6000-03-01 00:00
	diff = dateTupleSubtract(&dta, &dtb)
	if !(-86400*29 == diff) { // Feb 6000 is leap, 29 day
		t.Fatal(diff)
	}
}

func TestDateTupleNormalize(t *testing.T) {
	var dt dateTuple

	// 00:00
	dt = dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{2000, 1, 1, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 23:45
	dt = dateTuple{2000, 1, 1, 23*3600 + 45*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{2000, 1, 1, 23*3600 + 45*60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 24:00
	dt = dateTuple{2000, 1, 1, 24 * 3600, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{2000, 1, 2, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// 24:15
	dt = dateTuple{2000, 1, 1, 24*3600 + 15*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{2000, 1, 2, 15 * 60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// -24:00
	dt = dateTuple{2000, 1, 1, -24 * 3600, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{1999, 12, 31, 0, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}

	// -24:15
	dt = dateTuple{2000, 1, 1, -24*3600 - 15*60, zoneinfo.SuffixW}
	dateTupleNormalize(&dt)
	if !(dt == dateTuple{1999, 12, 31, -15 * 60, zoneinfo.SuffixW}) {
		t.Fatal(dt)
	}
}

func TestDateTupleExpand(t *testing.T) {
	var tt dateTuple
	var ttw dateTuple
	var tts dateTuple
	var ttu dateTuple

	const offsetSeconds = 2 * 60 * 60
	const deltaSeconds = 1 * 60 * 60

	tt = dateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW} // 04:00
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == dateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == dateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == dateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}

	tt = dateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == dateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == dateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == dateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}

	tt = dateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}
	dateTupleExpand(&tt, offsetSeconds, deltaSeconds, &ttw, &tts, &ttu)
	if !(ttw == dateTuple{2000, 1, 30, 4 * 3600, zoneinfo.SuffixW}) {
		t.Fatal(ttw)
	}
	if !(tts == dateTuple{2000, 1, 30, 3 * 3600, zoneinfo.SuffixS}) {
		t.Fatal(tts)
	}
	if !(ttu == dateTuple{2000, 1, 30, 1 * 3600, zoneinfo.SuffixU}) {
		t.Fatal(ttu)
	}
}

func TestDateTupleCompareFuzzy(t *testing.T) {
	status := dateTupleCompareFuzzy(
		&dateTuple{2000, 10, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusPrior == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&dateTuple{2000, 11, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&dateTuple{2002, 2, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&dateTuple{2002, 3, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusWithinMatch == status) {
		t.Fatal(status)
	}

	status = dateTupleCompareFuzzy(
		&dateTuple{2002, 4, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusFarFuture == status) {
		t.Fatal(status)
	}

	// Verify dates whose delta months is greater than 32767. In
	// other words, delta years is greater than 2730.
	status = dateTupleCompareFuzzy(
		&dateTuple{5000, 4, 1, 1, 0},
		&dateTuple{2000, 12, 1, 1, 0},
		&dateTuple{2002, 2, 1, 1, 0})
	if !(compareStatusFarFuture == status) {
		t.Fatal(status)
	}
	status = dateTupleCompareFuzzy(
		&dateTuple{1000, 4, 1, 1, 0},
		&dateTuple{4000, 12, 1, 1, 0},
		&dateTuple{4002, 2, 1, 1, 0})
	if !(compareStatusPrior == status) {
		t.Fatal(status)
	}
}
