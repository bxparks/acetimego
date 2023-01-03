package acetime

// This is a sample zone_info.go created by hand to help with developing the
// code that will parse and utilize these data structures. It will eventually be
// programmatically generated.

const TzDatabaseVersion string = "2022g"

//---------------------------------------------------------------------------
// Zone name: Africa/Abidjan
// Zone Eras: 1
// Strings (bytes): 19 (originally 19)
// Memory (8-bit): 41
// Memory (32-bit): 55
//---------------------------------------------------------------------------

var ZoneEraAfrica_Abidjan = []ZoneEra{
	//              0:00    -    GMT
	{
		nil,   /*zonePolicy*/
		"GMT", /*format*/
		0,     /*offsetCode*/
		4,     /*deltaCode (((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		10000, /*untilYear*/
		1,     /*untilMonth*/
		1,     /*untilDay*/
		0,     /*untilTimeCode*/
		0,     /*untilTimeModifier (suffixW + minute=0)*/
	},
}

const ZoneNameAfrica_Abidjan = "Africa/Abidjan"

var ZoneAfrica_Abidjan = ZoneInfo{
	ZoneNameAfrica_Abidjan, /*name*/
	0xc21305a3,             /*zone_id*/
	2000,                   /*startYear*/
	10000,                  /*untilYear*/
	ZoneEraAfrica_Abidjan,  /*eras*/
	nil,
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
		&ZonePolicyUS, /*zonePolicy*/
		"P%T",         /*format*/
		-32,           /*offsetCode*/
		4,             /*deltaCode (((offsetMinutes=0) << 4) + ((deltaMinutes=0)/15 + 4))*/
		10000,         /*untilYear*/
		1,             /*untilMonth*/
		1,             /*untilDay*/
		0,             /*untilTimeCode*/
		0,             /*untilTimeModifier (suffixW + minute=0)*/
	},
}

const ZoneNameAmerica_Los_Angeles = "America/Los_Angeles"

var ZoneAmerica_Los_Angeles = ZoneInfo{
	ZoneNameAmerica_Los_Angeles, /*name*/
	0xb7f7e8f2,                  /*zone_id*/
	2000,                        /*startYear*/
	10000,                       /*untilYear*/
	ZoneEraAmerica_Los_Angeles,  /*eras*/
	nil,                         /*targetInfo*/
}
