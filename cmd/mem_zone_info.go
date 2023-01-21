// Print the memory usage of various acetime data structs, such as ZonedDateTime
// and ZoneAndLinkRegistry.
//
//$ go run printmemory.go

package main

import (
	"fmt"
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"runtime"
)

func main() {
	fmt.Println("---- Initial memory usage")
	PrintMemUsage()

	fmt.Println("---- Create ZonedDateTime using ZoneInfo")
	tz := acetime.NewTimeZoneFromZoneInfo(&zonedb.ZoneAmerica_Los_Angeles)
	ldt := acetime.LocalDateTime{2023, 1, 19, 18, 36, 0, 0 /*Fold*/}
	zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	fmt.Println("zdt:", zdt.String())
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
