// Sample program to determine how Golang time package handles leap seconds.
//
// The package document explicitly says that leap seconds are not handled. Since
// the Time type stores only the number of seconds, it would be unable to
// distinguish between 23:59:59 and 23:59:60 during a leap second.
// It seems to normalize the 2016-12-31T23:59:60 into 2017-01-01T00:00:00, which
// makes sense because Time stores only the epochSeconds component.
//
// In acetimego/acetime, the ZonedDateTime struct stores the broken down
// date-time components, so in theory it may support leap seconds. However,
// acetimego also explicitly does not support leap seconds. The retains the
// broken down 23:59:60 time, but when converted to epochSeconds, it returns a
// value identical to 00:00:00.
//
// $ go run leapsecond.go
// ==== 2016 Leap second by Go time package
// 2016-12-31 23:59:59 +0000 UTC ; seconds= 1483228799
// 2017-01-01 00:00:00 +0000 UTC ; seconds= 1483228800
// 2017-01-01 00:00:00 +0000 UTC ; seconds= 1483228800
// ==== 2016 Leap second by acetime package
// 2016-12-31T23:59:59+00:00[UTC] ; seconds= 1483228799
// 2016-12-31T23:59:60+00:00[UTC] ; seconds= 1483228800
// 2017-01-01T00:00:00+00:00[UTC] ; seconds= 1483228800

package main

import (
	"github.com/bxparks/acetimego/acetime"
	"time"
)

const (
	name = "UTC"
	// name = "America/Los_Angeles"
)

func main() {
	leapGoTime()
	leapAceTime()
}

func leapAceTime() {
	println("==== 2016 Leap second by acetime package")
	atz := acetime.TimeZoneUTC

	ldt := acetime.LocalDateTime{2016, 12, 31, 23, 59, 59, 0 /*Fold*/}
	zdt := acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &atz)
	if zdt.IsError() {
		println("ERROR: ", name, ": Unable to create ZonedDateTime for ",
			ldt.String())
		return
	}
	seconds := zdt.EpochSeconds()
	println(zdt.String(), "; seconds=", seconds)

	ldt = acetime.LocalDateTime{2016, 12, 31, 23, 59, 60, 0 /*Fold*/}
	zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &atz)
	if zdt.IsError() {
		println("ERROR: ", name, ": Unable to create ZonedDateTime for ",
			ldt.String())
		return
	}
	seconds = zdt.EpochSeconds()
	println(zdt.String(), "; seconds=", seconds)

	ldt = acetime.LocalDateTime{2017, 1, 1, 0, 0, 0, 0 /*Fold*/}
	zdt = acetime.NewZonedDateTimeFromLocalDateTime(&ldt, &atz)
	if zdt.IsError() {
		println("ERROR: ", name, ": Unable to create ZonedDateTime for ",
			ldt.String())
		return
	}
	seconds = zdt.EpochSeconds()
	println(zdt.String(), "; seconds=", seconds)
}

func leapGoTime() {
	println("==== 2016 Leap second by Go time package")

	gotime := time.Date(2016, 12, 31, 23, 59, 59, 0, time.UTC)
	seconds := gotime.Unix()
	println(gotime.String(), "; seconds=", seconds)

	gotime = time.Date(2016, 12, 31, 23, 59, 60, 0, time.UTC)
	seconds = gotime.Unix()
	println(gotime.String(), "; seconds=", seconds)

	gotime = time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	seconds = gotime.Unix()
	println(gotime.String(), "; seconds=", seconds)
}
