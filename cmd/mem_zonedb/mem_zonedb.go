// Print the memory usage of acetime library when using the zonedb database.
//
//$ go run mem_zone_registry.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"runtime"
	"strings"
)

func main() {
	print("---- Initial memory usage\n")
	PrintMemUsage()

	print("---- Create ZonedDateTime using zonedb\n")
	zm := acetime.NewZoneManager(&zonedb.DataContext)
	name := "America/Los_Angeles"
	tz := zm.TimeZoneFromName(name)
	if tz.IsError() {
		print("ERROR: Could not find TimeZone for ")
		print(name)
		print("\n")
	}
	ldt := acetime.LocalDateTime{2023, 1, 19, 18, 36, 0, 0 /*Fold*/}
	zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	print("zdt:")
	print(zdt.String())
	print("\n")
	PrintMemUsage()

	print("---- Run GC()\n")
	runtime.GC()
	PrintMemUsage()
}

func PrintMemUsage() {
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var b strings.Builder
	print("TotalAlloc = ")
	acetime.BuildUint64(&b, m.TotalAlloc)
	print(b.String())
	print("\tSys = ")
	b.Reset()
	acetime.BuildUint64(&b, m.Sys)
	print(b.String())
	print("\n")

	// These are not found on tinygo.
	//fmt.Printf("Alloc = %v B", m.Alloc)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
