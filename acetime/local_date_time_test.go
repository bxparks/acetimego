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
	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0, 0 /*Fold*/}
	if ldt.IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be false")
	}

	ldt = LocalDateTime{InvalidYear, 1, 1, 0, 0, 0, 0 /*Fold*/}
	if !ldt.IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be true")
	}
}

func TestLocalDateTimeEpochSeconds(t *testing.T) {
	ldt := LocalDateTime{1970, 1, 1, 0, 0, 0, 0 /*Fold*/}
	if ldt.EpochSeconds() != 0 {
		t.Fatal("LocalDateTime{1970, 1, 1, 0, 0, 0}.EpochSeconds() should be 0")
	}

	ldt = LocalDateTime{1970, 1, 1, 0, 0, 1, 0 /*Fold*/}
	if ldt.EpochSeconds() != 1 {
		t.Fatal("LocalDateTime{1970, 1, 1, 0, 0, 1}.EpochSeconds() should be 1")
	}

	ldt = LocalDateTime{1971, 1, 1, 0, 0, 1, 0 /*Fold*/}
	if ldt.EpochSeconds() != 86400*365+1 {
		t.Fatal(
			"LocalDateTime{1971, 1, 1, 0, 0, 1}.EpochSeconds() should be 31536001")
	}

	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0, 0 /*Fold*/}
	if ldt.EpochSeconds() != 2524608000 {
		t.Fatal(
			"LocalDateTime{1970, 1, 1, 0, 0, 0}.EpochSeconds() should be 2524608000")
	}
}

func TestNewLocalDateTimeFromEpochSeconds(t *testing.T) {
	ldt := NewLocalDateTimeFromEpochSeconds(0)
	if ldt.Year != 1970 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 0 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(0) should be "+
			"(1970, 1, 1, 0, 0, 0 but was", ldt)
	}
	ldt = NewLocalDateTimeFromEpochSeconds(1)
	if ldt.Year != 1970 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(1) should be "+
			"(1970, 1, 1, 0, 0, 1) but was", ldt)
	}
	ldt = NewLocalDateTimeFromEpochSeconds(86400*365 + 1)
	if ldt.Year != 1971 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(86400*365+1) should be "+
			"(1971, 1, 1, 0, 0, 1) but was", ldt)
	}
	ldt = NewLocalDateTimeFromEpochSeconds(2524608000 + 1)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("NewLocalDateTimeFromEpochSeconds(2524608000) should be "+
			"(2050, 1, 1, 0, 0, 1) but was", ldt)
	}
}

func TestLocalDateTimeToString(t *testing.T) {
	ldt := LocalDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/}
	s := ldt.String()
	if !(s == "2023-01-19T16:09:01") {
		t.Fatal(s, ldt)
	}
}
