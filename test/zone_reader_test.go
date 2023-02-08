// Test for `zoneinfo/zone_reader.go` should normally be in the `zoneinfo/`
// directory. Unfortunately this has a dependency to the `zonedbtesting` and
// `zonedb` packages, which has a circular dependency to the `zoneinfo` package.

package acetime

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
			t.Fatal(i, record.ZoneID)
		}
		if record.NameIndex != expected.NameIndex {
			t.Fatal(i, record.NameIndex)
		}
		if record.EraIndex != expected.EraIndex {
			t.Fatal(i, record.EraIndex)
		}
		if record.EraCount != expected.EraCount {
			t.Fatal(i, record.EraCount)
		}
		if record.TargetIndex != expected.TargetIndex {
			t.Fatal(i, record.TargetIndex)
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
		if record.OffsetSecondsCode != expected.OffsetSecondsCode {
			t.Fatal(i, record.OffsetSecondsCode, expected.OffsetSecondsCode)
		}
		if record.DeltaCode != expected.DeltaCode {
			t.Fatal(i, record.DeltaCode, expected.DeltaCode)
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
		if record.UntilTimeCode != expected.UntilTimeCode {
			t.Fatal(i, record.UntilTimeCode, expected.UntilTimeCode)
		}
		if record.UntilTimeModifier != expected.UntilTimeModifier {
			t.Fatal(i, record.UntilTimeModifier, expected.UntilTimeModifier)
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
			t.Fatal(i, record.RuleIndex)
		}
		if record.RuleCount != expected.RuleCount {
			t.Fatal(i, record.RuleCount)
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
			t.Fatal(i, record.FromYear)
		}
		if record.ToYear != expected.ToYear {
			t.Fatal(i, record.ToYear)
		}
		if record.InMonth != expected.InMonth {
			t.Fatal(i, record.InMonth)
		}
		if record.OnDayOfWeek != expected.OnDayOfWeek {
			t.Fatal(i, record.OnDayOfWeek)
		}
		if record.OnDayOfMonth != expected.OnDayOfMonth {
			t.Fatal(i, record.OnDayOfMonth)
		}
		if record.AtTimeCode != expected.AtTimeCode {
			t.Fatal(i, record.AtTimeCode)
		}
		if record.AtTimeModifier != expected.AtTimeModifier {
			t.Fatal(i, record.AtTimeModifier)
		}
		if record.DeltaCode != expected.DeltaCode {
			t.Fatal(i, record.DeltaCode)
		}
		if record.LetterIndex != expected.LetterIndex {
			t.Fatal(i, record.LetterIndex)
		}
	}
}
