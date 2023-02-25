// Package ds3231 provides a driver for the DS3231 RTC
//
// Datasheet:
// https://datasheets.maximintegrated.com/en/ds/DS3231.pdf
package ds3231

import (
	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to a DS3231 device.
type Device struct {
	bus     drivers.I2C
	address uint8
}

// The date (year, month, day) and time (hour, minute, second) fields supported
// by the DS3231 RTC chip.
//
// The `Century` field corresponds to bit-7 of the 05h register of the DS3231
// chip. According to the datasheet, this bit is set when the 2-digit year
// component overflows from 99 to 00. It can be set to 1 and the RTC will retain
// this value when battery backup is enabled.
//
// The practical utility of this bit is severely limited because it does not
// affect the leap year calculation. The year 00 is assumed to be a leap year,
// regardless of the value of the `Century` bit. This works for the year 2000
// because 2000 was a special case in the Gregorian calendar system. But the
// year 2100 is *not* a leap year, and the `Century` bit has no effect on the
// calculation, so the DS3231 increments from Feb 28 to Feb 29 even when this
// bit is set. Therefore, the `Century` bit cannot interpreted as being the year
// 2100. Most implementations shoulr prbably ignore the `Century` bit. It is
// exposed by this package for completeness in case it is useful to the calling
// program.
type DateTime struct {
	Year    uint8 // [0, 99]
	Month   uint8 // [1, 12]
	Day     uint8 // [1, 31]
	Hour    uint8 // [0, 23]
	Minute  uint8 // [0, 59]
	Second  uint8 // [0, 59]
	Weekday uint8 // [1, 7], interpretation undefined, increments every day
	Century uint8 // [0, 1], set when Year overflows from 99 to 00
}

// New creates a new DS3231 connection. The I2C bus must already be
// configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{
		bus:     bus,
		address: Address,
	}
}

// Configure sets up the device for communication
func (d *Device) Configure() bool {
	return true
}

// SetTime sets the date and time in the DS3231
func (d *Device) SetTime(dt DateTime) error {
	data := [1]byte{0}
	err := d.bus.ReadRegister(uint8(d.address), REG_STATUS, data[:])
	if err != nil {
		return err
	}
	data[0] &^= 1 << OSF
	err = d.bus.WriteRegister(uint8(d.address), REG_STATUS, data[:])
	if err != nil {
		return err
	}

	var century uint8
	if dt.Century != 0 {
		century = 0x80
	}
	var tdata = [7]uint8{
		uint8ToBCD(dt.Second),
		uint8ToBCD(dt.Minute),
		uint8ToBCD(dt.Hour),
		uint8ToBCD(dt.Weekday),
		uint8ToBCD(dt.Day),
		uint8ToBCD(dt.Month) | century,
		uint8ToBCD(dt.Year),
	}

	err = d.bus.WriteRegister(uint8(d.address), REG_TIMEDATE, tdata[:])
	return err
}

// ReadTime returns the date and time
func (d *Device) ReadTime() (dt DateTime, err error) {
	var data [7]uint8
	err = d.bus.ReadRegister(uint8(d.address), REG_TIMEDATE, data[:])
	if err != nil {
		return
	}

	var century uint8
	if data[5] & 0x80 != 0 {
		century = 1
	} else {
		century = 0
	}
	dt = DateTime{
		Second:  bcdToUint8(data[0] & 0x7F),
		Minute:  bcdToUint8(data[1]),
		Hour:    bcdToUint8(data[2] & 0x3F),
		Weekday: data[3],
		Day:     bcdToUint8(data[4]),
		Month:   bcdToUint8(data[5] & 0x7F),
		Year:    bcdToUint8(data[6]),
		Century: century,
	}
	return
}

// ReadTemperature returns the temperature in millicelsius (mC)
func (d *Device) ReadTemperature() (int32, error) {
	var data [2]uint8
	err := d.bus.ReadRegister(uint8(d.address), REG_TEMP, data[:])
	if err != nil {
		return 0, err
	}
	return toMilliCelsius(data[0], data[1]), nil
}

// Convert raw temperature bytes to Celsius in units of 1/1000.
//
// According to the DS3231 datasheet, "The temperature is encoded in two's
// complement format. The upper 8 bits, the integer portion, are at location 11h
// and the lower 2 bits, the fractional portion, are in the upper nibble at
// location 12h. For example, 00011001 01b = +25.25C."
//
// One way to think about this format is a signed 8.8 fixed point type, where
// the `msb` represents the integer portion, and the `lsb` represents the
// fractional portion. The second, equivalent interpretation is to consider the
// (msb, lsb) pair as a signed 16-bit integer representing temperature in units
// of (1/256) degrees Celsius. We can convert this into an integer in units of
// (1/1000) degrees Celsius without loss of information because the DS3231 only
// uses the top 2 bits of the `lsb` portion.
func toMilliCelsius(msb uint8, lsb uint8) int32 {
	t256 := int16((uint16(msb) << 8) | uint16(lsb)) // units of (1/256) Celsius
	t1000 := int32(t256) * 1000 / 256 // always integral, same as (* 250 / 64)
	return t1000 // units of (1/1000) Celsius
}

// uint8ToBCD converts a byte to BCD for the DS3231
func uint8ToBCD(value uint8) uint8 {
	return value + 6*(value/10)
}

// bcdToUint8 converts BCD from the DS3231 to int
func bcdToUint8(value uint8) uint8 {
	return value - 6*(value>>4)
}
