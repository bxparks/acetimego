// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/acetimego/zonedbtesting/tzfiles
//     --output_dir /home/brian/src/acetimego/zonedbtesting
//     --tz_version 2024a
//     --actions zonedb
//     --languages go
//     --scope complete
//     --db_namespace zonedbtesting
//     --include_list include_list.txt
//     --start_year 1980
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
// from https://github.com/eggert/tz/releases/tag/2024a
//
// Supported Zones: 5 (4 zones, 1 links)
// Unsupported Zones: 591 (347 zones, 244 links)
//
// Requested Years: [1980,2200]
// Accurate Years: [1980,32767]
//
// Original Years:  [1844,2087]
// Generated Years: [1967,2012]
// Lower/Upper Truncated: [True,False]
//
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
	"github.com/bxparks/acetimego/zoneinfo"
)

// ---------------------------------------------------------------------------
// Zone Context
// ---------------------------------------------------------------------------

// RecordContext contains references to the various arrays of ZoneRuleRecord,
// ZonePolicyRecord, ZoneEraRecord, and ZoneInfoRecord objects, as well as the
// strings used by those objects.
//
// The `acetime` package uses the encoded XxxData objects, not the XxxRecord
// objects referenced here. These XxxRecord objects are used only for testing
// purposes, to verify that the XxxData objects were properly generated, and can
// be read back and reconstructed to be identical to the XxxRecord objects.
var RecordContext = zoneinfo.ZoneRecordContext{
	TzDatabaseVersion: TzDatabaseVersion,
	StartYear: 1980,
	UntilYear: 2200,
	StartYearAccurate: 1980,
	UntilYearAccurate: 32767,
	MaxTransitions: 6,
	LetterData: LetterData,
	LetterOffsets: LetterOffsets,
	FormatData: FormatData,
	FormatOffsets: FormatOffsets,
	NameData: NameData,
	NameOffsets: NameOffsets,
	ZoneRuleRecords: ZoneRuleRecords,
	ZonePolicyRecords: ZonePolicyRecords,
	ZoneEraRecords: ZoneEraRecords,
	ZoneInfoRecords: ZoneInfoRecords,
}

// ---------------------------------------------------------------------------
// Zone Indexes. Index into the ZoneInfoRecords array. Intended for unit tests
// which need direct access to the zoneinfo.ZoneInfo struct.
//
// Total: 5 (4 zones, 1 links)
// ---------------------------------------------------------------------------

const (
	ZoneInfoIndexAmerica_Los_Angeles uint16 = 3 // America/Los_Angeles
	ZoneInfoIndexAmerica_New_York uint16 = 0 // America/New_York
	ZoneInfoIndexEtc_UTC uint16 = 4 // Etc/UTC
	ZoneInfoIndexPacific_Apia uint16 = 1 // Pacific/Apia
	ZoneInfoIndexUS_Pacific uint16 = 2 // US/Pacific

)
