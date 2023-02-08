// Validate AceTimeGo/acetime package against std.time package.
//
// $ go run validatetime.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"time"
)

const (
	startYear = 1974
	untilYear = 2150
	samplingInterval = 22 // hours
)

func main() {
	println("Validating from", startYear, "until", untilYear)

	context := &zonedb.DataContext
	zm := acetime.NewZoneManager(context)
	names := zm.ZoneNames()
	for i, name := range names {
		print("[", i, "] Zone ", name, ": ")
		validateZoneName(&zm, name)
	}
}

func validateZoneName(zm *acetime.ZoneManager, name string) {
	atz := zm.NewTimeZoneFromName(name)
	if atz.IsError() {
		println("ERROR: acetime package: Zone", name, "not found")
		return
	}

	stz, err := time.LoadLocation(name)
	if err != nil {
		println("ERROR: time package: Zone", name, "not found")
		return
	}

	// Validate the before and after DST transitions.
	transitions := findTransitions(startYear, untilYear, samplingInterval, stz)
	for _, transition := range transitions {
		validateAtTime(transition.before, &atz)
		validateAtTime(transition.after, &atz)
	}

	// Validate some samples
	samples := 0
	for year := startYear; year < untilYear; year++ {
		for month := 1; month <= 12; month++ {
			for day := 1; day <= 28; day++ {
				gt := time.Date(year, time.Month(month), day, 2, 0, 0, 0, stz)
				validateAtTime(gt, &atz)
				samples++
			}
		}
	}

	println("Transitions:", len(transitions), "; Samples:", samples)
}

func validateAtTime(t time.Time, atz *acetime.TimeZone) {
	name := atz.Name()

	// Create acetime.ZonedDateTime based on the EpochSeconds, which uniquely
	// identifies the time. We can't use the components because those can be
	// ambiguous during an overlap.
	unixSeconds := t.Unix()
	zdt := acetime.NewZonedDateTimeFromEpochSeconds(
		acetime.ATime(unixSeconds), atz)
	if zdt.IsError() {
		println("ERROR: ", name, ": Unable to create ZonedDateTime for ",
			unixSeconds)
		return
	}

	year := int16(t.Year())
	month := uint8(t.Month())
	day := uint8(t.Day())
	hour := uint8(t.Hour())
	minute := uint8(t.Minute())
	second := uint8(t.Second())

	// Validate UnixSeconds. This should always succeed.
	acetimeEpochSeconds := int64(zdt.EpochSeconds())
	if acetimeEpochSeconds != unixSeconds {
		println("ERROR: ", name, ": EpochSeconds not equal",
			acetimeEpochSeconds, unixSeconds)
		return
	}

	// Validate components. The time.Time struct holds just a counter
	// (nanoseconds since a specific date-time), and the timezone. It does
	// NOT hold the date-time components like acetime.ZonedDateTime.
	// Therefore, each call to one of the component methods (e.g. Year(),
	// Month(), etc) causes a conversion from this counter to the
	// human-readable date-time components, which is a relatively slow
	// process. Each of the following if-statement causes the program to
	// become slower and slower.
	if int16(year) != zdt.Year {
		println("ERROR: ", name, ": Year not equal: ",
			t.String(), zdt.String())
		return
	}
	if uint8(month) != zdt.Month {
		println("ERROR: ", name, ": Month not equal: ",
			t.String(), zdt.String())
		return
	}
	if uint8(day) != zdt.Day {
		println("ERROR: ", name, ": Day not equal: ",
			t.String(), zdt.String())
		return
	}
	if uint8(hour) != zdt.Hour {
		println("ERROR: ", name, ": Hour not equal: ",
			t.String(), zdt.String())
		return
	}
	if uint8(minute) != zdt.Minute {
		println("ERROR: ", name, ": Minute not equal: ",
			t.String(), zdt.String())
		return
	}
	if uint8(second) != zdt.Second {
		println("ERROR: ", name, ": Second not equal: ",
			t.String(), zdt.String())
		return
	}
}
