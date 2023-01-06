package zonedbtesting

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// This is a sample zone_policies.go created by hand to help with developing the
// code that will parse and utilize these data structures. It will eventually be
// programmatically generated.

//---------------------------------------------------------------------------
// Policy name: US
// Rules: 5
// Memory (8-bit): 51
// Memory (32-bit): 72
//---------------------------------------------------------------------------

var ZoneRulesUS = []zoneinfo.ZoneRule{
	// Rule    US    1967    2006    -    Oct    lastSun    2:00    0    S
	{
		FromYear: 1967,
		ToYear: 2006,
		InMonth: 10,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 8,
		AtTimeModifier: 0, /*(SuffixW + minute=0)*/
		DeltaCode: 4, /*((deltaMinutes=0)/15 + 4)*/
		Letter: "S",
	},
	// Rule    US    1976    1986    -    Apr    lastSun    2:00    1:00    D
	{
		FromYear: 1976,
		ToYear: 1986,
		InMonth: 4,
		OnDayOfWeek: 7,
		OnDayOfMonth: 0,
		AtTimeCode: 8,
		AtTimeModifier: 0, /*(SuffixW + minute=0)*/
		DeltaCode: 8, /*((deltaMinutes=60)/15 + 4)*/
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
		AtTimeModifier: 0, /*(SuffixW + minute=0)*/
		DeltaCode: 8, /*((deltaMinutes=60)/15 + 4)*/
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
		AtTimeModifier: 0, /*(SuffixW + minute=0)*/
		DeltaCode: 8, /*((deltaMinutes=60)/15 + 4)*/
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
		AtTimeModifier: 0, /*(SuffixW + minute=0)*/
		DeltaCode: 4, /*((deltaMinutes=0)/15 + 4)*/
		Letter: "S",
	},
}

var ZonePolicyUS = zoneinfo.ZonePolicy{
	Rules: ZoneRulesUS,
	Letters: nil,
}
