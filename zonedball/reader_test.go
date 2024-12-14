// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/acetimego/zonedball/tzfiles
//     --output_dir /home/brian/src/acetimego/zonedball
//     --tz_version 2024b
//     --actions zonedb
//     --languages go
//     --scope complete
//     --db_namespace zonedball
//     --start_year 1800
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
// from https://github.com/eggert/tz/releases/tag/2024b
//
// Supported Zones: 596 (339 zones, 257 links)
// Unsupported Zones: 0 (0 zones, 0 links)
//
// Requested Years: [1800,2200]
// Accurate Years: [-32767,32767]
//
// Original Years:  [1844,2087]
// Generated Years: [1844,2087]
// Lower/Upper Truncated: [False,False]
//
// Estimator Years: [1800,2090]
// Max Buffer Size: 8
//
// Records:
//   Infos: 596
//   Eras: 1941
//   Policies: 134
//   Rules: 2231
//
// Memory:
//   Rules: 26772
//   Policies: 536
//   Eras: 27174
//   Zones: 4068
//   Links: 3084
//   Registry: 0
//   Formats: 600
//   Letters: 106
//   Fragments: 0
//   Names: 9675
//   TOTAL: 72015
//
// DO NOT EDIT

package zonedball

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
