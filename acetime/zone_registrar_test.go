package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
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

func TestFindByName(t *testing.T) {
	registrar := ZoneRegistrar{zonedbtesting.ZoneAndLinkRegistry}
	zoneInfo := registrar.FindZoneInfoByName("America/Los_Angeles")
	if zoneInfo == nil {
		t.Fatal("Not found")
	}
}

func TestFindByNameNotFound(t *testing.T) {
	registrar := ZoneRegistrar{zonedbtesting.ZoneAndLinkRegistry}
	zoneInfo := registrar.FindZoneInfoByName("America/DoesNotExist")
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestIsZoneRegistrySorted_Sorted(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_New_York, // 0x1e2a7654
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneEtc_UTC, // 0xd8e31abc
	}
	isSorted := IsZoneRegistrySorted(zis)
	if !isSorted {
		t.Fatal(isSorted)
	}
}

func TestIsZoneRegistrySorted_NotSorted(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneAmerica_New_York, // 0x1e2a7654
		&zonedbtesting.ZoneEtc_UTC, // 0xd8e31abc
	}
	isSorted := IsZoneRegistrySorted(zis)
	if isSorted {
		t.Fatal(isSorted)
	}
}
