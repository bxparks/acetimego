//-----------------------------------------------------------------------------
// Compare date-time conversions using:
//	* acetime.ZonedDateTime methods
//	* go standard library time.Date methods
//
// Usage:
//
// $ go test -run=^$ -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/bxparks/acetimego/test
// cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
// BenchmarkZonedDateTimeFromLocalDateTime-4           3037        378853 ns/op
// BenchmarkZonedDateTimeToEpochSeconds-4              2800        414800 ns/op
// BenchmarkZonedDateTimeRoundTrip-4                   1225        964319 ns/op
// BenchmarkGoTimeDate-4                                781       1534750 ns/op
// BenchmarkGoTimeUnixSeconds-4                        4754        245158 ns/op
// BenchmarkGoTimeRoundTrip-4                           751       1552672 ns/op
// PASS
// ok      github.com/bxparks/acetimego/test       7.557s
//-----------------------------------------------------------------------------

package test

import (
	"github.com/bxparks/acetimego/acetime"
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
	"time"
)

var (
	zoneManager = acetime.ZoneManagerFromDataContext(&zonedbtesting.DataContext)

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
	epochSeconds acetime.Time
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
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0}
					zdt = acetime.ZonedDateTimeFromLocalDateTime(
						&ldt, &tz, acetime.DisambiguateCompatible)
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
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0}
					zdt = acetime.ZonedDateTimeFromLocalDateTime(
						&ldt, &tz, acetime.DisambiguateCompatible)
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
					ldt = acetime.LocalDateTime{year, month, day, 9, 0, 0}
					zdt = acetime.ZonedDateTimeFromLocalDateTime(
						&ldt, &tz, acetime.DisambiguateCompatible)
					atYear = zdt.Year
					atMonth = zdt.Month
					atDay = zdt.Day
					atHour = zdt.Hour
					atMinute = zdt.Minute
					atSecond = zdt.Second

					epochSeconds = zdt.EpochSeconds()
					zdt = acetime.ZonedDateTimeFromEpochSeconds(epochSeconds, &tz)
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
