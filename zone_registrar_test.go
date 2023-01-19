package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"testing"
)

func TestFindById(t *testing.T) {
	registrar := ZoneRegistrar{zonedbtesting.ZoneAndLinkRegistry}
	zoneInfo := registrar.FindZoneInfoByID(
		zonedbtesting.ZoneAmerica_Los_Angeles.ZoneID)
	if zoneInfo == nil {
		t.Fatalf("%d (%s) not found",
			zonedbtesting.ZoneAmerica_Los_Angeles.ZoneID,
			zonedbtesting.ZoneAmerica_Los_Angeles.Name)
	}
	if !(zoneInfo.Name == "America/Los_Angeles") {
		t.Fatal(zoneInfo.Name)
	}
	if !(zoneInfo.ZoneID == zonedbtesting.ZoneAmerica_Los_Angeles.ZoneID) {
		t.Fatal(zoneInfo.ZoneID)
	}
}

func TestFindByIdNotFound(t *testing.T) {
	registrar := ZoneRegistrar{zonedbtesting.ZoneAndLinkRegistry}
	zoneInfo := registrar.FindZoneInfoByID(0)
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}
