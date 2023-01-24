package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

func TestTimeZoneToString(t *testing.T) {
	zoneInfo := &zonedbtesting.ZoneInfos[
		zonedbtesting.ZoneInfoIndexAmerica_Los_Angeles]
	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.Context, zoneInfo)
	if !(tz.Name() == "America/Los_Angeles") {
		t.Fatal(tz.Name(), tz)
	}
}

func TestTimeZoneUTC(t *testing.T) {
	tz := NewTimeZoneUTC()
	if !(tz.Name() == "UTC") {
		t.Fatal(tz.Name(), tz)
	}

	if !tz.IsUTC() {
		t.Fatal(tz.Name(), tz)
	}
}
