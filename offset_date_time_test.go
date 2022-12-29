package acetime

import (
	"testing"
)

func TestOffsetDateTimeIsError(t *testing.T) {
	if (&OffsetDateTime{2000, 1, 1, 0, 0, 0, 0, 0}).IsError() {
		t.Fatalf(`OffsetDateTime{2000, 1, 1, 0, 0, 0, 0, 0}.IsError() ` +
			`should be false`)
	}
	if !(&OffsetDateTime{InvalidYear, 1, 1, 0, 0, 0, 0, 0}).IsError() {
		t.Fatalf(`OffsetDateTime{2000, 1, 1, 0, 0, 0, 0, 0}.IsError() ` +
			`should be true`)
	}
}

func TestToEpochSeconds(t *testing.T) {
	if (&OffsetDateTime{2050, 1, 1, 0, 0, 0, 0, 0}).ToEpochSeconds() != 0 {
		t.Fatalf(`OffsetDateTime{2050, 1, 1, 0, 0, 0}.ToEpochSeconds() should be 0`)
	}
	if (&OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, 0}).ToEpochSeconds() != 1 {
		t.Fatalf(`OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, 0}.ToEpochSeconds() ` +
			` should be 1`)
	}
	if (&OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, -1}).ToEpochSeconds() != 61 {
		t.Fatalf(`OffsetDateTime{2050, 1, 1, 0, 0, 1, 0, -1}.ToEpochSeconds() ` +
			`should be 61`)
	}
}

func TestOffsetDateTimeFromEpochSeconds(t *testing.T) {
	odt := OffsetDateTimeFromEpochSeconds(0, 0)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 0 {
		t.Fatalf(`OffsetDateTimeFromEpochSeconds(0, 0) ` +
			`should be (2050, 1, 1, 0, 0, 0, 0, 0)`)
	}
	odt = OffsetDateTimeFromEpochSeconds(0, 60)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 1 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 60 {
		t.Fatalf(`OffsetDateTimeFromEpochSeconds(0, 60) ` +
			`should be (2050, 1, 1, 1, 0, 0, 0, 60)`)
	}
	odt = OffsetDateTimeFromEpochSeconds(-3600, 60)
	if odt.Year != 2050 || odt.Month != 1 || odt.Day != 1 ||
		odt.Hour != 0 || odt.Minute != 0 || odt.Second != 0 ||
		odt.OffsetMinutes != 60 {
		t.Fatalf(`OffsetDateTimeFromEpochSeconds(-3600, 60) ` +
			`should be (2050, 1, 1, 0, 0, 0, 0, 60)`)
	}
}
