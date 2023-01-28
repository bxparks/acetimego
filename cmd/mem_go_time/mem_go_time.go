// Print the memory usage when using the standard Go time package.
//
//$ go run mem_go_time.go

package main

import (
	"github.com/bxparks/AceTimeGo/acetime"
	"runtime"
	"strings"
	"time"
)

func main() {
	print("---- Initial memory usage\n")
	PrintMemUsage()

	print("---- Create America/Los_Angeles using time package\n")
	name := "America/Los_Angeles"
	tz, err := time.LoadLocation(name)
	if err != nil {
		print("ERROR: Zone not found: ")
		print(name)
		print("\n")
		return
	}
	t := time.Date(2023, 1, 19, 18, 36, 0, 0 /*nanos*/, tz)
	print("t:")
	print(t.String())
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
	acetime.WriteUint64(&b, m.TotalAlloc)
	print(b.String())
	print("\tSys = ")
	b.Reset()
	acetime.WriteUint64(&b, m.Sys)
	print(b.String())
	print("\n")

	// These are not found on tinygo.
	//fmt.Printf("Alloc = %v B", m.Alloc)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
