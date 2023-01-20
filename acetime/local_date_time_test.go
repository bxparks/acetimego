package acetime

import (
	"testing"
	"unsafe"
)

func TestLocalDateTimeSize(t *testing.T) {
	ldt := LocalDateTime{2000, 1, 1, 1, 2, 3, 0 /*Fold*/}
	size := unsafe.Sizeof(ldt)
	if !(size == 8) {
		t.Fatal("Sizeof(LocalDateTime): ", size)
	}
}

func TestLocalDateTimeIsError(t *testing.T) {
	if (&LocalDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/}).IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be false")
	}
	if !(&LocalDateTime{InvalidYear, 1, 1, 0, 0, 0, 0 /*Fold*/}).IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be true")
	}
}

func TestLocalDateTimeToEpochSeconds(t *testing.T) {
	if (&LocalDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/}).ToEpochSeconds() != 0 {
		t.Fatal("LocalDateTime{2050, 1, 1, 0, 0, 0}.ToEpochSeconds() should be 0")
	}
	if (&LocalDateTime{2050, 1, 1, 0, 0, 1, 0 /*Fold*/}).ToEpochSeconds() != 1 {
		t.Fatal("LocalDateTime{2050, 1, 1, 0, 0, 1}.ToEpochSeconds() should be 1")
	}
	if (&LocalDateTime{2051, 1, 1, 0, 0, 1, 0 /*Fold*/}).ToEpochSeconds() !=
		86400*365+1 {

		t.Fatal(
			"LocalDateTime{2051, 1, 1, 0, 0, 1}.ToEpochSeconds() should be 31536001")
	}
}

func TestNewLocalDateTimeFromEpochSeconds(t *testing.T) {
	ldt := NewLocalDateTimeFromEpochSeconds(0)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 0 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(0) should be " +
			"(2050, 1, 1, 0, 0, 0")
	}
	ldt = NewLocalDateTimeFromEpochSeconds(1)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(1) should be " +
			"(2050, 1, 1, 0, 0, 1)")
	}
	ldt = NewLocalDateTimeFromEpochSeconds(86400*365 + 1)
	if ldt.Year != 2051 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(86400*365+1) should be " +
			"(2051, 1, 1, 0, 0, 1)")
	}
}

func TestLocalDateTimeToString(t *testing.T) {
	ldt := LocalDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/}
	s := ldt.String()
	if !(s == "2023-01-19T16:09:01") {
		t.Fatal(s, ldt)
	}
}