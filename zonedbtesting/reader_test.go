// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/acetimego/zonedbtesting/tzfiles
//     --output_dir /home/brian/src/acetimego/zonedbtesting
//     --tz_version 2023c
//     --action zonedb
//     --language go
//     --scope complete
//     --db_namespace zonedbtesting
//     --include_list include_list.txt
//     --start_year 1980
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
// from https://github.com/eggert/tz/releases/tag/2023c
//
// Supported Zones: 5 (4 zones, 1 links)
// Unsupported Zones: 591 (346 zones, 245 links)
//
// Original Years:  [1844,2087]
// Generated Years: [1967,2012]
// Estimator Years: [1967,2024]
// Max Buffer Size: 6
//
// Records:
//   Infos: 5
//   Eras: 5
//   Policies: 2
//   Rules: 13
//
// Memory:
//   Rules: 156
//   Policies: 8
//   Eras: 70
//   Zones: 48
//   Links: 12
//   Registry: 0
//   Formats: 38
//   Letters: 7
//   Fragments: 0
//   Names: 77
//   TOTAL: 416
//
// DO NOT EDIT

package zonedbtesting

import (
	"github.com/bxparks/acetimego/internal"
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
