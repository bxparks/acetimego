package ds3231

// uint8ToBCD converts a byte to BCD for the DS3231
func uint8ToBCD(value uint8) uint8 {
	return value + 6*(value/10)
}

// bcdToUint8 converts BCD from the DS3231 to a byte
func bcdToUint8(value uint8) uint8 {
	return value - 6*(value>>4)
}
