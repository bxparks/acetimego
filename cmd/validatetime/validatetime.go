// Validate AceTimeGo/acetime package against std.time package.
//
// $ go run validatetime.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"time"
)

func main() {
	println("Validating from 2000 to 2100")

	context := &zonedb.DataContext
	zm := acetime.NewZoneManager(context)
	names := zm.ZoneNames()
	for i, name := range names {
		println("[", i, "] Zone", name)
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

	var year int16
	var month uint8
	var day uint8
	for year = 2000; year < 2100; year++ {
		for month = 1; month <= 12; month++ {
			for day = 1; day <= 28; day++ {
				// Create acetime.ZonedDateTime
				ldt := acetime.LocalDateTime{year, month, day, 2, 3, 4, 0 /*Fold*/}
				zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &atz)
				if zdt.IsError() {
					println("ERROR: ", name, ": Unable to create ZonedDateTime for ",
						ldt.String())
					return
				}

				// Create time.Time
				unixSeconds64 := zdt.UnixSeconds64()
				st := time.Unix(unixSeconds64, 0 /*nanos*/).In(stz)

				// Validate components. The time.Time struct holds just a counter
				// (nanoseconds since a specific date-time), and the timezone. It does
				// NOT hold the date-time components like acetime.ZonedDateTime.
				// Therefore, each call to one of the component methods (e.g. Year(),
				// Month(), etc) causes a conversion from this counter to the
				// human-readable date-time components, which is a relatively slow
				// process. Each of the following if-statement causes the program to
				// become slower and slower.
				if int16(st.Year()) != zdt.Year {
					println("ERROR: ", name, ": Year not equal: ",
						st.String(), zdt.String())
					return
				}
				if uint8(st.Month()) != zdt.Month {
					println("ERROR: ", name, ": Month not equal: ",
						st.String(), zdt.String())
					return
				}
				if uint8(st.Day()) != zdt.Day {
					println("ERROR: ", name, ": Day not equal: ",
						st.String(), zdt.String())
					return
				}
				if uint8(st.Hour()) != zdt.Hour {
					println("ERROR: ", name, ": Hour not equal: ",
						st.String(), zdt.String())
					return
				}
				if uint8(st.Minute()) != zdt.Minute {
					println("ERROR: ", name, ": Minute not equal: ",
						st.String(), zdt.String())
					return
				}
				if uint8(st.Second()) != zdt.Second {
					println("ERROR: ", name, ": Second not equal: ",
						st.String(), zdt.String())
					return
				}
			}
		}
	}
}
