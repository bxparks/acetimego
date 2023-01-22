package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestFindById(t *testing.T) {
	registrar := NewZoneRegistrar(zonedbtesting.ZoneAndLinkRegistry)
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
	registrar := NewZoneRegistrar(zonedbtesting.ZoneAndLinkRegistry)
	zoneInfo := registrar.FindZoneInfoByID(0)
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestFindByName(t *testing.T) {
	registrar := NewZoneRegistrar(zonedbtesting.ZoneAndLinkRegistry)
	zoneInfo := registrar.FindZoneInfoByName("America/Los_Angeles")
	if zoneInfo == nil {
		t.Fatal("Not found")
	}
}

func TestFindByNameNotFound(t *testing.T) {
	registrar := NewZoneRegistrar(zonedbtesting.ZoneAndLinkRegistry)
	zoneInfo := registrar.FindZoneInfoByName("America/DoesNotExist")
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestIsZoneRegistrySorted_Sorted(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	isSorted := IsZoneRegistrySorted(zis)
	if !isSorted {
		t.Fatal(isSorted)
	}
}

func TestIsZoneRegistrySorted_NotSorted(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	isSorted := IsZoneRegistrySorted(zis)
	if isSorted {
		t.Fatal(isSorted)
	}
}

func TestLinearSearch_NotFound(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	i := FindByIdLinear(zis, 0x0)
	if i != InvalidRegistryIndex {
		t.Fatal(i)
	}
}

func TestLinearSearch_Found(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	i := FindByIdLinear(zis, 0x1e2a7654)
	if !(i == 1) {
		t.Fatal(i)
	}
}

func TestBinarySearch_NotFound(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	if !IsZoneRegistrySorted(zis) {
		t.Fatal("Not sorted")
	}
	i := FindByIdBinary(zis, 0x11111111) // random zoneId, should not be there
	if !(i == InvalidRegistryIndex) {
		t.Fatal(i)
	}
}

func TestBinarySearch_Found(t *testing.T) {
	zis := []*zoneinfo.ZoneInfo{
		&zonedbtesting.ZoneAmerica_New_York,    // 0x1e2a7654
		&zonedbtesting.ZoneAmerica_Los_Angeles, // 0xb7f7e8f2
		&zonedbtesting.ZoneEtc_UTC,             // 0xd8e31abc
	}
	if !IsZoneRegistrySorted(zis) {
		t.Fatal("Not sorted")
	}
	i := FindByIdBinary(zis, 0x1e2a7654)
	if !(i == 0) {
		t.Fatal(i)
	}
	i = FindByIdBinary(zis, 0xb7f7e8f2)
	if !(i == 1) {
		t.Fatal(i)
	}
	i = FindByIdBinary(zis, 0xd8e31abc)
	if !(i == 2) {
		t.Fatal(i)
	}
}
