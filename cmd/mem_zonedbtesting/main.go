// Print the memory usage of acetime library when using the zonedbtesting
// database.
//
//$ go run mem_zone_registry.go

package main

import (
	"github.com/bxparks/acetimego/acetime"
	"github.com/bxparks/acetimego/internal/strbuild"
	"github.com/bxparks/acetimego/zonedbtesting"
	"runtime"
	"strings"
)

func main() {
	print("---- Initial memory usage\n")
	PrintMemUsage()

	print("---- Create ZonedDateTime using zonedbtesting\n")
	zm := acetime.NewZoneManager(&zonedbtesting.DataContext)
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
	strbuild.Uint64(&b, m.TotalAlloc)
	print(b.String())
	print("\tSys = ")
	b.Reset()
	strbuild.Uint64(&b, m.Sys)
	print(b.String())
	print("\n")

	// These are not found on tinygo.
	//fmt.Printf("Alloc = %v B", m.Alloc)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
