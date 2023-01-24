// Validate AceTimeGo/acetime package against std.time package.
//
// $ go run validatetime.go

package main

import (
	"fmt"
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"time"
)

func main() {
	fmt.Println("Validating from 2000 to 2100")

	context := &zonedb.Context
	zm := acetime.NewZoneManager(context)
	var index int
	for _, zi := range zonedb.Context.ZoneRegistry {
		name := zi.Name(context.NameData, context.NameOffsets)
		fmt.Printf("[%3d] Zone: %s\n", index, name)
		validateZoneName(&zm, name)
		index++
	}
}

func validateZoneName(zm *acetime.ZoneManager, name string) {
	atz := zm.NewTimeZoneFromName(name)
	if atz.IsError() {
		fmt.Println("ERROR: acetime package: Zone", name, "not found")
		return
	}

	stz, err := time.LoadLocation(name)
	if err != nil {
		fmt.Println("ERROR: time package: Zone", name, "not found")
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
					fmt.Printf("ERROR: %s: Unable to create ZonedDateTime for %s\n",
						name, ldt)
					return
				}

				// Create time.Time
				unixSeconds64 := zdt.ToUnixSeconds64()
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
					fmt.Printf("ERROR: %s: %s: Year not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
				if uint8(st.Month()) != zdt.Month {
					fmt.Printf("ERROR: %s: %s: Month not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
				if uint8(st.Day()) != zdt.Day {
					fmt.Printf("ERROR: %s: %s: Day not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
				if uint8(st.Hour()) != zdt.Hour {
					fmt.Printf("ERROR: %s: %s: Hour not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
				if uint8(st.Minute()) != zdt.Minute {
					fmt.Printf("ERROR: %s: %s: Minute not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
				if uint8(st.Second()) != zdt.Second {
					fmt.Printf("ERROR: %s: %s: Second not equal: %s, %s\n",
						name, st, st, zdt.String())
					return
				}
			}
		}
	}
}
