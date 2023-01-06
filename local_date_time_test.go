package acetime

import (
	"testing"
	"unsafe"
)

func TestLocalDateTimeSize(t *testing.T) {
	ldt := LocalDateTime{2000, 1, 1, 1, 2, 3}
	size := unsafe.Sizeof(ldt)
	if !(size == 8) {
		t.Fatal("Sizeof(LocalDateTime): ", size)
	}
}

func TestLocalDateTimeIsError(t *testing.T) {
	if (&LocalDateTime{2000, 1, 1, 0, 0, 0}).IsError() {
		t.Fatalf(`LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be false`)
	}
	if !(&LocalDateTime{InvalidYear, 1, 1, 0, 0, 0}).IsError() {
		t.Fatalf(`LocalDateTime{2000, 1, 1, 0, 0, 0}.IsError() should be true`)
	}
}

func TestLocalDateTimeToEpochSeconds(t *testing.T) {
	if (&LocalDateTime{2050, 1, 1, 0, 0, 0}).ToEpochSeconds() != 0 {
		t.Fatalf(`LocalDateTime{2050, 1, 1, 0, 0, 0}.ToEpochSeconds() should be 0`)
	}
	if (&LocalDateTime{2050, 1, 1, 0, 0, 1}).ToEpochSeconds() != 1 {
		t.Fatalf(`LocalDateTime{2050, 1, 1, 0, 0, 1}.ToEpochSeconds() should be 1`)
	}
	if (&LocalDateTime{2051, 1, 1, 0, 0, 1}).ToEpochSeconds() != 86400*365+1 {
		t.Fatalf(
			`LocalDateTime{2051, 1, 1, 0, 0, 1}.ToEpochSeconds() should be 31536001`)
	}
}

func TestLocalDateTimeFromEpochSeconds(t *testing.T) {
	ldt := LocalDateTimeFromEpochSeconds(0)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 0 {
		t.Fatalf(`LocalDateTimeFromEpochSeconds(0) should be (2050, 1, 1, 0, 0, 0`)
	}
	ldt = LocalDateTimeFromEpochSeconds(1)
	if ldt.Year != 2050 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatalf(`LocalDateTimeFromEpochSeconds(1) should be (2050, 1, 1, 0, 0, 1)`)
	}
	ldt = LocalDateTimeFromEpochSeconds(86400*365 + 1)
	if ldt.Year != 2051 || ldt.Month != 1 || ldt.Day != 1 ||
		ldt.Hour != 0 || ldt.Minute != 0 || ldt.Second != 1 {
		t.Fatalf(`LocalDateTime.ToEpochSeconds(86400*365+1) should be ` +
			`(2051, 1, 1, 0, 0, 1)`)
	}
}
