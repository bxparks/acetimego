package acetime

import (
	"math"
	"testing"
	"unsafe"
)

// Verify my understanding of Golang's integer arithematics.
func TestInvalidUnixSeconds(t *testing.T) {
	if InvalidUnixSeconds != math.MinInt64 {
		t.Fatal(InvalidYear)
	}
}

func TestPlainDateTimeSize(t *testing.T) {
	pdt := PlainDateTime{2000, 1, 1, 1, 2, 3}
	size := unsafe.Sizeof(pdt)
	if !(size == 8) {
		t.Fatal("Sizeof(PlainDateTime): ", size)
	}
}

func TestPlainDateTimeIsError(t *testing.T) {
	pdt := PlainDateTime{2000, 1, 1, 0, 0, 0}
	if pdt.IsError() {
		t.Fatal("PlainDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be false")
	}

	pdt = PlainDateTime{InvalidYear, 1, 1, 0, 0, 0}
	if !pdt.IsError() {
		t.Fatal("PlainDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be true")
	}
}

func TestPlainDateTimeUnixSeconds(t *testing.T) {
	pdt := PlainDateTime{1970, 1, 1, 0, 0, 0}
	if pdt.UnixSeconds() != 0 {
		t.Fatal("PlainDateTime{1970, 1, 1, 0, 0, 0}.UnixSeconds() should be 0")
	}

	pdt = PlainDateTime{1970, 1, 1, 0, 0, 1}
	if pdt.UnixSeconds() != 1 {
		t.Fatal("PlainDateTime{1970, 1, 1, 0, 0, 1}.UnixSeconds() should be 1")
	}

	pdt = PlainDateTime{1971, 1, 1, 0, 0, 1}
	if pdt.UnixSeconds() != 86400*365+1 {
		t.Fatal(
			"PlainDateTime{1971, 1, 1, 0, 0, 1}.UnixSeconds() should be 31536001")
	}

	pdt = PlainDateTime{2050, 1, 1, 0, 0, 0}
	if pdt.UnixSeconds() != 2524608000 {
		t.Fatal(
			"PlainDateTime{1970, 1, 1, 0, 0, 0}.UnixSeconds() should be 2524608000")
	}
}

func TestPlainDateTimeFromUnixSeconds(t *testing.T) {
	pdt := PlainDateTimeFromUnixSeconds(0)
	if pdt.Year != 1970 || pdt.Month != 1 || pdt.Day != 1 ||
		pdt.Hour != 0 || pdt.Minute != 0 || pdt.Second != 0 {
		t.Fatal("PlainDateTimeFromUnixSeconds(0) should be "+
			"(1970, 1, 1, 0, 0, 0 but was", pdt)
	}
	pdt = PlainDateTimeFromUnixSeconds(1)
	if pdt.Year != 1970 || pdt.Month != 1 || pdt.Day != 1 ||
		pdt.Hour != 0 || pdt.Minute != 0 || pdt.Second != 1 {
		t.Fatal("PlainDateTimeFromUnixSeconds(1) should be "+
			"(1970, 1, 1, 0, 0, 1) but was", pdt)
	}
	pdt = PlainDateTimeFromUnixSeconds(86400*365 + 1)
	if pdt.Year != 1971 || pdt.Month != 1 || pdt.Day != 1 ||
		pdt.Hour != 0 || pdt.Minute != 0 || pdt.Second != 1 {
		t.Fatal("PlainDateTimeFromUnixSeconds(86400*365+1) should be "+
			"(1971, 1, 1, 0, 0, 1) but was", pdt)
	}
	pdt = PlainDateTimeFromUnixSeconds(2524608000 + 1)
	if pdt.Year != 2050 || pdt.Month != 1 || pdt.Day != 1 ||
		pdt.Hour != 0 || pdt.Minute != 0 || pdt.Second != 1 {
		t.Fatal("PlainDateTimeFromUnixSeconds(2524608000) should be "+
			"(2050, 1, 1, 0, 0, 1) but was", pdt)
	}
}

func TestPlainDateTimeToString(t *testing.T) {
	pdt := PlainDateTime{2023, 1, 19, 16, 9, 1}
	s := pdt.String()
	if !(s == "2023-01-19T16:09:01") {
		t.Fatal(s, pdt)
	}
}

func TestPlainDateTimeEquals(t *testing.T) {
	pdt1 := PlainDateTime{2023, 1, 19, 16, 9, 1}
	pdt2 := PlainDateTime{2023, 1, 19, 16, 9, 1}
	if !(pdt1 == pdt2) {
		t.Fatal("pdt1 != pdt2", pdt2)
	}
}
