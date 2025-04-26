package acetime

import (
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
)

func TestZoneManagerZoneCount(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	if !(zm.ZoneCount() == zonedbtesting.DataContext.ZoneInfoCount) {
		t.Fatal(zm.ZoneCount())
	}
}

func TestZoneManagerZoneIDs(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	zoneIDs := zm.ZoneIDs()
	if !(uint16(len(zoneIDs)) == zm.ZoneCount()) {
		t.Fatal(len(zoneIDs))
	}
}

func TestZoneManagerZoneNames(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	zoneNames := zm.ZoneNames()
	if !(uint16(len(zoneNames)) == zm.ZoneCount()) {
		t.Fatal(len(zoneNames))
	}
}

func TestZoneManagerNewTimeZoneFromID(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromZoneID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	if !(tz.Name() == "America/Los_Angeles") {
		t.Fatal(tz.Name())
	}
}

func TestZoneManagerNewTimeZoneFromID_Error(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromZoneID(0 /*should not exist*/)
	if !(tz.IsError()) {
		t.Fatal(tz)
	}
}

func TestZoneManagerNewTimeZoneFromName(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("America/Los_Angeles")
	if !(tz.Name() == "America/Los_Angeles") {
		t.Fatal(tz.Name())
	}
}

func TestZoneManagerNewTimeZoneFromName_Error(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.TimeZoneFromName("ShouldNotExist")
	if !(tz.IsError()) {
		t.Fatal(tz)
	}
}

func TestZoneManagerNewTimeZoneFromIndex_Error(t *testing.T) {
	zm := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zm.timeZoneFromIndex(zm.ZoneCount()) // one past the end
	if !(tz.IsError()) {
		t.Fatal(tz)
	}
}
