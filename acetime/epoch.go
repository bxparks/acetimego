package acetime

// The type of the "seconds from epoch" in this library. This type is the
// equivalent of C lang 'time_t', or the Go lang 'time.Time'.
//
// A 32-bit integer is used because this library targets small microcontrollers
// which are supported by TinyGo where 64-bit integer operations are expensive.
// Normally, using a 32-bit integer type suffers from an overflow in the year
// 2038 due to the use of 1970-01-01 Unix epoch.
//
// AceTimeGo avoids this problem by setting the default epoch to be 2050-01-01
// Furthermore, the epoch of the AceTimeGo library is adjustable at runtime. The
// functions in this library will produce valid results within at least +/- 50
// years (and probably +/- 60 years) from the epoch year.
type ATime int32

const (
	// The base epoch year used by the ConvertToDays() and ConvertFromDays()
	// functions below. This must be a multiple of 400.
	converterEpochYear = 2000

	// Number of days from 1970-01-01 to 2000-01-01. This is a constant because
	// the converterEpochYear is a constant.
	daysToConverterEpochFromUnixEpoch = 10957
)

var (
	// Current epoch year. This is adjustable by the library caller. It is
	// expected to be set once by the application near the start of the app.
	currentEpochYear int16 = 2050

	// Number of days from 2000-01-01 to {currentEpochYear}-01-01. This is
	// derived from currentEpochYear and stored here for convenience.
	daysToCurrentEpochFromConverterEpoch int32 = 18263
)

func SetCurrentEpochYear(year int16) {
	currentEpochYear = year
	daysToCurrentEpochFromConverterEpoch = ConvertToDays(year, 1, 1)
}

func GetCurrentEpochYear() int16 {
	return currentEpochYear
}

func GetDaysToCurrentEpochFromConverterEpoch() int32 {
	return daysToCurrentEpochFromConverterEpoch
}

func GetDaysToCurrentEpochFromUnixEpoch() int32 {
	return daysToCurrentEpochFromConverterEpoch +
		daysToConverterEpochFromUnixEpoch
}

func GetSecondsToCurrentEpochFromUnixEpoch64() int64 {
	return 86400 * int64(GetDaysToCurrentEpochFromUnixEpoch())
}

// Convert to days relative to "converter epoch"
func ConvertToDays(year int16, month uint8, day uint8) int32 {
	var yearPrime int16 = year // [0, 10000], begins on Mar 1
	var monthPrime uint8       // [0,11], Mar = 0
	if month <= 2 {
		yearPrime--
		monthPrime = month + 9
	} else {
		monthPrime = month - 3
	}

	var era uint16 = uint16(yearPrime) / 400           // [0,24]
	var yearOfEra uint16 = uint16(yearPrime) - 400*era // [0,399]
	var daysUntilMonthPrime uint16 = calcDaysUntilMonthPrime(monthPrime)
	var dayOfYearPrime uint16 = daysUntilMonthPrime + uint16(day) - 1 // [0,365]
	var dayOfEra uint32 = 365*uint32(yearOfEra) +
		(uint32(yearOfEra) / 4) - (uint32(yearOfEra) / 100) +
		uint32(dayOfYearPrime) // [0,146096]

	// epoch_prime days is relative to 0000-03-01
	var dayOfEpochPrime int32 = int32(dayOfEra + 146097*uint32(era))
	// relative to 2000-03-01
	dayOfEpochPrime -= (converterEpochYear / 400) * 146097
	// relative to 2000-01-01, 2000 is a leap year
	dayOfEpochPrime += 60

	return dayOfEpochPrime
}

// Convert from days relative to "converter epoch".
func ConvertFromDays(epochDays int32) (year int16, month uint8, day uint8) {
	// epoch_prime days is relative to 0000-03-01
	var dayOfEpochPrime int32 = epochDays +
		(converterEpochYear/400)*146097 - 60

	var era uint16 = uint16(uint32(dayOfEpochPrime) / 146097)          // [0,24]
	var dayOfEra uint32 = uint32(dayOfEpochPrime) - 146097*uint32(era) // [0,146096]
	var yearOfEra uint16 = uint16((dayOfEra - dayOfEra/1460 +
		dayOfEra/36524 -
		dayOfEra/146096) / 365) // [0,399]
	var yearPrime uint16 = yearOfEra + 400*era // [0,9999]
	var dayOfYearPrime uint16 = uint16(dayOfEra -
		(365*uint32(yearOfEra) + uint32(yearOfEra)/4 -
			uint32(yearOfEra)/100)) // [0,365]
	var monthPrime uint8 = uint8((5*dayOfYearPrime + 2) / 153)
	var daysUntilMonthPrime uint16 = calcDaysUntilMonthPrime(monthPrime)

	day = uint8(dayOfYearPrime - daysUntilMonthPrime + 1) // [1,31]
	if monthPrime < 10 {
		month = monthPrime + 3 // [1,12]
	} else {
		month = monthPrime - 9 // [1,12]
	}
	year = int16(yearPrime)
	if month <= 2 {
		year++ // [1,9999]
	}
	return
}

// convertToDaysUntilMonthPrime returns the number days in the year before the
// given monthPrime, where March=0, and the year ends with February=11.
func calcDaysUntilMonthPrime(monthPrime uint8) uint16 {
	return (153*uint16(monthPrime) + 2) / 5
}
