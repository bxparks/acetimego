package acetime

import (
	"testing"
	"unsafe"
)

func TestOffsetDateTimeSize(t *testing.T) {
	odt := OffsetDateTime{2000, 1, 1, 1, 2, 3, 0 /*Fold*/, -8 * 3600}
	size := unsafe.Sizeof(odt)
	if !(size == 12) {
		t.Fatal("Sizeof(OffsetDateTime): ", size)
	}
}

func TestOffsetDateTimeIsError(t *testing.T) {
	odt := OffsetDateTime{2000, 1, 1, 0, 0, 0, 0, 0}
	if odt.IsError() {
		t.Fatal(odt)
	}
}

func TestOffsetDateTimeEpochSeconds(t *testing.T) {
	odt := OffsetDateTime{1970, 1, 1, 0, 0, 0, 0 /*Fold*/, 0 /*OffsetSeconds*/}
	if !(odt.EpochSeconds() == 0) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{1970, 1, 1, 0, 0, 1, 0 /*Fold*/, 0 /*OffsetSeconds*/}
	if !(odt.EpochSeconds() == 1) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{1970, 1, 1, 0, 0, 1, 0 /*Fold*/, -1 /*OffsetSeconds*/}
	if !(odt.EpochSeconds() == 2) {
		t.Fatal(odt)
	}
}

func TestNewOffsetDateTimeFromEpochSeconds(t *testing.T) {
	odt := NewOffsetDateTimeFromEpochSeconds(0, 0)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 0 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(0, 0) " +
			"should be (1970, 1, 1, 0, 0, 0, 0, 0)")
	}
	odt = NewOffsetDateTimeFromEpochSeconds(0, 3600)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 1 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 3600 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(0, 3600) " +
			"should be (1970, 1, 1, 1, 0, 0, 0, 3600)")
	}
	odt = NewOffsetDateTimeFromEpochSeconds(-3600, 3600)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 3600 {
		t.Fatalf("NewOffsetDateTimeFromEpochSeconds(-3600, 3600) " +
			"should be (1970, 1, 1, 0, 0, 0, 0, 3600)")
	}
}

func TestOffsetDateTimeToString(t *testing.T) {
	odt := OffsetDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/, -8*3600 - 30*60}
	s := odt.String()
	if !(s == "2023-01-19T16:09:01-08:30") {
		t.Fatal(s, odt)
	}

	odt = OffsetDateTime{2023, 1, 19, 16, 9, 1, 0 /*Fold*/, 8*3600 + 15*60}
	s = odt.String()
	if !(s == "2023-01-19T16:09:01+08:15") {
		t.Fatal(s, odt)
	}
}

func TestSecondsToHMS(t *testing.T) {
	sign, h, m, s := secondsToHMS(0)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 0) {
		t.Fatal(m)
	}
	if !(s == 0) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(1)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 0) {
		t.Fatal(m)
	}
	if !(s == 1) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(62)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 1) {
		t.Fatal(m)
	}
	if !(s == 2) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(-3663)
	if !(sign == -1) {
		t.Fatal(sign)
	}
	if !(h == 1) {
		t.Fatal(h)
	}
	if !(m == 1) {
		t.Fatal(m)
	}
	if !(s == 3) {
		t.Fatal(s)
	}
}
