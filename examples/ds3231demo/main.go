// A demo program for the ds3231 package that knows how to set and read
// date/time fields from a DS3231 RTC module.

package main

import (
	"gitlab.com/bxparks/coding/tinygo/ds3231"
	"gitlab.com/bxparks/coding/tinygo/segwriter"
	"gitlab.com/bxparks/coding/tinygo/tm1637"
	"machine"
	"time"
	"tinygo.org/x/drivers/i2csoft"
)

const (
	numDigits   = 4
	delayMicros = 1
	brightness  = 2
)

func main() {
	tm := tm1637.New(machine.GPIO33, machine.GPIO32, delayMicros, numDigits)
	tm.Configure()
	tm.SetBrightness(brightness)
	numWriter := segwriter.NewNumberWriter(&tm)

	i2c := i2csoft.New(machine.SCL_PIN, machine.SDA_PIN)
	i2c.Configure(i2csoft.I2CConfig{Frequency: 400e3})
	rtc := ds3231.New(i2c)
	rtc.Configure()

	// Set Date
	dt := ds3231.DateTime{
		Year:   32,
		Month:  2,
		Day:    21,
		Hour:   21,
		Minute: 26,
		Second: 0,
	}
	rtc.SetTime(dt)

	for {
		dt, err := rtc.ReadTime()
		if err != nil {
			numWriter.WriteHexChar(0, segwriter.HexCharMinus)
			numWriter.WriteHexChar(1, segwriter.HexCharMinus)
			numWriter.WriteHexChar(2, segwriter.HexCharMinus)
			numWriter.WriteHexChar(3, segwriter.HexCharMinus)
			continue
		}

		numWriter.WriteDec2(0, dt.Hour, segwriter.HexChar(0))
		numWriter.WriteDec2(2, dt.Minute, segwriter.HexChar(0))
		tm.SetDecimalPoint(1, true)
		tm.Flush()
		time.Sleep(time.Millisecond * 1000)
	}
}
