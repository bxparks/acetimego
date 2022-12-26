package acetime

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	if ! IsLeapYear(2000) {
		t.Fatalf(`IsLeapYear(2000) should is a leap year`)
	}
	if IsLeapYear(2001) {
		t.Fatalf(`IsLeapYear(2001) should not be a leap year`)
	}
	if ! IsLeapYear(2004) {
		t.Fatalf(`IsLeapYear(2004) should be a leap year`)
	}
	if IsLeapYear(2100) {
		t.Fatalf(`IsLeapYear(2100) should not be a leap year`)
	}
	if ! IsLeapYear(2400) {
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

func TestLocalDateForEpochDays(t *testing.T) {
	ld := LocalDateForEpochDays(0)
	if ld.Year != 2050 || ld.Month != 1 || ld.Day != 1 {
		t.Fatalf(`LocalDateForEpochDays(0) should be (2050, 1, 1)`)
	}

	ld = LocalDateForEpochDays(1)
	if ld.Year != 2050 || ld.Month != 1 || ld.Day != 2 {
		t.Fatalf(`LocalDateForEpochDays(0) should be (2050, 1, 2)`)
	}

	ld = LocalDateForEpochDays(-1)
	if ld.Year != 2049 || ld.Month != 12 || ld.Day != 31 {
		t.Fatalf(`LocalDateForEpochDays(0) should be (2049, 12, 31)`)
	}

	ld = LocalDateForEpochDays(365)
	if ld.Year != 2051 || ld.Month != 1 || ld.Day != 1 {
		t.Fatalf(`LocalDateForEpochDays(0) should be (2051, 1, 1)`)
	}
}

func TestLocalDateToEpochDays(t *testing.T) {
	days := LocalDate{2050, 1, 1}.ToEpochDays()
	if days != 0 {
		t.Fatalf("LocalDate{2050, 1, 1}.ToEpochDays() should be 0")
	}

	days = LocalDate{2050, 1, 2}.ToEpochDays()
	if days != 1 {
		t.Fatalf("LocalDate{2050, 1, 2}.ToEpochDays() should be 1")
	}

	days = LocalDate{2049, 12, 31}.ToEpochDays()
	if days != -1 {
		t.Fatalf("LocalDate{2049, 12, 31}.ToEpochDays() should be -1")
	}

	days = LocalDate{2051, 1, 1}.ToEpochDays()
	if days != 365 {
		t.Fatalf("LocalDate{2051, 1, 1}.ToEpochDays() should be 365")
	}
}
