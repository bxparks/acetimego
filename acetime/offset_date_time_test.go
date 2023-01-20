package acetime

import (
	"testing"
	"unsafe"
)

func TestOffsetDateTimeSize(t *testing.T) {
	odt := OffsetDateTime{2000, 1, 1, 1, 2, 3, 0 /*Fold*/, -8 * 60}
	size := unsafe.Sizeof(odt)
	if !(size == 10) {
		t.Fatal("Sizeof(OffsetDateTime): ", size)
	}
}

func TestOffsetDateTimeIsError(t *testing.T) {
	odt := OffsetDateTime{2000, 1, 1, 0, 0, 0, 0, 0}
	if odt.IsError() {
		t.Fatal(odt)
	}
}

func TestNewOffsetDateTimeError(t *testing.T) {
	odt := NewOffsetDateTimeError()
	if !odt.IsError() {
		t.Fatal(odt)
	}
}

func TestOffsetDateTimeToEpochSeconds(t *testing.T) {
	odt := OffsetDateTime{2050, 1, 1, 0, 0, 0, 0, 0}
	if !(odt.ToEpochSeconds() == 0) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, 0}
	if !(odt.ToEpochSeconds() == 1) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, -1}
	if !(odt.ToEpochSeconds() == 61) {
		t.Fatal(odt)
	}
}

func TestNewOffsetDateTimeFromEpochSeconds(t *testing.T) {
	odt := NewOffsetDateTimeFromEpochSeconds(0, 0)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 0 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(0, 0) " +
			"should be (2050, 1, 1, 0, 0, 0, 0, 0)")
	}
	odt = NewOffsetDateTimeFromEpochSeconds(0, 60)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 1 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 60 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(0, 60) " +
			"should be (2050, 1, 1, 1, 0, 0, 0, 60)")
	}
	odt = NewOffsetDateTimeFromEpochSeconds(-3600, 60)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 60 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(-3600, 60) " +
			"should be (2050, 1, 1, 0, 0, 0, 0, 60)")
	}
}
