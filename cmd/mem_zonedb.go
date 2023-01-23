// Print the memory usage of acetime library when using the zonedb database.
//
//$ go run mem_zone_registry.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"github.com/bxparks/AceTimeGo/zonedb"
	"os"
	"runtime"
	"strings"
)

func main() {
	os.Stdout.WriteString("---- Initial memory usage\n")
	PrintMemUsage()

	os.Stdout.WriteString("---- Create ZonedDateTime using zonedb\n")
	zm := acetime.NewZoneManager(&zonedb.Context)
	name := "America/Los_Angeles"
	tz := zm.NewTimeZoneFromName(name)
	if tz.IsError() {
		os.Stdout.WriteString("ERROR: Could not find TimeZone for ")
		os.Stdout.WriteString(name)
		os.Stdout.WriteString("\n")
	}
	ldt := acetime.LocalDateTime{2023, 1, 19, 18, 36, 0, 0 /*Fold*/}
	zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &tz)
	os.Stdout.WriteString("zdt:")
	os.Stdout.WriteString(zdt.String())
	os.Stdout.WriteString("\n")
	PrintMemUsage()

	os.Stdout.WriteString("---- Run GC()\n")
	runtime.GC()
	PrintMemUsage()
}

func PrintMemUsage() {
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var b strings.Builder
	os.Stdout.WriteString("TotalAlloc = ")
	acetime.WriteUint64(&b, m.TotalAlloc)
	os.Stdout.WriteString(b.String())
	os.Stdout.WriteString("\tSys = ")
	b.Reset()
	acetime.WriteUint64(&b, m.Sys)
	os.Stdout.WriteString(b.String())
	os.Stdout.WriteString("\n")

	// These are not found on tinygo.
	//fmt.Printf("Alloc = %v B", m.Alloc)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
