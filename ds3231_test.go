package ds3231

import (
	"testing"
)

func TestPositiveTemperatures(t *testing.T) {
	temp := toMilliCelsius(0, 0)
	if !(temp == 0) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0, 0b01000000)
	if !(temp == 250) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0, 0b10000000)
	if !(temp == 500) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0, 0b11000000)
	if !(temp == 750) {
		t.Error(temp)
	}

	temp = toMilliCelsius(1, 0)
	if !(temp == 1000) {
		t.Error(temp)
	}

	temp = toMilliCelsius(1, 0b01000000)
	if !(temp == 1250) {
		t.Error(temp)
	}
}

func TestNegativeTemperatures(t *testing.T) {
	temp := toMilliCelsius(0xff, 0b11000000)
	if !(temp == -250) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0xff, 0b10000000)
	if !(temp == -500) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0xff, 0b01000000)
	if !(temp == -750) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0xff, 0)
	if !(temp == -1000) {
		t.Error(temp)
	}

	temp = toMilliCelsius(0xfe, 0b11000000)
	if !(temp == -1250) {
		t.Error(temp)
	}
}

func TestUnit8ToBCD(t *testing.T) {
	bcd := uint8ToBCD(0)
	if !(bcd == 0) {
		t.Error(bcd)
	}

	bcd = uint8ToBCD(11)
	if !(bcd == 0x11) {
		t.Error(bcd)
	}

	bcd = uint8ToBCD(99)
	if !(bcd == 0x99) {
		t.Error(bcd)
	}
}

func TestBCDToUint8(t *testing.T) {
	u := bcdToUint8(0)
	if !(u == 0) {
		t.Error(u)
	}

	u = bcdToUint8(0x11)
	if !(u == 11) {
		t.Error(u)
	}

	u = bcdToUint8(0x99)
	if !(u == 99) {
		t.Error(u)
	}
}
