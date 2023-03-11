package acetime

import (
	"testing"
)

func TestCalcDaysUntilMonthPrime(t *testing.T) {
	var days uint16 = calcDaysUntilMonthPrime(0)
	if days != 0 {
		t.Fatalf(`calcDaysUntilMonthPrime(0) should be 0 but was %d`, days)
	}

	days = calcDaysUntilMonthPrime(1)
	if days != 31 {
		t.Fatalf(`calcDaysUntilMonthPrime(1) should be 31 but was %d`, days)
	}

	days = calcDaysUntilMonthPrime(11)
	if days != 337 {
		t.Fatalf(`calcDaysUntilMonthPrime(11) should be 334 but was %d`, days)
	}
}

func TestConvertToDays(t *testing.T) {
	var days int32 = convertToDays(2000, 1, 1)
	var expected int32 = 0
	if days != expected {
		t.Fatalf(`convertToDays(2000, 1, 1) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(2000, 1, 2)
	expected = 1
	if days != expected {
		t.Fatalf(`convertToDays(2000, 1, 2) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(1999, 12, 31)
	expected = -1
	if days != expected {
		t.Fatalf(`convertToDays(1999, 12, 31) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(2050, 1, 1)
	expected = 18263
	if days != expected {
		t.Fatalf(`convertToDays(2050, 1, 1) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(2100, 1, 1)
	expected = 36525
	if days != expected {
		t.Fatalf(`convertToDays(2100, 1, 1) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(1900, 1, 1)
	expected = -36524
	if days != expected {
		t.Fatalf(`convertToDays(1900, 1, 1) should be %d but was %d`,
			expected, days)
	}

	days = convertToDays(2400, 1, 1)
	expected = 146097
	if days != expected {
		t.Fatalf(`convertToDays(2400, 1, 1) should be %d but was %d`,
			expected, days)
	}
}

func TestConvertFromDays(t *testing.T) {
	var days int32 = 0
	year, month, day := convertFromDays(days)
	if year != 2000 || month != 1 || day != 1 {
		t.Fatalf(
			`convertFromDays(%d) should return (2000, 1, 1) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = 1
	year, month, day = convertFromDays(days)
	if year != 2000 || month != 1 || day != 2 {
		t.Fatalf(
			`convertFromDays(%d) should return (2000, 1, 2) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = -1
	year, month, day = convertFromDays(days)
	if year != 1999 || month != 12 || day != 31 {
		t.Fatalf(
			`convertFromDays(%d) should return (1999, 12, 31) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = 18263
	year, month, day = convertFromDays(days)
	if year != 2050 || month != 1 || day != 1 {
		t.Fatalf(
			`convertFromDays(%d) should return (2050, 1, 1) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = 36525
	year, month, day = convertFromDays(days)
	if year != 2100 || month != 1 || day != 1 {
		t.Fatalf(
			`convertFromDays(%d) should return (2100, 1, 1) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = -36524
	year, month, day = convertFromDays(days)
	if year != 1900 || month != 1 || day != 1 {
		t.Fatalf(
			`convertFromDays(%d) should return (1900, 1, 2) but was (%d, %d, %d)`,
			days, year, month, day)
	}

	days = 146097
	year, month, day = convertFromDays(days)
	if year != 2400 || month != 1 || day != 1 {
		t.Fatalf(
			`convertFromDays(%d) should return (2400, 1, 1) but was (%d, %d, %d)`,
			days, year, month, day)
	}
}
