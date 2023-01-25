package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestFindById(t *testing.T) {
	context := &zonedbtesting.Context
	registrar := NewZoneRegistrar(context)
	zoneID := zonedbtesting.ZoneIDAmerica_Los_Angeles
	zoneInfo := registrar.FindZoneInfoByID(zoneID)
	if zoneInfo == nil {
		t.Fatalf("%d not found", zoneID)
	}
	zoneName := zoneInfo.Name(context.NameData, context.NameOffsets)
	if !(zoneName == "America/Los_Angeles") {
		t.Fatal(zoneName)
	}
	if !(zoneInfo.ZoneID == zoneID) {
		t.Fatal(zoneInfo.ZoneID)
	}
}

func TestFindByIdNotFound(t *testing.T) {
	context := &zonedbtesting.Context
	registrar := NewZoneRegistrar(context)
	zoneInfo := registrar.FindZoneInfoByID(0)
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestFindByName(t *testing.T) {
	context := &zonedbtesting.Context
	registrar := NewZoneRegistrar(context)
	zoneInfo := registrar.FindZoneInfoByName("America/Los_Angeles")
	if zoneInfo == nil {
		t.Fatal("Not found")
	}
}

func TestFindByNameNotFound(t *testing.T) {
	context := &zonedbtesting.Context
	registrar := NewZoneRegistrar(context)
	zoneInfo := registrar.FindZoneInfoByName("America/DoesNotExist")
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestIsZoneInfosSorted_Sorted(t *testing.T) {
	zis := []zoneinfo.ZoneInfo{
		zonedbtesting.ZoneInfos[0],
		zonedbtesting.ZoneInfos[2],
		zonedbtesting.ZoneInfos[4],
	}
	isSorted := IsZoneInfosSorted(zis)
	if !isSorted {
		t.Fatal(isSorted)
	}
}

func TestIsZoneInfosSorted_NotSorted(t *testing.T) {
	zis := []zoneinfo.ZoneInfo{
		zonedbtesting.ZoneInfos[2],
		zonedbtesting.ZoneInfos[0],
		zonedbtesting.ZoneInfos[4],
	}
	isSorted := IsZoneInfosSorted(zis)
	if isSorted {
		t.Fatal(isSorted)
	}
}

func TestLinearSearch_NotFound(t *testing.T) {
	i := FindByIdLinear(zonedbtesting.ZoneInfos, 0x0)
	if i != InvalidRegistryIndex {
		t.Fatal(i)
	}
}

func TestLinearSearch_Found(t *testing.T) {
	i := FindByIdLinear(zonedbtesting.ZoneInfos, 0xa950f6ab) // Los_Angeles
	if i == InvalidRegistryIndex {
		t.Fatal(i)
	}
}

func TestBinarySearch_NotFound(t *testing.T) {
	zis := zonedbtesting.ZoneInfos
	if !IsZoneInfosSorted(zis) {
		t.Fatal("Not sorted")
	}
	i := FindByIdBinary(zis, 0x11111111) // random zoneId, should not be there
	if i != InvalidRegistryIndex {
		t.Fatal(i)
	}
}

func TestBinarySearch_Found(t *testing.T) {
	zis := zonedbtesting.ZoneInfos
	if !IsZoneInfosSorted(zis) {
		t.Fatal("Not sorted")
	}
	i := FindByIdBinary(zis, 0x1e2a7654)
	if i == InvalidRegistryIndex {
		t.Fatal(i)
	}
	i = FindByIdBinary(zis, 0xb7f7e8f2)
	if i == InvalidRegistryIndex {
		t.Fatal(i)
	}
	i = FindByIdBinary(zis, 0xd8e31abc)
	if i == InvalidRegistryIndex {
		t.Fatal(i)
	}
}
