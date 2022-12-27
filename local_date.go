package acetime

import (
	"math"
)

const InvalidYear int16 = math.MinInt16

// Offsets used to calculate the day of the week of a particular (year, month,
// day). The element represents the number of days that the first of month of
// the given index was shifted by the cummulative days from the previous months.
// To determine the "day of the week", we must normalize the resulting "day of
// the week" modulo 7.
//
// January is index 0, but we also use a modified year, where the year starts in
// March to make leap years easier to handle, so the shift for March=3 is 0.
//
// For example:
//    * atc_days_of_week[3] is 3 because April (index=3) 1st is shifted by 3
//      days because March has 31 days (28 + 3).
//    * atc_days_of_week[4] is 5 because May (index=4) 1st is shifted by 2
//      additional days from April, because April has 30 days (28 + 2).
var daysOfWeek = [12]uint8{
  5 /*Jan=31*/,
  1 /*Feb=28*/,
  0 /*Mar=31, start of "year"*/,
  3 /*Apr=30*/,
  5 /*May=31*/,
  1 /*Jun=30*/,
  3 /*Jul=31*/,
  6 /*Aug=31*/,
  2 /*Sep=30*/,
  4 /*Oct=31*/,
  0 /*Nov=30*/,
  2 /*Dec=31*/,
}

var daysInMonth = [12]uint8{
  31 /*Jan=31*/,
  28 /*Feb=28*/,
  31 /*Mar=31*/,
  30 /*Apr=30*/,
  31 /*May=31*/,
  30 /*Jun=30*/,
  31 /*Jul=31*/,
  31 /*Aug=31*/,
  30 /*Sep=30*/,
  31 /*Oct=31*/,
  30 /*Nov=30*/,
  31 /*Dec=31*/,
}

func IsLeapYear(year int16) bool {
  return ((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0)
}

func DaysInYearMonth(year int16, month uint8) uint8 {
  days := daysInMonth[month - 1]
  if month == 2 && IsLeapYear(year) {
		return days + 1
	} else {
		return days
	}
}

func DayOfWeek(year int16, month uint8, day uint8) uint8 {
  // The "y" starts in March to shift leap year calculation to end.
  var y int16 = year
	if month < 3 {
		y--
	}

  var d int16 = y + y/4 - y/100 + y/400 +
		int16(daysOfWeek[month-1]) + int16(day)

  // 2000-01-01 was a Saturday=6, so set the offsets accordingly
  if d < -1 {
		return uint8((d + 1) % 7 + 8)
	} else {
		return uint8((d + 1) % 7 + 1)
	}
}

// Convert epoch days to (y, m, d).
func LocalDateFromEpochDays(days int32) (year int16, month uint8, day uint8) {
	// shift relative to Converter Epoch
	days += GetDaysToCurrentEpochFromConverterEpoch();
	year, month, day = ConvertFromDays(days)
	return
}

// Convert (y, m, d) to epoch days.
func LocalDateToEpochDays(year int16, month uint8, day uint8) int32 {
  converterDays := ConvertToDays(year, month, day)
  return converterDays - GetDaysToCurrentEpochFromConverterEpoch()
}
