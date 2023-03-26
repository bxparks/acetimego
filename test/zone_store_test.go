// Test for `zoneinfo/zone_store.go` should normally be in the `zoneinfo/`
// directory. Unfortunately this has a dependency to the `zonedbtesting`
// package, which has a circular dependency to the `zoneinfo` package. I could
// copy over most of the data in `zonedbtesting/zone*.go` files, but that would
// defeat the purpose of having that package be auto-generated. The only
// solution that I can think of right now is to move this test file into the
// `acetime` package.

package test

import (
	"github.com/bxparks/AceTimeGo/zonedb"
	"github.com/bxparks/AceTimeGo/zonedball"
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

func TestZoneStoreZoneCount(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	if !(store.ZoneCount() == zonedbtesting.DataContext.ZoneInfoCount) {
		t.Fatal(store.ZoneCount(), zonedbtesting.DataContext.ZoneInfoCount)
	}

	store = zoneinfo.NewZoneStore(&zonedb.DataContext)
	if !(store.ZoneCount() == zonedb.DataContext.ZoneInfoCount) {
		t.Fatal(store.ZoneCount(), zonedb.DataContext.ZoneInfoCount)
	}

	store = zoneinfo.NewZoneStore(&zonedball.DataContext)
	if !(store.ZoneCount() == zonedball.DataContext.ZoneInfoCount) {
		t.Fatal(store.ZoneCount(), zonedball.DataContext.ZoneInfoCount)
	}
}

func TestZoneStoreIsSorted(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	if !store.IsSorted() {
		t.Fatal("zonedbtesting should be sorted")
	}

	store = zoneinfo.NewZoneStore(&zonedb.DataContext)
	if !store.IsSorted() {
		t.Fatal("zonedb should be sorted")
	}

	store = zoneinfo.NewZoneStore(&zonedball.DataContext)
	if !store.IsSorted() {
		t.Fatal("zonedball should be sorted")
	}
}

func TestZoneStoreFindByID_Found(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneID := zonedbtesting.ZoneIDAmerica_Los_Angeles
	idL := store.FindByIDLinear(zoneID)
	idB := store.FindByIDBinary(zoneID)
	id := store.FindByID(zoneID)

	if !(idL == idB) {
		t.Fatal(idL, idB)
	}
	if !(idL == id) {
		t.Fatal(idL, id)
	}
}

func TestZoneStoreFindByID_NotFound(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)

	index := store.FindByIDLinear(0)
	if index != zoneinfo.InvalidIndex {
		t.Fatal("FindByIDLinear() should have returned InvalidIndex")
	}

	index = store.FindByIDBinary(0)
	if index != zoneinfo.InvalidIndex {
		t.Fatal("FindByIDBinary() should have returned InvalidIndex")
	}

	index = store.FindByID(0)
	if index != zoneinfo.InvalidIndex {
		t.Fatal("FindByID() should have returned InvalidIndex")
	}
}

func TestZoneStoreZoneInfoByID_Zone(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneID := zonedbtesting.ZoneIDAmerica_Los_Angeles

	zoneInfo := store.ZoneInfoByID(zoneID)
	if zoneInfo == nil {
		t.Fatalf("%d not found", zoneID)
	}
	if !(zoneInfo.Name == "America/Los_Angeles") {
		t.Fatal(zoneInfo.Name)
	}
	if !(zoneInfo.ZoneID == zoneID) {
		t.Fatal(zoneInfo.ZoneID)
	}
	if zoneInfo.IsLink() {
		t.Fatal("Should not be a Link")
	}
}

func TestZoneStoreZoneInfoByID_Link(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneID := zonedbtesting.ZoneIDUS_Pacific

	zoneInfo := store.ZoneInfoByID(zoneID)
	if zoneInfo == nil {
		t.Fatalf("%d not found", zoneID)
	}
	if !(zoneInfo.Name == "US/Pacific") {
		t.Fatal(zoneInfo.Name)
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

func TestZoneStoreZoneInfoByName_Found(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneInfo := store.ZoneInfoByName("America/Los_Angeles")
	if zoneInfo == nil {
		t.Fatal("Not found")
	}
	if !(zoneInfo.Name == "America/Los_Angeles") {
		t.Fatal(zoneInfo.Name)
	}
}

func TestZoneStoreZoneInfoByName_NotFound(t *testing.T) {
	store := zoneinfo.NewZoneStore(&zonedbtesting.DataContext)
	zoneInfo := store.ZoneInfoByName("America/DoesNotExist")
	if zoneInfo != nil {
		t.Fatal("Should have returned nil")
	}
}
