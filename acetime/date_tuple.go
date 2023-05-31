package acetime

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

// A dateTuple is an internal version of [LocalDateTime] which also tracks the
// `s`, `w` or `u` suffixes given in the TZ database files.
//
// TODO: We only need about 24-bits (3-bytes) the seconds field. It should be
// possible to reduce this struct by 1-byte so that the object fits entirely
// within 8-bytes, a multiple of 4 or 8 bytes which reduces memory consumption
// on 32-bit and 64-bit processors.
type dateTuple struct {
	year    int16 // [0,10000]
	month   uint8 // [1-12]
	day     uint8 // [1-31]
	seconds int32 // negative values allowed
	suffix  uint8 // zoneinfo.SuffixS, zoneinfo.SuffixW, zoneinfo.SuffixU
}

// dateTupleCompare compare 2 dateTuple instances (a, b) and returns -1, 0, 1
// depending on whether a is less than, equal, or greater than b, respectively.
func dateTupleCompare(a *dateTuple, b *dateTuple) int8 {
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
	if a.seconds < b.seconds {
		return -1
	}
	if a.seconds > b.seconds {
		return 1
	}
	return 0
}

// dateTupleSubtract returns the number of seconds of (a - b).
func dateTupleSubtract(a *dateTuple, b *dateTuple) Time {
	da := LocalDateToEpochDays(a.year, a.month, a.day)
	db := LocalDateToEpochDays(b.year, b.month, b.day)

	return Time(da-db)*86400 + Time(a.seconds-b.seconds)
}

func dateTupleNormalize(dt *dateTuple) {
	const oneDayAsSeconds = 60 * 60 * 24

	if dt.seconds <= -oneDayAsSeconds {
		dt.year, dt.month, dt.day = LocalDateDecrementOneDay(
			dt.year, dt.month, dt.day)
		dt.seconds += oneDayAsSeconds
	} else if oneDayAsSeconds <= dt.seconds {
		dt.year, dt.month, dt.day = LocalDateIncrementOneDay(
			dt.year, dt.month, dt.day)
		dt.seconds -= oneDayAsSeconds
	} else {
		// do nothing
	}
}

// dateTupleExpand converts the given 'tt', offsetSeconds, and deltaSeconds into
// the 'w', 's' and 'u' versions of the dateTuple. It is allowed for 'ttw' to
// be an alias of 'tt'.
func dateTupleExpand(
	tt *dateTuple,
	offsetSeconds int32,
	deltaSeconds int32,
	ttw *dateTuple,
	tts *dateTuple,
	ttu *dateTuple) {

	if tt.suffix == zoneinfo.SuffixS {
		*tts = *tt

		ttu.year = tt.year
		ttu.month = tt.month
		ttu.day = tt.day
		ttu.seconds = tt.seconds - offsetSeconds
		ttu.suffix = zoneinfo.SuffixU

		ttw.year = tt.year
		ttw.month = tt.month
		ttw.day = tt.day
		ttw.seconds = tt.seconds + deltaSeconds
		ttw.suffix = zoneinfo.SuffixW
	} else if tt.suffix == zoneinfo.SuffixU {
		*ttu = *tt

		tts.year = tt.year
		tts.month = tt.month
		tts.day = tt.day
		tts.seconds = tt.seconds + offsetSeconds
		tts.suffix = zoneinfo.SuffixS

		ttw.year = tt.year
		ttw.month = tt.month
		ttw.day = tt.day
		ttw.seconds = tt.seconds + (offsetSeconds + deltaSeconds)
		ttw.suffix = zoneinfo.SuffixW
	} else {
		// Explicit set the suffix to 'w' in case it was something else.
		*ttw = *tt
		ttw.suffix = zoneinfo.SuffixW

		tts.year = tt.year
		tts.month = tt.month
		tts.day = tt.day
		tts.seconds = tt.seconds - deltaSeconds
		tts.suffix = zoneinfo.SuffixS

		ttu.year = tt.year
		ttu.month = tt.month
		ttu.day = tt.day
		ttu.seconds = tt.seconds - (deltaSeconds + offsetSeconds)
		ttu.suffix = zoneinfo.SuffixU
	}

	dateTupleNormalize(ttw)
	dateTupleNormalize(tts)
	dateTupleNormalize(ttu)
}

const (
	compareStatusFarPast     = iota // 0
	compareStatusPrior              // 1
	compareStatusExactMatch         // 2
	compareStatusWithinMatch        // 3
	compareStatusFarFuture          // 4
)

// dateTupleCompareFuzzy compares the given 't' with the interval defined by
// [start, until). The comparison is fuzzy, with a slop of about one month so
// that we can ignore the day and minutes fields.
//
// The following values are returned:
//
//   - compareStatusPrior if 't' is less than 'start' by at least one month,
//   - compareStatusFarFuture if 't' is greater than 'until' by at least one
//     month,
//   - compareStatusWithinMatch if 't' is within [start, until) with a one
//     month slop,
//   - compareStatusExactMatch is never returned.
func dateTupleCompareFuzzy(
	t *dateTuple, start *dateTuple, until *dateTuple) uint8 {

	// Use int32 because a delta year of 2730 or greater will exceed
	// the range of an int16.
	var tMonths int32 = int32(t.year)*12 + int32(t.month)
	var startMonths int32 = int32(start.year)*12 + int32(start.month)
	if tMonths < startMonths-1 {
		return compareStatusPrior
	}
	var untilMonths int32 = int32(until.year)*12 + int32(until.month)
	if untilMonths+1 < tMonths {
		return compareStatusFarFuture
	}
	return compareStatusWithinMatch
}
