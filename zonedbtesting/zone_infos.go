package zonedbtesting

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// This is a sample zone_infos.go created manually to allow development of the
// code that parses and utilizes these data structures. This will eventually be
// programmatically generated.

const TzDatabaseVersion string = "2022g"

//---------------------------------------------------------------------------
// Zone name: Africa/Abidjan
// Zone Eras: 1
//---------------------------------------------------------------------------

var ZoneEraAfrica_Abidjan = []zoneinfo.ZoneEra{
	//              0:00    -    GMT
	{
		ZonePolicy: nil,
		Format: "GMT",
		OffsetCode: 0,
		DeltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		UntilYear: 10000,
		UntilMonth: 1,
		UntilDay: 1,
		UntilTimeCode: 0,
		UntilTimeModifier: 0, /*(SuffixW + minute=0)*/
	},
}

const ZoneNameAfrica_Abidjan = "Africa/Abidjan"

var ZoneAfrica_Abidjan = zoneinfo.ZoneInfo{
	Name: ZoneNameAfrica_Abidjan,
	ZoneID: 0xc21305a3,
	StartYear: 2000,
	UntilYear: 10000,
	Eras: ZoneEraAfrica_Abidjan,
	Target: nil,
}

//---------------------------------------------------------------------------
// Zone name: America/Los_Angeles
// Zone Eras: 1
// Strings (bytes): 24 (originally 24)
// Memory (8-bit): 46
// Memory (32-bit): 60
//---------------------------------------------------------------------------

var ZoneEraAmerica_Los_Angeles = []zoneinfo.ZoneEra{
	//             -8:00    US    P%sT
	{
		ZonePolicy: &ZonePolicyUS,
		Format: "P%T",
		OffsetCode: -32,
		DeltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		UntilYear: 10000,
		UntilMonth: 1,
		UntilDay: 1,
		UntilTimeCode: 0,
		UntilTimeModifier: 0, /*(SuffixW + minute=0)*/
	},
}

const ZoneNameAmerica_Los_Angeles = "America/Los_Angeles"

var ZoneAmerica_Los_Angeles = zoneinfo.ZoneInfo{
	Name: ZoneNameAmerica_Los_Angeles,
	ZoneID: 0xb7f7e8f2,
	StartYear: 2000,
	UntilYear: 10000,
	Eras: ZoneEraAmerica_Los_Angeles,
	Target: nil,
}

//---------------------------------------------------------------------------
// Zone name: America/New_York
// Zone Eras: 1
// Strings (bytes): 21 (originally 21)
// Memory (8-bit): 43
// Memory (32-bit): 57
//---------------------------------------------------------------------------

var ZoneEraAmerica_New_York = []zoneinfo.ZoneEra{
	//             -5:00    US    E%sT
	{
		ZonePolicy: &ZonePolicyUS,
		Format: "E%T",
		OffsetCode: -20,
		DeltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		UntilYear: 10000,
		UntilMonth: 1,
		UntilDay: 1,
		UntilTimeCode: 0,
		UntilTimeModifier: 0, /*(SuffixW + minute=0)*/
	},
}

const ZoneNameAmerica_New_York = "America/New_York"

var ZoneAmerica_New_York = zoneinfo.ZoneInfo{
	Name: ZoneNameAmerica_New_York,
	ZoneID: 0x1e2a7654,
	StartYear: 2000,
	UntilYear: 10000,
	Eras: ZoneEraAmerica_New_York,
	Target: nil,
}

//---------------------------------------------------------------------------
// Zone name: Etc/UTC
// Zone Eras: 1
// Strings (bytes): 12 (originally 12)
// Memory (8-bit): 34
// Memory (32-bit): 48
//---------------------------------------------------------------------------

var ZoneEraEtc_UTC = []zoneinfo.ZoneEra{
	// 0 - UTC
	{
		ZonePolicy: nil,
		Format: "UTC",
		OffsetCode: 0,
		DeltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		UntilYear: 10000,
		UntilMonth: 1,
		UntilDay: 1,
		UntilTimeCode: 0,
		UntilTimeModifier: 0, /*(SuffixW + minute=0)*/
	},
}

const ZoneNameEtc_UTC = "Etc/UTC"

var ZoneEtc_UTC = zoneinfo.ZoneInfo{
	Name: ZoneNameEtc_UTC,
	ZoneID: 0xd8e31abc,
	StartYear: 2000,
	UntilYear: 10000,
	Eras: ZoneEraEtc_UTC,
	Target: nil,
}

//---------------------------------------------------------------------------
// Link name: US/Pacific -> America/Los_Angeles
// Strings (bytes): 11 (originally 11)
// Memory (8-bit): 22
// Memory (32-bit): 31
//---------------------------------------------------------------------------

const ZoneNameUS_Pacific = "US/Pacific"

var ZoneUS_Pacific  = zoneinfo.ZoneInfo{
  Name: ZoneNameUS_Pacific,
  ZoneID: 0xa950f6ab,
	StartYear: 2000,
	UntilYear: 10000,
	Target: &ZoneAmerica_Los_Angeles,
}
