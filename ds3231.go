// Package ds3231 provides a driver for the DS3231 RTC
//
// Datasheet:
// https://datasheets.maximintegrated.com/en/ds/DS3231.pdf
package ds3231 // import "tinygo.org/x/drivers/ds3231"

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
type DateTime struct {
	Year    uint8
	Month   uint8 // [1, 12]
	Day     uint8 // [1, 31]
	Hour    uint8 // [0, 23]
	Minute  uint8 // [0, 59]
	Second  uint8 // [0, 59]
	Weekday uint8 // [1, 7], interpretation undefined, increments every day
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

	var tdata = [7]uint8{
		uint8ToBCD(dt.Second),
		uint8ToBCD(dt.Minute),
		uint8ToBCD(dt.Hour),
		uint8ToBCD(dt.Weekday),
		uint8ToBCD(dt.Day),
		uint8ToBCD(dt.Month),
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

	dt = DateTime{
		Second:  bcdToUint8(data[0] & 0x7F),
		Minute:  bcdToUint8(data[1]),
		Hour:    bcdToUint8(data[2] & 0x3F),
		Weekday: data[3],
		Day:     bcdToUint8(data[4]),
		Month:   bcdToUint8(data[5] & 0x7F),
		Year:    bcdToUint8(data[6]),
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
	return int32(data[0])*1000 + int32((data[1]>>6)*25)*10, nil
}

// uint8ToBCD converts a byte to BCD for the DS3231
func uint8ToBCD(value uint8) uint8 {
	return value + 6*(value/10)
}

// bcdToUint8 converts BCD from the DS3231 to int
func bcdToUint8(value uint8) uint8 {
	return value - 6*(value>>4)
}
