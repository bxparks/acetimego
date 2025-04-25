// A demo program to check the 2 special TimeZone objects: TimeZoneUTC and
// TimeZoneError.

package main

import (
	"github.com/bxparks/acetimego/acetime"
)

func main() {
	// TimeZoneUTC
	tzUTC := acetime.TimeZoneUTC

	print("Created TimeZoneUTC\n")
	print("- IsUTC():")
	isUTC := tzUTC.IsUTC()
	print(isUTC)
	print("\n")

	print("- IsLink():")
	isLink := tzUTC.IsLink()
	print(isLink)
	print("\n")

	// TimeZoneError
	tzErr := acetime.TimeZoneError

	print("Created TimeZoneError\n")
	print("- IsUTC():")
	isUTC = tzErr.IsUTC()
	print(isUTC)
	print("\n")

	print("- IsLink():")
	isLink = tzErr.IsLink()
	print(isLink)
	print("\n")

	return
}
