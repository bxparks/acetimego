package acetime

import (
	"github.com/bxparks/acetimego/zonedbtesting"
	"testing"
)

func TestTimeZoneNormalZone(t *testing.T) {
	zoneManager := ZoneManagerFromDataContext(&zonedbtesting.DataContext)
	tz := zoneManager.TimeZoneFromZoneID(zonedbtesting.ZoneIDAmerica_Los_Angeles)

	if !(tz.Name() == "America/Los_Angeles") {
		t.Fatal(tz)
	}
	if !(tz.ZoneID() == zonedbtesting.ZoneIDAmerica_Los_Angeles) {
		t.Fatal(tz)
	}
	if tz.IsLink() {
		t.Fatal(tz)
	}
	if tz.IsUTC() {
		t.Fatal(tz)
	}
	if tz.IsError() {
		t.Fatal(tz)
	}
}

func TestTimeZoneUTC(t *testing.T) {
	tz := TimeZoneUTC
	if !(tz.Name() == "UTC") {
		t.Fatal(tz.Name(), tz)
	}
	if !(tz.ZoneID() == 0) {
		t.Fatal(tz)
	}
	if tz.IsLink() {
		t.Fatal(tz)
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
	if !(tz.Name() == "Err") {
		t.Fatal(tz)
	}
	if !(tz.ZoneID() == 0) {
		t.Fatal(tz)
	}
	if tz.IsLink() {
		t.Fatal(tz)
	}
	if !tz.IsUTC() {
		t.Fatal(tz)
	}
	if !tz.IsError() {
		t.Fatal(tz)
	}
}
