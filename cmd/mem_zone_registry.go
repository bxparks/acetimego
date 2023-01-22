// Print the memory usage of various acetime data structs, such as ZonedDateTime
// and ZoneAndLinkRegistry.
//
//$ go run printmemory.go

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

	os.Stdout.WriteString("---- Load the ZoneAndLinkRegistry\n")
	zr := acetime.ZoneRegistrar{zonedb.ZoneAndLinkRegistry}
	PrintMemUsage()

	os.Stdout.WriteString("---- Create ZonedDateTime using Registry\n")
	zi := zr.FindZoneInfoByName("US/Pacific")
	if zi == nil {
		os.Stdout.WriteString("US/Pacific not found\n")
		return
	}
	tz := acetime.NewTimeZoneFromZoneInfo(zi)
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
