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
// DO NOT EDIT

package zonedbtesting

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// ---------------------------------------------------------------------------
// Supported zone policies: 1
// numRules: 6
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Policy name: US
// Rule count: 6
// ---------------------------------------------------------------------------
var ZoneRulesUS = []zoneinfo.ZoneRule{
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
		Letter: "S",
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
		Letter: "D",
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
		Letter: "D",
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
		Letter: "D",
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
		Letter: "D",
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
		Letter: "S",
	},

}

var ZonePolicyUS = zoneinfo.ZonePolicy{
	Rules: ZoneRulesUS,
	Letters: nil,
}



// ---------------------------------------------------------------------------
// Unsupported zone policies: 133
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
// WS {unused}
// Winn {unused}
// Yukon {unused}
// Zion {unused}


// ---------------------------------------------------------------------------
// Notable zone policies: 0
// ---------------------------------------------------------------------------


