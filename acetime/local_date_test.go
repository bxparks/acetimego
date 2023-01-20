package acetime

import (
	"testing"
)

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

func TestDayOfWeeek(t *testing.T) {
	if DayOfWeek(2000, 1, 1) != 6 {
		t.Fatalf(`(2000, 1, 1) should be a Saturday`)
	}
	if DayOfWeek(2000, 1, 2) != 7 {
		t.Fatalf(`(2000, 1, 2) should be a Sunday`)
	}
	if DayOfWeek(2000, 12, 31) != 7 {
		t.Fatalf(`(2000, 12, 31) should be a Sunday`)
	}
	if DayOfWeek(2022, 12, 23) != 5 {
		t.Fatalf(`(2022, 12, 23) should be a Friday`)
	}
	if DayOfWeek(2022, 12, 25) != 7 {
		t.Fatalf(`(2022, 12, 25) should be a Sunday`)
	}
}

func TestLocalDateFromEpochDays(t *testing.T) {
	y, m, d := LocalDateFromEpochDays(0)
	if y != 2050 || m != 1 || d != 1 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (2050, 1, 1)`)
	}

	y, m, d = LocalDateFromEpochDays(1)
	if y != 2050 || m != 1 || d != 2 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (2050, 1, 2)`)
	}

	y, m, d = LocalDateFromEpochDays(-1)
	if y != 2049 || m != 12 || d != 31 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (2049, 12, 31)`)
	}

	y, m, d = LocalDateFromEpochDays(365)
	if y != 2051 || m != 1 || d != 1 {
		t.Fatalf(`LocalDateFromEpochDays(0) should be (2051, 1, 1)`)
	}
}

func TestLocalDateToEpochDays(t *testing.T) {
	days := LocalDateToEpochDays(2050, 1, 1)
	if days != 0 {
		t.Fatalf("LocalDateToEpochDays(2050, 1, 1) should be 0")
	}

	days = LocalDateToEpochDays(2050, 1, 2)
	if days != 1 {
		t.Fatalf("LocalDateToEpochDays(2050, 1, 2) should be 1")
	}

	days = LocalDateToEpochDays(2049, 12, 31)
	if days != -1 {
		t.Fatalf("LocalDateToEpochDays(2049, 12, 31) should be -1")
	}

	days = LocalDateToEpochDays(2051, 1, 1)
	if days != 365 {
		t.Fatalf("LocalDateToEpochDays(2051, 1, 1) should be 365")
	}
}

func TestLocalDateIncrementOneDay(t *testing.T) {
	var year int16
	var month uint8
	var day uint8

	year, month, day = LocalDateIncrementOneDay(2050, 1, 1)
	if year != 2050 || month != 1 || day != 2 {
		t.Fatalf("LocalDateIncrementOneDay(2050, 1, 1) should be (2050, 1, 2)")
	}

	year, month, day = LocalDateIncrementOneDay(2050, 2, 28)
	if year != 2050 || month != 3 || day != 1 {
		t.Fatalf("LocalDateIncrementOneDay(2050, 2, 28) should be (2050, 3, 1)")
	}

	year, month, day = LocalDateIncrementOneDay(2050, 12, 31)
	if year != 2051 || month != 1 || day != 1 {
		t.Fatalf("LocalDateIncrementOneDay(2050, 12, 31) should be (2051, 1, 1)")
	}
}

func TestLocalDateDecrementOneDay(t *testing.T) {
	var year int16
	var month uint8
	var day uint8

	year, month, day = LocalDateDecrementOneDay(2050, 1, 1)
	if year != 2049 || month != 12 || day != 31 {
		t.Fatalf("LocalDateDecrementOneDay(2050, 1, 1) should be (2049, 12, 31)")
	}

	year, month, day = LocalDateDecrementOneDay(2052, 3, 1)
	if year != 2052 || month != 2 || day != 29 {
		t.Fatalf("LocalDateDecrementOneDay(2052, 3, 1) should be (2052, 2, 29)")
	}

	year, month, day = LocalDateDecrementOneDay(2050, 2, 1)
	if year != 2050 || month != 1 || day != 31 {
		t.Fatalf("LocalDateDecrementOneDay(2050, 2, 1) should be (2050, 1, 31)")
	}
}