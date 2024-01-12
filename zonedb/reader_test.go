// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/acetimego/zonedb/tzfiles
//     --output_dir /home/brian/src/acetimego/zonedb
//     --tz_version 2023d
//     --actions zonedb
//     --languages go
//     --scope complete
//     --db_namespace zonedb
//     --start_year 2000
//     --until_year 2200
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
// from https://github.com/eggert/tz/releases/tag/2023d
//
// Supported Zones: 596 (351 zones, 245 links)
// Unsupported Zones: 0 (0 zones, 0 links)
//
// Requested Years: [2000,2200]
// Accurate Years: [2000,32767]
//
// Original Years:  [1844,2087]
// Generated Years: [1950,2087]
// Lower/Upper Truncated: [True,False]
//
// Estimator Years: [1950,2090]
// Max Buffer Size: 7
//
// Records:
//   Infos: 596
//   Eras: 655
//   Policies: 83
//   Rules: 735
//
// Memory:
//   Rules: 8820
//   Policies: 332
//   Eras: 9170
//   Zones: 4212
//   Links: 2940
//   Registry: 0
//   Formats: 712
//   Letters: 30
//   Fragments: 0
//   Names: 9675
//   TOTAL: 35891
//
// DO NOT EDIT

package zonedb

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
