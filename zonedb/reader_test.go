// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/AceTimeGo/zonedb/tzfiles
//     --output_dir /home/brian/src/AceTimeGo/zonedb
//     --tz_version 2022g
//     --action zonedb
//     --language go
//     --scope extended
//     --offset_granularity 1
//     --delta_granularity 60
//     --until_at_granularity 1
//     --db_namespace zonedb
//     --generate_int16_years
//     --start_year 1800
//     --until_year 10000
//
// using the TZ Database files
//
//   africa
//   antarctica
//   asia
//   australasia
//   backward
//   etcetera
//   europe
//   northamerica
//   southamerica
//
// from https://github.com/eggert/tz/releases/tag/2022g
//
// Supported Zones: 596 (351 zones, 245 links)
// Unsupported Zones: 0 (0 zones, 0 links)
//
// Original Years:  [1844,2087]
// Generated Years: [1844,2087]
// Estimator Years: [1800,2090]
// Max Buffer Size: 8
//
// Records:
//   Infos: 596
//   Eras: 1952
//   Policies: 134
//   Rules: 2158
//
// Memory:
//   Rules: 25896
//   Policies: 536
//   Eras: 27328
//   Zones: 4212
//   Links: 2940
//   Registry: 0
//   Formats: 1228
//   Letters: 106
//   Fragments: 0
//   Names: 9675
//   TOTAL: 71921
//
// DO NOT EDIT

package zonedb

import (
	"github.com/bxparks/AceTimeGo/internal"
	"testing"
)

// Verify that ZoneInfosData is identical to ZoneInfoRecords.
func TestZoneInfoReader(t *testing.T) {
	internal.VerifyZoneInfoReader(t, &DataContext, &RecordContext)
}

// Verify that ZoneErasData is identical to ZoneEraRecords.
func TestZoneEraReader(t *testing.T) {
	internal.VerifyZoneEraReader(t, &DataContext, &RecordContext)
}

// Verify that ZonePoliciesData is identical to ZonePolicyRecords.
func TestZonePolicyReader(t *testing.T) {
	internal.VerifyZonePolicyReader(t, &DataContext, &RecordContext)
}

// Verify that ZoneRulesData is identical to ZoneRuleRecords.
func TestZoneRuleReader(t *testing.T) {
	internal.VerifyZoneRuleReader(t, &DataContext, &RecordContext)
}
