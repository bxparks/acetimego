package acetime

// This is a sample zone_infos.go created manually to allow development of the
// code that parses and utilizes these data structures. This will eventually be
// programmatically generated.

const TzDatabaseVersion string = "2022g"

//---------------------------------------------------------------------------
// Zone name: Africa/Abidjan
// Zone Eras: 1
//---------------------------------------------------------------------------

var ZoneEraAfrica_Abidjan = []ZoneEra{
	//              0:00    -    GMT
	{
		zonePolicy: nil,
		format: "GMT",
		offsetCode: 0,
		deltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		untilYear: 10000,
		untilMonth: 1,
		untilDay: 1,
		untilTimeCode: 0,
		untilTimeModifier: 0, /*(suffixW + minute=0)*/
	},
}

const ZoneNameAfrica_Abidjan = "Africa/Abidjan"

var ZoneAfrica_Abidjan = ZoneInfo{
	name: ZoneNameAfrica_Abidjan,
	zoneID: 0xc21305a3,
	startYear: 2000,
	untilYear: 10000,
	eras: ZoneEraAfrica_Abidjan,
	target: nil,
}

//---------------------------------------------------------------------------
// Zone name: America/Los_Angeles
// Zone Eras: 1
// Strings (bytes): 24 (originally 24)
// Memory (8-bit): 46
// Memory (32-bit): 60
//---------------------------------------------------------------------------

var ZoneEraAmerica_Los_Angeles = []ZoneEra{
	//             -8:00    US    P%sT
	{
		zonePolicy: &ZonePolicyUS,
		format: "P%T",
		offsetCode: -32,
		deltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		untilYear: 10000,
		untilMonth: 1,
		untilDay: 1,
		untilTimeCode: 0,
		untilTimeModifier: 0, /*(suffixW + minute=0)*/
	},
}

const ZoneNameAmerica_Los_Angeles = "America/Los_Angeles"

var ZoneAmerica_Los_Angeles = ZoneInfo{
	name: ZoneNameAmerica_Los_Angeles,
	zoneID: 0xb7f7e8f2,
	startYear: 2000,
	untilYear: 10000,
	eras: ZoneEraAmerica_Los_Angeles,
	target: nil,
}

//---------------------------------------------------------------------------
// Zone name: America/New_York
// Zone Eras: 1
// Strings (bytes): 21 (originally 21)
// Memory (8-bit): 43
// Memory (32-bit): 57
//---------------------------------------------------------------------------

var ZoneEraAmerica_New_York = []ZoneEra{
	//             -5:00    US    E%sT
	{
		zonePolicy: &ZonePolicyUS,
		format: "E%T",
		offsetCode: -20,
		deltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		untilYear: 10000,
		untilMonth: 1,
		untilDay: 1,
		untilTimeCode: 0,
		untilTimeModifier: 0, /*(suffixW + minute=0)*/
	},
}

const ZoneNameAmerica_New_York = "America/New_York"

var ZoneAmerica_New_York = ZoneInfo{
	name: ZoneNameAmerica_New_York,
	zoneID: 0x1e2a7654,
	startYear: 2000,
	untilYear: 10000,
	eras: ZoneEraAmerica_New_York,
	target: nil,
}

//---------------------------------------------------------------------------
// Zone name: Etc/UTC
// Zone Eras: 1
// Strings (bytes): 12 (originally 12)
// Memory (8-bit): 34
// Memory (32-bit): 48
//---------------------------------------------------------------------------

var ZoneEraEtc_UTC = []ZoneEra{
	// 0 - UTC
	{
		zonePolicy: nil,
		format: "UTC",
		offsetCode: 0,
		deltaCode: 4, /*(((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		untilYear: 10000,
		untilMonth: 1,
		untilDay: 1,
		untilTimeCode: 0,
		untilTimeModifier: 0, /*(suffixW + minute=0)*/
	},
}

const ZoneNameEtc_UTC = "Etc/UTC"

var ZoneEtc_UTC = ZoneInfo{
	name: ZoneNameEtc_UTC,
	zoneID: 0xd8e31abc,
	startYear: 2000,
	untilYear: 10000,
	eras: ZoneEraEtc_UTC,
	target: nil,
}
