package acetime

// The type of the "seconds from epoch" in this library. This type is the
// equivalent of C lang `time_t`, or the Go lang [time.Time].
//
// Other versions of AceTimeXXX uses a 32-bit integer for the `Time` type
// because they want to support small 8-bit microcontrollers. For TinyGo, we
// assume we assume that acetimego will normally be used on 32-bit processors
// (e.g. ESP32), so we will use 64-bit integers. It increases the flash usage by
// only 100-200 bytes on the ESP32.
type Time int64

const (
	// The base epoch year used by the convertToDays() and convertFromDays()
	// functions below. This must be a multiple of 400.
	converterEpochYear = 2000

	// Number of days from 1970-01-01 to 2000-01-01. This is a constant because
	// the converterEpochYear is a constant.
	daysToConverterEpochFromUnixEpoch = 10957
)

// Convert to days relative to "converter epoch". From
// https://howardhinnant.github.io/date_algorithms.html and
// AceTime/EpochConverterHinnant.h.
func convertToDays(year int16, month uint8, day uint8) int32 {
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

// Convert from days relative to "converter epoch". From
// https://howardhinnant.github.io/date_algorithms.html and
// AceTime/EpochConverterHinnant.h.
func convertFromDays(converterDays int32) (year int16, month uint8, day uint8) {
	// epoch_prime days is relative to 0000-03-01
	var dayOfEpochPrime int32 = converterDays +
		(converterEpochYear/400)*146097 - 60

	var era uint16 = uint16(uint32(dayOfEpochPrime) / 146097) // [0,24]
	var dayOfEra uint32 = uint32(dayOfEpochPrime) -
		146097*uint32(era) // [0,146096]
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
