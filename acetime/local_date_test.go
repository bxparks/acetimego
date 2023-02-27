package acetime

import (
	"math"
	"testing"
)

// Verify my understanding of Golang's integer arithematics.
func TestInvalidYear(t *testing.T) {
	if InvalidYear != math.MinInt16 {
		t.Fatal(InvalidYear)
	}
}

func TestIsLeapYear(t *testing.T) {
	if !IsLeapYear(2000) {
		t.Fatalf(`IsLeapYear(2000) should is a leap year`)
	}
	if IsLeapYear(2001) {
		t.Fatalf(`IsLeapYear(2001) should not be a leap year`)
	}
	if !IsLeapYear(2004) {
		t.Fatalf(`IsLeapYear(2004) should be a leap year`)
	}
	if IsLeapYear(2100) {
		t.Fatalf(`IsLeapYear(2100) should not be a leap year`)
	}
	if !IsLeapYear(2400) {
		t.Fatalf(`IsLeapYear(2400) should be a leap year`)
	}
	if IsLeapYear(2401) {
		t.Fatalf(`IsLeapYear(2401) should not be a leap year`)
	}
}

func TestDaysInYearMonth(t *testing.T) {
	if DaysInYearMonth(2000, 1) != 31 {
		t.Fatalf(`(2000, 1) should have 31 days`)
	}
	if DaysInYearMonth(2000, 2) != 29 {
		t.Fatalf(`(2000, 1) should have 29 days`)
	}
	if DaysInYearMonth(2000, 12) != 31 {
		t.Fatalf(`(2000, 12) should have 31 days`)
	}
	if DaysInYearMonth(2004, 2) != 29 {
		t.Fatalf(`(2004, 1) should have 29 days`)
	}
	if DaysInYearMonth(2004, 11) != 30 {
		t.Fatalf(`(2004, 11) should have 30 days`)
	}
}

func TestLocalDateToWeekday(t *testing.T) {
	if LocalDateToWeekday(2000, 1, 1) != Saturday {
		t.Fatalf(`(2000, 1, 1) should be a Saturday`)
	}
	if LocalDateToWeekday(2000, 1, 2) != Sunday {
		t.Fatalf(`(2000, 1, 2) should be a Sunday`)
	}
	if LocalDateToWeekday(2000, 12, 31) != Sunday {
		t.Fatalf(`(2000, 12, 31) should be a Sunday`)
	}
	if LocalDateToWeekday(2022, 12, 23) != Friday {
		t.Fatalf(`(2022, 12, 23) should be a Friday`)
	}
	if LocalDateToWeekday(2022, 12, 25) != Sunday {
		t.Fatalf(`(2022, 12, 25) should be a Sunday`)
	}
}

func TestLocalDateFromEpochDays(t *testing.T) {
	y, m, d := LocalDateFromEpochDays(0)
	if y != 1970 || m != 1 || d != 1 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (1970, 1, 1)`)
	}

	y, m, d = LocalDateFromEpochDays(1)
	if y != 1970 || m != 1 || d != 2 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (1970, 1, 2)`)
	}

	y, m, d = LocalDateFromEpochDays(-1)
	if y != 1969 || m != 12 || d != 31 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (1969, 12, 31)`)
	}

	y, m, d = LocalDateFromEpochDays(365)
	if y != 1971 || m != 1 || d != 1 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (1971, 1, 1)`)
	}
}

func TestLocalDateToEpochDays(t *testing.T) {
	days := LocalDateToEpochDays(1970, 1, 1)
	if days != 0 {
		t.Fatalf("LocalDateToEpochDays(1970, 1, 1) should be 0")
	}

	days = LocalDateToEpochDays(1970, 1, 2)
	if days != 1 {
		t.Fatalf("LocalDateToEpochDays(1970, 1, 2) should be 1")
	}

	days = LocalDateToEpochDays(1969, 12, 31)
	if days != -1 {
		t.Fatalf("LocalDateToEpochDays(1969, 12, 31) should be -1")
	}

	days = LocalDateToEpochDays(1971, 1, 1)
	if days != 365 {
		t.Fatalf("LocalDateToEpochDays(1971, 1, 1) should be 365")
	}
}

func TestLocalDateIncrementOneDay(t *testing.T) {
	var year int16
	var month uint8
	var day uint8

	year, month, day = LocalDateIncrementOneDay(1970, 1, 1)
	if year != 1970 || month != 1 || day != 2 {
		t.Fatalf("LocalDateIncrementOneDay(1970, 1, 1) should be (1970, 1, 2)")
	}

	year, month, day = LocalDateIncrementOneDay(1970, 2, 28)
	if year != 1970 || month != 3 || day != 1 {
		t.Fatalf("LocalDateIncrementOneDay(1970, 2, 28) should be (1970, 3, 1)")
	}

	year, month, day = LocalDateIncrementOneDay(1970, 12, 31)
	if year != 1971 || month != 1 || day != 1 {
		t.Fatalf("LocalDateIncrementOneDay(1970, 12, 31) should be (1971, 1, 1)")
	}
}

func TestLocalDateDecrementOneDay(t *testing.T) {
	var year int16
	var month uint8
	var day uint8

	year, month, day = LocalDateDecrementOneDay(1970, 1, 1)
	if year != 1969 || month != 12 || day != 31 {
		t.Fatalf("LocalDateDecrementOneDay(1970, 1, 1) should be (1969, 12, 31)")
	}

	year, month, day = LocalDateDecrementOneDay(1972, 3, 1)
	if year != 1972 || month != 2 || day != 29 {
		t.Fatalf("LocalDateDecrementOneDay(1972, 3, 1) should be (1972, 2, 29)")
	}

	year, month, day = LocalDateDecrementOneDay(1970, 2, 1)
	if year != 1970 || month != 1 || day != 31 {
		t.Fatalf("LocalDateDecrementOneDay(1970, 2, 1) should be (1970, 1, 31)")
	}
}
