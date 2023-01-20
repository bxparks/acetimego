package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

func TestTimeZoneToString(t *testing.T) {
	tz := NewTimeZoneFromZoneInfo(&zonedbtesting.ZoneAmerica_Los_Angeles)
	if !(tz.String() == "America/Los_Angeles") {
		t.Fatal(tz.String(), tz)
	}
}
