//go:build tinygo

//
// A demo program for the ds3231 package to set and read the date/time fields
// from a DS3231 RTC module. The date-time is printed to the serial monitor
// every second.

package main

import (
	"github.com/bxparks/AceTimeGo/ds3231"
	"time"
)

var (
	rtc = ds3231.New(i2c)
)

func main() {
	setupI2C()
	rtc.Configure()

	// Set Date
	dt := ds3231.DateTime{
		Year:   23,
		Month:  2,
		Day:    21,
		Hour:   21,
		Minute: 26,
		Second: 0,
	}
	rtc.SetTime(dt)

	year := 2000 + int16(dt.Year)
	for {
		dt, err := rtc.ReadTime()

		if err == nil {
			println(year, "-", dt.Month, "-", dt.Day, " ",
				dt.Hour, ":", dt.Minute, ":", dt.Second)
		} else {
			println("Err")
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
