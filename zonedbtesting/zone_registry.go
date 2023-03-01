// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/AceTimeGo/zonedbtesting/tzfiles
//     --output_dir /home/brian/src/AceTimeGo/zonedbtesting
//     --tz_version 2022g
//     --action zonedb
//     --language go
//     --scope extended
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
// from https://github.com/eggert/tz/releases/tag/2022g
//
// Supported Zones: 5 (4 zones, 1 links)
// Unsupported Zones: 591 (347 zones, 244 links)
// Original Years:  [1844,2087]
// Generated Years: [1967,2012]
// Estimator Years: [1967,2024]
// Max Buffer Size: 6
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
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// ---------------------------------------------------------------------------
// Zone Context
// ---------------------------------------------------------------------------

const TzDatabaseVersion string = "2022g"

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
	UntilYear: 10000,
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

// DataContext contains references to various XxxData objects and strings. These
// are the binary encoded versions of the various XxxRecord objects. This object
// is passed to the ZoneManager.
//
// The encoding to a binary string is performed because the Go language is able
// to treat strings as constants, and the TinyGo compiler can place them in
// flash memory, saving tremendous amounts of random memory.
var DataContext = zoneinfo.ZoneDataContext{
	TzDatabaseVersion: TzDatabaseVersion,
	StartYear: 1980,
	UntilYear: 10000,
	MaxTransitions: 6,
	LetterData: LetterData,
	LetterOffsets: LetterOffsets,
	FormatData: FormatData,
	FormatOffsets: FormatOffsets,
	NameData: NameData,
	NameOffsets: NameOffsets,
	ZoneRuleChunkSize: ZoneRuleChunkSize,
	ZonePolicyChunkSize: ZonePolicyChunkSize,
	ZoneEraChunkSize: ZoneEraChunkSize,
	ZoneInfoChunkSize: ZoneInfoChunkSize,
	ZoneRuleCount: ZoneRuleCount,
	ZonePolicyCount: ZonePolicyCount,
	ZoneEraCount: ZoneEraCount,
	ZoneInfoCount: ZoneInfoCount,
	ZoneRulesData: ZoneRulesData,
	ZonePoliciesData: ZonePoliciesData,
	ZoneErasData: ZoneErasData,
	ZoneInfosData: ZoneInfosData,
}

// ---------------------------------------------------------------------------
// Zone IDs. Unique stable uint32 identifier for each zone which can be given to
// ZoneManager.NewTimeZoneFromID(). Useful for microcontroller environments
// where saving variable length strings is more difficult than a fixed width
// integer.
//
// Total: 5 (4 zones, 1 links)
// ---------------------------------------------------------------------------

const (
	ZoneIDAmerica_Los_Angeles uint32 = 0xb7f7e8f2 // America/Los_Angeles
	ZoneIDAmerica_New_York uint32 = 0x1e2a7654 // America/New_York
	ZoneIDEtc_UTC uint32 = 0xd8e31abc // Etc/UTC
	ZoneIDPacific_Apia uint32 = 0x23359b5e // Pacific/Apia
	ZoneIDUS_Pacific uint32 = 0xa950f6ab // US/Pacific

)

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
