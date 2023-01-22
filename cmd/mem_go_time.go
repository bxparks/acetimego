// Print the memory usage when using the standard Go time package.
//
//$ go run mem_go_time.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	os.Stdout.WriteString("---- Initial memory usage\n")
	PrintMemUsage()

	os.Stdout.WriteString("---- Create America/Los_Angeles using time package\n")
	name := "America/Los_Angeles"
	tz, err := time.LoadLocation(name)
	if err != nil {
		os.Stdout.WriteString("ERROR: Zone not found: ")
		os.Stdout.WriteString(name)
		os.Stdout.WriteString("\n")
		return
	}
	t := time.Date(2023, 1, 19, 18, 36, 0, 0 /*nanos*/, tz)
	os.Stdout.WriteString("t:")
	os.Stdout.WriteString(t.String())
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
