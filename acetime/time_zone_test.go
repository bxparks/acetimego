package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

func TestTimeZoneToString(t *testing.T) {
	zoneManager := NewZoneManager(&zonedbtesting.DataContext)
	tz := zoneManager.NewTimeZoneFromID(zonedbtesting.ZoneIDAmerica_Los_Angeles)
	if !(tz.Name() == "America/Los_Angeles") {
		t.Fatal(tz)
	}
}

func TestTimeZoneUTC(t *testing.T) {
	tz := TimeZoneUTC
	if !(tz.Name() == "UTC") {
		t.Fatal(tz.Name(), tz)
	}
	if !tz.IsUTC() {
		t.Fatal(tz)
	}
	if tz.IsError() {
		t.Fatal(tz)
	}
}

func TestTimeZoneError(t *testing.T) {
	tz := TimeZoneError
	if !(tz.Name() == "<Error>") {
		t.Fatal(tz)
	}
	if !tz.IsUTC() {
		t.Fatal(tz)
	}
	if !tz.IsError() {
		t.Fatal(tz)
	}
}
