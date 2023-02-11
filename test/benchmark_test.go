//-----------------------------------------------------------------------------
// Compare date-time conversions using:
//	* acetime.ZonedDateTime methods
//	* go standard library time.Date methods
//
// $ go test -run=^$ -bench=.
//-----------------------------------------------------------------------------

package test

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
	"time"
)

var (
	zoneManager = acetime.NewZoneManager(&zonedbtesting.DataContext)

	// Time zones for acetime and time packages.
	tz     = zoneManager.TimeZoneFromName("America/Los_Angeles")
	loc, _ = time.LoadLocation("America/Los_Angeles")
	//tz     = acetime.TimeZoneUTC
	//loc    = time.UTC

	// These define the limits of the nested for-loops, which controls how long
	// each "iteration" runs for.
	startYear  int16 = 1950
	untilYear  int16 = 2000
	startMonth uint8 = 2
	endMonth   uint8 = 6

	// Temporary variables to prevent the compiler from optimizing away the
	// for-loops that do nothing.
	zdt          acetime.ZonedDateTime
	ldt          acetime.LocalDateTime
	epochSeconds acetime.ATime
	atYear       int16
	atMonth      uint8
	atDay        uint8
	atHour       uint8
	atMinute     uint8
	atSecond     uint8

	// Temporary variables to prevent the compiler from optimizing away the
	// for-loops that do nothing.
	got           time.Time
	goYear        int
	goMonth       int
	goDay         int
	goHour        int
	goMinute      int
	goSecond      int
	goUnixSeconds int64
)

//-----------------------------------------------------------------------------

// Test date-time -> ZonedDateTime
func BenchmarkZonedDateTimeFromLocalDateTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := startYear; year < untilYear; year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0, 0 /*Fold*/}
					zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
					atYear = zdt.Year
					atMonth = zdt.Month
					atDay = zdt.Day
					atHour = zdt.Hour
					atMinute = zdt.Minute
					atSecond = zdt.Second
				}
			}
		}
	}
}

// Test date-time -> ZonedDateTime -> epochSeconds
func BenchmarkZonedDateTimeToEpochSeconds(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := startYear; year < untilYear; year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0, 0 /*Fold*/}
					zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
					atYear = zdt.Year
					atMonth = zdt.Month
					atDay = zdt.Day
					atHour = zdt.Hour
					atMinute = zdt.Minute
					atSecond = zdt.Second

					epochSeconds = zdt.EpochSeconds()
				}
			}
		}
	}
}

// Test date-time -> ZonedDateTime -> epochSeconds -> ZonedDateTime
func BenchmarkZonedDateTimeRoundTrip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := startYear; year < untilYear; year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0, 0 /*Fold*/}
					zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
					atYear = zdt.Year
					atMonth = zdt.Month
					atDay = zdt.Day
					atHour = zdt.Hour
					atMinute = zdt.Minute
					atSecond = zdt.Second

					epochSeconds = zdt.EpochSeconds()
					zdt = acetime.NewZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
				}
			}
		}
	}
}

//-----------------------------------------------------------------------------

// Test date-time -> Time
func BenchmarkGoTimeDate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := int(startYear); year < int(untilYear); year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					got = time.Date(
						int(year), time.Month(month), int(day), 9, 0, 0,
						0 /*nsec*/, loc)
					goYear = got.Year()
					goMonth = int(got.Month())
					goDay = got.Day()
					goHour = got.Hour()
					goMinute = got.Minute()
					goSecond = got.Second()
				}
			}
		}
	}
}

// Test date-time -> Time -> unixSeconds
func BenchmarkGoTimeUnixSeconds(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := int(startYear); year < int(untilYear); year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					got = time.Date(
						int(year), time.Month(month), int(day), 9, 0, 0,
						0 /*nsec*/, loc)
					goUnixSeconds = got.Unix()
				}
			}
		}
	}
}

// Test date-time -> Time -> unixSeconds -> Time
func BenchmarkGoTimeRoundTrip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for year := int(startYear); year < int(untilYear); year++ {
			for month := uint8(startMonth); month <= endMonth; month++ {
				for day := uint8(1); day <= 28; day++ {
					got = time.Date(
						int(year), time.Month(month), int(day), 9, 0, 0,
						0 /*nsec*/, loc)

					// Extract year and hours components. time.Time holds only the
					// epochSeconds, so each call to one of the component methods performs
					// a recalculation of all the breakout components, which is quite
					// inefficient.
					goYear = got.Year()
					goMonth = int(got.Month())
					goDay = got.Day()
					goHour = got.Hour()
					goMinute = got.Minute()
					goSecond = got.Second()

					goUnixSeconds = got.Unix()
					got = time.Unix(goUnixSeconds, 0)
				}
			}
		}
	}
}
