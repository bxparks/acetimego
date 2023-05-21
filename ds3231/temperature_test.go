package ds3231

import (
	"testing"
)

func TestPositiveTemperature(t *testing.T) {
	rawTemp := NewTemperature(0, 0)
	temp := rawTemp.CentiC()
	if !(temp == 0) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3200) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == 25) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3245) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0, 0b10000000)
	temp = rawTemp.CentiC()
	if !(temp == 50) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3290) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == 75) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3335) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(1, 0)
	temp = rawTemp.CentiC()
	if !(temp == 100) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3380) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(1, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == 125) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3425) {
		t.Error(temp)
	}

	// 127.75C, highest temp possible on DS3231
	rawTemp = NewTemperature(0x7f, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == 12775) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 26195) {
		t.Error(temp)
	}
}

func TestNegativeTemperature(t *testing.T) {
	rawTemp := NewTemperature(0xff, 0b11000000)
	temp := rawTemp.CentiC()
	if !(temp == -25) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3155) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0xff, 0b10000000)
	temp = rawTemp.CentiC()
	if !(temp == -50) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3110) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0xff, 0b01000000)
	temp = rawTemp.CentiC()
	if !(temp == -75) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3065) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0xff, 0)
	temp = rawTemp.CentiC()
	if !(temp == -100) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 3020) {
		t.Error(temp)
	}

	rawTemp = NewTemperature(0xfe, 0b11000000)
	temp = rawTemp.CentiC()
	if !(temp == -125) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == 2975) {
		t.Error(temp)
	}

	// -128.00C, lowest temp possible on DS3231
	rawTemp = NewTemperature(0x80, 0b00000000)
	temp = rawTemp.CentiC()
	if !(temp == -12800) {
		t.Error(temp)
	}
	temp = rawTemp.CentiF()
	if !(temp == -19840) {
		t.Error(temp)
	}
}
