// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/acetimego/zonedb/tzfiles
//     --output_dir /home/brian/src/acetimego/zonedb
//     --tz_version 2023c
//     --action zonedb
//     --language go
//     --scope complete
//     --db_namespace zonedb
//     --start_year 2000
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
// from https://github.com/eggert/tz/releases/tag/2023c
//
// Supported Zones: 596 (350 zones, 246 links)
// Unsupported Zones: 0 (0 zones, 0 links)
//
// Requested Years: [2000,2200]
// Accurate Years: [2000,32767]
//
// Original Years:  [1844,2087]
// Generated Years: [1950,2087]
// Lower/Upper Truncated: [True,False]
//
// Estimator Years: [1950,2090]
// Max Buffer Size: 7
//
// Records:
//   Infos: 596
//   Eras: 646
//   Policies: 83
//   Rules: 735
//
// Memory:
//   Rules: 8820
//   Policies: 332
//   Eras: 9044
//   Zones: 4200
//   Links: 2952
//   Registry: 0
//   Formats: 712
//   Letters: 30
//   Fragments: 0
//   Names: 9675
//   TOTAL: 35765
//
// DO NOT EDIT

package zonedb

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

// ---------------------------------------------------------------------------
// Zone Context
// ---------------------------------------------------------------------------

const TzDatabaseVersion string = "2023c"

