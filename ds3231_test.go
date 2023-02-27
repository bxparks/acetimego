package ds3231

import (
	"testing"
)

func TestPositiveTemp(t *testing.T) {
	rawTemp := NewTemp(0, 0)
	temp := rawTemp.CentiC()
	if !(temp == 0) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3200) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == 25) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3245) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0, 0b10000000)
	temp = rawTemp.CentiC()
	if !(temp == 50) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3290) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == 75) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3335) {
		t.Error(temp)
	}

	rawTemp = NewTemp(1, 0)
	temp = rawTemp.CentiC()
	if !(temp == 100) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3380) {
		t.Error(temp)
	}

	rawTemp = NewTemp(1, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == 125) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3425) {
		t.Error(temp)
	}

	// 127.75C, highest temp possible on DS3231
	rawTemp = NewTemp(0x7f, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == 12775) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 26195) {
		t.Error(temp)
	}
}

func TestNegativeTemp(t *testing.T) {
	rawTemp := NewTemp(0xff, 0b11000000)
	temp := rawTemp.CentiC()
	if !(temp == -25) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3155) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0xff, 0b10000000)
	temp = rawTemp.CentiC()
	if !(temp == -50) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3110) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0xff, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == -75) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3065) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0xff, 0)
	temp = rawTemp.CentiC()
	if !(temp == -100) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3020) {
		t.Error(temp)
	}

	rawTemp = NewTemp(0xfe, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == -125) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 2975) {
		t.Error(temp)
	}

	// -128.00C, lowest temp possible on DS3231
	rawTemp = NewTemp(0x80, 0b00000000)
	temp = rawTemp.CentiC()
	if !(temp == -12800) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
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
