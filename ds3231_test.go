package ds3231

import (
	"testing"
)

func TestPositiveTemp(t *testing.T) {
	var rawTemp uint16 = 0 << 8 | 0
	temp := ToCentiC(rawTemp)
	if !(temp == 0) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3200) {
		t.Error(temp)
	}

	rawTemp = 0 << 8 | 0b01000000
	temp = ToCentiC(rawTemp)
	if !(temp == 25) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3245) {
		t.Error(temp)
	}

	rawTemp = 0 << 8 | 0b10000000
	temp = ToCentiC(rawTemp)
	if !(temp == 50) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3290) {
		t.Error(temp)
	}

	rawTemp = 0 << 8 | 0b11000000
	temp = ToCentiC(rawTemp)
	if !(temp == 75) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3335) {
		t.Error(temp)
	}

	rawTemp = 1 << 8 | 0
	temp = ToCentiC(rawTemp)
	if !(temp == 100) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3380) {
		t.Error(temp)
	}

	rawTemp = 1 << 8 | 0b01000000
	temp = ToCentiC(rawTemp)
	if !(temp == 125) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3425) {
		t.Error(temp)
	}

	// 127.75C, highest temp possible on DS3231
	rawTemp = 0x7f << 8 | 0b11000000
	temp = ToCentiC(rawTemp)
	if !(temp == 12775) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 26195) {
		t.Error(temp)
	}
}

func TestNegativeTemp(t *testing.T) {
	var rawTemp uint16 = 0xff << 8 | 0b11000000
	temp := ToCentiC(rawTemp)
	if !(temp == -25) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3155) {
		t.Error(temp)
	}

	rawTemp = 0xff << 8 | 0b10000000
	temp = ToCentiC(rawTemp)
	if !(temp == -50) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3110) {
		t.Error(temp)
	}

	rawTemp = 0xff << 8 | 0b01000000
	temp = ToCentiC(rawTemp)
	if !(temp == -75) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3065) {
		t.Error(temp)
	}

	rawTemp = 0xff << 8 | 0
	temp = ToCentiC(rawTemp)
	if !(temp == -100) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 3020) {
		t.Error(temp)
	}

	rawTemp = 0xfe << 8 | 0b11000000
	temp = ToCentiC(rawTemp)
	if !(temp == -125) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == 2975) {
		t.Error(temp)
	}

	// -128.00C, lowest temp possible on DS3231
	rawTemp = 0x80 << 8 | 0b00000000
	temp = ToCentiC(rawTemp)
	if !(temp == -12800) {
		t.Error(temp)
	}
	temp = ToCentiF(rawTemp)
	if !(temp == -19840) {
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
