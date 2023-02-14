// Test for `zoneinfo/zone_reader.go` should normally be in the `zoneinfo/`
// directory. Unfortunately this has a dependency to the `zonedbtesting` and
// `zonedb` packages, which has a circular dependency to the `zoneinfo` package.

package test

import (
	"github.com/bxparks/AceTimeGo/zonedb"
	"github.com/bxparks/AceTimeGo/zonedbtesting"
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"testing"
)

// Verify that data encoded in ZoneInfosData is identical to ZoneInfoRecords.
func TestZoneInfoReader(t *testing.T) {
	verifyZoneInfoReader(
		t, &zonedbtesting.DataContext, &zonedbtesting.RecordContext)
	verifyZoneInfoReader(
		t, &zonedb.DataContext, &zonedb.RecordContext)
}

func verifyZoneInfoReader(
	t *testing.T, c *zoneinfo.ZoneDataContext, cc *zoneinfo.ZoneRecordContext) {

	reader := zoneinfo.NewZoneInfoReader(
		zoneinfo.NewDataIO(c.ZoneInfosData), c.ZoneInfoChunkSize)
	for i := uint16(0); i < c.ZoneInfoCount; i++ {
		reader.Seek(i) // Use random Seek() to verify the chunk size
		record := reader.Read()
		expected := &cc.ZoneInfoRecords[i]
		if record.ZoneID != expected.ZoneID {
			t.Fatal(i, record.ZoneID, expected.ZoneID)
		}
		if record.NameIndex != expected.NameIndex {
			t.Fatal(i, record.NameIndex, expected.NameIndex)
		}
		if record.EraIndex != expected.EraIndex {
			t.Fatal(i, record.EraIndex, expected.EraIndex)
		}
		if record.EraCount != expected.EraCount {
			t.Fatal(i, record.EraCount, expected.EraCount)
		}
		if record.TargetIndex != expected.TargetIndex {
			t.Fatal(i, record.TargetIndex, expected.TargetIndex)
		}
	}
}

// Verify that data encoded in ZoneErasData is identical to ZoneEraRecords.
func TestZoneEraReader(t *testing.T) {
	verifyZoneEraReader(
		t, &zonedbtesting.DataContext, &zonedbtesting.RecordContext)
	verifyZoneEraReader(
		t, &zonedb.DataContext, &zonedb.RecordContext)
}

func verifyZoneEraReader(
	t *testing.T, c *zoneinfo.ZoneDataContext, cc *zoneinfo.ZoneRecordContext) {

	reader := zoneinfo.NewZoneEraReader(
		zoneinfo.NewDataIO(c.ZoneErasData), c.ZoneEraChunkSize)
	for i := uint16(0); i < c.ZoneEraCount; i++ {
		reader.Seek(i) // Use random Seek() to verify the chunk size
		record := reader.Read()
		expected := &cc.ZoneEraRecords[i]
		if record.FormatIndex != expected.FormatIndex {
			t.Fatal(i, record.FormatIndex, expected.FormatIndex)
		}
		if record.PolicyIndex != expected.PolicyIndex {
			t.Fatal(i, record.PolicyIndex, expected.PolicyIndex)
		}
		if record.OffsetSecondsRemainder != expected.OffsetSecondsRemainder {
			t.Fatal(i, record.OffsetSecondsRemainder, expected.OffsetSecondsRemainder)
		}
		if record.OffsetSecondsCode != expected.OffsetSecondsCode {
			t.Fatal(i, record.OffsetSecondsCode, expected.OffsetSecondsCode)
		}
		if record.DeltaMinutes != expected.DeltaMinutes {
			t.Fatal(i, record.DeltaMinutes, expected.DeltaMinutes)
		}
		if record.UntilYear != expected.UntilYear {
			t.Fatal(i, record.UntilYear, expected.UntilYear)
		}
		if record.UntilMonth != expected.UntilMonth {
			t.Fatal(i, record.UntilMonth, expected.UntilMonth)
		}
		if record.UntilDay != expected.UntilDay {
			t.Fatal(i, record.UntilDay, expected.UntilDay)
		}
		if record.UntilSecondsCode != expected.UntilSecondsCode {
			t.Fatal(i, record.UntilSecondsCode, expected.UntilSecondsCode)
		}
		if record.UntilSecondsModifier != expected.UntilSecondsModifier {
			t.Fatal(i, record.UntilSecondsModifier, expected.UntilSecondsModifier)
		}
	}
}

