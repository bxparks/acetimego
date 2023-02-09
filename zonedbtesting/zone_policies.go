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
//
// DO NOT EDIT

package zonedbtesting

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// ---------------------------------------------------------------------------
// String constants.
// ---------------------------------------------------------------------------

const (
	// All ZoneRule.Letter entries concatenated together.
	LetterData = "" +
		"D" +
		"S" +
		"~"
)

var (
	// Byte offset into LetterData for each index. The actual Letter string
	// at index `i` given by the `ZoneRule.Letter` field is
	// `LetterData[LetterOffsets[i]:LetterOffsets[i+1]]`.
	LetterOffsets = []uint8{
		0, 0, 1, 2,
	}
)

// ---------------------------------------------------------------------------
// ZoneRuleRecords is a concatenated array of zoneinfo.ZoneInfoRecord objects
// from all ZonePolicyRecords.
//
// Supported zone policies: 2
// numRules: 12
// ---------------------------------------------------------------------------

var ZoneRuleRecords = []zoneinfo.ZoneRuleRecord{
	// ---------------------------------------------------------------------------
	// PolicyName: US
	// RuleIndex: 0
	// RuleCount: 6
	// ---------------------------------------------------------------------------

	// Rule    US    1967    2006    -    Oct    lastSun    2:00    0    S
	{
		FromYear: 1967,
		ToYear: 2006,
		InMonth: 10,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 4, // (delta_minutes=0)/15 + 4
		LetterIndex: 2, // "S"
	},
	// Rule    US    1975    only    -    Feb    lastSun    2:00    1:00    D
	{
		FromYear: 1975,
		ToYear: 1975,
		InMonth: 2,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 1, // "D"
	},
	// Rule    US    1976    1986    -    Apr    lastSun    2:00    1:00    D
	{
		FromYear: 1976,
		ToYear: 1986,
		InMonth: 4,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 1, // "D"
	},
	// Rule    US    1987    2006    -    Apr    Sun>=1    2:00    1:00    D
	{
		FromYear: 1987,
		ToYear: 2006,
		InMonth: 4,
		OnDayOfWeek: 7,
		OnDayOfMonth: 1,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 1, // "D"
	},
	// Rule    US    2007    max    -    Mar    Sun>=8    2:00    1:00    D
	{
		FromYear: 2007,
		ToYear: 9999,
		InMonth: 3,
		OnDayOfWeek: 7,
		OnDayOfMonth: 8,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 1, // "D"
	},
	// Rule    US    2007    max    -    Nov    Sun>=1    2:00    0    S
	{
		FromYear: 2007,
		ToYear: 9999,
		InMonth: 11,
		OnDayOfWeek: 7,
		OnDayOfMonth: 1,
		AtTimeCode: 8,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 4, // (delta_minutes=0)/15 + 4
		LetterIndex: 2, // "S"
	},

	// ---------------------------------------------------------------------------
	// PolicyName: WS
	// RuleIndex: 6
	// RuleCount: 6
	// ---------------------------------------------------------------------------

	// Anchor: Rule    WS    2011    only    -    Apr    Sat>=1    4:00    0    -
	{
		FromYear: 0,
		ToYear: 0,
		InMonth: 1,
		OnDayOfWeek: 0,
		OnDayOfMonth: 1,
		AtTimeCode: 0,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 4, // (delta_minutes=0)/15 + 4
		LetterIndex: 0, // ""
	},
	// Rule    WS    2010    only    -    Sep    lastSun    0:00    1    -
	{
		FromYear: 2010,
		ToYear: 2010,
		InMonth: 9,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 0,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 0, // ""
	},
	// Rule    WS    2011    only    -    Apr    Sat>=1    4:00    0    -
	{
		FromYear: 2011,
		ToYear: 2011,
		InMonth: 4,
		OnDayOfWeek: 6,
		OnDayOfMonth: 1,
		AtTimeCode: 16,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 4, // (delta_minutes=0)/15 + 4
		LetterIndex: 0, // ""
	},
	// Rule    WS    2011    only    -    Sep    lastSat    3:00    1    -
	{
		FromYear: 2011,
		ToYear: 2011,
		InMonth: 9,
		OnDayOfWeek: 6,
		OnDayOfMonth: 0,
		AtTimeCode: 12,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 0, // ""
	},
	// Rule    WS    2012    2021    -    Apr    Sun>=1    4:00    0    -
	{
		FromYear: 2012,
		ToYear: 2021,
		InMonth: 4,
		OnDayOfWeek: 7,
		OnDayOfMonth: 1,
		AtTimeCode: 16,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 4, // (delta_minutes=0)/15 + 4
		LetterIndex: 0, // ""
	},
	// Rule    WS    2012    2020    -    Sep    lastSun    3:00    1    -
	{
		FromYear: 2012,
		ToYear: 2020,
		InMonth: 9,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 12,
		AtTimeModifier: 0, // SuffixW + minute=0
		DeltaCode: 8, // (delta_minutes=60)/15 + 4
		LetterIndex: 0, // ""
	},


}

const ZoneRuleCount = 12

const ZoneRuleChunkSize = 11

