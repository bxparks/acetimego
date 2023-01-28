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

func TestOffsetDateTimeEpochSeconds(t *testing.T) {
	odt := OffsetDateTime{2050, 1, 1, 0, 0, 0, 0, 0}
	if !(odt.EpochSeconds() == 0) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, 0}
	if !(odt.EpochSeconds() == 1) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, -1}
	if !(odt.EpochSeconds() == 61) {
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

func TestOffsetDateTimeToString(t *testing.T) {
	odt := OffsetDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/, -8*60 - 30}
	s := odt.String()
	if !(s == "2023-01-19T16:09:01-08:30") {
		t.Fatal(s, odt)
	}

	odt = OffsetDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/, 8*60 + 15}
	s = odt.String()
	if !(s == "2023-01-19T16:09:01+08:15") {
		t.Fatal(s, odt)
	}
}

func TestMinutesToHM(t *testing.T) {
	s, h, m := minutesToHM(0)
	if !(s == 1) {
		t.Fatal(s)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 0) {
		t.Fatal(m)
	}

	s, h, m = minutesToHM(1)
	if !(s == 1) {
		t.Fatal(s)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 1) {
		t.Fatal(m)
	}

	s, h, m = minutesToHM(62)
	if !(s == 1) {
		t.Fatal(s)
	}
	if !(h == 1) {
		t.Fatal(h)
	}
	if !(m == 2) {
		t.Fatal(m)
	}

	s, h, m = minutesToHM(-123)
	if !(s == -1) {
		t.Fatal(s)
	}
	if !(h == 2) {
		t.Fatal(h)
	}
	if !(m == 3) {
		t.Fatal(m)
	}
}