// Verify that data encoded in ZoneErasData is identical to ZoneEraRecords.
func TestZonePolicyReader(t *testing.T) {
	verifyZonePolicyReader(
		t, &zonedbtesting.DataContext, &zonedbtesting.RecordContext)
	verifyZonePolicyReader(
		t, &zonedb.DataContext, &zonedb.RecordContext)
}

func verifyZonePolicyReader(
	t *testing.T, c *zoneinfo.ZoneDataContext, cc *zoneinfo.ZoneRecordContext) {

	reader := zoneinfo.NewZonePolicyReader(
		zoneinfo.NewDataIO(c.ZonePoliciesData), c.ZonePolicyChunkSize)
	for i := uint16(0); i < c.ZonePolicyCount; i++ {
		reader.Seek(i) // Use random Seek() to verify the chunk size
		record := reader.Read()
		expected := &cc.ZonePolicyRecords[i]
		if record.RuleIndex != expected.RuleIndex {
			t.Fatal(i, record.RuleIndex, expected.RuleIndex)
		}
		if record.RuleCount != expected.RuleCount {
			t.Fatal(i, record.RuleCount, expected.RuleCount)
		}
	}
}

// Verify that data encoded in ZoneErasData is identical to ZoneEraRecords.
func TestZoneRuleReader(t *testing.T) {
	verifyZoneRuleReader(
		t, &zonedbtesting.DataContext, &zonedbtesting.RecordContext)
	verifyZoneRuleReader(
		t, &zonedb.DataContext, &zonedb.RecordContext)
}

func verifyZoneRuleReader(
	t *testing.T, c *zoneinfo.ZoneDataContext, cc *zoneinfo.ZoneRecordContext) {

	reader := zoneinfo.NewZoneRuleReader(
		zoneinfo.NewDataIO(c.ZoneRulesData), c.ZoneRuleChunkSize)
	for i := uint16(0); i < c.ZoneRuleCount; i++ {
		reader.Seek(i) // Use random Seek() to verify the chunk size
		record := reader.Read()
		expected := &cc.ZoneRuleRecords[i]
		if record.FromYear != expected.FromYear {
			t.Fatal(i, record.FromYear, expected.FromYear)
		}
		if record.ToYear != expected.ToYear {
			t.Fatal(i, record.ToYear, expected.ToYear)
		}
		if record.InMonth != expected.InMonth {
			t.Fatal(i, record.InMonth, expected.InMonth)
		}
		if record.OnDayOfWeek != expected.OnDayOfWeek {
			t.Fatal(i, record.OnDayOfWeek, expected.OnDayOfWeek)
		}
		if record.OnDayOfMonth != expected.OnDayOfMonth {
			t.Fatal(i, record.OnDayOfMonth, expected.OnDayOfMonth)
		}
		if record.AtSecondsCode != expected.AtSecondsCode {
			t.Fatal(i, record.AtSecondsCode, expected.AtSecondsCode)
		}
		if record.AtSecondsModifier != expected.AtSecondsModifier {
			t.Fatal(i, record.AtSecondsModifier, expected.AtSecondsModifier)
		}
		if record.DeltaMinutes != expected.DeltaMinutes {
			t.Fatal(i, record.DeltaMinutes, expected.DeltaMinutes)
		}
		if record.LetterIndex != expected.LetterIndex {
			t.Fatal(i, record.LetterIndex, expected.LetterIndex)
		}
	}
}