// DataContext contains references to various XxxData objects and strings. These
// are the binary encoded versions of the various XxxRecord objects. This object
// is passed to the ZoneManager.
//
// The encoding to a binary string is performed because the Go language is able
// to treat strings as constants, and the TinyGo compiler can place them in
// flash memory, saving tremendous amounts of random memory.
var DataContext = zoneinfo.ZoneDataContext{
	TzDatabaseVersion: TzDatabaseVersion,
	StartYear: 2000,
	UntilYear: 2200,
	StartYearAccurate: 2000,
	UntilYearAccurate: 32767,
	MaxTransitions: 7,
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
// Total: 596 (350 zones, 246 links)
// ---------------------------------------------------------------------------

const (
	ZoneIDAfrica_Abidjan uint32 = 0xc21305a3 // Africa/Abidjan
	ZoneIDAfrica_Accra uint32 = 0x77d5b054 // Africa/Accra
	ZoneIDAfrica_Addis_Ababa uint32 = 0x05ae1e65 // Africa/Addis_Ababa
	ZoneIDAfrica_Algiers uint32 = 0xd94515c1 // Africa/Algiers
	ZoneIDAfrica_Asmara uint32 = 0x73b278ef // Africa/Asmara
	ZoneIDAfrica_Asmera uint32 = 0x73b289f3 // Africa/Asmera
	ZoneIDAfrica_Bamako uint32 = 0x74c1e7a5 // Africa/Bamako
	ZoneIDAfrica_Bangui uint32 = 0x74c28ed0 // Africa/Bangui
	ZoneIDAfrica_Banjul uint32 = 0x74c29b96 // Africa/Banjul
	ZoneIDAfrica_Bissau uint32 = 0x75564141 // Africa/Bissau
	ZoneIDAfrica_Blantyre uint32 = 0xe08d813b // Africa/Blantyre
	ZoneIDAfrica_Brazzaville uint32 = 0x39cda760 // Africa/Brazzaville
	ZoneIDAfrica_Bujumbura uint32 = 0x05232a47 // Africa/Bujumbura
	ZoneIDAfrica_Cairo uint32 = 0x77f8e228 // Africa/Cairo
	ZoneIDAfrica_Casablanca uint32 = 0xc59f1b33 // Africa/Casablanca
	ZoneIDAfrica_Ceuta uint32 = 0x77fb46ec // Africa/Ceuta
	ZoneIDAfrica_Conakry uint32 = 0x7ab36b31 // Africa/Conakry
	ZoneIDAfrica_Dakar uint32 = 0x780b00fd // Africa/Dakar
	ZoneIDAfrica_Dar_es_Salaam uint32 = 0xa04c47b6 // Africa/Dar_es_Salaam
	ZoneIDAfrica_Djibouti uint32 = 0x30ea01d4 // Africa/Djibouti
	ZoneIDAfrica_Douala uint32 = 0x7a6df310 // Africa/Douala
	ZoneIDAfrica_El_Aaiun uint32 = 0x9d6fb118 // Africa/El_Aaiun
	ZoneIDAfrica_Freetown uint32 = 0x6823dd64 // Africa/Freetown
	ZoneIDAfrica_Gaborone uint32 = 0x317c0aa7 // Africa/Gaborone
	ZoneIDAfrica_Harare uint32 = 0x82c39a2d // Africa/Harare
	ZoneIDAfrica_Johannesburg uint32 = 0xd5d157a0 // Africa/Johannesburg
	ZoneIDAfrica_Juba uint32 = 0xd51b395c // Africa/Juba
	ZoneIDAfrica_Kampala uint32 = 0xc1d30e31 // Africa/Kampala
	ZoneIDAfrica_Khartoum uint32 = 0xfb3d4205 // Africa/Khartoum
	ZoneIDAfrica_Kigali uint32 = 0x8a4dcf2b // Africa/Kigali
	ZoneIDAfrica_Kinshasa uint32 = 0x6695d70c // Africa/Kinshasa
	ZoneIDAfrica_Lagos uint32 = 0x789bb5d0 // Africa/Lagos
	ZoneIDAfrica_Libreville uint32 = 0x01d96de4 // Africa/Libreville
	ZoneIDAfrica_Lome uint32 = 0xd51c3a07 // Africa/Lome
	ZoneIDAfrica_Luanda uint32 = 0x8d7909cf // Africa/Luanda
	ZoneIDAfrica_Lubumbashi uint32 = 0x6fd88566 // Africa/Lubumbashi
	ZoneIDAfrica_Lusaka uint32 = 0x8d82b23b // Africa/Lusaka
	ZoneIDAfrica_Malabo uint32 = 0x8e6a1906 // Africa/Malabo
	ZoneIDAfrica_Maputo uint32 = 0x8e6ca1f0 // Africa/Maputo
	ZoneIDAfrica_Maseru uint32 = 0x8e6e02c7 // Africa/Maseru
	ZoneIDAfrica_Mbabane uint32 = 0x5d3bdd40 // Africa/Mbabane
	ZoneIDAfrica_Mogadishu uint32 = 0x66bc159b // Africa/Mogadishu
	ZoneIDAfrica_Monrovia uint32 = 0x0ce90385 // Africa/Monrovia
	ZoneIDAfrica_Nairobi uint32 = 0xa87ab57e // Africa/Nairobi
	ZoneIDAfrica_Ndjamena uint32 = 0x9fe09898 // Africa/Ndjamena
	ZoneIDAfrica_Niamey uint32 = 0x914a30fd // Africa/Niamey
	ZoneIDAfrica_Nouakchott uint32 = 0x3dc49dba // Africa/Nouakchott
	ZoneIDAfrica_Ouagadougou uint32 = 0x04d7219a // Africa/Ouagadougou
	ZoneIDAfrica_Porto_Novo uint32 = 0x3d1bf95d // Africa/Porto-Novo
	ZoneIDAfrica_Sao_Tome uint32 = 0x61b319d1 // Africa/Sao_Tome
	ZoneIDAfrica_Timbuktu uint32 = 0xb164d56f // Africa/Timbuktu
	ZoneIDAfrica_Tripoli uint32 = 0x9dfebd3d // Africa/Tripoli
	ZoneIDAfrica_Tunis uint32 = 0x79378e6d // Africa/Tunis
	ZoneIDAfrica_Windhoek uint32 = 0x789c9bd3 // Africa/Windhoek
	ZoneIDAmerica_Adak uint32 = 0x97fe49d7 // America/Adak
	ZoneIDAmerica_Anchorage uint32 = 0x5a79260e // America/Anchorage
	ZoneIDAmerica_Anguilla uint32 = 0xafe31333 // America/Anguilla
	ZoneIDAmerica_Antigua uint32 = 0xc067a32f // America/Antigua
	ZoneIDAmerica_Araguaina uint32 = 0x6f9a3aef // America/Araguaina
	ZoneIDAmerica_Argentina_Buenos_Aires uint32 = 0xd43b4c0d // America/Argentina/Buenos_Aires
	ZoneIDAmerica_Argentina_Catamarca uint32 = 0x8d40986b // America/Argentina/Catamarca
	ZoneIDAmerica_Argentina_ComodRivadavia uint32 = 0x22758877 // America/Argentina/ComodRivadavia
	ZoneIDAmerica_Argentina_Cordoba uint32 = 0xbfccc308 // America/Argentina/Cordoba
	ZoneIDAmerica_Argentina_Jujuy uint32 = 0x5f2f46c5 // America/Argentina/Jujuy
	ZoneIDAmerica_Argentina_La_Rioja uint32 = 0xa46b7eef // America/Argentina/La_Rioja
	ZoneIDAmerica_Argentina_Mendoza uint32 = 0xa9f72d5c // America/Argentina/Mendoza
	ZoneIDAmerica_Argentina_Rio_Gallegos uint32 = 0xc5b0f565 // America/Argentina/Rio_Gallegos
	ZoneIDAmerica_Argentina_Salta uint32 = 0x5fc73403 // America/Argentina/Salta
	ZoneIDAmerica_Argentina_San_Juan uint32 = 0x3e1009bd // America/Argentina/San_Juan
	ZoneIDAmerica_Argentina_San_Luis uint32 = 0x3e11238c // America/Argentina/San_Luis
	ZoneIDAmerica_Argentina_Tucuman uint32 = 0xe96399eb // America/Argentina/Tucuman
	ZoneIDAmerica_Argentina_Ushuaia uint32 = 0x320dcdde // America/Argentina/Ushuaia
	ZoneIDAmerica_Aruba uint32 = 0x97cf8651 // America/Aruba
	ZoneIDAmerica_Asuncion uint32 = 0x50ec79a6 // America/Asuncion
	ZoneIDAmerica_Atikokan uint32 = 0x81b92098 // America/Atikokan
	ZoneIDAmerica_Atka uint32 = 0x97fe8f27 // America/Atka
	ZoneIDAmerica_Bahia uint32 = 0x97d815fb // America/Bahia
	ZoneIDAmerica_Bahia_Banderas uint32 = 0x14f6329a // America/Bahia_Banderas
	ZoneIDAmerica_Barbados uint32 = 0xcbbc3b04 // America/Barbados
	ZoneIDAmerica_Belem uint32 = 0x97da580b // America/Belem
	ZoneIDAmerica_Belize uint32 = 0x93256c81 // America/Belize
	ZoneIDAmerica_Blanc_Sablon uint32 = 0x6e299892 // America/Blanc-Sablon
	ZoneIDAmerica_Boa_Vista uint32 = 0x0a7b7efe // America/Boa_Vista
	ZoneIDAmerica_Bogota uint32 = 0x93d7bc62 // America/Bogota
	ZoneIDAmerica_Boise uint32 = 0x97dfc8d8 // America/Boise
	ZoneIDAmerica_Buenos_Aires uint32 = 0x67d79a05 // America/Buenos_Aires
	ZoneIDAmerica_Cambridge_Bay uint32 = 0xd5a44aff // America/Cambridge_Bay
	ZoneIDAmerica_Campo_Grande uint32 = 0xfec3e7a6 // America/Campo_Grande
	ZoneIDAmerica_Cancun uint32 = 0x953331be // America/Cancun
	ZoneIDAmerica_Caracas uint32 = 0x3be064f4 // America/Caracas
	ZoneIDAmerica_Catamarca uint32 = 0x5036e963 // America/Catamarca
	ZoneIDAmerica_Cayenne uint32 = 0x3c617269 // America/Cayenne
	ZoneIDAmerica_Cayman uint32 = 0x953961df // America/Cayman
	ZoneIDAmerica_Chicago uint32 = 0x4b92b5d4 // America/Chicago
	ZoneIDAmerica_Chihuahua uint32 = 0x8827d776 // America/Chihuahua
	ZoneIDAmerica_Ciudad_Juarez uint32 = 0x7347fc60 // America/Ciudad_Juarez
	ZoneIDAmerica_Coral_Harbour uint32 = 0xabcb7569 // America/Coral_Harbour
	ZoneIDAmerica_Cordoba uint32 = 0x5c8a7600 // America/Cordoba
	ZoneIDAmerica_Costa_Rica uint32 = 0x63ff66be // America/Costa_Rica
	ZoneIDAmerica_Creston uint32 = 0x62a70204 // America/Creston
	ZoneIDAmerica_Cuiaba uint32 = 0x969a52eb // America/Cuiaba
	ZoneIDAmerica_Curacao uint32 = 0x6a879184 // America/Curacao
	ZoneIDAmerica_Danmarkshavn uint32 = 0xf554d204 // America/Danmarkshavn
	ZoneIDAmerica_Dawson uint32 = 0x978d8d12 // America/Dawson
	ZoneIDAmerica_Dawson_Creek uint32 = 0x6cf24e5b // America/Dawson_Creek
	ZoneIDAmerica_Denver uint32 = 0x97d10b2a // America/Denver
	ZoneIDAmerica_Detroit uint32 = 0x925cfbc1 // America/Detroit
	ZoneIDAmerica_Dominica uint32 = 0xcecb4c4a // America/Dominica
	ZoneIDAmerica_Edmonton uint32 = 0x6cb9484a // America/Edmonton
	ZoneIDAmerica_Eirunepe uint32 = 0xf9b29683 // America/Eirunepe
	ZoneIDAmerica_El_Salvador uint32 = 0x752ad652 // America/El_Salvador
	ZoneIDAmerica_Ensenada uint32 = 0x7bc95445 // America/Ensenada
	ZoneIDAmerica_Fort_Nelson uint32 = 0x3f437e0f // America/Fort_Nelson
	ZoneIDAmerica_Fort_Wayne uint32 = 0x7eaaaf24 // America/Fort_Wayne
	ZoneIDAmerica_Fortaleza uint32 = 0x2ad018ee // America/Fortaleza
	ZoneIDAmerica_Glace_Bay uint32 = 0x9681f8dd // America/Glace_Bay
	ZoneIDAmerica_Godthab uint32 = 0x8f7eba1f // America/Godthab
	ZoneIDAmerica_Goose_Bay uint32 = 0xb649541e // America/Goose_Bay
	ZoneIDAmerica_Grand_Turk uint32 = 0x6e216197 // America/Grand_Turk
	ZoneIDAmerica_Grenada uint32 = 0x968ce4d8 // America/Grenada
	ZoneIDAmerica_Guadeloupe uint32 = 0xcd1f8a31 // America/Guadeloupe
	ZoneIDAmerica_Guatemala uint32 = 0x0c8259f7 // America/Guatemala
	ZoneIDAmerica_Guayaquil uint32 = 0x17e64958 // America/Guayaquil
	ZoneIDAmerica_Guyana uint32 = 0x9ff7bd0b // America/Guyana
	ZoneIDAmerica_Halifax uint32 = 0xbc5b7183 // America/Halifax
	ZoneIDAmerica_Havana uint32 = 0xa0e15675 // America/Havana
	ZoneIDAmerica_Hermosillo uint32 = 0x065d21c4 // America/Hermosillo
	ZoneIDAmerica_Indiana_Indianapolis uint32 = 0x28a669a4 // America/Indiana/Indianapolis
	ZoneIDAmerica_Indiana_Knox uint32 = 0x6554adc9 // America/Indiana/Knox
	ZoneIDAmerica_Indiana_Marengo uint32 = 0x2feeee72 // America/Indiana/Marengo
	ZoneIDAmerica_Indiana_Petersburg uint32 = 0x94ac7acc // America/Indiana/Petersburg
	ZoneIDAmerica_Indiana_Tell_City uint32 = 0x09263612 // America/Indiana/Tell_City
	ZoneIDAmerica_Indiana_Vevay uint32 = 0x10aca054 // America/Indiana/Vevay
	ZoneIDAmerica_Indiana_Vincennes uint32 = 0x28a0b212 // America/Indiana/Vincennes
	ZoneIDAmerica_Indiana_Winamac uint32 = 0x4413fa69 // America/Indiana/Winamac
	ZoneIDAmerica_Indianapolis uint32 = 0x6a009ae1 // America/Indianapolis
	ZoneIDAmerica_Inuvik uint32 = 0xa42189fc // America/Inuvik
	ZoneIDAmerica_Iqaluit uint32 = 0x2de310bf // America/Iqaluit
	ZoneIDAmerica_Jamaica uint32 = 0x565dad6c // America/Jamaica
	ZoneIDAmerica_Jujuy uint32 = 0x9873dbbd // America/Jujuy
	ZoneIDAmerica_Juneau uint32 = 0xa6f13e2e // America/Juneau
	ZoneIDAmerica_Kentucky_Louisville uint32 = 0x1a21024b // America/Kentucky/Louisville
	ZoneIDAmerica_Kentucky_Monticello uint32 = 0xde71c439 // America/Kentucky/Monticello
	ZoneIDAmerica_Knox_IN uint32 = 0xc1db9a1c // America/Knox_IN
	ZoneIDAmerica_Kralendijk uint32 = 0xe7c456c5 // America/Kralendijk
	ZoneIDAmerica_La_Paz uint32 = 0xaa29125d // America/La_Paz
	ZoneIDAmerica_Lima uint32 = 0x980468c9 // America/Lima
	ZoneIDAmerica_Los_Angeles uint32 = 0xb7f7e8f2 // America/Los_Angeles
	ZoneIDAmerica_Louisville uint32 = 0x3dcb47ee // America/Louisville
	ZoneIDAmerica_Lower_Princes uint32 = 0x6ae45b62 // America/Lower_Princes
	ZoneIDAmerica_Maceio uint32 = 0xac80c6d4 // America/Maceio
	ZoneIDAmerica_Managua uint32 = 0x3d5e7600 // America/Managua
	ZoneIDAmerica_Manaus uint32 = 0xac86bf8b // America/Manaus
	ZoneIDAmerica_Marigot uint32 = 0x3dab3a59 // America/Marigot
	ZoneIDAmerica_Martinique uint32 = 0x551e84c5 // America/Martinique
	ZoneIDAmerica_Matamoros uint32 = 0xdd1b0259 // America/Matamoros
	ZoneIDAmerica_Mazatlan uint32 = 0x0532189e // America/Mazatlan
	ZoneIDAmerica_Mendoza uint32 = 0x46b4e054 // America/Mendoza
	ZoneIDAmerica_Menominee uint32 = 0xe0e9c583 // America/Menominee
	ZoneIDAmerica_Merida uint32 = 0xacd172d8 // America/Merida
	ZoneIDAmerica_Metlakatla uint32 = 0x84de2686 // America/Metlakatla
	ZoneIDAmerica_Mexico_City uint32 = 0xd0d93f43 // America/Mexico_City
	ZoneIDAmerica_Miquelon uint32 = 0x59674330 // America/Miquelon
	ZoneIDAmerica_Moncton uint32 = 0x5e07fe24 // America/Moncton
	ZoneIDAmerica_Monterrey uint32 = 0x269a1deb // America/Monterrey
	ZoneIDAmerica_Montevideo uint32 = 0xfa214780 // America/Montevideo
	ZoneIDAmerica_Montreal uint32 = 0x203a1ea8 // America/Montreal
	ZoneIDAmerica_Montserrat uint32 = 0x199b0a35 // America/Montserrat
	ZoneIDAmerica_Nassau uint32 = 0xaedef011 // America/Nassau
	ZoneIDAmerica_New_York uint32 = 0x1e2a7654 // America/New_York
	ZoneIDAmerica_Nipigon uint32 = 0x9d2a8b1a // America/Nipigon
	ZoneIDAmerica_Nome uint32 = 0x98059b15 // America/Nome
	ZoneIDAmerica_Noronha uint32 = 0xab5116fb // America/Noronha
	ZoneIDAmerica_North_Dakota_Beulah uint32 = 0x9b52b384 // America/North_Dakota/Beulah
	ZoneIDAmerica_North_Dakota_Center uint32 = 0x9da42814 // America/North_Dakota/Center
	ZoneIDAmerica_North_Dakota_New_Salem uint32 = 0x04f9958e // America/North_Dakota/New_Salem
	ZoneIDAmerica_Nuuk uint32 = 0x9805b5a9 // America/Nuuk
	ZoneIDAmerica_Ojinaga uint32 = 0xebfde83f // America/Ojinaga
	ZoneIDAmerica_Panama uint32 = 0xb3863854 // America/Panama
	ZoneIDAmerica_Pangnirtung uint32 = 0x2d999193 // America/Pangnirtung
	ZoneIDAmerica_Paramaribo uint32 = 0xb319e4c4 // America/Paramaribo
	ZoneIDAmerica_Phoenix uint32 = 0x34b5af01 // America/Phoenix
	ZoneIDAmerica_Port_au_Prince uint32 = 0x8e4a7bdc // America/Port-au-Prince
	ZoneIDAmerica_Port_of_Spain uint32 = 0xd8b28d59 // America/Port_of_Spain
	ZoneIDAmerica_Porto_Acre uint32 = 0xcce5bf54 // America/Porto_Acre
	ZoneIDAmerica_Porto_Velho uint32 = 0x6b1aac77 // America/Porto_Velho
	ZoneIDAmerica_Puerto_Rico uint32 = 0x6752ca31 // America/Puerto_Rico
	ZoneIDAmerica_Punta_Arenas uint32 = 0xc2c3bce7 // America/Punta_Arenas
	ZoneIDAmerica_Rainy_River uint32 = 0x9cd58a10 // America/Rainy_River
	ZoneIDAmerica_Rankin_Inlet uint32 = 0xc8de4984 // America/Rankin_Inlet
	ZoneIDAmerica_Recife uint32 = 0xb8730494 // America/Recife
	ZoneIDAmerica_Regina uint32 = 0xb875371c // America/Regina
	ZoneIDAmerica_Resolute uint32 = 0xc7093459 // America/Resolute
	ZoneIDAmerica_Rio_Branco uint32 = 0x9d352764 // America/Rio_Branco
	ZoneIDAmerica_Rosario uint32 = 0xdf448665 // America/Rosario
	ZoneIDAmerica_Santa_Isabel uint32 = 0xfd18a56c // America/Santa_Isabel
	ZoneIDAmerica_Santarem uint32 = 0x740caec1 // America/Santarem
	ZoneIDAmerica_Santiago uint32 = 0x7410c9bc // America/Santiago
	ZoneIDAmerica_Santo_Domingo uint32 = 0x75a0d177 // America/Santo_Domingo
	ZoneIDAmerica_Sao_Paulo uint32 = 0x1063bfc9 // America/Sao_Paulo
	ZoneIDAmerica_Scoresbysund uint32 = 0x123f8d2a // America/Scoresbysund
	ZoneIDAmerica_Shiprock uint32 = 0x82fb7049 // America/Shiprock
	ZoneIDAmerica_Sitka uint32 = 0x99104ce2 // America/Sitka
	ZoneIDAmerica_St_Barthelemy uint32 = 0x054e6a79 // America/St_Barthelemy
	ZoneIDAmerica_St_Johns uint32 = 0x04b14e6e // America/St_Johns
	ZoneIDAmerica_St_Kitts uint32 = 0x04c0507b // America/St_Kitts
	ZoneIDAmerica_St_Lucia uint32 = 0x04d8b3ba // America/St_Lucia
	ZoneIDAmerica_St_Thomas uint32 = 0xb1b3d778 // America/St_Thomas
	ZoneIDAmerica_St_Vincent uint32 = 0x8460e523 // America/St_Vincent
	ZoneIDAmerica_Swift_Current uint32 = 0xdef98e55 // America/Swift_Current
	ZoneIDAmerica_Tegucigalpa uint32 = 0xbfd6fd4c // America/Tegucigalpa
	ZoneIDAmerica_Thule uint32 = 0x9921dd68 // America/Thule
	ZoneIDAmerica_Thunder_Bay uint32 = 0xf962e71b // America/Thunder_Bay
	ZoneIDAmerica_Tijuana uint32 = 0x6aa1df72 // America/Tijuana
	ZoneIDAmerica_Toronto uint32 = 0x792e851b // America/Toronto
	ZoneIDAmerica_Tortola uint32 = 0x7931462b // America/Tortola
	ZoneIDAmerica_Vancouver uint32 = 0x2c6f6b1f // America/Vancouver
	ZoneIDAmerica_Virgin uint32 = 0xc2183ab5 // America/Virgin
	ZoneIDAmerica_Whitehorse uint32 = 0x54e0e3e8 // America/Whitehorse
	ZoneIDAmerica_Winnipeg uint32 = 0x8c7dafc7 // America/Winnipeg
	ZoneIDAmerica_Yakutat uint32 = 0xd8ee31e9 // America/Yakutat
	ZoneIDAmerica_Yellowknife uint32 = 0x0f76c76f // America/Yellowknife
	ZoneIDAntarctica_Casey uint32 = 0xe2022583 // Antarctica/Casey
	ZoneIDAntarctica_Davis uint32 = 0xe2144b45 // Antarctica/Davis
	ZoneIDAntarctica_DumontDUrville uint32 = 0x5a3c656c // Antarctica/DumontDUrville
	ZoneIDAntarctica_Macquarie uint32 = 0x92f47626 // Antarctica/Macquarie
	ZoneIDAntarctica_Mawson uint32 = 0x399cd863 // Antarctica/Mawson
	ZoneIDAntarctica_McMurdo uint32 = 0x6eeb5585 // Antarctica/McMurdo
	ZoneIDAntarctica_Palmer uint32 = 0x40962f4f // Antarctica/Palmer
	ZoneIDAntarctica_Rothera uint32 = 0x0e86d203 // Antarctica/Rothera
	ZoneIDAntarctica_South_Pole uint32 = 0xcd96b290 // Antarctica/South_Pole
	ZoneIDAntarctica_Syowa uint32 = 0xe330c7e1 // Antarctica/Syowa
	ZoneIDAntarctica_Troll uint32 = 0xe33f085b // Antarctica/Troll
	ZoneIDAntarctica_Vostok uint32 = 0x4f966fd4 // Antarctica/Vostok
	ZoneIDArctic_Longyearbyen uint32 = 0xd23e7859 // Arctic/Longyearbyen
	ZoneIDAsia_Aden uint32 = 0x1fa7084a // Asia/Aden
	ZoneIDAsia_Almaty uint32 = 0xa61f41fa // Asia/Almaty
	ZoneIDAsia_Amman uint32 = 0x148d21bc // Asia/Amman
	ZoneIDAsia_Anadyr uint32 = 0xa63cebd1 // Asia/Anadyr
	ZoneIDAsia_Aqtau uint32 = 0x148f710e // Asia/Aqtau
	ZoneIDAsia_Aqtobe uint32 = 0xa67dcc4e // Asia/Aqtobe
	ZoneIDAsia_Ashgabat uint32 = 0xba87598d // Asia/Ashgabat
	ZoneIDAsia_Ashkhabad uint32 = 0x15454f09 // Asia/Ashkhabad
	ZoneIDAsia_Atyrau uint32 = 0xa6b6e068 // Asia/Atyrau
	ZoneIDAsia_Baghdad uint32 = 0x9ceffbed // Asia/Baghdad
	ZoneIDAsia_Bahrain uint32 = 0x9d078487 // Asia/Bahrain
	ZoneIDAsia_Baku uint32 = 0x1fa788b5 // Asia/Baku
	ZoneIDAsia_Bangkok uint32 = 0x9d6e3aaf // Asia/Bangkok
	ZoneIDAsia_Barnaul uint32 = 0x9dba4997 // Asia/Barnaul
	ZoneIDAsia_Beirut uint32 = 0xa7f3d5fd // Asia/Beirut
	ZoneIDAsia_Bishkek uint32 = 0xb0728553 // Asia/Bishkek
	ZoneIDAsia_Brunei uint32 = 0xa8e595f7 // Asia/Brunei
	ZoneIDAsia_Calcutta uint32 = 0x328a44c3 // Asia/Calcutta
	ZoneIDAsia_Chita uint32 = 0x14ae863b // Asia/Chita
	ZoneIDAsia_Choibalsan uint32 = 0x928aa4a6 // Asia/Choibalsan
	ZoneIDAsia_Chongqing uint32 = 0xf937fb90 // Asia/Chongqing
	ZoneIDAsia_Chungking uint32 = 0xc7121dd0 // Asia/Chungking
	ZoneIDAsia_Colombo uint32 = 0x0af0e91d // Asia/Colombo
	ZoneIDAsia_Dacca uint32 = 0x14bcac5e // Asia/Dacca
	ZoneIDAsia_Damascus uint32 = 0x20fbb063 // Asia/Damascus
	ZoneIDAsia_Dhaka uint32 = 0x14c07b8b // Asia/Dhaka
	ZoneIDAsia_Dili uint32 = 0x1fa8c394 // Asia/Dili
	ZoneIDAsia_Dubai uint32 = 0x14c79f77 // Asia/Dubai
	ZoneIDAsia_Dushanbe uint32 = 0x32fc5c3c // Asia/Dushanbe
	ZoneIDAsia_Famagusta uint32 = 0x289b4f8b // Asia/Famagusta
	ZoneIDAsia_Gaza uint32 = 0x1faa4875 // Asia/Gaza
	ZoneIDAsia_Harbin uint32 = 0xb5af1186 // Asia/Harbin
	ZoneIDAsia_Hebron uint32 = 0xb5eef250 // Asia/Hebron
	ZoneIDAsia_Ho_Chi_Minh uint32 = 0x20f2d127 // Asia/Ho_Chi_Minh
	ZoneIDAsia_Hong_Kong uint32 = 0x577f28ac // Asia/Hong_Kong
	ZoneIDAsia_Hovd uint32 = 0x1fab0fe3 // Asia/Hovd
	ZoneIDAsia_Irkutsk uint32 = 0xdfbf213f // Asia/Irkutsk
	ZoneIDAsia_Istanbul uint32 = 0x382e7894 // Asia/Istanbul
	ZoneIDAsia_Jakarta uint32 = 0x0506ab50 // Asia/Jakarta
	ZoneIDAsia_Jayapura uint32 = 0xc6833c2f // Asia/Jayapura
	ZoneIDAsia_Jerusalem uint32 = 0x5becd23a // Asia/Jerusalem
	ZoneIDAsia_Kabul uint32 = 0x153b5601 // Asia/Kabul
	ZoneIDAsia_Kamchatka uint32 = 0x73baf9d7 // Asia/Kamchatka
	ZoneIDAsia_Karachi uint32 = 0x527f5245 // Asia/Karachi
	ZoneIDAsia_Kashgar uint32 = 0x52955193 // Asia/Kashgar
	ZoneIDAsia_Kathmandu uint32 = 0x9a96ce6f // Asia/Kathmandu
	ZoneIDAsia_Katmandu uint32 = 0xa7ec12c7 // Asia/Katmandu
	ZoneIDAsia_Khandyga uint32 = 0x9685a4d9 // Asia/Khandyga
	ZoneIDAsia_Kolkata uint32 = 0x72c06cd9 // Asia/Kolkata
	ZoneIDAsia_Krasnoyarsk uint32 = 0xd0376c6a // Asia/Krasnoyarsk
	ZoneIDAsia_Kuala_Lumpur uint32 = 0x014763c4 // Asia/Kuala_Lumpur
	ZoneIDAsia_Kuching uint32 = 0x801b003b // Asia/Kuching
	ZoneIDAsia_Kuwait uint32 = 0xbe1b2f27 // Asia/Kuwait
	ZoneIDAsia_Macao uint32 = 0x155f88b3 // Asia/Macao
	ZoneIDAsia_Macau uint32 = 0x155f88b9 // Asia/Macau
	ZoneIDAsia_Magadan uint32 = 0xebacc19b // Asia/Magadan
	ZoneIDAsia_Makassar uint32 = 0x6aa21c85 // Asia/Makassar
	ZoneIDAsia_Manila uint32 = 0xc156c944 // Asia/Manila
	ZoneIDAsia_Muscat uint32 = 0xc2c3565f // Asia/Muscat
	ZoneIDAsia_Nicosia uint32 = 0x4b0fcf78 // Asia/Nicosia
	ZoneIDAsia_Novokuznetsk uint32 = 0x69264f93 // Asia/Novokuznetsk
	ZoneIDAsia_Novosibirsk uint32 = 0xa2a435cb // Asia/Novosibirsk
	ZoneIDAsia_Omsk uint32 = 0x1faeddac // Asia/Omsk
	ZoneIDAsia_Oral uint32 = 0x1faef0a0 // Asia/Oral
	ZoneIDAsia_Phnom_Penh uint32 = 0xc224945e // Asia/Phnom_Penh
	ZoneIDAsia_Pontianak uint32 = 0x1a76c057 // Asia/Pontianak
	ZoneIDAsia_Pyongyang uint32 = 0x93ed1c8e // Asia/Pyongyang
	ZoneIDAsia_Qatar uint32 = 0x15a8330b // Asia/Qatar
	ZoneIDAsia_Qostanay uint32 = 0x654fe522 // Asia/Qostanay
	ZoneIDAsia_Qyzylorda uint32 = 0x71282e81 // Asia/Qyzylorda
	ZoneIDAsia_Rangoon uint32 = 0x6d1217c6 // Asia/Rangoon
	ZoneIDAsia_Riyadh uint32 = 0xcd973d93 // Asia/Riyadh
	ZoneIDAsia_Saigon uint32 = 0xcf52f713 // Asia/Saigon
	ZoneIDAsia_Sakhalin uint32 = 0xf4a1c9bd // Asia/Sakhalin
	ZoneIDAsia_Samarkand uint32 = 0x13ae5104 // Asia/Samarkand
	ZoneIDAsia_Seoul uint32 = 0x15ce82da // Asia/Seoul
	ZoneIDAsia_Shanghai uint32 = 0xf895a7f5 // Asia/Shanghai
	ZoneIDAsia_Singapore uint32 = 0xcf8581fa // Asia/Singapore
	ZoneIDAsia_Srednekolymsk uint32 = 0xbf8e337d // Asia/Srednekolymsk
	ZoneIDAsia_Taipei uint32 = 0xd1a844ae // Asia/Taipei
	ZoneIDAsia_Tashkent uint32 = 0xf3924254 // Asia/Tashkent
	ZoneIDAsia_Tbilisi uint32 = 0x0903e442 // Asia/Tbilisi
	ZoneIDAsia_Tehran uint32 = 0xd1f02254 // Asia/Tehran
	ZoneIDAsia_Tel_Aviv uint32 = 0x166d7c2c // Asia/Tel_Aviv
	ZoneIDAsia_Thimbu uint32 = 0xd226e31b // Asia/Thimbu
	ZoneIDAsia_Thimphu uint32 = 0x170380d1 // Asia/Thimphu
	ZoneIDAsia_Tokyo uint32 = 0x15e606a8 // Asia/Tokyo
	ZoneIDAsia_Tomsk uint32 = 0x15e60e60 // Asia/Tomsk
	ZoneIDAsia_Ujung_Pandang uint32 = 0x5d001eb3 // Asia/Ujung_Pandang
	ZoneIDAsia_Ulaanbaatar uint32 = 0x30f0cc4e // Asia/Ulaanbaatar
	ZoneIDAsia_Ulan_Bator uint32 = 0x394db4d9 // Asia/Ulan_Bator
	ZoneIDAsia_Urumqi uint32 = 0xd5379735 // Asia/Urumqi
	ZoneIDAsia_Ust_Nera uint32 = 0x4785f921 // Asia/Ust-Nera
	ZoneIDAsia_Vientiane uint32 = 0x89d68d75 // Asia/Vientiane
	ZoneIDAsia_Vladivostok uint32 = 0x29de34a8 // Asia/Vladivostok
	ZoneIDAsia_Yakutsk uint32 = 0x87bb3a9e // Asia/Yakutsk
	ZoneIDAsia_Yangon uint32 = 0xdd54a8be // Asia/Yangon
	ZoneIDAsia_Yekaterinburg uint32 = 0xfb544c6e // Asia/Yekaterinburg
	ZoneIDAsia_Yerevan uint32 = 0x9185c8cc // Asia/Yerevan
	ZoneIDAtlantic_Azores uint32 = 0xf93ed918 // Atlantic/Azores
	ZoneIDAtlantic_Bermuda uint32 = 0x3d4bb1c4 // Atlantic/Bermuda
	ZoneIDAtlantic_Canary uint32 = 0xfc23f2c2 // Atlantic/Canary
	ZoneIDAtlantic_Cape_Verde uint32 = 0x5c5e1772 // Atlantic/Cape_Verde
	ZoneIDAtlantic_Faeroe uint32 = 0x031ec516 // Atlantic/Faeroe
	ZoneIDAtlantic_Faroe uint32 = 0xe110a971 // Atlantic/Faroe
	ZoneIDAtlantic_Jan_Mayen uint32 = 0x5a7535b6 // Atlantic/Jan_Mayen
	ZoneIDAtlantic_Madeira uint32 = 0x81b5c037 // Atlantic/Madeira
	ZoneIDAtlantic_Reykjavik uint32 = 0x1c2b4f74 // Atlantic/Reykjavik
	ZoneIDAtlantic_South_Georgia uint32 = 0x33013174 // Atlantic/South_Georgia
	ZoneIDAtlantic_St_Helena uint32 = 0x451fc5f7 // Atlantic/St_Helena
	ZoneIDAtlantic_Stanley uint32 = 0x7bb3e1c4 // Atlantic/Stanley
	ZoneIDAustralia_ACT uint32 = 0x8a970eb2 // Australia/ACT
	ZoneIDAustralia_Adelaide uint32 = 0x2428e8a3 // Australia/Adelaide
	ZoneIDAustralia_Brisbane uint32 = 0x4fedc9c0 // Australia/Brisbane
	ZoneIDAustralia_Broken_Hill uint32 = 0xb06eada3 // Australia/Broken_Hill
	ZoneIDAustralia_Canberra uint32 = 0x2a09ae58 // Australia/Canberra
	ZoneIDAustralia_Currie uint32 = 0x278b6a24 // Australia/Currie
	ZoneIDAustralia_Darwin uint32 = 0x2876bdff // Australia/Darwin
	ZoneIDAustralia_Eucla uint32 = 0x8cf99e44 // Australia/Eucla
	ZoneIDAustralia_Hobart uint32 = 0x32bf951a // Australia/Hobart
	ZoneIDAustralia_LHI uint32 = 0x8a973e17 // Australia/LHI
	ZoneIDAustralia_Lindeman uint32 = 0xe05029e2 // Australia/Lindeman
	ZoneIDAustralia_Lord_Howe uint32 = 0xa748b67d // Australia/Lord_Howe
	ZoneIDAustralia_Melbourne uint32 = 0x0fe559a3 // Australia/Melbourne
	ZoneIDAustralia_NSW uint32 = 0x8a974812 // Australia/NSW
	ZoneIDAustralia_North uint32 = 0x8d997165 // Australia/North
	ZoneIDAustralia_Perth uint32 = 0x8db8269d // Australia/Perth
	ZoneIDAustralia_Queensland uint32 = 0xd326ed0a // Australia/Queensland
	ZoneIDAustralia_South uint32 = 0x8df3f8ad // Australia/South
	ZoneIDAustralia_Sydney uint32 = 0x4d1e9776 // Australia/Sydney
	ZoneIDAustralia_Tasmania uint32 = 0xe6d76648 // Australia/Tasmania
	ZoneIDAustralia_Victoria uint32 = 0x0260d5db // Australia/Victoria
	ZoneIDAustralia_West uint32 = 0xdd858a5d // Australia/West
	ZoneIDAustralia_Yancowinna uint32 = 0x90bac131 // Australia/Yancowinna
	ZoneIDBrazil_Acre uint32 = 0x66934f93 // Brazil/Acre
	ZoneIDBrazil_DeNoronha uint32 = 0x9b4cb496 // Brazil/DeNoronha
	ZoneIDBrazil_East uint32 = 0x669578c5 // Brazil/East
	ZoneIDBrazil_West uint32 = 0x669f689b // Brazil/West
	ZoneIDCET uint32 = 0x0b87d921 // CET
	ZoneIDCST6CDT uint32 = 0xf0e87d00 // CST6CDT
	ZoneIDCanada_Atlantic uint32 = 0x536b119c // Canada/Atlantic
	ZoneIDCanada_Central uint32 = 0x626710f5 // Canada/Central
	ZoneIDCanada_Eastern uint32 = 0xf3612d5e // Canada/Eastern
	ZoneIDCanada_Mountain uint32 = 0xfb8a8217 // Canada/Mountain
	ZoneIDCanada_Newfoundland uint32 = 0xb396e991 // Canada/Newfoundland
	ZoneIDCanada_Pacific uint32 = 0x40fa3c7b // Canada/Pacific
	ZoneIDCanada_Saskatchewan uint32 = 0x77311f49 // Canada/Saskatchewan
	ZoneIDCanada_Yukon uint32 = 0x78dd35c2 // Canada/Yukon
	ZoneIDChile_Continental uint32 = 0x7e2bdb18 // Chile/Continental
	ZoneIDChile_EasterIsland uint32 = 0xb0982af8 // Chile/EasterIsland
	ZoneIDCuba uint32 = 0x7c83cba0 // Cuba
	ZoneIDEET uint32 = 0x0b87e1a3 // EET
	ZoneIDEST uint32 = 0x0b87e371 // EST
	ZoneIDEST5EDT uint32 = 0x8adc72a3 // EST5EDT
	ZoneIDEgypt uint32 = 0x0d1a278e // Egypt
	ZoneIDEire uint32 = 0x7c84b36a // Eire
	ZoneIDEtc_GMT uint32 = 0xd8e2de58 // Etc/GMT
	ZoneIDEtc_GMT_PLUS_0 uint32 = 0x9d13da13 // Etc/GMT+0
	ZoneIDEtc_GMT_PLUS_1 uint32 = 0x9d13da14 // Etc/GMT+1
	ZoneIDEtc_GMT_PLUS_10 uint32 = 0x3f8f1cc4 // Etc/GMT+10
	ZoneIDEtc_GMT_PLUS_11 uint32 = 0x3f8f1cc5 // Etc/GMT+11
	ZoneIDEtc_GMT_PLUS_12 uint32 = 0x3f8f1cc6 // Etc/GMT+12
	ZoneIDEtc_GMT_PLUS_2 uint32 = 0x9d13da15 // Etc/GMT+2
	ZoneIDEtc_GMT_PLUS_3 uint32 = 0x9d13da16 // Etc/GMT+3
	ZoneIDEtc_GMT_PLUS_4 uint32 = 0x9d13da17 // Etc/GMT+4
	ZoneIDEtc_GMT_PLUS_5 uint32 = 0x9d13da18 // Etc/GMT+5
	ZoneIDEtc_GMT_PLUS_6 uint32 = 0x9d13da19 // Etc/GMT+6
	ZoneIDEtc_GMT_PLUS_7 uint32 = 0x9d13da1a // Etc/GMT+7
	ZoneIDEtc_GMT_PLUS_8 uint32 = 0x9d13da1b // Etc/GMT+8
	ZoneIDEtc_GMT_PLUS_9 uint32 = 0x9d13da1c // Etc/GMT+9
	ZoneIDEtc_GMT_0 uint32 = 0x9d13da55 // Etc/GMT-0
	ZoneIDEtc_GMT_1 uint32 = 0x9d13da56 // Etc/GMT-1
	ZoneIDEtc_GMT_10 uint32 = 0x3f8f2546 // Etc/GMT-10
	ZoneIDEtc_GMT_11 uint32 = 0x3f8f2547 // Etc/GMT-11
	ZoneIDEtc_GMT_12 uint32 = 0x3f8f2548 // Etc/GMT-12
	ZoneIDEtc_GMT_13 uint32 = 0x3f8f2549 // Etc/GMT-13
	ZoneIDEtc_GMT_14 uint32 = 0x3f8f254a // Etc/GMT-14
	ZoneIDEtc_GMT_2 uint32 = 0x9d13da57 // Etc/GMT-2
	ZoneIDEtc_GMT_3 uint32 = 0x9d13da58 // Etc/GMT-3
	ZoneIDEtc_GMT_4 uint32 = 0x9d13da59 // Etc/GMT-4
	ZoneIDEtc_GMT_5 uint32 = 0x9d13da5a // Etc/GMT-5
	ZoneIDEtc_GMT_6 uint32 = 0x9d13da5b // Etc/GMT-6
	ZoneIDEtc_GMT_7 uint32 = 0x9d13da5c // Etc/GMT-7
	ZoneIDEtc_GMT_8 uint32 = 0x9d13da5d // Etc/GMT-8
	ZoneIDEtc_GMT_9 uint32 = 0x9d13da5e // Etc/GMT-9
	ZoneIDEtc_GMT0 uint32 = 0xf53ea988 // Etc/GMT0
	ZoneIDEtc_Greenwich uint32 = 0x26daa98c // Etc/Greenwich
	ZoneIDEtc_UCT uint32 = 0xd8e3189c // Etc/UCT
	ZoneIDEtc_UTC uint32 = 0xd8e31abc // Etc/UTC
	ZoneIDEtc_Universal uint32 = 0x2f8cb9a9 // Etc/Universal
	ZoneIDEtc_Zulu uint32 = 0xf549c240 // Etc/Zulu
	ZoneIDEurope_Amsterdam uint32 = 0x109395c2 // Europe/Amsterdam
	ZoneIDEurope_Andorra uint32 = 0x97f6764b // Europe/Andorra
	ZoneIDEurope_Astrakhan uint32 = 0xe22256e1 // Europe/Astrakhan
	ZoneIDEurope_Athens uint32 = 0x4318fa27 // Europe/Athens
	ZoneIDEurope_Belfast uint32 = 0xd07dd1e5 // Europe/Belfast
	ZoneIDEurope_Belgrade uint32 = 0xe0532b3a // Europe/Belgrade
	ZoneIDEurope_Berlin uint32 = 0x44644c20 // Europe/Berlin
	ZoneIDEurope_Bratislava uint32 = 0xda493bed // Europe/Bratislava
	ZoneIDEurope_Brussels uint32 = 0xdee07337 // Europe/Brussels
	ZoneIDEurope_Bucharest uint32 = 0xfb349ec5 // Europe/Bucharest
	ZoneIDEurope_Budapest uint32 = 0x9ce0197c // Europe/Budapest
	ZoneIDEurope_Busingen uint32 = 0xc06d2cdf // Europe/Busingen
	ZoneIDEurope_Chisinau uint32 = 0xad58aa18 // Europe/Chisinau
	ZoneIDEurope_Copenhagen uint32 = 0xe0ed30bc // Europe/Copenhagen
	ZoneIDEurope_Dublin uint32 = 0x4a275f62 // Europe/Dublin
	ZoneIDEurope_Gibraltar uint32 = 0xf8e325fc // Europe/Gibraltar
	ZoneIDEurope_Guernsey uint32 = 0x3db12c16 // Europe/Guernsey
	ZoneIDEurope_Helsinki uint32 = 0x6ab2975b // Europe/Helsinki
	ZoneIDEurope_Isle_of_Man uint32 = 0xeaf84580 // Europe/Isle_of_Man
	ZoneIDEurope_Istanbul uint32 = 0x9e09d6e6 // Europe/Istanbul
	ZoneIDEurope_Jersey uint32 = 0x570dae76 // Europe/Jersey
	ZoneIDEurope_Kaliningrad uint32 = 0xd33b2f28 // Europe/Kaliningrad
	ZoneIDEurope_Kiev uint32 = 0xa2c19eb3 // Europe/Kiev
	ZoneIDEurope_Kirov uint32 = 0xfaf5abef // Europe/Kirov
	ZoneIDEurope_Kyiv uint32 = 0xa2c1e347 // Europe/Kyiv
	ZoneIDEurope_Lisbon uint32 = 0x5c00a70b // Europe/Lisbon
	ZoneIDEurope_Ljubljana uint32 = 0xbd98cdb7 // Europe/Ljubljana
	ZoneIDEurope_London uint32 = 0x5c6a84ae // Europe/London
	ZoneIDEurope_Luxembourg uint32 = 0x1f8bc6ce // Europe/Luxembourg
	ZoneIDEurope_Madrid uint32 = 0x5dbd1535 // Europe/Madrid
	ZoneIDEurope_Malta uint32 = 0xfb1560f3 // Europe/Malta
	ZoneIDEurope_Mariehamn uint32 = 0x0caa6496 // Europe/Mariehamn
	ZoneIDEurope_Minsk uint32 = 0xfb19cc66 // Europe/Minsk
	ZoneIDEurope_Monaco uint32 = 0x5ebf9f01 // Europe/Monaco
	ZoneIDEurope_Moscow uint32 = 0x5ec266fc // Europe/Moscow
	ZoneIDEurope_Nicosia uint32 = 0x74efab8a // Europe/Nicosia
	ZoneIDEurope_Oslo uint32 = 0xa2c3fba1 // Europe/Oslo
	ZoneIDEurope_Paris uint32 = 0xfb4bc2a3 // Europe/Paris
	ZoneIDEurope_Podgorica uint32 = 0x1c1a499c // Europe/Podgorica
	ZoneIDEurope_Prague uint32 = 0x65ee5d48 // Europe/Prague
	ZoneIDEurope_Riga uint32 = 0xa2c57587 // Europe/Riga
	ZoneIDEurope_Rome uint32 = 0xa2c58fd7 // Europe/Rome
	ZoneIDEurope_Samara uint32 = 0x6bc0b139 // Europe/Samara
	ZoneIDEurope_San_Marino uint32 = 0xcef7724b // Europe/San_Marino
	ZoneIDEurope_Sarajevo uint32 = 0x6a576c3f // Europe/Sarajevo
	ZoneIDEurope_Saratov uint32 = 0xe4315da4 // Europe/Saratov
	ZoneIDEurope_Simferopol uint32 = 0xda9eb724 // Europe/Simferopol
	ZoneIDEurope_Skopje uint32 = 0x6c76fdd0 // Europe/Skopje
	ZoneIDEurope_Sofia uint32 = 0xfb898656 // Europe/Sofia
	ZoneIDEurope_Stockholm uint32 = 0x5bf6fbb8 // Europe/Stockholm
	ZoneIDEurope_Tallinn uint32 = 0x30c4e096 // Europe/Tallinn
	ZoneIDEurope_Tirane uint32 = 0x6ea95b47 // Europe/Tirane
	ZoneIDEurope_Tiraspol uint32 = 0xbe704472 // Europe/Tiraspol
	ZoneIDEurope_Ulyanovsk uint32 = 0xe03783d0 // Europe/Ulyanovsk
	ZoneIDEurope_Uzhgorod uint32 = 0xb066f5d6 // Europe/Uzhgorod
	ZoneIDEurope_Vaduz uint32 = 0xfbb81bae // Europe/Vaduz
	ZoneIDEurope_Vatican uint32 = 0xcb485dca // Europe/Vatican
	ZoneIDEurope_Vienna uint32 = 0x734cc2e5 // Europe/Vienna
	ZoneIDEurope_Vilnius uint32 = 0xdd63b8ce // Europe/Vilnius
	ZoneIDEurope_Volgograd uint32 = 0x3ed0f389 // Europe/Volgograd
	ZoneIDEurope_Warsaw uint32 = 0x75185c19 // Europe/Warsaw
	ZoneIDEurope_Zagreb uint32 = 0x7c11c9ff // Europe/Zagreb
	ZoneIDEurope_Zaporozhye uint32 = 0xeab9767f // Europe/Zaporozhye
	ZoneIDEurope_Zurich uint32 = 0x7d8195b9 // Europe/Zurich
	ZoneIDGB uint32 = 0x005973ae // GB
	ZoneIDGB_Eire uint32 = 0xfa70e300 // GB-Eire
	ZoneIDGMT uint32 = 0x0b87eb2d // GMT
	ZoneIDGMT_PLUS_0 uint32 = 0x0d2f7028 // GMT+0
	ZoneIDGMT_0 uint32 = 0x0d2f706a // GMT-0
	ZoneIDGMT0 uint32 = 0x7c8550fd // GMT0
	ZoneIDGreenwich uint32 = 0xc84d4221 // Greenwich
	ZoneIDHST uint32 = 0x0b87f034 // HST
	ZoneIDHongkong uint32 = 0x56d36560 // Hongkong
	ZoneIDIceland uint32 = 0xe56a35b5 // Iceland
	ZoneIDIndian_Antananarivo uint32 = 0x9ebf5289 // Indian/Antananarivo
	ZoneIDIndian_Chagos uint32 = 0x456f7c3c // Indian/Chagos
	ZoneIDIndian_Christmas uint32 = 0x68c207d5 // Indian/Christmas
	ZoneIDIndian_Cocos uint32 = 0x021e86de // Indian/Cocos
	ZoneIDIndian_Comoro uint32 = 0x45f4deb6 // Indian/Comoro
	ZoneIDIndian_Kerguelen uint32 = 0x4351b389 // Indian/Kerguelen
	ZoneIDIndian_Mahe uint32 = 0x45e725e2 // Indian/Mahe
	ZoneIDIndian_Maldives uint32 = 0x9869681c // Indian/Maldives
	ZoneIDIndian_Mauritius uint32 = 0x7b09c02a // Indian/Mauritius
	ZoneIDIndian_Mayotte uint32 = 0xe6880bca // Indian/Mayotte
	ZoneIDIndian_Reunion uint32 = 0x7076c047 // Indian/Reunion
	ZoneIDIran uint32 = 0x7c87090f // Iran
	ZoneIDIsrael uint32 = 0xba88c9e5 // Israel
	ZoneIDJamaica uint32 = 0x2e44fdab // Jamaica
	ZoneIDJapan uint32 = 0x0d712f8f // Japan
	ZoneIDKwajalein uint32 = 0x0e57afbb // Kwajalein
	ZoneIDLibya uint32 = 0x0d998b16 // Libya
	ZoneIDMET uint32 = 0x0b8803ab // MET
	ZoneIDMST uint32 = 0x0b880579 // MST
	ZoneIDMST7MDT uint32 = 0xf2af9375 // MST7MDT
	ZoneIDMexico_BajaNorte uint32 = 0xfcf7150f // Mexico/BajaNorte
	ZoneIDMexico_BajaSur uint32 = 0x08ee3641 // Mexico/BajaSur
	ZoneIDMexico_General uint32 = 0x93711d57 // Mexico/General
	ZoneIDNZ uint32 = 0x005974ad // NZ
	ZoneIDNZ_CHAT uint32 = 0x4d42afda // NZ-CHAT
	ZoneIDNavajo uint32 = 0xc4ef0e24 // Navajo
	ZoneIDPRC uint32 = 0x0b88120a // PRC
	ZoneIDPST8PDT uint32 = 0xd99ee2dc // PST8PDT
	ZoneIDPacific_Apia uint32 = 0x23359b5e // Pacific/Apia
	ZoneIDPacific_Auckland uint32 = 0x25062f86 // Pacific/Auckland
	ZoneIDPacific_Bougainville uint32 = 0x5e10f7a4 // Pacific/Bougainville
	ZoneIDPacific_Chatham uint32 = 0x2f0de999 // Pacific/Chatham
	ZoneIDPacific_Chuuk uint32 = 0x8a090b23 // Pacific/Chuuk
	ZoneIDPacific_Easter uint32 = 0xcf54f7e7 // Pacific/Easter
	ZoneIDPacific_Efate uint32 = 0x8a2bce28 // Pacific/Efate
	ZoneIDPacific_Enderbury uint32 = 0x61599a93 // Pacific/Enderbury
	ZoneIDPacific_Fakaofo uint32 = 0x06532bba // Pacific/Fakaofo
	ZoneIDPacific_Fiji uint32 = 0x23383ba5 // Pacific/Fiji
	ZoneIDPacific_Funafuti uint32 = 0xdb402d65 // Pacific/Funafuti
	ZoneIDPacific_Galapagos uint32 = 0xa952f752 // Pacific/Galapagos
	ZoneIDPacific_Gambier uint32 = 0x53720c3a // Pacific/Gambier
	ZoneIDPacific_Guadalcanal uint32 = 0xf4dd25f0 // Pacific/Guadalcanal
	ZoneIDPacific_Guam uint32 = 0x2338f9ed // Pacific/Guam
	ZoneIDPacific_Honolulu uint32 = 0xe6e70af9 // Pacific/Honolulu
	ZoneIDPacific_Johnston uint32 = 0xb15d7b36 // Pacific/Johnston
	ZoneIDPacific_Kanton uint32 = 0xdd512f0e // Pacific/Kanton
	ZoneIDPacific_Kiritimati uint32 = 0x8305073a // Pacific/Kiritimati
	ZoneIDPacific_Kosrae uint32 = 0xde5139a8 // Pacific/Kosrae
	ZoneIDPacific_Kwajalein uint32 = 0x8e216759 // Pacific/Kwajalein
	ZoneIDPacific_Majuro uint32 = 0xe1f95371 // Pacific/Majuro
	ZoneIDPacific_Marquesas uint32 = 0x57ca7135 // Pacific/Marquesas
	ZoneIDPacific_Midway uint32 = 0xe286d38e // Pacific/Midway
	ZoneIDPacific_Nauru uint32 = 0x8acc41ae // Pacific/Nauru
	ZoneIDPacific_Niue uint32 = 0x233ca014 // Pacific/Niue
	ZoneIDPacific_Norfolk uint32 = 0x8f4eb4be // Pacific/Norfolk
	ZoneIDPacific_Noumea uint32 = 0xe551b788 // Pacific/Noumea
	ZoneIDPacific_Pago_Pago uint32 = 0x603aebd0 // Pacific/Pago_Pago
	ZoneIDPacific_Palau uint32 = 0x8af04a36 // Pacific/Palau
	ZoneIDPacific_Pitcairn uint32 = 0x8837d8bd // Pacific/Pitcairn
	ZoneIDPacific_Pohnpei uint32 = 0x28929f96 // Pacific/Pohnpei
	ZoneIDPacific_Ponape uint32 = 0xe9f80086 // Pacific/Ponape
	ZoneIDPacific_Port_Moresby uint32 = 0xa7ba7f68 // Pacific/Port_Moresby
	ZoneIDPacific_Rarotonga uint32 = 0x9981a3b0 // Pacific/Rarotonga
	ZoneIDPacific_Saipan uint32 = 0xeff7a35f // Pacific/Saipan
	ZoneIDPacific_Samoa uint32 = 0x8b2699b4 // Pacific/Samoa
	ZoneIDPacific_Tahiti uint32 = 0xf24c2446 // Pacific/Tahiti
	ZoneIDPacific_Tarawa uint32 = 0xf2517e63 // Pacific/Tarawa
	ZoneIDPacific_Tongatapu uint32 = 0x262ca836 // Pacific/Tongatapu
	ZoneIDPacific_Truk uint32 = 0x234010a9 // Pacific/Truk
	ZoneIDPacific_Wake uint32 = 0x23416c2b // Pacific/Wake
	ZoneIDPacific_Wallis uint32 = 0xf94ddb0f // Pacific/Wallis
	ZoneIDPacific_Yap uint32 = 0xbb40138d // Pacific/Yap
	ZoneIDPoland uint32 = 0xca913b23 // Poland
	ZoneIDPortugal uint32 = 0xc3274593 // Portugal
	ZoneIDROC uint32 = 0x0b881a29 // ROC
	ZoneIDROK uint32 = 0x0b881a31 // ROK
	ZoneIDSingapore uint32 = 0xa8598c8d // Singapore
	ZoneIDTurkey uint32 = 0xd455e469 // Turkey
	ZoneIDUCT uint32 = 0x0b882571 // UCT
	ZoneIDUS_Alaska uint32 = 0xfa300bc9 // US/Alaska
	ZoneIDUS_Aleutian uint32 = 0x4fe013ef // US/Aleutian
	ZoneIDUS_Arizona uint32 = 0x4ec52670 // US/Arizona
	ZoneIDUS_Central uint32 = 0xcabdcb25 // US/Central
	ZoneIDUS_East_Indiana uint32 = 0x6dcf558a // US/East-Indiana
	ZoneIDUS_Eastern uint32 = 0x5bb7e78e // US/Eastern
	ZoneIDUS_Hawaii uint32 = 0x09c8de2f // US/Hawaii
	ZoneIDUS_Indiana_Starke uint32 = 0x67977be7 // US/Indiana-Starke
	ZoneIDUS_Michigan uint32 = 0x766bb7bc // US/Michigan
	ZoneIDUS_Mountain uint32 = 0x6eb88247 // US/Mountain
	ZoneIDUS_Pacific uint32 = 0xa950f6ab // US/Pacific
	ZoneIDUS_Samoa uint32 = 0x566821cd // US/Samoa
	ZoneIDUTC uint32 = 0x0b882791 // UTC
	ZoneIDUniversal uint32 = 0xd0ff523e // Universal
	ZoneIDW_SU uint32 = 0x7c8d8ef1 // W-SU
	ZoneIDWET uint32 = 0x0b882e35 // WET
	ZoneIDZulu uint32 = 0x7c9069b5 // Zulu

)
