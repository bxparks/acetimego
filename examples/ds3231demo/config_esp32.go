//go:build esp32

package main

import (
	"machine"
	"tinygo.org/x/drivers/i2csoft"
)

// I2C
var i2c = i2csoft.New(machine.SCL_PIN, machine.SDA_PIN)

func setupI2C() {
	i2c.Configure(i2csoft.I2CConfig{Frequency: 400e3})
}
