// Test for `zoneinfo/zone_store.go` should normally be in the `zoneinfo/`
// directory. Unfortunately this has a dependency to the `zonedbtesting`
// package, which has a circular dependency to the `zoneinfo` package. I could
// copy over most of the data in `zonedbtesting/zone*.go` files, but that would
// defeat the purpose of having that package be auto-generated. The only
// solution that I can think of right now is to move this test file into the
// `acetime` package.

package acetime

import (
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestFindByID_Found(t *testing.T) {
	context := &zonedbtesting.DataContext
	store := zoneinfo.NewZoneStore(context)
	zoneID := zonedbtesting.ZoneIDAmerica_Los_Angeles
	zoneInfo := store.ZoneInfoByID(zoneID)
	if zoneInfo == nil {
		t.Fatalf("%d not found", zoneID)
	}
	zoneName := zoneInfo.Name
	if !(zoneName == "America/Los_Angeles") {
		t.Fatal(zoneName)
	}
	if !(zoneInfo.ZoneID == zoneID) {
		t.Fatal(zoneInfo.ZoneID)
	}
}

func TestFindByID_Link(t *testing.T) {
	context := &zonedbtesting.DataContext
	store := zoneinfo.NewZoneStore(context)
	zoneID := zonedbtesting.ZoneIDUS_Pacific
	zoneInfo := store.ZoneInfoByID(zoneID)
	if zoneInfo == nil {
		t.Fatalf("%d not found", zoneID)
	}
	zoneName := zoneInfo.Name
	if !(zoneName == "US/Pacific") {
		t.Fatal(zoneName)
	}
	if !(zoneInfo.ZoneID == zoneID) {
		t.Fatal(zoneInfo.ZoneID)
	}

	if !zoneInfo.IsLink() {
		t.Fatal("US/Pacific should be a Link")
	}
	target := zoneInfo.Target
	if target.Name != "America/Los_Angeles" {
		t.Fatal("US/Pacific should point to America/Los_Angeles")
	}
}

func TestFindByID_NotFound(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneInfo := store.ZoneInfoByID(0)
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestFindByName_Found(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneInfo := store.ZoneInfoByName("America/Los_Angeles")
	if zoneInfo == nil {
		t.Fatal("Not found")
	}
	zoneName := zoneInfo.Name
	if !(zoneName == "America/Los_Angeles") {
		t.Fatal(zoneName)
	}
}

func TestFindByName_NotFound(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneInfo := store.ZoneInfoByName("America/DoesNotExist")
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}

func TestIsSorted(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	if !store.IsSorted() {
		t.Fatal("zonedbtesting should be sorted")
	}
}
