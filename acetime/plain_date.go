package acetime

const (
	// Sentinel year value that is guaranteed to not appear in a zonedb entry.
	// Used by internal functions to indicate that something was not found.
	InvalidYear = int16(-(1 << 15)) // math.MinInt16
)

type IsoWeekday uint8

// ISO Weekday starts with Monday=1, Sunday=7.
const (
	Monday IsoWeekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var weekdayStrings = []string{
	"Err",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

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
//   - weekdayOffset[3] is 3 because April (index=3) 1st is shifted by 3
//     days because March has 31 days (28 + 3).
//   - weekdayOffset[4] is 5 because May (index=4) 1st is shifted by 2
//     additional days from April, because April has 30 days (28 + 2).
var weekdayOffset = [12]uint8{
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

// Number of cummulative days in all months prior to the month at given index.
// Assume non-leap year. 0=Jan, 11=Dec.
var daysPriorToMonth = [12]uint16{
	0,   /*Jan=31*/
	31,  /*Feb=28*/
	59,  /*Mar=31*/
	90,  /*Apr=30*/
	120, /*May=31*/
	151, /*Jun=30*/
	181, /*Jul=31*/
	212, /*Aug=31*/
	243, /*Sep=30*/
	273, /*Oct=31*/
	304, /*Nov=30*/
	334, /*Dec=31*/
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

// PlainDateToWeekday returns the ISO week day of the given (year, month, day).
func PlainDateToWeekday(year int16, month uint8, day uint8) IsoWeekday {
	// The "y" starts in March to shift leap year calculation to end.
	var y int16 = year
	if month < 3 {
		y--
	}

	var d int16 = y + y/4 - y/100 + y/400 +
		int16(weekdayOffset[month-1]) + int16(day)

	// 2000-01-01 was a Saturday=6, so set the offsets accordingly
	if d < -1 {
		return IsoWeekday((d+1)%7 + 8)
	} else {
		return IsoWeekday((d+1)%7 + 1)
	}
}

// PlainDateToYearday returns the day of the year for the given (year, month,
// day). Jan 1 returns 1.
func PlainDateToYearday(year int16, month uint8, day uint8) uint16 {
	daysPrior := daysPriorToMonth[month-1]
	if IsLeapYear(year) && month > 2 {
		daysPrior++
	}
	return daysPrior + uint16(day)
}

// PlainDateFromUnixDays converts Unix epoch days to (y, m, d).
func PlainDateFromUnixDays(days int32) (year int16, month uint8, day uint8) {
	days -= daysToConverterEpochFromUnixEpoch
	year, month, day = convertFromDays(days)
	return
}

// PlainDateToUnixDays converts (y, m, d) to Unix epoch days.
func PlainDateToUnixDays(year int16, month uint8, day uint8) int32 {
	converterDays := convertToDays(year, month, day)
	return converterDays + daysToConverterEpochFromUnixEpoch
}

// PlainDateIncrementOneDay returns the given (year, month, day) incremented by
// one day, taking proper account of leap years.
func PlainDateIncrementOneDay(y int16, m uint8, d uint8) (
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

// PlainDateIncrementOneDay returns the given (year, month, day) decremented by
// one day, taking proper account of leap years.
func PlainDateDecrementOneDay(y int16, m uint8, d uint8) (
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

// Return the human readable English name for the weekday (e.g. "Sunday").
// The 3-letter abbreviation can be retrieved with a string slice.
func (wd IsoWeekday) Name() string {
	return weekdayStrings[wd]
}
