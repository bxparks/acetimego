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
// from https://github.com/eggert/tz/releases/tag/2023c
//
// Supported Zones: 596 (350 zones, 246 links)
// Unsupported Zones: 0 (0 zones, 0 links)
//
// Original Years:  [1844,2087]
// Generated Years: [1950,2087]
// Lower/Upper Truncated: [True, False]
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
	StartYear: 2000,
	UntilYear: 10000,
	MaxTransitions: 7,
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
// Total: 596 (350 zones, 246 links)
// ---------------------------------------------------------------------------

const (
	ZoneInfoIndexAfrica_Abidjan uint16 = 468 // Africa/Abidjan
	ZoneInfoIndexAfrica_Accra uint16 = 269 // Africa/Accra
	ZoneInfoIndexAfrica_Addis_Ababa uint16 = 16 // Africa/Addis_Ababa
	ZoneInfoIndexAfrica_Algiers uint16 = 517 // Africa/Algiers
	ZoneInfoIndexAfrica_Asmara uint16 = 254 // Africa/Asmara
	ZoneInfoIndexAfrica_Asmera uint16 = 255 // Africa/Asmera
	ZoneInfoIndexAfrica_Bamako uint16 = 259 // Africa/Bamako
	ZoneInfoIndexAfrica_Bangui uint16 = 260 // Africa/Bangui
	ZoneInfoIndexAfrica_Banjul uint16 = 261 // Africa/Banjul
	ZoneInfoIndexAfrica_Bissau uint16 = 265 // Africa/Bissau
	ZoneInfoIndexAfrica_Blantyre uint16 = 536 // Africa/Blantyre
	ZoneInfoIndexAfrica_Brazzaville uint16 = 131 // Africa/Brazzaville
	ZoneInfoIndexAfrica_Bujumbura uint16 = 13 // Africa/Bujumbura
	ZoneInfoIndexAfrica_Cairo uint16 = 270 // Africa/Cairo
	ZoneInfoIndexAfrica_Casablanca uint16 = 475 // Africa/Casablanca
	ZoneInfoIndexAfrica_Ceuta uint16 = 271 // Africa/Ceuta
	ZoneInfoIndexAfrica_Conakry uint16 = 280 // Africa/Conakry
	ZoneInfoIndexAfrica_Dakar uint16 = 272 // Africa/Dakar
	ZoneInfoIndexAfrica_Dar_es_Salaam uint16 = 403 // Africa/Dar_es_Salaam
	ZoneInfoIndexAfrica_Djibouti uint16 = 119 // Africa/Djibouti
	ZoneInfoIndexAfrica_Douala uint16 = 279 // Africa/Douala
	ZoneInfoIndexAfrica_El_Aaiun uint16 = 395 // Africa/El_Aaiun
	ZoneInfoIndexAfrica_Freetown uint16 = 225 // Africa/Freetown
	ZoneInfoIndexAfrica_Gaborone uint16 = 121 // Africa/Gaborone
	ZoneInfoIndexAfrica_Harare uint16 = 297 // Africa/Harare
	ZoneInfoIndexAfrica_Johannesburg uint16 = 511 // Africa/Johannesburg
	ZoneInfoIndexAfrica_Juba uint16 = 507 // Africa/Juba
	ZoneInfoIndexAfrica_Kampala uint16 = 466 // Africa/Kampala
	ZoneInfoIndexAfrica_Khartoum uint16 = 586 // Africa/Khartoum
	ZoneInfoIndexAfrica_Kigali uint16 = 308 // Africa/Kigali
	ZoneInfoIndexAfrica_Kinshasa uint16 = 219 // Africa/Kinshasa
	ZoneInfoIndexAfrica_Lagos uint16 = 273 // Africa/Lagos
	ZoneInfoIndexAfrica_Libreville uint16 = 3 // Africa/Libreville
	ZoneInfoIndexAfrica_Lome uint16 = 508 // Africa/Lome
	ZoneInfoIndexAfrica_Luanda uint16 = 319 // Africa/Luanda
	ZoneInfoIndexAfrica_Lubumbashi uint16 = 248 // Africa/Lubumbashi
	ZoneInfoIndexAfrica_Lusaka uint16 = 320 // Africa/Lusaka
	ZoneInfoIndexAfrica_Malabo uint16 = 326 // Africa/Malabo
	ZoneInfoIndexAfrica_Maputo uint16 = 327 // Africa/Maputo
	ZoneInfoIndexAfrica_Maseru uint16 = 328 // Africa/Maseru
	ZoneInfoIndexAfrica_Mbabane uint16 = 200 // Africa/Mbabane
	ZoneInfoIndexAfrica_Mogadishu uint16 = 221 // Africa/Mogadishu
	ZoneInfoIndexAfrica_Monrovia uint16 = 40 // Africa/Monrovia
	ZoneInfoIndexAfrica_Nairobi uint16 = 423 // Africa/Nairobi
	ZoneInfoIndexAfrica_Ndjamena uint16 = 401 // Africa/Ndjamena
	ZoneInfoIndexAfrica_Niamey uint16 = 332 // Africa/Niamey
	ZoneInfoIndexAfrica_Nouakchott uint16 = 139 // Africa/Nouakchott
	ZoneInfoIndexAfrica_Ouagadougou uint16 = 9 // Africa/Ouagadougou
	ZoneInfoIndexAfrica_Porto_Novo uint16 = 134 // Africa/Porto-Novo
	ZoneInfoIndexAfrica_Sao_Tome uint16 = 210 // Africa/Sao_Tome
	ZoneInfoIndexAfrica_Timbuktu uint16 = 442 // Africa/Timbuktu
	ZoneInfoIndexAfrica_Tripoli uint16 = 398 // Africa/Tripoli
	ZoneInfoIndexAfrica_Tunis uint16 = 278 // Africa/Tunis
	ZoneInfoIndexAfrica_Windhoek uint16 = 274 // Africa/Windhoek
	ZoneInfoIndexAmerica_Adak uint16 = 355 // America/Adak
	ZoneInfoIndexAmerica_Anchorage uint16 = 191 // America/Anchorage
	ZoneInfoIndexAmerica_Anguilla uint16 = 436 // America/Anguilla
	ZoneInfoIndexAmerica_Antigua uint16 = 463 // America/Antigua
	ZoneInfoIndexAmerica_Araguaina uint16 = 247 // America/Araguaina
	ZoneInfoIndexAmerica_Argentina_Buenos_Aires uint16 = 505 // America/Argentina/Buenos_Aires
	ZoneInfoIndexAmerica_Argentina_Catamarca uint16 = 318 // America/Argentina/Catamarca
	ZoneInfoIndexAmerica_Argentina_ComodRivadavia uint16 = 90 // America/Argentina/ComodRivadavia
	ZoneInfoIndexAmerica_Argentina_Cordoba uint16 = 461 // America/Argentina/Cordoba
	ZoneInfoIndexAmerica_Argentina_Jujuy uint16 = 206 // America/Argentina/Jujuy
	ZoneInfoIndexAmerica_Argentina_La_Rioja uint16 = 412 // America/Argentina/La_Rioja
	ZoneInfoIndexAmerica_Argentina_Mendoza uint16 = 427 // America/Argentina/Mendoza
	ZoneInfoIndexAmerica_Argentina_Rio_Gallegos uint16 = 476 // America/Argentina/Rio_Gallegos
	ZoneInfoIndexAmerica_Argentina_Salta uint16 = 207 // America/Argentina/Salta
	ZoneInfoIndexAmerica_Argentina_San_Juan uint16 = 141 // America/Argentina/San_Juan
	ZoneInfoIndexAmerica_Argentina_San_Luis uint16 = 142 // America/Argentina/San_Luis
	ZoneInfoIndexAmerica_Argentina_Tucuman uint16 = 554 // America/Argentina/Tucuman
	ZoneInfoIndexAmerica_Argentina_Ushuaia uint16 = 122 // America/Argentina/Ushuaia
	ZoneInfoIndexAmerica_Aruba uint16 = 349 // America/Aruba
	ZoneInfoIndexAmerica_Asuncion uint16 = 175 // America/Asuncion
	ZoneInfoIndexAmerica_Atikokan uint16 = 296 // America/Atikokan
	ZoneInfoIndexAmerica_Atka uint16 = 356 // America/Atka
	ZoneInfoIndexAmerica_Bahia uint16 = 351 // America/Bahia
	ZoneInfoIndexAmerica_Bahia_Banderas uint16 = 61 // America/Bahia_Banderas
	ZoneInfoIndexAmerica_Barbados uint16 = 485 // America/Barbados
	ZoneInfoIndexAmerica_Belem uint16 = 352 // America/Belem
	ZoneInfoIndexAmerica_Belize uint16 = 337 // America/Belize
	ZoneInfoIndexAmerica_Blanc_Sablon uint16 = 243 // America/Blanc-Sablon
	ZoneInfoIndexAmerica_Boa_Vista uint16 = 23 // America/Boa_Vista
	ZoneInfoIndexAmerica_Bogota uint16 = 339 // America/Bogota
	ZoneInfoIndexAmerica_Boise uint16 = 353 // America/Boise
	ZoneInfoIndexAmerica_Buenos_Aires uint16 = 224 // America/Buenos_Aires
	ZoneInfoIndexAmerica_Cambridge_Bay uint16 = 510 // America/Cambridge_Bay
	ZoneInfoIndexAmerica_Campo_Grande uint16 = 595 // America/Campo_Grande
	ZoneInfoIndexAmerica_Cancun uint16 = 342 // America/Cancun
	ZoneInfoIndexAmerica_Caracas uint16 = 132 // America/Caracas
	ZoneInfoIndexAmerica_Catamarca uint16 = 174 // America/Catamarca
	ZoneInfoIndexAmerica_Cayenne uint16 = 133 // America/Cayenne
	ZoneInfoIndexAmerica_Cayman uint16 = 343 // America/Cayman
	ZoneInfoIndexAmerica_Chicago uint16 = 167 // America/Chicago
	ZoneInfoIndexAmerica_Chihuahua uint16 = 303 // America/Chihuahua
	ZoneInfoIndexAmerica_Ciudad_Juarez uint16 = 252 // America/Ciudad_Juarez
	ZoneInfoIndexAmerica_Coral_Harbour uint16 = 430 // America/Coral_Harbour
	ZoneInfoIndexAmerica_Cordoba uint16 = 198 // America/Cordoba
	ZoneInfoIndexAmerica_Costa_Rica uint16 = 213 // America/Costa_Rica
	ZoneInfoIndexAmerica_Creston uint16 = 212 // America/Creston
	ZoneInfoIndexAmerica_Cuiaba uint16 = 347 // America/Cuiaba
	ZoneInfoIndexAmerica_Curacao uint16 = 230 // America/Curacao
	ZoneInfoIndexAmerica_Danmarkshavn uint16 = 571 // America/Danmarkshavn
	ZoneInfoIndexAmerica_Dawson uint16 = 348 // America/Dawson
	ZoneInfoIndexAmerica_Dawson_Creek uint16 = 239 // America/Dawson_Creek
	ZoneInfoIndexAmerica_Denver uint16 = 350 // America/Denver
	ZoneInfoIndexAmerica_Detroit uint16 = 334 // America/Detroit
	ZoneInfoIndexAmerica_Dominica uint16 = 490 // America/Dominica
	ZoneInfoIndexAmerica_Edmonton uint16 = 238 // America/Edmonton
	ZoneInfoIndexAmerica_Eirunepe uint16 = 578 // America/Eirunepe
	ZoneInfoIndexAmerica_El_Salvador uint16 = 264 // America/El_Salvador
	ZoneInfoIndexAmerica_Ensenada uint16 = 283 // America/Ensenada
	ZoneInfoIndexAmerica_Fort_Nelson uint16 = 144 // America/Fort_Nelson
	ZoneInfoIndexAmerica_Fort_Wayne uint16 = 293 // America/Fort_Wayne
	ZoneInfoIndexAmerica_Fortaleza uint16 = 110 // America/Fortaleza
	ZoneInfoIndexAmerica_Glace_Bay uint16 = 344 // America/Glace_Bay
	ZoneInfoIndexAmerica_Godthab uint16 = 330 // America/Godthab
	ZoneInfoIndexAmerica_Goose_Bay uint16 = 449 // America/Goose_Bay
	ZoneInfoIndexAmerica_Grand_Turk uint16 = 242 // America/Grand_Turk
	ZoneInfoIndexAmerica_Grenada uint16 = 346 // America/Grenada
	ZoneInfoIndexAmerica_Guadeloupe uint16 = 487 // America/Guadeloupe
	ZoneInfoIndexAmerica_Guatemala uint16 = 38 // America/Guatemala
	ZoneInfoIndexAmerica_Guayaquil uint16 = 72 // America/Guayaquil
	ZoneInfoIndexAmerica_Guyana uint16 = 402 // America/Guyana
	ZoneInfoIndexAmerica_Halifax uint16 = 456 // America/Halifax
	ZoneInfoIndexAmerica_Havana uint16 = 404 // America/Havana
	ZoneInfoIndexAmerica_Hermosillo uint16 = 18 // America/Hermosillo
	ZoneInfoIndexAmerica_Indiana_Indianapolis uint16 = 107 // America/Indiana/Indianapolis
	ZoneInfoIndexAmerica_Indiana_Knox uint16 = 215 // America/Indiana/Knox
	ZoneInfoIndexAmerica_Indiana_Marengo uint16 = 117 // America/Indiana/Marengo
	ZoneInfoIndexAmerica_Indiana_Petersburg uint16 = 341 // America/Indiana/Petersburg
	ZoneInfoIndexAmerica_Indiana_Tell_City uint16 = 21 // America/Indiana/Tell_City
	ZoneInfoIndexAmerica_Indiana_Vevay uint16 = 52 // America/Indiana/Vevay
	ZoneInfoIndexAmerica_Indiana_Vincennes uint16 = 106 // America/Indiana/Vincennes
	ZoneInfoIndexAmerica_Indiana_Winamac uint16 = 157 // America/Indiana/Winamac
	ZoneInfoIndexAmerica_Indianapolis uint16 = 228 // America/Indianapolis
	ZoneInfoIndexAmerica_Inuvik uint16 = 411 // America/Inuvik
	ZoneInfoIndexAmerica_Iqaluit uint16 = 113 // America/Iqaluit
	ZoneInfoIndexAmerica_Jamaica uint16 = 182 // America/Jamaica
	ZoneInfoIndexAmerica_Jujuy uint16 = 361 // America/Jujuy
	ZoneInfoIndexAmerica_Juneau uint16 = 417 // America/Juneau
	ZoneInfoIndexAmerica_Kentucky_Louisville uint16 = 74 // America/Kentucky/Louisville
	ZoneInfoIndexAmerica_Kentucky_Monticello uint16 = 528 // America/Kentucky/Monticello
	ZoneInfoIndexAmerica_Knox_IN uint16 = 467 // America/Knox_IN
	ZoneInfoIndexAmerica_Kralendijk uint16 = 553 // America/Kralendijk
	ZoneInfoIndexAmerica_La_Paz uint16 = 428 // America/La_Paz
	ZoneInfoIndexAmerica_Lima uint16 = 357 // America/Lima
	ZoneInfoIndexAmerica_Los_Angeles uint16 = 450 // America/Los_Angeles
	ZoneInfoIndexAmerica_Louisville uint16 = 140 // America/Louisville
	ZoneInfoIndexAmerica_Lower_Princes uint16 = 234 // America/Lower_Princes
	ZoneInfoIndexAmerica_Maceio uint16 = 431 // America/Maceio
	ZoneInfoIndexAmerica_Managua uint16 = 136 // America/Managua
	ZoneInfoIndexAmerica_Manaus uint16 = 432 // America/Manaus
	ZoneInfoIndexAmerica_Marigot uint16 = 137 // America/Marigot
	ZoneInfoIndexAmerica_Martinique uint16 = 181 // America/Martinique
	ZoneInfoIndexAmerica_Matamoros uint16 = 522 // America/Matamoros
	ZoneInfoIndexAmerica_Mazatlan uint16 = 14 // America/Mazatlan
	ZoneInfoIndexAmerica_Mendoza uint16 = 163 // America/Mendoza
	ZoneInfoIndexAmerica_Menominee uint16 = 537 // America/Menominee
	ZoneInfoIndexAmerica_Merida uint16 = 433 // America/Merida
	ZoneInfoIndexAmerica_Metlakatla uint16 = 301 // America/Metlakatla
	ZoneInfoIndexAmerica_Mexico_City uint16 = 497 // America/Mexico_City
	ZoneInfoIndexAmerica_Miquelon uint16 = 188 // America/Miquelon
	ZoneInfoIndexAmerica_Moncton uint16 = 202 // America/Moncton
	ZoneInfoIndexAmerica_Monterrey uint16 = 100 // America/Monterrey
	ZoneInfoIndexAmerica_Montevideo uint16 = 579 // America/Montevideo
	ZoneInfoIndexAmerica_Montreal uint16 = 87 // America/Montreal
	ZoneInfoIndexAmerica_Montserrat uint16 = 73 // America/Montserrat
	ZoneInfoIndexAmerica_Nassau uint16 = 435 // America/Nassau
	ZoneInfoIndexAmerica_New_York uint16 = 78 // America/New_York
	ZoneInfoIndexAmerica_Nipigon uint16 = 392 // America/Nipigon
	ZoneInfoIndexAmerica_Nome uint16 = 358 // America/Nome
	ZoneInfoIndexAmerica_Noronha uint16 = 429 // America/Noronha
	ZoneInfoIndexAmerica_North_Dakota_Beulah uint16 = 367 // America/North_Dakota/Beulah
	ZoneInfoIndexAmerica_North_Dakota_Center uint16 = 396 // America/North_Dakota/Center
	ZoneInfoIndexAmerica_North_Dakota_New_Salem uint16 = 11 // America/North_Dakota/New_Salem
	ZoneInfoIndexAmerica_Nuuk uint16 = 359 // America/Nuuk
	ZoneInfoIndexAmerica_Ojinaga uint16 = 559 // America/Ojinaga
	ZoneInfoIndexAmerica_Panama uint16 = 445 // America/Panama
	ZoneInfoIndexAmerica_Pangnirtung uint16 = 112 // America/Pangnirtung
	ZoneInfoIndexAmerica_Paramaribo uint16 = 444 // America/Paramaribo
	ZoneInfoIndexAmerica_Phoenix uint16 = 127 // America/Phoenix
	ZoneInfoIndexAmerica_Port_au_Prince uint16 = 325 // America/Port-au-Prince
	ZoneInfoIndexAmerica_Port_of_Spain uint16 = 512 // America/Port_of_Spain
	ZoneInfoIndexAmerica_Porto_Acre uint16 = 486 // America/Porto_Acre
	ZoneInfoIndexAmerica_Porto_Velho uint16 = 235 // America/Porto_Velho
	ZoneInfoIndexAmerica_Puerto_Rico uint16 = 222 // America/Puerto_Rico
	ZoneInfoIndexAmerica_Punta_Arenas uint16 = 472 // America/Punta_Arenas
	ZoneInfoIndexAmerica_Rainy_River uint16 = 368 // America/Rainy_River
	ZoneInfoIndexAmerica_Rankin_Inlet uint16 = 481 // America/Rankin_Inlet
	ZoneInfoIndexAmerica_Recife uint16 = 451 // America/Recife
	ZoneInfoIndexAmerica_Regina uint16 = 452 // America/Regina
	ZoneInfoIndexAmerica_Resolute uint16 = 478 // America/Resolute
	ZoneInfoIndexAmerica_Rio_Branco uint16 = 393 // America/Rio_Branco
	ZoneInfoIndexAmerica_Rosario uint16 = 531 // America/Rosario
	ZoneInfoIndexAmerica_Santa_Isabel uint16 = 594 // America/Santa_Isabel
	ZoneInfoIndexAmerica_Santarem uint16 = 257 // America/Santarem
	ZoneInfoIndexAmerica_Santiago uint16 = 258 // America/Santiago
	ZoneInfoIndexAmerica_Santo_Domingo uint16 = 266 // America/Santo_Domingo
	ZoneInfoIndexAmerica_Sao_Paulo uint16 = 50 // America/Sao_Paulo
	ZoneInfoIndexAmerica_Scoresbysund uint16 = 53 // America/Scoresbysund
	ZoneInfoIndexAmerica_Shiprock uint16 = 298 // America/Shiprock
	ZoneInfoIndexAmerica_Sitka uint16 = 362 // America/Sitka
	ZoneInfoIndexAmerica_St_Barthelemy uint16 = 15 // America/St_Barthelemy
	ZoneInfoIndexAmerica_St_Johns uint16 = 7 // America/St_Johns
	ZoneInfoIndexAmerica_St_Kitts uint16 = 8 // America/St_Kitts
	ZoneInfoIndexAmerica_St_Lucia uint16 = 10 // America/St_Lucia
	ZoneInfoIndexAmerica_St_Thomas uint16 = 443 // America/St_Thomas
	ZoneInfoIndexAmerica_St_Vincent uint16 = 300 // America/St_Vincent
	ZoneInfoIndexAmerica_Swift_Current uint16 = 530 // America/Swift_Current
	ZoneInfoIndexAmerica_Tegucigalpa uint16 = 462 // America/Tegucigalpa
	ZoneInfoIndexAmerica_Thule uint16 = 363 // America/Thule
	ZoneInfoIndexAmerica_Thunder_Bay uint16 = 577 // America/Thunder_Bay
	ZoneInfoIndexAmerica_Tijuana uint16 = 231 // America/Tijuana
	ZoneInfoIndexAmerica_Toronto uint16 = 276 // America/Toronto
	ZoneInfoIndexAmerica_Tortola uint16 = 277 // America/Tortola
	ZoneInfoIndexAmerica_Vancouver uint16 = 111 // America/Vancouver
	ZoneInfoIndexAmerica_Virgin uint16 = 469 // America/Virgin
	ZoneInfoIndexAmerica_Whitehorse uint16 = 180 // America/Whitehorse
	ZoneInfoIndexAmerica_Winnipeg uint16 = 316 // America/Winnipeg
	ZoneInfoIndexAmerica_Yakutat uint16 = 516 // America/Yakutat
	ZoneInfoIndexAmerica_Yellowknife uint16 = 48 // America/Yellowknife
	ZoneInfoIndexAntarctica_Casey uint16 = 541 // Antarctica/Casey
	ZoneInfoIndexAntarctica_Davis uint16 = 542 // Antarctica/Davis
	ZoneInfoIndexAntarctica_DumontDUrville uint16 = 189 // Antarctica/DumontDUrville
	ZoneInfoIndexAntarctica_Macquarie uint16 = 336 // Antarctica/Macquarie
	ZoneInfoIndexAntarctica_Mawson uint16 = 130 // Antarctica/Mawson
	ZoneInfoIndexAntarctica_McMurdo uint16 = 246 // Antarctica/McMurdo
	ZoneInfoIndexAntarctica_Palmer uint16 = 153 // Antarctica/Palmer
	ZoneInfoIndexAntarctica_Rothera uint16 = 47 // Antarctica/Rothera
	ZoneInfoIndexAntarctica_South_Pole uint16 = 488 // Antarctica/South_Pole
	ZoneInfoIndexAntarctica_Syowa uint16 = 545 // Antarctica/Syowa
	ZoneInfoIndexAntarctica_Troll uint16 = 546 // Antarctica/Troll
	ZoneInfoIndexAntarctica_Vostok uint16 = 171 // Antarctica/Vostok
	ZoneInfoIndexArctic_Longyearbyen uint16 = 502 // Arctic/Longyearbyen
	ZoneInfoIndexAsia_Aden uint16 = 80 // Asia/Aden
	ZoneInfoIndexAsia_Almaty uint16 = 413 // Asia/Almaty
	ZoneInfoIndexAsia_Amman uint16 = 55 // Asia/Amman
	ZoneInfoIndexAsia_Anadyr uint16 = 414 // Asia/Anadyr
	ZoneInfoIndexAsia_Aqtau uint16 = 56 // Asia/Aqtau
	ZoneInfoIndexAsia_Aqtobe uint16 = 415 // Asia/Aqtobe
	ZoneInfoIndexAsia_Ashgabat uint16 = 453 // Asia/Ashgabat
	ZoneInfoIndexAsia_Ashkhabad uint16 = 63 // Asia/Ashkhabad
	ZoneInfoIndexAsia_Atyrau uint16 = 416 // Asia/Atyrau
	ZoneInfoIndexAsia_Baghdad uint16 = 370 // Asia/Baghdad
	ZoneInfoIndexAsia_Bahrain uint16 = 371 // Asia/Bahrain
	ZoneInfoIndexAsia_Baku uint16 = 81 // Asia/Baku
	ZoneInfoIndexAsia_Bangkok uint16 = 394 // Asia/Bangkok
	ZoneInfoIndexAsia_Barnaul uint16 = 397 // Asia/Barnaul
	ZoneInfoIndexAsia_Beirut uint16 = 421 // Asia/Beirut
	ZoneInfoIndexAsia_Bishkek uint16 = 439 // Asia/Bishkek
	ZoneInfoIndexAsia_Brunei uint16 = 424 // Asia/Brunei
	ZoneInfoIndexAsia_Calcutta uint16 = 123 // Asia/Calcutta
	ZoneInfoIndexAsia_Chita uint16 = 57 // Asia/Chita
	ZoneInfoIndexAsia_Choibalsan uint16 = 335 // Asia/Choibalsan
	ZoneInfoIndexAsia_Chongqing uint16 = 574 // Asia/Chongqing
	ZoneInfoIndexAsia_Chungking uint16 = 479 // Asia/Chungking
	ZoneInfoIndexAsia_Colombo uint16 = 24 // Asia/Colombo
	ZoneInfoIndexAsia_Dacca uint16 = 58 // Asia/Dacca
	ZoneInfoIndexAsia_Damascus uint16 = 89 // Asia/Damascus
	ZoneInfoIndexAsia_Dhaka uint16 = 59 // Asia/Dhaka
	ZoneInfoIndexAsia_Dili uint16 = 82 // Asia/Dili
	ZoneInfoIndexAsia_Dubai uint16 = 60 // Asia/Dubai
	ZoneInfoIndexAsia_Dushanbe uint16 = 125 // Asia/Dushanbe
	ZoneInfoIndexAsia_Famagusta uint16 = 105 // Asia/Famagusta
	ZoneInfoIndexAsia_Gaza uint16 = 83 // Asia/Gaza
	ZoneInfoIndexAsia_Harbin uint16 = 447 // Asia/Harbin
	ZoneInfoIndexAsia_Hebron uint16 = 448 // Asia/Hebron
	ZoneInfoIndexAsia_Ho_Chi_Minh uint16 = 88 // Asia/Ho_Chi_Minh
	ZoneInfoIndexAsia_Hong_Kong uint16 = 186 // Asia/Hong_Kong
	ZoneInfoIndexAsia_Hovd uint16 = 84 // Asia/Hovd
	ZoneInfoIndexAsia_Irkutsk uint16 = 532 // Asia/Irkutsk
	ZoneInfoIndexAsia_Istanbul uint16 = 128 // Asia/Istanbul
	ZoneInfoIndexAsia_Jakarta uint16 = 12 // Asia/Jakarta
	ZoneInfoIndexAsia_Jayapura uint16 = 477 // Asia/Jayapura
	ZoneInfoIndexAsia_Jerusalem uint16 = 193 // Asia/Jerusalem
	ZoneInfoIndexAsia_Kabul uint16 = 62 // Asia/Kabul
	ZoneInfoIndexAsia_Kamchatka uint16 = 256 // Asia/Kamchatka
	ZoneInfoIndexAsia_Karachi uint16 = 176 // Asia/Karachi
	ZoneInfoIndexAsia_Kashgar uint16 = 177 // Asia/Kashgar
	ZoneInfoIndexAsia_Kathmandu uint16 = 365 // Asia/Kathmandu
	ZoneInfoIndexAsia_Katmandu uint16 = 420 // Asia/Katmandu
	ZoneInfoIndexAsia_Khandyga uint16 = 345 // Asia/Khandyga
	ZoneInfoIndexAsia_Kolkata uint16 = 251 // Asia/Kolkata
	ZoneInfoIndexAsia_Krasnoyarsk uint16 = 495 // Asia/Krasnoyarsk
	ZoneInfoIndexAsia_Kuala_Lumpur uint16 = 2 // Asia/Kuala_Lumpur
	ZoneInfoIndexAsia_Kuching uint16 = 294 // Asia/Kuching
	ZoneInfoIndexAsia_Kuwait uint16 = 458 // Asia/Kuwait
	ZoneInfoIndexAsia_Macao uint16 = 64 // Asia/Macao
	ZoneInfoIndexAsia_Macau uint16 = 65 // Asia/Macau
	ZoneInfoIndexAsia_Magadan uint16 = 558 // Asia/Magadan
	ZoneInfoIndexAsia_Makassar uint16 = 232 // Asia/Makassar
	ZoneInfoIndexAsia_Manila uint16 = 465 // Asia/Manila
	ZoneInfoIndexAsia_Muscat uint16 = 471 // Asia/Muscat
	ZoneInfoIndexAsia_Nicosia uint16 = 166 // Asia/Nicosia
	ZoneInfoIndexAsia_Novokuznetsk uint16 = 227 // Asia/Novokuznetsk
	ZoneInfoIndexAsia_Novosibirsk uint16 = 405 // Asia/Novosibirsk
	ZoneInfoIndexAsia_Omsk uint16 = 85 // Asia/Omsk
	ZoneInfoIndexAsia_Oral uint16 = 86 // Asia/Oral
	ZoneInfoIndexAsia_Phnom_Penh uint16 = 470 // Asia/Phnom_Penh
	ZoneInfoIndexAsia_Pontianak uint16 = 75 // Asia/Pontianak
	ZoneInfoIndexAsia_Pyongyang uint16 = 340 // Asia/Pyongyang
	ZoneInfoIndexAsia_Qatar uint16 = 66 // Asia/Qatar
	ZoneInfoIndexAsia_Qostanay uint16 = 214 // Asia/Qostanay
	ZoneInfoIndexAsia_Qyzylorda uint16 = 250 // Asia/Qyzylorda
	ZoneInfoIndexAsia_Rangoon uint16 = 240 // Asia/Rangoon
	ZoneInfoIndexAsia_Riyadh uint16 = 489 // Asia/Riyadh
	ZoneInfoIndexAsia_Saigon uint16 = 492 // Asia/Saigon
	ZoneInfoIndexAsia_Sakhalin uint16 = 567 // Asia/Sakhalin
	ZoneInfoIndexAsia_Samarkand uint16 = 54 // Asia/Samarkand
	ZoneInfoIndexAsia_Seoul uint16 = 67 // Asia/Seoul
	ZoneInfoIndexAsia_Shanghai uint16 = 572 // Asia/Shanghai
	ZoneInfoIndexAsia_Singapore uint16 = 494 // Asia/Singapore
	ZoneInfoIndexAsia_Srednekolymsk uint16 = 460 // Asia/Srednekolymsk
	ZoneInfoIndexAsia_Taipei uint16 = 499 // Asia/Taipei
	ZoneInfoIndexAsia_Tashkent uint16 = 566 // Asia/Tashkent
	ZoneInfoIndexAsia_Tbilisi uint16 = 20 // Asia/Tbilisi
	ZoneInfoIndexAsia_Tehran uint16 = 500 // Asia/Tehran
	ZoneInfoIndexAsia_Tel_Aviv uint16 = 70 // Asia/Tel_Aviv
	ZoneInfoIndexAsia_Thimbu uint16 = 501 // Asia/Thimbu
	ZoneInfoIndexAsia_Thimphu uint16 = 71 // Asia/Thimphu
	ZoneInfoIndexAsia_Tokyo uint16 = 68 // Asia/Tokyo
	ZoneInfoIndexAsia_Tomsk uint16 = 69 // Asia/Tomsk
	ZoneInfoIndexAsia_Ujung_Pandang uint16 = 199 // Asia/Ujung_Pandang
	ZoneInfoIndexAsia_Ulaanbaatar uint16 = 120 // Asia/Ulaanbaatar
	ZoneInfoIndexAsia_Ulan_Bator uint16 = 129 // Asia/Ulan_Bator
	ZoneInfoIndexAsia_Urumqi uint16 = 509 // Asia/Urumqi
	ZoneInfoIndexAsia_Ust_Nera uint16 = 164 // Asia/Ust-Nera
	ZoneInfoIndexAsia_Vientiane uint16 = 305 // Asia/Vientiane
	ZoneInfoIndexAsia_Vladivostok uint16 = 108 // Asia/Vladivostok
	ZoneInfoIndexAsia_Yakutsk uint16 = 302 // Asia/Yakutsk
	ZoneInfoIndexAsia_Yangon uint16 = 524 // Asia/Yangon
	ZoneInfoIndexAsia_Yekaterinburg uint16 = 588 // Asia/Yekaterinburg
	ZoneInfoIndexAsia_Yerevan uint16 = 333 // Asia/Yerevan
	ZoneInfoIndexAtlantic_Azores uint16 = 575 // Atlantic/Azores
	ZoneInfoIndexAtlantic_Bermuda uint16 = 135 // Atlantic/Bermuda
	ZoneInfoIndexAtlantic_Canary uint16 = 592 // Atlantic/Canary
	ZoneInfoIndexAtlantic_Cape_Verde uint16 = 196 // Atlantic/Cape_Verde
	ZoneInfoIndexAtlantic_Faeroe uint16 = 6 // Atlantic/Faeroe
	ZoneInfoIndexAtlantic_Faroe uint16 = 539 // Atlantic/Faroe
	ZoneInfoIndexAtlantic_Jan_Mayen uint16 = 190 // Atlantic/Jan_Mayen
	ZoneInfoIndexAtlantic_Madeira uint16 = 295 // Atlantic/Madeira
	ZoneInfoIndexAtlantic_Reykjavik uint16 = 77 // Atlantic/Reykjavik
	ZoneInfoIndexAtlantic_South_Georgia uint16 = 126 // Atlantic/South_Georgia
	ZoneInfoIndexAtlantic_St_Helena uint16 = 159 // Atlantic/St_Helena
	ZoneInfoIndexAtlantic_Stanley uint16 = 282 // Atlantic/Stanley
	ZoneInfoIndexAustralia_ACT uint16 = 309 // Australia/ACT
	ZoneInfoIndexAustralia_Adelaide uint16 = 97 // Australia/Adelaide
	ZoneInfoIndexAustralia_Brisbane uint16 = 173 // Australia/Brisbane
	ZoneInfoIndexAustralia_Broken_Hill uint16 = 438 // Australia/Broken_Hill
	ZoneInfoIndexAustralia_Canberra uint16 = 109 // Australia/Canberra
	ZoneInfoIndexAustralia_Currie uint16 = 102 // Australia/Currie
	ZoneInfoIndexAustralia_Darwin uint16 = 103 // Australia/Darwin
	ZoneInfoIndexAustralia_Eucla uint16 = 317 // Australia/Eucla
	ZoneInfoIndexAustralia_Hobart uint16 = 124 // Australia/Hobart
	ZoneInfoIndexAustralia_LHI uint16 = 310 // Australia/LHI
	ZoneInfoIndexAustralia_Lindeman uint16 = 534 // Australia/Lindeman
	ZoneInfoIndexAustralia_Lord_Howe uint16 = 418 // Australia/Lord_Howe
	ZoneInfoIndexAustralia_Melbourne uint16 = 49 // Australia/Melbourne
	ZoneInfoIndexAustralia_NSW uint16 = 311 // Australia/NSW
	ZoneInfoIndexAustralia_North uint16 = 321 // Australia/North
	ZoneInfoIndexAustralia_Perth uint16 = 322 // Australia/Perth
	ZoneInfoIndexAustralia_Queensland uint16 = 503 // Australia/Queensland
	ZoneInfoIndexAustralia_South uint16 = 323 // Australia/South
	ZoneInfoIndexAustralia_Sydney uint16 = 168 // Australia/Sydney
	ZoneInfoIndexAustralia_Tasmania uint16 = 551 // Australia/Tasmania
	ZoneInfoIndexAustralia_Victoria uint16 = 5 // Australia/Victoria
	ZoneInfoIndexAustralia_West uint16 = 526 // Australia/West
	ZoneInfoIndexAustralia_Yancowinna uint16 = 331 // Australia/Yancowinna
	ZoneInfoIndexBrazil_Acre uint16 = 217 // Brazil/Acre
	ZoneInfoIndexBrazil_DeNoronha uint16 = 366 // Brazil/DeNoronha
	ZoneInfoIndexBrazil_East uint16 = 218 // Brazil/East
	ZoneInfoIndexBrazil_West uint16 = 220 // Brazil/West
	ZoneInfoIndexCET uint16 = 25 // CET
	ZoneInfoIndexCST6CDT uint16 = 561 // CST6CDT
	ZoneInfoIndexCanada_Atlantic uint16 = 178 // Canada/Atlantic
	ZoneInfoIndexCanada_Central uint16 = 211 // Canada/Central
	ZoneInfoIndexCanada_Eastern uint16 = 565 // Canada/Eastern
	ZoneInfoIndexCanada_Mountain uint16 = 590 // Canada/Mountain
	ZoneInfoIndexCanada_Newfoundland uint16 = 446 // Canada/Newfoundland
	ZoneInfoIndexCanada_Pacific uint16 = 154 // Canada/Pacific
	ZoneInfoIndexCanada_Saskatchewan uint16 = 268 // Canada/Saskatchewan
	ZoneInfoIndexCanada_Yukon uint16 = 275 // Canada/Yukon
	ZoneInfoIndexChile_Continental uint16 = 292 // Chile/Continental
	ZoneInfoIndexChile_EasterIsland uint16 = 440 // Chile/EasterIsland
	ZoneInfoIndexCuba uint16 = 285 // Cuba
	ZoneInfoIndexEET uint16 = 26 // EET
	ZoneInfoIndexEST uint16 = 27 // EST
	ZoneInfoIndexEST5EDT uint16 = 313 // EST5EDT
	ZoneInfoIndexEgypt uint16 = 41 // Egypt
	ZoneInfoIndexEire uint16 = 286 // Eire
	ZoneInfoIndexEtc_GMT uint16 = 513 // Etc/GMT
	ZoneInfoIndexEtc_GMT_PLUS_0 uint16 = 372 // Etc/GMT+0
	ZoneInfoIndexEtc_GMT_PLUS_1 uint16 = 373 // Etc/GMT+1
	ZoneInfoIndexEtc_GMT_PLUS_10 uint16 = 145 // Etc/GMT+10
	ZoneInfoIndexEtc_GMT_PLUS_11 uint16 = 146 // Etc/GMT+11
	ZoneInfoIndexEtc_GMT_PLUS_12 uint16 = 147 // Etc/GMT+12
	ZoneInfoIndexEtc_GMT_PLUS_2 uint16 = 374 // Etc/GMT+2
	ZoneInfoIndexEtc_GMT_PLUS_3 uint16 = 375 // Etc/GMT+3
	ZoneInfoIndexEtc_GMT_PLUS_4 uint16 = 376 // Etc/GMT+4
	ZoneInfoIndexEtc_GMT_PLUS_5 uint16 = 377 // Etc/GMT+5
	ZoneInfoIndexEtc_GMT_PLUS_6 uint16 = 378 // Etc/GMT+6
	ZoneInfoIndexEtc_GMT_PLUS_7 uint16 = 379 // Etc/GMT+7
	ZoneInfoIndexEtc_GMT_PLUS_8 uint16 = 380 // Etc/GMT+8
	ZoneInfoIndexEtc_GMT_PLUS_9 uint16 = 381 // Etc/GMT+9
	ZoneInfoIndexEtc_GMT_0 uint16 = 382 // Etc/GMT-0
	ZoneInfoIndexEtc_GMT_1 uint16 = 383 // Etc/GMT-1
	ZoneInfoIndexEtc_GMT_10 uint16 = 148 // Etc/GMT-10
	ZoneInfoIndexEtc_GMT_11 uint16 = 149 // Etc/GMT-11
	ZoneInfoIndexEtc_GMT_12 uint16 = 150 // Etc/GMT-12
	ZoneInfoIndexEtc_GMT_13 uint16 = 151 // Etc/GMT-13
	ZoneInfoIndexEtc_GMT_14 uint16 = 152 // Etc/GMT-14
	ZoneInfoIndexEtc_GMT_2 uint16 = 384 // Etc/GMT-2
	ZoneInfoIndexEtc_GMT_3 uint16 = 385 // Etc/GMT-3
	ZoneInfoIndexEtc_GMT_4 uint16 = 386 // Etc/GMT-4
	ZoneInfoIndexEtc_GMT_5 uint16 = 387 // Etc/GMT-5
	ZoneInfoIndexEtc_GMT_6 uint16 = 388 // Etc/GMT-6
	ZoneInfoIndexEtc_GMT_7 uint16 = 389 // Etc/GMT-7
	ZoneInfoIndexEtc_GMT_8 uint16 = 390 // Etc/GMT-8
	ZoneInfoIndexEtc_GMT_9 uint16 = 391 // Etc/GMT-9
	ZoneInfoIndexEtc_GMT0 uint16 = 569 // Etc/GMT0
	ZoneInfoIndexEtc_Greenwich uint16 = 101 // Etc/Greenwich
	ZoneInfoIndexEtc_UCT uint16 = 514 // Etc/UCT
	ZoneInfoIndexEtc_UTC uint16 = 515 // Etc/UTC
	ZoneInfoIndexEtc_Universal uint16 = 116 // Etc/Universal
	ZoneInfoIndexEtc_Zulu uint16 = 570 // Etc/Zulu
	ZoneInfoIndexEurope_Amsterdam uint16 = 51 // Europe/Amsterdam
	ZoneInfoIndexEurope_Andorra uint16 = 354 // Europe/Andorra
	ZoneInfoIndexEurope_Astrakhan uint16 = 543 // Europe/Astrakhan
	ZoneInfoIndexEurope_Athens uint16 = 155 // Europe/Athens
	ZoneInfoIndexEurope_Belfast uint16 = 496 // Europe/Belfast
	ZoneInfoIndexEurope_Belgrade uint16 = 535 // Europe/Belgrade
	ZoneInfoIndexEurope_Berlin uint16 = 158 // Europe/Berlin
	ZoneInfoIndexEurope_Bratislava uint16 = 519 // Europe/Bratislava
	ZoneInfoIndexEurope_Brussels uint16 = 529 // Europe/Brussels
	ZoneInfoIndexEurope_Bucharest uint16 = 585 // Europe/Bucharest
	ZoneInfoIndexEurope_Budapest uint16 = 369 // Europe/Budapest
	ZoneInfoIndexEurope_Busingen uint16 = 464 // Europe/Busingen
	ZoneInfoIndexEurope_Chisinau uint16 = 434 // Europe/Chisinau
	ZoneInfoIndexEurope_Copenhagen uint16 = 538 // Europe/Copenhagen
	ZoneInfoIndexEurope_Dublin uint16 = 165 // Europe/Dublin
	ZoneInfoIndexEurope_Gibraltar uint16 = 573 // Europe/Gibraltar
	ZoneInfoIndexEurope_Guernsey uint16 = 138 // Europe/Guernsey
	ZoneInfoIndexEurope_Helsinki uint16 = 233 // Europe/Helsinki
	ZoneInfoIndexEurope_Isle_of_Man uint16 = 557 // Europe/Isle_of_Man
	ZoneInfoIndexEurope_Istanbul uint16 = 399 // Europe/Istanbul
	ZoneInfoIndexEurope_Jersey uint16 = 185 // Europe/Jersey
	ZoneInfoIndexEurope_Kaliningrad uint16 = 504 // Europe/Kaliningrad
	ZoneInfoIndexEurope_Kiev uint16 = 406 // Europe/Kiev
	ZoneInfoIndexEurope_Kirov uint16 = 582 // Europe/Kirov
	ZoneInfoIndexEurope_Kyiv uint16 = 407 // Europe/Kyiv
	ZoneInfoIndexEurope_Lisbon uint16 = 195 // Europe/Lisbon
	ZoneInfoIndexEurope_Ljubljana uint16 = 457 // Europe/Ljubljana
	ZoneInfoIndexEurope_London uint16 = 197 // Europe/London
	ZoneInfoIndexEurope_Luxembourg uint16 = 79 // Europe/Luxembourg
	ZoneInfoIndexEurope_Madrid uint16 = 201 // Europe/Madrid
	ZoneInfoIndexEurope_Malta uint16 = 583 // Europe/Malta
	ZoneInfoIndexEurope_Mariehamn uint16 = 39 // Europe/Mariehamn
	ZoneInfoIndexEurope_Minsk uint16 = 584 // Europe/Minsk
	ZoneInfoIndexEurope_Monaco uint16 = 204 // Europe/Monaco
	ZoneInfoIndexEurope_Moscow uint16 = 205 // Europe/Moscow
	ZoneInfoIndexEurope_Nicosia uint16 = 262 // Europe/Nicosia
	ZoneInfoIndexEurope_Oslo uint16 = 408 // Europe/Oslo
	ZoneInfoIndexEurope_Paris uint16 = 587 // Europe/Paris
	ZoneInfoIndexEurope_Podgorica uint16 = 76 // Europe/Podgorica
	ZoneInfoIndexEurope_Prague uint16 = 216 // Europe/Prague
	ZoneInfoIndexEurope_Riga uint16 = 409 // Europe/Riga
	ZoneInfoIndexEurope_Rome uint16 = 410 // Europe/Rome
	ZoneInfoIndexEurope_Samara uint16 = 236 // Europe/Samara
	ZoneInfoIndexEurope_San_Marino uint16 = 491 // Europe/San_Marino
	ZoneInfoIndexEurope_Sarajevo uint16 = 229 // Europe/Sarajevo
	ZoneInfoIndexEurope_Saratov uint16 = 547 // Europe/Saratov
	ZoneInfoIndexEurope_Simferopol uint16 = 520 // Europe/Simferopol
	ZoneInfoIndexEurope_Skopje uint16 = 237 // Europe/Skopje
	ZoneInfoIndexEurope_Sofia uint16 = 589 // Europe/Sofia
	ZoneInfoIndexEurope_Stockholm uint16 = 194 // Europe/Stockholm
	ZoneInfoIndexEurope_Tallinn uint16 = 118 // Europe/Tallinn
	ZoneInfoIndexEurope_Tirane uint16 = 244 // Europe/Tirane
	ZoneInfoIndexEurope_Tiraspol uint16 = 459 // Europe/Tiraspol
	ZoneInfoIndexEurope_Ulyanovsk uint16 = 533 // Europe/Ulyanovsk
	ZoneInfoIndexEurope_Uzhgorod uint16 = 437 // Europe/Uzhgorod
	ZoneInfoIndexEurope_Vaduz uint16 = 591 // Europe/Vaduz
	ZoneInfoIndexEurope_Vatican uint16 = 484 // Europe/Vatican
	ZoneInfoIndexEurope_Vienna uint16 = 253 // Europe/Vienna
	ZoneInfoIndexEurope_Vilnius uint16 = 525 // Europe/Vilnius
	ZoneInfoIndexEurope_Volgograd uint16 = 143 // Europe/Volgograd
	ZoneInfoIndexEurope_Warsaw uint16 = 263 // Europe/Warsaw
	ZoneInfoIndexEurope_Zagreb uint16 = 284 // Europe/Zagreb
	ZoneInfoIndexEurope_Zaporozhye uint16 = 556 // Europe/Zaporozhye
	ZoneInfoIndexEurope_Zurich uint16 = 291 // Europe/Zurich
	ZoneInfoIndexGB uint16 = 0 // GB
	ZoneInfoIndexGB_Eire uint16 = 581 // GB-Eire
	ZoneInfoIndexGMT uint16 = 28 // GMT
	ZoneInfoIndexGMT_PLUS_0 uint16 = 42 // GMT+0
	ZoneInfoIndexGMT_0 uint16 = 43 // GMT-0
	ZoneInfoIndexGMT0 uint16 = 287 // GMT0
	ZoneInfoIndexGreenwich uint16 = 480 // Greenwich
	ZoneInfoIndexHST uint16 = 29 // HST
	ZoneInfoIndexHongkong uint16 = 184 // Hongkong
	ZoneInfoIndexIceland uint16 = 549 // Iceland
	ZoneInfoIndexIndian_Antananarivo uint16 = 400 // Indian/Antananarivo
	ZoneInfoIndexIndian_Chagos uint16 = 160 // Indian/Chagos
	ZoneInfoIndexIndian_Christmas uint16 = 226 // Indian/Christmas
	ZoneInfoIndexIndian_Cocos uint16 = 4 // Indian/Cocos
	ZoneInfoIndexIndian_Comoro uint16 = 162 // Indian/Comoro
	ZoneInfoIndexIndian_Kerguelen uint16 = 156 // Indian/Kerguelen
	ZoneInfoIndexIndian_Mahe uint16 = 161 // Indian/Mahe
	ZoneInfoIndexIndian_Maldives uint16 = 360 // Indian/Maldives
	ZoneInfoIndexIndian_Mauritius uint16 = 281 // Indian/Mauritius
	ZoneInfoIndexIndian_Mayotte uint16 = 550 // Indian/Mayotte
	ZoneInfoIndexIndian_Reunion uint16 = 249 // Indian/Reunion
	ZoneInfoIndexIran uint16 = 288 // Iran
	ZoneInfoIndexIsrael uint16 = 454 // Israel
	ZoneInfoIndexJamaica uint16 = 114 // Jamaica
	ZoneInfoIndexJapan uint16 = 44 // Japan
	ZoneInfoIndexKwajalein uint16 = 46 // Kwajalein
	ZoneInfoIndexLibya uint16 = 45 // Libya
	ZoneInfoIndexMET uint16 = 30 // MET
	ZoneInfoIndexMST uint16 = 31 // MST
	ZoneInfoIndexMST7MDT uint16 = 564 // MST7MDT
	ZoneInfoIndexMexico_BajaNorte uint16 = 593 // Mexico/BajaNorte
	ZoneInfoIndexMexico_BajaSur uint16 = 19 // Mexico/BajaSur
	ZoneInfoIndexMexico_General uint16 = 338 // Mexico/General
	ZoneInfoIndexNZ uint16 = 1 // NZ
	ZoneInfoIndexNZ_CHAT uint16 = 169 // NZ-CHAT
	ZoneInfoIndexNavajo uint16 = 474 // Navajo
	ZoneInfoIndexPRC uint16 = 32 // PRC
	ZoneInfoIndexPST8PDT uint16 = 518 // PST8PDT
	ZoneInfoIndexPacific_Apia uint16 = 91 // Pacific/Apia
	ZoneInfoIndexPacific_Auckland uint16 = 98 // Pacific/Auckland
	ZoneInfoIndexPacific_Bougainville uint16 = 203 // Pacific/Bougainville
	ZoneInfoIndexPacific_Chatham uint16 = 115 // Pacific/Chatham
	ZoneInfoIndexPacific_Chuuk uint16 = 306 // Pacific/Chuuk
	ZoneInfoIndexPacific_Easter uint16 = 493 // Pacific/Easter
	ZoneInfoIndexPacific_Efate uint16 = 307 // Pacific/Efate
	ZoneInfoIndexPacific_Enderbury uint16 = 209 // Pacific/Enderbury
	ZoneInfoIndexPacific_Fakaofo uint16 = 17 // Pacific/Fakaofo
	ZoneInfoIndexPacific_Fiji uint16 = 92 // Pacific/Fiji
	ZoneInfoIndexPacific_Funafuti uint16 = 521 // Pacific/Funafuti
	ZoneInfoIndexPacific_Galapagos uint16 = 426 // Pacific/Galapagos
	ZoneInfoIndexPacific_Gambier uint16 = 179 // Pacific/Gambier
	ZoneInfoIndexPacific_Guadalcanal uint16 = 568 // Pacific/Guadalcanal
	ZoneInfoIndexPacific_Guam uint16 = 93 // Pacific/Guam
	ZoneInfoIndexPacific_Honolulu uint16 = 552 // Pacific/Honolulu
	ZoneInfoIndexPacific_Johnston uint16 = 441 // Pacific/Johnston
	ZoneInfoIndexPacific_Kanton uint16 = 523 // Pacific/Kanton
	ZoneInfoIndexPacific_Kiritimati uint16 = 299 // Pacific/Kiritimati
	ZoneInfoIndexPacific_Kosrae uint16 = 527 // Pacific/Kosrae
	ZoneInfoIndexPacific_Kwajalein uint16 = 324 // Pacific/Kwajalein
	ZoneInfoIndexPacific_Majuro uint16 = 540 // Pacific/Majuro
	ZoneInfoIndexPacific_Marquesas uint16 = 187 // Pacific/Marquesas
	ZoneInfoIndexPacific_Midway uint16 = 544 // Pacific/Midway
	ZoneInfoIndexPacific_Nauru uint16 = 312 // Pacific/Nauru
	ZoneInfoIndexPacific_Niue uint16 = 94 // Pacific/Niue
	ZoneInfoIndexPacific_Norfolk uint16 = 329 // Pacific/Norfolk
	ZoneInfoIndexPacific_Noumea uint16 = 548 // Pacific/Noumea
	ZoneInfoIndexPacific_Pago_Pago uint16 = 208 // Pacific/Pago_Pago
	ZoneInfoIndexPacific_Palau uint16 = 314 // Pacific/Palau
	ZoneInfoIndexPacific_Pitcairn uint16 = 304 // Pacific/Pitcairn
	ZoneInfoIndexPacific_Pohnpei uint16 = 104 // Pacific/Pohnpei
	ZoneInfoIndexPacific_Ponape uint16 = 555 // Pacific/Ponape
	ZoneInfoIndexPacific_Port_Moresby uint16 = 419 // Pacific/Port_Moresby
	ZoneInfoIndexPacific_Rarotonga uint16 = 364 // Pacific/Rarotonga
	ZoneInfoIndexPacific_Saipan uint16 = 560 // Pacific/Saipan
	ZoneInfoIndexPacific_Samoa uint16 = 315 // Pacific/Samoa
	ZoneInfoIndexPacific_Tahiti uint16 = 562 // Pacific/Tahiti
	ZoneInfoIndexPacific_Tarawa uint16 = 563 // Pacific/Tarawa
	ZoneInfoIndexPacific_Tongatapu uint16 = 99 // Pacific/Tongatapu
	ZoneInfoIndexPacific_Truk uint16 = 95 // Pacific/Truk
	ZoneInfoIndexPacific_Wake uint16 = 96 // Pacific/Wake
	ZoneInfoIndexPacific_Wallis uint16 = 576 // Pacific/Wallis
	ZoneInfoIndexPacific_Yap uint16 = 455 // Pacific/Yap
	ZoneInfoIndexPoland uint16 = 482 // Poland
	ZoneInfoIndexPortugal uint16 = 473 // Portugal
	ZoneInfoIndexROC uint16 = 33 // ROC
	ZoneInfoIndexROK uint16 = 34 // ROK
	ZoneInfoIndexSingapore uint16 = 422 // Singapore
	ZoneInfoIndexTurkey uint16 = 506 // Turkey
	ZoneInfoIndexUCT uint16 = 35 // UCT
	ZoneInfoIndexUS_Alaska uint16 = 580 // US/Alaska
	ZoneInfoIndexUS_Aleutian uint16 = 172 // US/Aleutian
	ZoneInfoIndexUS_Arizona uint16 = 170 // US/Arizona
	ZoneInfoIndexUS_Central uint16 = 483 // US/Central
	ZoneInfoIndexUS_East_Indiana uint16 = 241 // US/East-Indiana
	ZoneInfoIndexUS_Eastern uint16 = 192 // US/Eastern
	ZoneInfoIndexUS_Hawaii uint16 = 22 // US/Hawaii
	ZoneInfoIndexUS_Indiana_Starke uint16 = 223 // US/Indiana-Starke
	ZoneInfoIndexUS_Michigan uint16 = 267 // US/Michigan
	ZoneInfoIndexUS_Mountain uint16 = 245 // US/Mountain
	ZoneInfoIndexUS_Pacific uint16 = 425 // US/Pacific
	ZoneInfoIndexUS_Samoa uint16 = 183 // US/Samoa
	ZoneInfoIndexUTC uint16 = 36 // UTC
	ZoneInfoIndexUniversal uint16 = 498 // Universal
	ZoneInfoIndexW_SU uint16 = 289 // W-SU
	ZoneInfoIndexWET uint16 = 37 // WET
	ZoneInfoIndexZulu uint16 = 290 // Zulu

)