// ZoneRulesData contains the ZoneRuleRecords data as a hex encoded string.
const ZoneRulesData = "\xaf\x07\xd6\x07\x0a\x07\x00\x08\x00\x04\x02" +
		"\xb7\x07\xb7\x07\x02\x07\x00\x08\x00\x08\x01" +
		"\xb8\x07\xc2\x07\x04\x07\x00\x08\x00\x08\x01" +
		"\xc3\x07\xd6\x07\x04\x07\x01\x08\x00\x08\x01" +
		"\xd7\x07\x0f\x27\x03\x07\x08\x08\x00\x08\x01" +
		"\xd7\x07\x0f\x27\x0b\x07\x01\x08\x00\x04\x02" +
		"\x00\x00\x00\x00\x01\x00\x01\x00\x00\x04\x00" +
		"\xda\x07\xda\x07\x09\x07\x00\x00\x00\x08\x00" +
		"\xdb\x07\xdb\x07\x04\x06\x01\x10\x00\x04\x00" +
		"\xdb\x07\xdb\x07\x09\x06\x00\x0c\x00\x08\x00" +
		"\xdc\x07\xe5\x07\x04\x07\x01\x10\x00\x04\x00" +
		"\xdc\x07\xe4\x07\x09\x07\x00\x0c\x00\x08\x00"

// ---------------------------------------------------------------------------
// ZonePolicyRecords contain indexes into the ZoneRuleRecords.
// Supported zone policies: 2
// ---------------------------------------------------------------------------

var ZonePolicyRecords = []zoneinfo.ZonePolicyRecord{
	{RuleIndex: 0, RuleCount: 0}, // 0: PolicyName: (None)
	{RuleIndex: 0, RuleCount: 6}, // 1: PolicyName: US
	{RuleIndex: 6, RuleCount: 6}, // 2: PolicyName: WS

}

const ZonePolicyCount = 3

const ZonePolicyChunkSize = 4

// ZonePoliciesData contains the ZonePolicyRecords data as a hex encoded string.
const ZonePoliciesData = "\x00\x00\x00\x00" +
		"\x00\x00\x06\x00" +
		"\x06\x00\x06\x00"

// ---------------------------------------------------------------------------
// Unsupported zone policies: 132
// ---------------------------------------------------------------------------

// AN {unused}
// AQ {unused}
// AS {unused}
// AT {unused}
// AV {unused}
// AW {unused}
// Albania {unused}
// Algeria {unused}
// Arg {unused}
// Armenia {unused}
// Aus {unused}
// Austria {unused}
// Azer {unused}
// Barb {unused}
// Belgium {unused}
// Belize {unused}
// Bermuda {unused}
// Brazil {unused}
// Bulg {unused}
// C-Eur {unused}
// CA {unused}
// CO {unused}
// CR {unused}
// Canada {unused}
// Chatham {unused}
// Chicago {unused}
// Chile {unused}
// Cook {unused}
// Cuba {unused}
// Cyprus {unused}
// Czech {unused}
// DR {unused}
// Denver {unused}
// Detroit {unused}
// Dhaka {unused}
// E-Eur {unused}
// E-EurAsia {unused}
// EU {unused}
// EUAsia {unused}
// Ecuador {unused}
// Edm {unused}
// Egypt {unused}
// EgyptAsia {unused}
// Eire {unused}
// Falk {unused}
// Fiji {unused}
// Finland {unused}
// France {unused}
// GB-Eire {unused}
// Germany {unused}
// Greece {unused}
// Guam {unused}
// Guat {unused}
// HK {unused}
// Haiti {unused}
// Halifax {unused}
// Holiday {unused}
// Hond {unused}
// Hungary {unused}
// Indianapolis {unused}
// Iran {unused}
// Iraq {unused}
// Italy {unused}
// Japan {unused}
// Jordan {unused}
// Kyrgyz {unused}
// LH {unused}
// Latvia {unused}
// Lebanon {unused}
// Libya {unused}
// Louisville {unused}
// Macau {unused}
// Malta {unused}
// Marengo {unused}
// Mauritius {unused}
// Menominee {unused}
// Mexico {unused}
// Moldova {unused}
// Moncton {unused}
// Mongol {unused}
// Morocco {unused}
// NBorneo {unused}
// NC {unused}
// NT_YK {unused}
// NYC {unused}
// NZ {unused}
// Namibia {unused}
// Nic {unused}
// PRC {unused}
// Pakistan {unused}
// Palestine {unused}
// Para {unused}
// Perry {unused}
// Peru {unused}
// Phil {unused}
// Pike {unused}
// Poland {unused}
// Port {unused}
// Pulaski {unused}
// ROK {unused}
// Regina {unused}
// Romania {unused}
// Russia {unused}
// RussiaAsia {unused}
// SA {unused}
// Salv {unused}
// SanLuis {unused}
// Shang {unused}
// SovietZone {unused}
// Spain {unused}
// SpainAfrica {unused}
// StJohns {unused}
// Starke {unused}
// Sudan {unused}
// Swift {unused}
// Swiss {unused}
// Syria {unused}
// Taiwan {unused}
// Thule {unused}
// Tonga {unused}
// Toronto {unused}
// Troll {unused}
// Tunisia {unused}
// Turkey {unused}
// Uruguay {unused}
// Vanc {unused}
// Vanuatu {unused}
// Vincennes {unused}
// W-Eur {unused}
// Winn {unused}
// Yukon {unused}
// Zion {unused}


// ---------------------------------------------------------------------------
// Notable zone policies: 1
// ---------------------------------------------------------------------------

// WS {Added anchor rule at year 0}


