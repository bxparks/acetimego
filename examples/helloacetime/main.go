// A demo program for the acetimego package to create a ZonedDateTime object for
// source timezone, converts it to the destination timezone, and prints both
// datetime objects in ISO8601 format.

package main

import (
	"github.com/bxparks/acetimego/acetime"
	"github.com/bxparks/acetimego/zonedb"
	"time"
)

const (
	srcName = "America/Los_Angeles"
	dstName = "Europe/Paris"
)

func main() {
	zm := acetime.ZoneManagerFromDataContext(&zonedb.DataContext)
	srcTz := zm.TimeZoneFromName(srcName)
	if srcTz.IsError() {
		print("ERROR: Could not find TimeZone for ")
		print(srcName)
		print("\n")
		return
	}

	dstTz := zm.TimeZoneFromName(dstName)
	if dstTz.IsError() {
		print("ERROR: Could not find TimeZone for ")
		print(dstName)
		print("\n")
		return
	}

	for {
		now := time.Now().Unix()
		aceNow := acetime.Time(now)
		srcZdt := acetime.ZonedDateTimeFromEpochSeconds(aceNow, &srcTz)
		dstZdt := srcZdt.ConvertToTimeZone(&dstTz)
		print("src:")
		print(srcZdt.String())
		print("; dst:")
		print(dstZdt.String())
		print("\n")

		time.Sleep(2 * time.Second)
	}
}
