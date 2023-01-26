package zoneinfo

import (
	"testing"
)

const letterData = "" +
	"D" +
	"S" +
	"~"

var letterOffsets = []uint8{
	0, 0, 1, 2,
}

const nameData = "" +
	"America/Los_Angeles" +
	"America/New_York" +
	"Etc/UTC" +
	"Pacific/Apia" +
	"US/Pacific" +
	"~"

var nameOffsets = []uint16{
	0, 0, 19, 35, 42, 54, 64,
}

func TestStringIO8(t *testing.T) {
	f := StringIO8{letterData, letterOffsets}
	if f.StringAt(0) != "" {
		t.Fatal(f.StringAt(0))
	}
	if f.StringAt(2) != "S" {
		t.Fatal(f.StringAt(2))
	}
}

func TestStringIO16(t *testing.T) {
	f := StringIO16{nameData, nameOffsets}
	if f.StringAt(0) != "" {
		t.Fatal(f.StringAt(0))
	}
	if f.StringAt(1) != "America/Los_Angeles" {
		t.Fatal(f.StringAt(0))
	}
	if f.StringAt(5) != "US/Pacific" {
		t.Fatal(f.StringAt(5))
	}
}
