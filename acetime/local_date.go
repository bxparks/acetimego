package acetime

const (
	InvalidYear = int16(-(1<<15)) // math.MinInt16
)

const (
	IsoWeekdayMonday = iota + 1
	IsoWeekdayTuesday
	IsoWeekdayWednesday
	IsoWeekdayThursday
	IsoWeekdayFriday
	IsoWeekdaySaturday
	IsoWeekdaySunday
)

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
//
//   - daysOfWeek[3] is 3 because April (index=3) 1st is shifted by 3
//     days because March has 31 days (28 + 3).
//   - daysOfWeek[4] is 5 because May (index=4) 1st is shifted by 2
//     additional days from April, because April has 30 days (28 + 2).
var daysOfWeek = [12]uint8{
	5, /*Jan=31*/
	1, /*Feb=28*/
	0, /*Mar=31, start of "year"*/
	3, /*Apr=30*/
	5, /*May=31*/
	1, /*Jun=30*/
	3, /*Jul=31*/
	6, /*Aug=31*/
	2, /*Sep=30*/
	4, /*Oct=31*/
	0, /*Nov=30*/
	2, /*Dec=31*/
}

// Number of days in each month in a non-leap year. 0=Jan, 11=Dec.
var daysInMonth = [12]uint8{
	31, /*Jan=31*/
	28, /*Feb=28*/
	31, /*Mar=31*/
	30, /*Apr=30*/
	31, /*May=31*/
	30, /*Jun=30*/
	31, /*Jul=31*/
	31, /*Aug=31*/
	30, /*Sep=30*/
	31, /*Oct=31*/
	30, /*Nov=30*/
	31, /*Dec=31*/
}

// IsLeapYear returns true if the given year is a leap year, false otherwise.
func IsLeapYear(year int16) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}

// DaysInYearMonth returns the number of days in the given (year, month) pair,
// properly accounting for leap years.
func DaysInYearMonth(year int16, month uint8) uint8 {
	days := daysInMonth[month-1]
	if month == 2 && IsLeapYear(year) {
		return days + 1
	} else {
		return days
	}
}

// DayOfWeek returns the ISO week number (Monday=1, Sunday=7) of the given
// (year, month, day).
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
		return uint8((d+1)%7 + 8)
	} else {
		return uint8((d+1)%7 + 1)
	}
}

// LocalDateFromEpochDays converts epoch days to (y, m, d).
func LocalDateFromEpochDays(days int32) (year int16, month uint8, day uint8) {
	days -= daysToConverterEpochFromUnixEpoch
	year, month, day = ConvertFromDays(days)
	return
}

// LocalDateToEpochDays converts (y, m, d) to epoch days.
func LocalDateToEpochDays(year int16, month uint8, day uint8) int32 {
	converterDays := ConvertToDays(year, month, day)
	return converterDays + daysToConverterEpochFromUnixEpoch
}

// LocalDateIncrementOneDay returns the given (year, month, day) incremented by
// one day, taking proper account of leap years.
func LocalDateIncrementOneDay(y int16, m uint8, d uint8) (
	yy int16, mm uint8, dd uint8) {

	dd = d + 1
	mm = m
	yy = y

	if dd > DaysInYearMonth(y, m) {
		dd = 1
		mm++
		if mm > 12 {
			mm = 1
			yy++
		}
	}
	return
}

// LocalDateIncrementOneDay returns the given (year, month, day) decremented by
// one day, taking proper account of leap years.
func LocalDateDecrementOneDay(y int16, m uint8, d uint8) (
	yy int16, mm uint8, dd uint8) {

	dd = d - 1
	mm = m
	yy = y

	if dd == 0 {
		mm--
		if mm == 0 {
			mm = 12
			yy--
			dd = 31
		} else {
			dd = DaysInYearMonth(yy, mm)
		}
	}
	return
}
