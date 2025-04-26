package acetime

import (
	"testing"
	"unsafe"
)

func TestOffsetDateTimeSize(t *testing.T) {
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2000, 1, 1, 1, 2, 3},
		OffsetSeconds: -8 * 3600,
	}
	size := unsafe.Sizeof(odt)
	if !(size == 12) {
		t.Fatal("Sizeof(OffsetDateTime): ", size)
	}
}

func TestOffsetDateTimeIsError(t *testing.T) {
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2000, 1, 1, 0, 0, 0},
		OffsetSeconds: 0,
	}
	if odt.IsError() {
		t.Fatal(odt)
	}
}

func TestOffsetDateTimeEpochSeconds(t *testing.T) {
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{1970, 1, 1, 0, 0, 0},
		OffsetSeconds: 0,
	}
	if !(odt.EpochSeconds() == 0) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{
		LocalDateTime: LocalDateTime{1970, 1, 1, 0, 0, 1},
		OffsetSeconds: 0,
	}
	if !(odt.EpochSeconds() == 1) {
		t.Fatal(odt)
	}
	odt = OffsetDateTime{
		LocalDateTime: LocalDateTime{1970, 1, 1, 0, 0, 1},
		OffsetSeconds: -1,
	}
	if !(odt.EpochSeconds() == 2) {
		t.Fatal(odt)
	}
}

func TestOffsetDateTimeFromEpochSeconds(t *testing.T) {
	odt := OffsetDateTimeFromEpochSeconds(0, 0)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 0 {
		t.Fatalf("OffsetDateTimeFromEpochSeconds(0, 0) " +
			"should be (1970, 1, 1, 0, 0, 0, 0, 0)")
	}
	odt = OffsetDateTimeFromEpochSeconds(0, 3600)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 1 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 3600 {
		t.Fatalf("OffsetDateTimeFromEpochSeconds(0, 3600) " +
			"should be (1970, 1, 1, 1, 0, 0, 0, 3600)")
	}
	odt = OffsetDateTimeFromEpochSeconds(-3600, 3600)
	if odt.Year != 1970 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetSeconds != 3600 {
		t.Fatalf("OffsetDateTimeFromEpochSeconds(-3600, 3600) " +
			"should be (1970, 1, 1, 0, 0, 0, 0, 3600)")
	}
}

func TestOffsetDateTimeToString(t *testing.T) {
	odt := OffsetDateTime{
		LocalDateTime: LocalDateTime{2023, 1, 19, 16, 9, 1},
		OffsetSeconds: -8*3600 - 30*60,
	}
	s := odt.String()
	if !(s == "2023-01-19T16:09:01-08:30") {
		t.Fatal(s, odt)
	}

	odt = OffsetDateTime{
		LocalDateTime: LocalDateTime{2023, 1, 19, 16, 9, 1},
		OffsetSeconds: 8*3600 + 15*60,
	}
	s = odt.String()
	if !(s == "2023-01-19T16:09:01+08:15") {
		t.Fatal(s, odt)
	}
}
