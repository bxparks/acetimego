package ds3231

// The raw temperature value from the DS3231. Two 8-bit integers merged into a
// single signed 16-bit integer in units of 1/256 degrees Celsius.
type Temperature int16

// Create a new Temperature instance from the msb and lsb bytes.
func NewTemperature(msb uint8, lsb uint8) Temperature {
	return Temperature(uint16(msb)<<8 | uint16(lsb))
}

// Convert the raw temperature readings (units of 1/256 Celsius) to centi
// Celsius (units of 0.01C). The DS3231 has a precision of 2 bits after the
// decimal point, in other words, 0.25C. The lowest temperature is -128.00C. The
// highest temperature is 127.75C.
//
// According to the DS3231 datasheet: "The temperature is encoded in two's
// complement format. The upper 8 bits, the integer portion, are at location 11h
// and the lower 2 bits, the fractional portion, are in the upper nibble at
// location 12h. For example, 00011001 01b = +25.25C."
//
// This format is a signed 8.8 fixed point type, where the `msb` represents the
// integer portion, and the `lsb` represents the fractional portion.
// Equivalently, we can consider the (msb, lsb) pair as a signed 16-bit integer
// representing temperature in units of (1/256) degrees Celsius. We can convert
// this into an integer in units of (1/100) degrees Celsius without loss of
// information because the DS3231 only uses the top 2 bits of the `lsb` portion.
func (temp Temperature) CentiC() int16 {
	c100 := int16(temp) / 64 * 25 // (*100/256), always integral
	return c100
}

// Convert raw temperature reading (units of 1/256 Celsius) into centi
// Fahrenheit (units of 0.01F). The DS3231 has a precision of 2 bits after the
// decimal point in Celsius, which corresponds to 0.45F. The lowest temperature
// is -198.40F. The highest temperature is 261.95F.
func (temp Temperature) CentiF() int16 {
	c100 := temp.CentiC()
	f100 := c100/5*9 + 3200 // always integral, with no loss of bits
	return f100
}
