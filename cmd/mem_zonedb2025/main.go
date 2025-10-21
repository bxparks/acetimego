// Print the memory usage of acetime library when using the zonedb database.
//
//$ go run mem_zone_registry.go

package main

import (
	"github.com/bxparks/acetimego/acetime"
	"github.com/bxparks/acetimego/internal/strbuild"
	"github.com/bxparks/acetimego/zonedb2025"
	"runtime"
	"strings"
)

func main() {
	print("---- Initial memory usage\n")
	PrintMemUsage()

	print("---- Create ZonedDateTime using zonedb\n")
	zm := acetime.ZoneManagerFromDataContext(&zonedb2025.DataContext)
	name := "America/Los_Angeles"
	tz := zm.TimeZoneFromName(name)
	if tz.IsError() {
		print("ERROR: Could not find TimeZone for ")
		print(name)
		print("\n")
	}
	pdt := acetime.PlainDateTime{2023, 1, 19, 18, 36, 0}
	zdt := acetime.ZonedDateTimeFromPlainDateTime(
		&pdt, &tz, acetime.DisambiguateCompatible)
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
