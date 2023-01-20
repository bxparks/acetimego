// Print the memory usage of various acetime data structs, such as ZonedDateTime
// and ZoneAndLinkRegistry.
//
//$ go run printmemory.go

package main

import (
	"fmt"
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"runtime"
)

func main() {
	fmt.Println("---- Initial memory usage")
	PrintMemUsage()

	fmt.Println("---- Create ZonedDateTime using ZoneInfo")
	tz := acetime.NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	ldt := acetime.LocalDateTime{2023, 1, 19, 18, 36, 0, 0 /*Fold*/}
	zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	fmt.Println("zdt1:", zdt.String())
	PrintMemUsage()

	fmt.Println("---- Load the ZoneAndLinkRegistry")
	zr := acetime.ZoneRegistrar{zonedb.ZoneAndLinkRegistry}
	PrintMemUsage()

	fmt.Println("---- Create ZonedDateTime using Registry")
	zi := zr.FindZoneInfoByName("US/Pacific")
	if zi == nil {
		fmt.Println("US/Pacific not found")
		return
	}
	tz = acetime.NewTimeZoneFromZoneInfo(zi)
	zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	fmt.Println("zdt2:", zdt.String())
	PrintMemUsage()

	fmt.Println("---- Run GC()")
	runtime.GC()
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("TotalAlloc = %v B", m.TotalAlloc)
	fmt.Printf("\tSys = %v B\n", m.Sys)

	// These are not found on tinygo.
	//fmt.Printf("Alloc = %v B", m.Alloc)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
