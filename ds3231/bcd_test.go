package ds3231

import (
	"testing"
)

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
