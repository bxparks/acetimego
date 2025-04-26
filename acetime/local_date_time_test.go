package acetime

import (
	"math"
	"testing"
	"unsafe"
)

// Verify my understanding of Golang's integer arithematics.
func TestInvalidEpochSeconds(t *testing.T) {
	if InvalidEpochSeconds != math.MinInt64 {
		t.Fatal(InvalidYear)
	}
}

func TestLocalDateTimeSize(t *testing.T) {
	ldt := LocalDateTime{2000, 1, 1, 1, 2, 3}
	size := unsafe.Sizeof(ldt)
	if !(size == 8) {
		t.Fatal("Sizeof(LocalDateTime): ", size)
	}
}

func TestLocalDateTimeIsError(t *testing.T) {
	ldt := LocalDateTime{2000, 1, 1, 0, 0, 0}
	if ldt.IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be false")
	}

	ldt = LocalDateTime{InvalidYear, 1, 1, 0, 0, 0}
	if !ldt.IsError() {
		t.Fatal("LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be true")
	}
}

func TestLocalDateTimeEpochSeconds(t *testing.T) {
	ldt := LocalDateTime{1970, 1, 1, 0, 0, 0}
	if ldt.EpochSeconds() != 0 {
		t.Fatal("LocalDateTime{1970, 1, 1, 0, 0, 0}.EpochSeconds() should be 0")
	}

	ldt = LocalDateTime{1970, 1, 1, 0, 0, 1}
	if ldt.EpochSeconds() != 1 {
		t.Fatal("LocalDateTime{1970, 1, 1, 0, 0, 1}.EpochSeconds() should be 1")
	}

	ldt = LocalDateTime{1971, 1, 1, 0, 0, 1}
	if ldt.EpochSeconds() != 86400*365+1 {
		t.Fatal(
			"LocalDateTime{1971, 1, 1, 0, 0, 1}.EpochSeconds() should be 31536001")
	}

	ldt = LocalDateTime{2050, 1, 1, 0, 0, 0}
	if ldt.EpochSeconds() != 2524608000 {
		t.Fatal(
			"LocalDateTime{1970, 1, 1, 0, 0, 0}.EpochSeconds() should be 2524608000")
	}
}

func TestLocalDateTimeFromEpochSeconds(t *testing.T) {
	ldt := LocalDateTimeFromEpochSeconds(0)
	if ldt.Year != 1970 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 0 {
		t.Fatal("LocalDateTimeFromEpochSeconds(0) should be "+
			"(1970, 1, 1, 0, 0, 0 but was", ldt)
	}
	ldt = LocalDateTimeFromEpochSeconds(1)
	if ldt.Year != 1970 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("LocalDateTimeFromEpochSeconds(1) should be "+
			"(1970, 1, 1, 0, 0, 1) but was", ldt)
	}
	ldt = LocalDateTimeFromEpochSeconds(86400*365 + 1)
	if ldt.Year != 1971 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("LocalDateTimeFromEpochSeconds(86400*365+1) should be "+
			"(1971, 1, 1, 0, 0, 1) but was", ldt)
	}
	ldt = LocalDateTimeFromEpochSeconds(2524608000 + 1)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatal("LocalDateTimeFromEpochSeconds(2524608000) should be "+
			"(2050, 1, 1, 0, 0, 1) but was", ldt)
	}
}

func TestLocalDateTimeToString(t *testing.T) {
	ldt := LocalDateTime{2023, 1, 19, 16, 9, 1}
	s := ldt.String()
	if !(s == "2023-01-19T16:09:01") {
		t.Fatal(s, ldt)
	}
}

func TestLocalDateTimeEquals(t *testing.T) {
	ldt1 := LocalDateTime{2023, 1, 19, 16, 9, 1}
	ldt2 := LocalDateTime{2023, 1, 19, 16, 9, 1}
	if !(ldt1 == ldt2) {
		t.Fatal("ldt1 != ldt2", ldt2)
	}
}
