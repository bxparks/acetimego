package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

func TestTimeZoneToString(t *testing.T) {
	tz := NewTimeZoneFromZoneInfo(
		&zonedbtesting.Context, &zonedbtesting.ZoneAmerica_Los_Angeles)
	if !(tz.String() == "America/Los_Angeles") {
		t.Fatal(tz.String(), tz)
	}
}

func TestTimeZoneUTC(t *testing.T) {
	tz := NewTimeZoneUTC()
	if !(tz.String() == "UTC") {
		t.Fatal(tz.String(), tz)
	}

	if !tz.IsUTC() {
		t.Fatal(tz.String(), tz)
	}
}
