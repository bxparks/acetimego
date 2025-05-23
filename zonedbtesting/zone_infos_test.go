// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeSuite/compiler/src/acetimecompiler/tzcompiler.py
//     --input_dir /home/brian/src/AceTimeSuite/libraries/acetimego/zonedbtesting/tzfiles
//     --output_dir /home/brian/src/AceTimeSuite/libraries/acetimego/zonedbtesting
//     --tz_version 2025b
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
// from https://github.com/eggert/tz/releases/tag/2025b
//
// Supported Zones: 5 (4 zones, 1 links)
// Unsupported Zones: 592 (336 zones, 256 links)
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
//   Formats: 24
//   Letters: 7
//   Fragments: 0
//   Names: 77
//   TOTAL: 402
//
// DO NOT EDIT

package zonedbtesting

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

// ---------------------------------------------------------------------------
// ZoneEraRecords is an array of ZoneEraRecord items concatenated together
// across all zones.
//
// Supported zones: 4
// numEras: 5
// ---------------------------------------------------------------------------

var ZoneEraRecords = []zoneinfo.ZoneEraRecord{
	// ---------------------------------------------------------------------------
	// ZoneName: America/Los_Angeles
	// EraIndex: 0
	// EraCount: 1
	// ---------------------------------------------------------------------------

	//             -8:00    US    P%sT
	{
		PolicyIndex: 1, // PolicyName: US
		FormatIndex: 3, // "P%T"
		DeltaMinutes: 0,
		OffsetSecondsCode: -1920, // -28800 / 15
		OffsetSecondsRemainder: 0,
		UntilYear: 32767,
		UntilMonth: 1,
		UntilDay: 1,
		UntilSecondsCode: 0, // 0 / 15
		UntilSecondsModifier: 0, // SuffixW + remainder=0
	},

	// ---------------------------------------------------------------------------
	// ZoneName: America/New_York
	// EraIndex: 1
	// EraCount: 1
	// ---------------------------------------------------------------------------

	//             -5:00    US    E%sT
	{
		PolicyIndex: 1, // PolicyName: US
		FormatIndex: 2, // "E%T"
		DeltaMinutes: 0,
		OffsetSecondsCode: -1200, // -18000 / 15
		OffsetSecondsRemainder: 0,
		UntilYear: 32767,
		UntilMonth: 1,
		UntilDay: 1,
		UntilSecondsCode: 0, // 0 / 15
		UntilSecondsModifier: 0, // SuffixW + remainder=0
	},

	// ---------------------------------------------------------------------------
	// ZoneName: Etc/UTC
	// EraIndex: 2
	// EraCount: 1
	// ---------------------------------------------------------------------------

	// 0 - UTC
	{
		PolicyIndex: 0, // PolicyName: (none)
		FormatIndex: 4, // "UTC"
		DeltaMinutes: 0,
		OffsetSecondsCode: 0, // 0 / 15
		OffsetSecondsRemainder: 0,
		UntilYear: 32767,
		UntilMonth: 1,
		UntilDay: 1,
		UntilSecondsCode: 0, // 0 / 15
		UntilSecondsModifier: 0, // SuffixW + remainder=0
	},

	// ---------------------------------------------------------------------------
	// ZoneName: Pacific/Apia
	// EraIndex: 3
	// EraCount: 2
	// ---------------------------------------------------------------------------

	//             -11:00    WS    %z    2011 Dec 29 24:00
	{
		PolicyIndex: 2, // PolicyName: WS
		FormatIndex: 0, // ""
		DeltaMinutes: 0,
		OffsetSecondsCode: -2640, // -39600 / 15
		OffsetSecondsRemainder: 0,
		UntilYear: 2011,
		UntilMonth: 12,
		UntilDay: 29,
		UntilSecondsCode: 5760, // 86400 / 15
		UntilSecondsModifier: 0, // SuffixW + remainder=0
	},

	//              13:00    WS    %z
	{
		PolicyIndex: 2, // PolicyName: WS
		FormatIndex: 0, // ""
		DeltaMinutes: 0,
		OffsetSecondsCode: 3120, // 46800 / 15
		OffsetSecondsRemainder: 0,
		UntilYear: 32767,
		UntilMonth: 1,
		UntilDay: 1,
		UntilSecondsCode: 0, // 0 / 15
		UntilSecondsModifier: 0, // SuffixW + remainder=0
	},


}

// ---------------------------------------------------------------------------
// ZoneInfoRecords is an array of ZoneInfoRecord items concatenated together
// across all zones.
//
// Total: 5 (4 zones, 1 links)
// ---------------------------------------------------------------------------

var ZoneInfoRecords = []zoneinfo.ZoneInfoRecord{
	// 0: Zone America/New_York
	{
		ZoneID: 0x1e2a7654,
		NameIndex: 1, // "America/New_York"
		EraIndex: 1,
		EraCount: 1,
		TargetIndex: 0,
	},
	// 1: Zone Pacific/Apia
	{
		ZoneID: 0x23359b5e,
		NameIndex: 3, // "Pacific/Apia"
		EraIndex: 3,
		EraCount: 2,
		TargetIndex: 0,
	},
	// 2: Link US/Pacific -> America/Los_Angeles
	{
		ZoneID: 0xa950f6ab,
		NameIndex: 4, // "US/Pacific"
		EraIndex: 0,
		EraCount: 0, // IsLink=true
		TargetIndex: 3, // America/Los_Angeles
	},
	// 3: Zone America/Los_Angeles
	{
		ZoneID: 0xb7f7e8f2,
		NameIndex: 0, // "America/Los_Angeles"
		EraIndex: 0,
		EraCount: 1,
		TargetIndex: 0,
	},
	// 4: Zone Etc/UTC
	{
		ZoneID: 0xd8e31abc,
		NameIndex: 2, // "Etc/UTC"
		EraIndex: 2,
		EraCount: 1,
		TargetIndex: 0,
	},

}
