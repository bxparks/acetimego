// This file was generated by the following script:
//
//   $ /home/brian/src/AceTimeTools/src/acetimetools/tzcompiler.py
//     --input_dir /home/brian/src/AceTimeGo/zonedb/tzfiles
//     --output_dir /home/brian/src/AceTimeGo/zonedb
//     --tz_version 2022g
//     --action zonedb
//     --language go
//     --scope extended
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
// from https://github.com/eggert/tz/releases/tag/2022g
//
// DO NOT EDIT

package zonedb

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

// ---------------------------------------------------------------------------
// Zone Context
// ---------------------------------------------------------------------------

const TzDatabaseVersion string = "2022g"

var Context = zoneinfo.ZoneContext{
	LetterBuffer: LetterBuffer,
	LetterOffsets: LetterOffsets,
	FormatBuffer: FormatBuffer,
	FormatOffsets: FormatOffsets,
	NameBuffer: NameBuffer,
	NameOffsets: NameOffsets,
	ZoneRegistry: ZoneAndLinkRegistry,
	TzDatabaseVersion: TzDatabaseVersion,
}

// ---------------------------------------------------------------------------
// Zone Registry
// Total: 596 (351 zones, 245 links)
// ---------------------------------------------------------------------------

var ZoneAndLinkRegistry = []*zoneinfo.ZoneInfo{
	&ZoneGB, // 0x005973ae, GB -> Europe/London
	&ZoneNZ, // 0x005974ad, NZ -> Pacific/Auckland
	&ZoneAsia_Kuala_Lumpur, // 0x014763c4, Asia/Kuala_Lumpur -> Asia/Singapore
	&ZoneAfrica_Libreville, // 0x01d96de4, Africa/Libreville -> Africa/Lagos
	&ZoneIndian_Cocos, // 0x021e86de, Indian/Cocos -> Asia/Yangon
	&ZoneAustralia_Victoria, // 0x0260d5db, Australia/Victoria -> Australia/Melbourne
	&ZoneAtlantic_Faeroe, // 0x031ec516, Atlantic/Faeroe -> Atlantic/Faroe
	&ZoneAmerica_St_Johns, // 0x04b14e6e, America/St_Johns
	&ZoneAmerica_St_Kitts, // 0x04c0507b, America/St_Kitts -> America/Puerto_Rico
	&ZoneAfrica_Ouagadougou, // 0x04d7219a, Africa/Ouagadougou -> Africa/Abidjan
	&ZoneAmerica_St_Lucia, // 0x04d8b3ba, America/St_Lucia -> America/Puerto_Rico
	&ZoneAmerica_North_Dakota_New_Salem, // 0x04f9958e, America/North_Dakota/New_Salem
	&ZoneAsia_Jakarta, // 0x0506ab50, Asia/Jakarta
	&ZoneAfrica_Bujumbura, // 0x05232a47, Africa/Bujumbura -> Africa/Maputo
	&ZoneAmerica_Mazatlan, // 0x0532189e, America/Mazatlan
	&ZoneAmerica_St_Barthelemy, // 0x054e6a79, America/St_Barthelemy -> America/Puerto_Rico
	&ZoneAfrica_Addis_Ababa, // 0x05ae1e65, Africa/Addis_Ababa -> Africa/Nairobi
	&ZonePacific_Fakaofo, // 0x06532bba, Pacific/Fakaofo
	&ZoneAmerica_Hermosillo, // 0x065d21c4, America/Hermosillo
	&ZoneMexico_BajaSur, // 0x08ee3641, Mexico/BajaSur -> America/Mazatlan
	&ZoneAsia_Tbilisi, // 0x0903e442, Asia/Tbilisi
	&ZoneAmerica_Indiana_Tell_City, // 0x09263612, America/Indiana/Tell_City
	&ZoneUS_Hawaii, // 0x09c8de2f, US/Hawaii -> Pacific/Honolulu
	&ZoneAmerica_Boa_Vista, // 0x0a7b7efe, America/Boa_Vista
	&ZoneAsia_Colombo, // 0x0af0e91d, Asia/Colombo
	&ZoneCET, // 0x0b87d921, CET
	&ZoneEET, // 0x0b87e1a3, EET
	&ZoneEST, // 0x0b87e371, EST
	&ZoneGMT, // 0x0b87eb2d, GMT -> Etc/GMT
	&ZoneHST, // 0x0b87f034, HST
	&ZoneMET, // 0x0b8803ab, MET
	&ZoneMST, // 0x0b880579, MST
	&ZonePRC, // 0x0b88120a, PRC -> Asia/Shanghai
	&ZoneROC, // 0x0b881a29, ROC -> Asia/Taipei
	&ZoneROK, // 0x0b881a31, ROK -> Asia/Seoul
	&ZoneUCT, // 0x0b882571, UCT -> Etc/UTC
	&ZoneUTC, // 0x0b882791, UTC -> Etc/UTC
	&ZoneWET, // 0x0b882e35, WET
	&ZoneAmerica_Guatemala, // 0x0c8259f7, America/Guatemala
	&ZoneEurope_Mariehamn, // 0x0caa6496, Europe/Mariehamn -> Europe/Helsinki
	&ZoneAfrica_Monrovia, // 0x0ce90385, Africa/Monrovia
	&ZoneEgypt, // 0x0d1a278e, Egypt -> Africa/Cairo
	&ZoneGMT_PLUS_0, // 0x0d2f7028, GMT+0 -> Etc/GMT
	&ZoneGMT_0, // 0x0d2f706a, GMT-0 -> Etc/GMT
	&ZoneJapan, // 0x0d712f8f, Japan -> Asia/Tokyo
	&ZoneLibya, // 0x0d998b16, Libya -> Africa/Tripoli
	&ZoneKwajalein, // 0x0e57afbb, Kwajalein -> Pacific/Kwajalein
	&ZoneAntarctica_Rothera, // 0x0e86d203, Antarctica/Rothera
	&ZoneAmerica_Yellowknife, // 0x0f76c76f, America/Yellowknife
	&ZoneAustralia_Melbourne, // 0x0fe559a3, Australia/Melbourne
	&ZoneAmerica_Sao_Paulo, // 0x1063bfc9, America/Sao_Paulo
	&ZoneEurope_Amsterdam, // 0x109395c2, Europe/Amsterdam -> Europe/Brussels
	&ZoneAmerica_Indiana_Vevay, // 0x10aca054, America/Indiana/Vevay
	&ZoneAmerica_Scoresbysund, // 0x123f8d2a, America/Scoresbysund
	&ZoneAsia_Samarkand, // 0x13ae5104, Asia/Samarkand
	&ZoneAsia_Amman, // 0x148d21bc, Asia/Amman
	&ZoneAsia_Aqtau, // 0x148f710e, Asia/Aqtau
	&ZoneAsia_Chita, // 0x14ae863b, Asia/Chita
	&ZoneAsia_Dacca, // 0x14bcac5e, Asia/Dacca -> Asia/Dhaka
	&ZoneAsia_Dhaka, // 0x14c07b8b, Asia/Dhaka
	&ZoneAsia_Dubai, // 0x14c79f77, Asia/Dubai
	&ZoneAmerica_Bahia_Banderas, // 0x14f6329a, America/Bahia_Banderas
	&ZoneAsia_Kabul, // 0x153b5601, Asia/Kabul
	&ZoneAsia_Ashkhabad, // 0x15454f09, Asia/Ashkhabad -> Asia/Ashgabat
	&ZoneAsia_Macao, // 0x155f88b3, Asia/Macao -> Asia/Macau
	&ZoneAsia_Macau, // 0x155f88b9, Asia/Macau
	&ZoneAsia_Qatar, // 0x15a8330b, Asia/Qatar
	&ZoneAsia_Seoul, // 0x15ce82da, Asia/Seoul
	&ZoneAsia_Tokyo, // 0x15e606a8, Asia/Tokyo
	&ZoneAsia_Tomsk, // 0x15e60e60, Asia/Tomsk
	&ZoneAsia_Tel_Aviv, // 0x166d7c2c, Asia/Tel_Aviv -> Asia/Jerusalem
	&ZoneAsia_Thimphu, // 0x170380d1, Asia/Thimphu
	&ZoneAmerica_Guayaquil, // 0x17e64958, America/Guayaquil
	&ZoneAmerica_Montserrat, // 0x199b0a35, America/Montserrat -> America/Puerto_Rico
	&ZoneAmerica_Kentucky_Louisville, // 0x1a21024b, America/Kentucky/Louisville
	&ZoneAsia_Pontianak, // 0x1a76c057, Asia/Pontianak
	&ZoneEurope_Podgorica, // 0x1c1a499c, Europe/Podgorica -> Europe/Belgrade
	&ZoneAtlantic_Reykjavik, // 0x1c2b4f74, Atlantic/Reykjavik -> Africa/Abidjan
	&ZoneAmerica_New_York, // 0x1e2a7654, America/New_York
	&ZoneEurope_Luxembourg, // 0x1f8bc6ce, Europe/Luxembourg -> Europe/Brussels
	&ZoneAsia_Aden, // 0x1fa7084a, Asia/Aden -> Asia/Riyadh
	&ZoneAsia_Baku, // 0x1fa788b5, Asia/Baku
	&ZoneAsia_Dili, // 0x1fa8c394, Asia/Dili
	&ZoneAsia_Gaza, // 0x1faa4875, Asia/Gaza
	&ZoneAsia_Hovd, // 0x1fab0fe3, Asia/Hovd
	&ZoneAsia_Omsk, // 0x1faeddac, Asia/Omsk
	&ZoneAsia_Oral, // 0x1faef0a0, Asia/Oral
	&ZoneAmerica_Montreal, // 0x203a1ea8, America/Montreal -> America/Toronto
	&ZoneAsia_Ho_Chi_Minh, // 0x20f2d127, Asia/Ho_Chi_Minh
	&ZoneAsia_Damascus, // 0x20fbb063, Asia/Damascus
	&ZoneAmerica_Argentina_ComodRivadavia, // 0x22758877, America/Argentina/ComodRivadavia -> America/Argentina/Catamarca
	&ZonePacific_Apia, // 0x23359b5e, Pacific/Apia
	&ZonePacific_Fiji, // 0x23383ba5, Pacific/Fiji
	&ZonePacific_Guam, // 0x2338f9ed, Pacific/Guam
	&ZonePacific_Niue, // 0x233ca014, Pacific/Niue
	&ZonePacific_Truk, // 0x234010a9, Pacific/Truk -> Pacific/Port_Moresby
	&ZonePacific_Wake, // 0x23416c2b, Pacific/Wake -> Pacific/Tarawa
	&ZoneAustralia_Adelaide, // 0x2428e8a3, Australia/Adelaide
	&ZonePacific_Auckland, // 0x25062f86, Pacific/Auckland
	&ZonePacific_Tongatapu, // 0x262ca836, Pacific/Tongatapu
	&ZoneAmerica_Monterrey, // 0x269a1deb, America/Monterrey
	&ZoneEtc_Greenwich, // 0x26daa98c, Etc/Greenwich -> Etc/GMT
	&ZoneAustralia_Currie, // 0x278b6a24, Australia/Currie -> Australia/Hobart
	&ZoneAustralia_Darwin, // 0x2876bdff, Australia/Darwin
	&ZonePacific_Pohnpei, // 0x28929f96, Pacific/Pohnpei -> Pacific/Guadalcanal
	&ZoneAsia_Famagusta, // 0x289b4f8b, Asia/Famagusta
	&ZoneAmerica_Indiana_Vincennes, // 0x28a0b212, America/Indiana/Vincennes
	&ZoneAmerica_Indiana_Indianapolis, // 0x28a669a4, America/Indiana/Indianapolis
	&ZoneAsia_Vladivostok, // 0x29de34a8, Asia/Vladivostok
	&ZoneAustralia_Canberra, // 0x2a09ae58, Australia/Canberra -> Australia/Sydney
	&ZoneAmerica_Fortaleza, // 0x2ad018ee, America/Fortaleza
	&ZoneAmerica_Vancouver, // 0x2c6f6b1f, America/Vancouver
	&ZoneAmerica_Pangnirtung, // 0x2d999193, America/Pangnirtung -> America/Iqaluit
	&ZoneAmerica_Iqaluit, // 0x2de310bf, America/Iqaluit
	&ZoneJamaica, // 0x2e44fdab, Jamaica -> America/Jamaica
	&ZonePacific_Chatham, // 0x2f0de999, Pacific/Chatham
	&ZoneEtc_Universal, // 0x2f8cb9a9, Etc/Universal -> Etc/UTC
	&ZoneAmerica_Indiana_Marengo, // 0x2feeee72, America/Indiana/Marengo
	&ZoneEurope_Tallinn, // 0x30c4e096, Europe/Tallinn
	&ZoneAfrica_Djibouti, // 0x30ea01d4, Africa/Djibouti -> Africa/Nairobi
	&ZoneAsia_Ulaanbaatar, // 0x30f0cc4e, Asia/Ulaanbaatar
	&ZoneAfrica_Gaborone, // 0x317c0aa7, Africa/Gaborone -> Africa/Maputo
	&ZoneAmerica_Argentina_Ushuaia, // 0x320dcdde, America/Argentina/Ushuaia
	&ZoneAsia_Calcutta, // 0x328a44c3, Asia/Calcutta -> Asia/Kolkata
	&ZoneAustralia_Hobart, // 0x32bf951a, Australia/Hobart
	&ZoneAsia_Dushanbe, // 0x32fc5c3c, Asia/Dushanbe
	&ZoneAtlantic_South_Georgia, // 0x33013174, Atlantic/South_Georgia
	&ZoneAmerica_Phoenix, // 0x34b5af01, America/Phoenix
	&ZoneAsia_Istanbul, // 0x382e7894, Asia/Istanbul -> Europe/Istanbul
	&ZoneAsia_Ulan_Bator, // 0x394db4d9, Asia/Ulan_Bator -> Asia/Ulaanbaatar
	&ZoneAntarctica_Mawson, // 0x399cd863, Antarctica/Mawson
	&ZoneAfrica_Brazzaville, // 0x39cda760, Africa/Brazzaville -> Africa/Lagos
	&ZoneAmerica_Caracas, // 0x3be064f4, America/Caracas
	&ZoneAmerica_Cayenne, // 0x3c617269, America/Cayenne
	&ZoneAfrica_Porto_Novo, // 0x3d1bf95d, Africa/Porto-Novo -> Africa/Lagos
	&ZoneAtlantic_Bermuda, // 0x3d4bb1c4, Atlantic/Bermuda
	&ZoneAmerica_Managua, // 0x3d5e7600, America/Managua
	&ZoneAmerica_Marigot, // 0x3dab3a59, America/Marigot -> America/Puerto_Rico
	&ZoneEurope_Guernsey, // 0x3db12c16, Europe/Guernsey -> Europe/London
	&ZoneAfrica_Nouakchott, // 0x3dc49dba, Africa/Nouakchott -> Africa/Abidjan
	&ZoneAmerica_Louisville, // 0x3dcb47ee, America/Louisville -> America/Kentucky/Louisville
	&ZoneAmerica_Argentina_San_Juan, // 0x3e1009bd, America/Argentina/San_Juan
	&ZoneAmerica_Argentina_San_Luis, // 0x3e11238c, America/Argentina/San_Luis
	&ZoneEurope_Volgograd, // 0x3ed0f389, Europe/Volgograd
	&ZoneAmerica_Fort_Nelson, // 0x3f437e0f, America/Fort_Nelson
	&ZoneEtc_GMT_PLUS_10, // 0x3f8f1cc4, Etc/GMT+10
	&ZoneEtc_GMT_PLUS_11, // 0x3f8f1cc5, Etc/GMT+11
	&ZoneEtc_GMT_PLUS_12, // 0x3f8f1cc6, Etc/GMT+12
	&ZoneEtc_GMT_10, // 0x3f8f2546, Etc/GMT-10
	&ZoneEtc_GMT_11, // 0x3f8f2547, Etc/GMT-11
	&ZoneEtc_GMT_12, // 0x3f8f2548, Etc/GMT-12
	&ZoneEtc_GMT_13, // 0x3f8f2549, Etc/GMT-13
	&ZoneEtc_GMT_14, // 0x3f8f254a, Etc/GMT-14
	&ZoneAntarctica_Palmer, // 0x40962f4f, Antarctica/Palmer
	&ZoneCanada_Pacific, // 0x40fa3c7b, Canada/Pacific -> America/Vancouver
	&ZoneEurope_Athens, // 0x4318fa27, Europe/Athens
	&ZoneIndian_Kerguelen, // 0x4351b389, Indian/Kerguelen -> Indian/Maldives
	&ZoneAmerica_Indiana_Winamac, // 0x4413fa69, America/Indiana/Winamac
	&ZoneEurope_Berlin, // 0x44644c20, Europe/Berlin
	&ZoneAtlantic_St_Helena, // 0x451fc5f7, Atlantic/St_Helena -> Africa/Abidjan
	&ZoneIndian_Chagos, // 0x456f7c3c, Indian/Chagos
	&ZoneIndian_Mahe, // 0x45e725e2, Indian/Mahe -> Asia/Dubai
	&ZoneIndian_Comoro, // 0x45f4deb6, Indian/Comoro -> Africa/Nairobi
	&ZoneAmerica_Mendoza, // 0x46b4e054, America/Mendoza -> America/Argentina/Mendoza
	&ZoneAsia_Ust_Nera, // 0x4785f921, Asia/Ust-Nera
	&ZoneEurope_Dublin, // 0x4a275f62, Europe/Dublin
	&ZoneAsia_Nicosia, // 0x4b0fcf78, Asia/Nicosia
	&ZoneAmerica_Chicago, // 0x4b92b5d4, America/Chicago
	&ZoneAustralia_Sydney, // 0x4d1e9776, Australia/Sydney
	&ZoneNZ_CHAT, // 0x4d42afda, NZ-CHAT -> Pacific/Chatham
	&ZoneUS_Arizona, // 0x4ec52670, US/Arizona -> America/Phoenix
	&ZoneAntarctica_Vostok, // 0x4f966fd4, Antarctica/Vostok -> Asia/Urumqi
	&ZoneUS_Aleutian, // 0x4fe013ef, US/Aleutian -> America/Adak
	&ZoneAustralia_Brisbane, // 0x4fedc9c0, Australia/Brisbane
	&ZoneAmerica_Catamarca, // 0x5036e963, America/Catamarca -> America/Argentina/Catamarca
	&ZoneAmerica_Asuncion, // 0x50ec79a6, America/Asuncion
	&ZoneAsia_Karachi, // 0x527f5245, Asia/Karachi
	&ZoneAsia_Kashgar, // 0x52955193, Asia/Kashgar -> Asia/Urumqi
	&ZoneCanada_Atlantic, // 0x536b119c, Canada/Atlantic -> America/Halifax
	&ZonePacific_Gambier, // 0x53720c3a, Pacific/Gambier
	&ZoneAmerica_Whitehorse, // 0x54e0e3e8, America/Whitehorse
	&ZoneAmerica_Martinique, // 0x551e84c5, America/Martinique
	&ZoneAmerica_Jamaica, // 0x565dad6c, America/Jamaica
	&ZoneUS_Samoa, // 0x566821cd, US/Samoa -> Pacific/Pago_Pago
	&ZoneHongkong, // 0x56d36560, Hongkong -> Asia/Hong_Kong
	&ZoneEurope_Jersey, // 0x570dae76, Europe/Jersey -> Europe/London
	&ZoneAsia_Hong_Kong, // 0x577f28ac, Asia/Hong_Kong
	&ZonePacific_Marquesas, // 0x57ca7135, Pacific/Marquesas
	&ZoneAmerica_Miquelon, // 0x59674330, America/Miquelon
	&ZoneAntarctica_DumontDUrville, // 0x5a3c656c, Antarctica/DumontDUrville -> Pacific/Port_Moresby
	&ZoneAtlantic_Jan_Mayen, // 0x5a7535b6, Atlantic/Jan_Mayen -> Europe/Berlin
	&ZoneAmerica_Anchorage, // 0x5a79260e, America/Anchorage
	&ZoneUS_Eastern, // 0x5bb7e78e, US/Eastern -> America/New_York
	&ZoneAsia_Jerusalem, // 0x5becd23a, Asia/Jerusalem
	&ZoneEurope_Stockholm, // 0x5bf6fbb8, Europe/Stockholm -> Europe/Berlin
	&ZoneEurope_Lisbon, // 0x5c00a70b, Europe/Lisbon
	&ZoneAtlantic_Cape_Verde, // 0x5c5e1772, Atlantic/Cape_Verde
	&ZoneEurope_London, // 0x5c6a84ae, Europe/London
	&ZoneAmerica_Cordoba, // 0x5c8a7600, America/Cordoba -> America/Argentina/Cordoba
	&ZoneAsia_Ujung_Pandang, // 0x5d001eb3, Asia/Ujung_Pandang -> Asia/Makassar
	&ZoneAfrica_Mbabane, // 0x5d3bdd40, Africa/Mbabane -> Africa/Johannesburg
	&ZoneEurope_Madrid, // 0x5dbd1535, Europe/Madrid
	&ZoneAmerica_Moncton, // 0x5e07fe24, America/Moncton
	&ZonePacific_Bougainville, // 0x5e10f7a4, Pacific/Bougainville
	&ZoneEurope_Monaco, // 0x5ebf9f01, Europe/Monaco -> Europe/Paris
	&ZoneEurope_Moscow, // 0x5ec266fc, Europe/Moscow
	&ZoneAmerica_Argentina_Jujuy, // 0x5f2f46c5, America/Argentina/Jujuy
	&ZoneAmerica_Argentina_Salta, // 0x5fc73403, America/Argentina/Salta
	&ZonePacific_Pago_Pago, // 0x603aebd0, Pacific/Pago_Pago
	&ZonePacific_Enderbury, // 0x61599a93, Pacific/Enderbury -> Pacific/Kanton
	&ZoneAfrica_Sao_Tome, // 0x61b319d1, Africa/Sao_Tome
	&ZoneCanada_Central, // 0x626710f5, Canada/Central -> America/Winnipeg
	&ZoneAmerica_Creston, // 0x62a70204, America/Creston -> America/Phoenix
	&ZoneAmerica_Costa_Rica, // 0x63ff66be, America/Costa_Rica
	&ZoneAsia_Qostanay, // 0x654fe522, Asia/Qostanay
	&ZoneAmerica_Indiana_Knox, // 0x6554adc9, America/Indiana/Knox
	&ZoneEurope_Prague, // 0x65ee5d48, Europe/Prague
	&ZoneBrazil_Acre, // 0x66934f93, Brazil/Acre -> America/Rio_Branco
	&ZoneBrazil_East, // 0x669578c5, Brazil/East -> America/Sao_Paulo
	&ZoneAfrica_Kinshasa, // 0x6695d70c, Africa/Kinshasa -> Africa/Lagos
	&ZoneBrazil_West, // 0x669f689b, Brazil/West -> America/Manaus
	&ZoneAfrica_Mogadishu, // 0x66bc159b, Africa/Mogadishu -> Africa/Nairobi
	&ZoneAmerica_Puerto_Rico, // 0x6752ca31, America/Puerto_Rico
	&ZoneUS_Indiana_Starke, // 0x67977be7, US/Indiana-Starke -> America/Indiana/Knox
	&ZoneAmerica_Buenos_Aires, // 0x67d79a05, America/Buenos_Aires -> America/Argentina/Buenos_Aires
	&ZoneAfrica_Freetown, // 0x6823dd64, Africa/Freetown -> Africa/Abidjan
	&ZoneIndian_Christmas, // 0x68c207d5, Indian/Christmas -> Asia/Bangkok
	&ZoneAsia_Novokuznetsk, // 0x69264f93, Asia/Novokuznetsk
	&ZoneAmerica_Indianapolis, // 0x6a009ae1, America/Indianapolis -> America/Indiana/Indianapolis
	&ZoneEurope_Sarajevo, // 0x6a576c3f, Europe/Sarajevo -> Europe/Belgrade
	&ZoneAmerica_Curacao, // 0x6a879184, America/Curacao -> America/Puerto_Rico
	&ZoneAmerica_Tijuana, // 0x6aa1df72, America/Tijuana
	&ZoneAsia_Makassar, // 0x6aa21c85, Asia/Makassar
	&ZoneEurope_Helsinki, // 0x6ab2975b, Europe/Helsinki
	&ZoneAmerica_Lower_Princes, // 0x6ae45b62, America/Lower_Princes -> America/Puerto_Rico
	&ZoneAmerica_Porto_Velho, // 0x6b1aac77, America/Porto_Velho
	&ZoneEurope_Samara, // 0x6bc0b139, Europe/Samara
	&ZoneEurope_Skopje, // 0x6c76fdd0, Europe/Skopje -> Europe/Belgrade
	&ZoneAmerica_Edmonton, // 0x6cb9484a, America/Edmonton
	&ZoneAmerica_Dawson_Creek, // 0x6cf24e5b, America/Dawson_Creek
	&ZoneAsia_Rangoon, // 0x6d1217c6, Asia/Rangoon -> Asia/Yangon
	&ZoneUS_East_Indiana, // 0x6dcf558a, US/East-Indiana -> America/Indiana/Indianapolis
	&ZoneAmerica_Grand_Turk, // 0x6e216197, America/Grand_Turk
	&ZoneAmerica_Blanc_Sablon, // 0x6e299892, America/Blanc-Sablon -> America/Puerto_Rico
	&ZoneEurope_Tirane, // 0x6ea95b47, Europe/Tirane
	&ZoneUS_Mountain, // 0x6eb88247, US/Mountain -> America/Denver
	&ZoneAntarctica_McMurdo, // 0x6eeb5585, Antarctica/McMurdo -> Pacific/Auckland
	&ZoneAmerica_Araguaina, // 0x6f9a3aef, America/Araguaina
	&ZoneAfrica_Lubumbashi, // 0x6fd88566, Africa/Lubumbashi -> Africa/Maputo
	&ZoneIndian_Reunion, // 0x7076c047, Indian/Reunion -> Asia/Dubai
	&ZoneAsia_Qyzylorda, // 0x71282e81, Asia/Qyzylorda
	&ZoneAsia_Kolkata, // 0x72c06cd9, Asia/Kolkata
	&ZoneAmerica_Ciudad_Juarez, // 0x7347fc60, America/Ciudad_Juarez
	&ZoneEurope_Vienna, // 0x734cc2e5, Europe/Vienna
	&ZoneAfrica_Asmara, // 0x73b278ef, Africa/Asmara -> Africa/Nairobi
	&ZoneAfrica_Asmera, // 0x73b289f3, Africa/Asmera -> Africa/Nairobi
	&ZoneAsia_Kamchatka, // 0x73baf9d7, Asia/Kamchatka
	&ZoneAmerica_Santarem, // 0x740caec1, America/Santarem
	&ZoneAmerica_Santiago, // 0x7410c9bc, America/Santiago
	&ZoneAfrica_Bamako, // 0x74c1e7a5, Africa/Bamako -> Africa/Abidjan
	&ZoneAfrica_Bangui, // 0x74c28ed0, Africa/Bangui -> Africa/Lagos
	&ZoneAfrica_Banjul, // 0x74c29b96, Africa/Banjul -> Africa/Abidjan
	&ZoneEurope_Nicosia, // 0x74efab8a, Europe/Nicosia -> Asia/Nicosia
	&ZoneEurope_Warsaw, // 0x75185c19, Europe/Warsaw
	&ZoneAmerica_El_Salvador, // 0x752ad652, America/El_Salvador
	&ZoneAfrica_Bissau, // 0x75564141, Africa/Bissau
	&ZoneAmerica_Santo_Domingo, // 0x75a0d177, America/Santo_Domingo
	&ZoneUS_Michigan, // 0x766bb7bc, US/Michigan -> America/Detroit
	&ZoneCanada_Saskatchewan, // 0x77311f49, Canada/Saskatchewan -> America/Regina
	&ZoneAfrica_Accra, // 0x77d5b054, Africa/Accra -> Africa/Abidjan
	&ZoneAfrica_Cairo, // 0x77f8e228, Africa/Cairo
	&ZoneAfrica_Ceuta, // 0x77fb46ec, Africa/Ceuta
	&ZoneAfrica_Dakar, // 0x780b00fd, Africa/Dakar -> Africa/Abidjan
	&ZoneAfrica_Lagos, // 0x789bb5d0, Africa/Lagos
	&ZoneAfrica_Windhoek, // 0x789c9bd3, Africa/Windhoek
	&ZoneCanada_Yukon, // 0x78dd35c2, Canada/Yukon -> America/Whitehorse
	&ZoneAmerica_Toronto, // 0x792e851b, America/Toronto
	&ZoneAmerica_Tortola, // 0x7931462b, America/Tortola -> America/Puerto_Rico
	&ZoneAfrica_Tunis, // 0x79378e6d, Africa/Tunis
	&ZoneAfrica_Douala, // 0x7a6df310, Africa/Douala -> Africa/Lagos
	&ZoneAfrica_Conakry, // 0x7ab36b31, Africa/Conakry -> Africa/Abidjan
	&ZoneIndian_Mauritius, // 0x7b09c02a, Indian/Mauritius
	&ZoneAtlantic_Stanley, // 0x7bb3e1c4, Atlantic/Stanley
	&ZoneAmerica_Ensenada, // 0x7bc95445, America/Ensenada -> America/Tijuana
	&ZoneEurope_Zagreb, // 0x7c11c9ff, Europe/Zagreb -> Europe/Belgrade
	&ZoneCuba, // 0x7c83cba0, Cuba -> America/Havana
	&ZoneEire, // 0x7c84b36a, Eire -> Europe/Dublin
	&ZoneGMT0, // 0x7c8550fd, GMT0 -> Etc/GMT
	&ZoneIran, // 0x7c87090f, Iran -> Asia/Tehran
	&ZoneW_SU, // 0x7c8d8ef1, W-SU -> Europe/Moscow
	&ZoneZulu, // 0x7c9069b5, Zulu -> Etc/UTC
	&ZoneEurope_Zurich, // 0x7d8195b9, Europe/Zurich
	&ZoneChile_Continental, // 0x7e2bdb18, Chile/Continental -> America/Santiago
	&ZoneAmerica_Fort_Wayne, // 0x7eaaaf24, America/Fort_Wayne -> America/Indiana/Indianapolis
	&ZoneAsia_Kuching, // 0x801b003b, Asia/Kuching
	&ZoneAtlantic_Madeira, // 0x81b5c037, Atlantic/Madeira
	&ZoneAmerica_Atikokan, // 0x81b92098, America/Atikokan -> America/Panama
	&ZoneAfrica_Harare, // 0x82c39a2d, Africa/Harare -> Africa/Maputo
	&ZoneAmerica_Shiprock, // 0x82fb7049, America/Shiprock -> America/Denver
	&ZonePacific_Kiritimati, // 0x8305073a, Pacific/Kiritimati
	&ZoneAmerica_St_Vincent, // 0x8460e523, America/St_Vincent -> America/Puerto_Rico
	&ZoneAmerica_Metlakatla, // 0x84de2686, America/Metlakatla
	&ZoneAsia_Yakutsk, // 0x87bb3a9e, Asia/Yakutsk
	&ZoneAmerica_Chihuahua, // 0x8827d776, America/Chihuahua
	&ZonePacific_Pitcairn, // 0x8837d8bd, Pacific/Pitcairn
	&ZoneAsia_Vientiane, // 0x89d68d75, Asia/Vientiane -> Asia/Bangkok
	&ZonePacific_Chuuk, // 0x8a090b23, Pacific/Chuuk -> Pacific/Port_Moresby
	&ZonePacific_Efate, // 0x8a2bce28, Pacific/Efate
	&ZoneAfrica_Kigali, // 0x8a4dcf2b, Africa/Kigali -> Africa/Maputo
	&ZoneAustralia_ACT, // 0x8a970eb2, Australia/ACT -> Australia/Sydney
	&ZoneAustralia_LHI, // 0x8a973e17, Australia/LHI -> Australia/Lord_Howe
	&ZoneAustralia_NSW, // 0x8a974812, Australia/NSW -> Australia/Sydney
	&ZonePacific_Nauru, // 0x8acc41ae, Pacific/Nauru
	&ZoneEST5EDT, // 0x8adc72a3, EST5EDT
	&ZonePacific_Palau, // 0x8af04a36, Pacific/Palau
	&ZonePacific_Samoa, // 0x8b2699b4, Pacific/Samoa -> Pacific/Pago_Pago
	&ZoneAmerica_Winnipeg, // 0x8c7dafc7, America/Winnipeg
	&ZoneAustralia_Eucla, // 0x8cf99e44, Australia/Eucla
	&ZoneAmerica_Argentina_Catamarca, // 0x8d40986b, America/Argentina/Catamarca
	&ZoneAfrica_Luanda, // 0x8d7909cf, Africa/Luanda -> Africa/Lagos
	&ZoneAfrica_Lusaka, // 0x8d82b23b, Africa/Lusaka -> Africa/Maputo
	&ZoneAustralia_North, // 0x8d997165, Australia/North -> Australia/Darwin
	&ZoneAustralia_Perth, // 0x8db8269d, Australia/Perth
	&ZoneAustralia_South, // 0x8df3f8ad, Australia/South -> Australia/Adelaide
	&ZonePacific_Kwajalein, // 0x8e216759, Pacific/Kwajalein
	&ZoneAmerica_Port_au_Prince, // 0x8e4a7bdc, America/Port-au-Prince
	&ZoneAfrica_Malabo, // 0x8e6a1906, Africa/Malabo -> Africa/Lagos
	&ZoneAfrica_Maputo, // 0x8e6ca1f0, Africa/Maputo
	&ZoneAfrica_Maseru, // 0x8e6e02c7, Africa/Maseru -> Africa/Johannesburg
	&ZonePacific_Norfolk, // 0x8f4eb4be, Pacific/Norfolk
	&ZoneAmerica_Godthab, // 0x8f7eba1f, America/Godthab -> America/Nuuk
	&ZoneAustralia_Yancowinna, // 0x90bac131, Australia/Yancowinna -> Australia/Broken_Hill
	&ZoneAfrica_Niamey, // 0x914a30fd, Africa/Niamey -> Africa/Lagos
	&ZoneAsia_Yerevan, // 0x9185c8cc, Asia/Yerevan
	&ZoneAmerica_Detroit, // 0x925cfbc1, America/Detroit
	&ZoneAsia_Choibalsan, // 0x928aa4a6, Asia/Choibalsan
	&ZoneAntarctica_Macquarie, // 0x92f47626, Antarctica/Macquarie
	&ZoneAmerica_Belize, // 0x93256c81, America/Belize
	&ZoneMexico_General, // 0x93711d57, Mexico/General -> America/Mexico_City
	&ZoneAmerica_Bogota, // 0x93d7bc62, America/Bogota
	&ZoneAsia_Pyongyang, // 0x93ed1c8e, Asia/Pyongyang
	&ZoneAmerica_Indiana_Petersburg, // 0x94ac7acc, America/Indiana/Petersburg
	&ZoneAmerica_Cancun, // 0x953331be, America/Cancun
	&ZoneAmerica_Cayman, // 0x953961df, America/Cayman -> America/Panama
	&ZoneAmerica_Glace_Bay, // 0x9681f8dd, America/Glace_Bay
	&ZoneAsia_Khandyga, // 0x9685a4d9, Asia/Khandyga
	&ZoneAmerica_Grenada, // 0x968ce4d8, America/Grenada -> America/Puerto_Rico
	&ZoneAmerica_Cuiaba, // 0x969a52eb, America/Cuiaba
	&ZoneAmerica_Dawson, // 0x978d8d12, America/Dawson
	&ZoneAmerica_Aruba, // 0x97cf8651, America/Aruba -> America/Puerto_Rico
	&ZoneAmerica_Denver, // 0x97d10b2a, America/Denver
	&ZoneAmerica_Bahia, // 0x97d815fb, America/Bahia
	&ZoneAmerica_Belem, // 0x97da580b, America/Belem
	&ZoneAmerica_Boise, // 0x97dfc8d8, America/Boise
	&ZoneEurope_Andorra, // 0x97f6764b, Europe/Andorra
	&ZoneAmerica_Adak, // 0x97fe49d7, America/Adak
	&ZoneAmerica_Atka, // 0x97fe8f27, America/Atka -> America/Adak
	&ZoneAmerica_Lima, // 0x980468c9, America/Lima
	&ZoneAmerica_Nome, // 0x98059b15, America/Nome
	&ZoneAmerica_Nuuk, // 0x9805b5a9, America/Nuuk
	&ZoneIndian_Maldives, // 0x9869681c, Indian/Maldives
	&ZoneAmerica_Jujuy, // 0x9873dbbd, America/Jujuy -> America/Argentina/Jujuy
	&ZoneAmerica_Sitka, // 0x99104ce2, America/Sitka
	&ZoneAmerica_Thule, // 0x9921dd68, America/Thule
	&ZonePacific_Rarotonga, // 0x9981a3b0, Pacific/Rarotonga
	&ZoneAsia_Kathmandu, // 0x9a96ce6f, Asia/Kathmandu
	&ZoneBrazil_DeNoronha, // 0x9b4cb496, Brazil/DeNoronha -> America/Noronha
	&ZoneAmerica_North_Dakota_Beulah, // 0x9b52b384, America/North_Dakota/Beulah
	&ZoneAmerica_Rainy_River, // 0x9cd58a10, America/Rainy_River -> America/Winnipeg
	&ZoneEurope_Budapest, // 0x9ce0197c, Europe/Budapest
	&ZoneAsia_Baghdad, // 0x9ceffbed, Asia/Baghdad
	&ZoneAsia_Bahrain, // 0x9d078487, Asia/Bahrain -> Asia/Qatar
	&ZoneEtc_GMT_PLUS_0, // 0x9d13da13, Etc/GMT+0 -> Etc/GMT
	&ZoneEtc_GMT_PLUS_1, // 0x9d13da14, Etc/GMT+1
	&ZoneEtc_GMT_PLUS_2, // 0x9d13da15, Etc/GMT+2
	&ZoneEtc_GMT_PLUS_3, // 0x9d13da16, Etc/GMT+3
	&ZoneEtc_GMT_PLUS_4, // 0x9d13da17, Etc/GMT+4
	&ZoneEtc_GMT_PLUS_5, // 0x9d13da18, Etc/GMT+5
	&ZoneEtc_GMT_PLUS_6, // 0x9d13da19, Etc/GMT+6
	&ZoneEtc_GMT_PLUS_7, // 0x9d13da1a, Etc/GMT+7
	&ZoneEtc_GMT_PLUS_8, // 0x9d13da1b, Etc/GMT+8
	&ZoneEtc_GMT_PLUS_9, // 0x9d13da1c, Etc/GMT+9
	&ZoneEtc_GMT_0, // 0x9d13da55, Etc/GMT-0 -> Etc/GMT
	&ZoneEtc_GMT_1, // 0x9d13da56, Etc/GMT-1
	&ZoneEtc_GMT_2, // 0x9d13da57, Etc/GMT-2
	&ZoneEtc_GMT_3, // 0x9d13da58, Etc/GMT-3
	&ZoneEtc_GMT_4, // 0x9d13da59, Etc/GMT-4
	&ZoneEtc_GMT_5, // 0x9d13da5a, Etc/GMT-5
	&ZoneEtc_GMT_6, // 0x9d13da5b, Etc/GMT-6
	&ZoneEtc_GMT_7, // 0x9d13da5c, Etc/GMT-7
	&ZoneEtc_GMT_8, // 0x9d13da5d, Etc/GMT-8
	&ZoneEtc_GMT_9, // 0x9d13da5e, Etc/GMT-9
	&ZoneAmerica_Nipigon, // 0x9d2a8b1a, America/Nipigon -> America/Toronto
	&ZoneAmerica_Rio_Branco, // 0x9d352764, America/Rio_Branco
	&ZoneAsia_Bangkok, // 0x9d6e3aaf, Asia/Bangkok
	&ZoneAfrica_El_Aaiun, // 0x9d6fb118, Africa/El_Aaiun
	&ZoneAmerica_North_Dakota_Center, // 0x9da42814, America/North_Dakota/Center
	&ZoneAsia_Barnaul, // 0x9dba4997, Asia/Barnaul
	&ZoneAfrica_Tripoli, // 0x9dfebd3d, Africa/Tripoli
	&ZoneEurope_Istanbul, // 0x9e09d6e6, Europe/Istanbul
	&ZoneIndian_Antananarivo, // 0x9ebf5289, Indian/Antananarivo -> Africa/Nairobi
	&ZoneAfrica_Ndjamena, // 0x9fe09898, Africa/Ndjamena
	&ZoneAmerica_Guyana, // 0x9ff7bd0b, America/Guyana
	&ZoneAfrica_Dar_es_Salaam, // 0xa04c47b6, Africa/Dar_es_Salaam -> Africa/Nairobi
	&ZoneAmerica_Havana, // 0xa0e15675, America/Havana
	&ZoneAsia_Novosibirsk, // 0xa2a435cb, Asia/Novosibirsk
	&ZoneEurope_Kiev, // 0xa2c19eb3, Europe/Kiev -> Europe/Kyiv
	&ZoneEurope_Kyiv, // 0xa2c1e347, Europe/Kyiv
	&ZoneEurope_Oslo, // 0xa2c3fba1, Europe/Oslo -> Europe/Berlin
	&ZoneEurope_Riga, // 0xa2c57587, Europe/Riga
	&ZoneEurope_Rome, // 0xa2c58fd7, Europe/Rome
	&ZoneAmerica_Inuvik, // 0xa42189fc, America/Inuvik
	&ZoneAmerica_Argentina_La_Rioja, // 0xa46b7eef, America/Argentina/La_Rioja
	&ZoneAsia_Almaty, // 0xa61f41fa, Asia/Almaty
	&ZoneAsia_Anadyr, // 0xa63cebd1, Asia/Anadyr
	&ZoneAsia_Aqtobe, // 0xa67dcc4e, Asia/Aqtobe
	&ZoneAsia_Atyrau, // 0xa6b6e068, Asia/Atyrau
	&ZoneAmerica_Juneau, // 0xa6f13e2e, America/Juneau
	&ZoneAustralia_Lord_Howe, // 0xa748b67d, Australia/Lord_Howe
	&ZonePacific_Port_Moresby, // 0xa7ba7f68, Pacific/Port_Moresby
	&ZoneAsia_Katmandu, // 0xa7ec12c7, Asia/Katmandu -> Asia/Kathmandu
	&ZoneAsia_Beirut, // 0xa7f3d5fd, Asia/Beirut
	&ZoneSingapore, // 0xa8598c8d, Singapore -> Asia/Singapore
	&ZoneAfrica_Nairobi, // 0xa87ab57e, Africa/Nairobi
	&ZoneAsia_Brunei, // 0xa8e595f7, Asia/Brunei -> Asia/Kuching
	&ZoneUS_Pacific, // 0xa950f6ab, US/Pacific -> America/Los_Angeles
	&ZonePacific_Galapagos, // 0xa952f752, Pacific/Galapagos
	&ZoneAmerica_Argentina_Mendoza, // 0xa9f72d5c, America/Argentina/Mendoza
	&ZoneAmerica_La_Paz, // 0xaa29125d, America/La_Paz
	&ZoneAmerica_Noronha, // 0xab5116fb, America/Noronha
	&ZoneAmerica_Coral_Harbour, // 0xabcb7569, America/Coral_Harbour -> America/Panama
	&ZoneAmerica_Maceio, // 0xac80c6d4, America/Maceio
	&ZoneAmerica_Manaus, // 0xac86bf8b, America/Manaus
	&ZoneAmerica_Merida, // 0xacd172d8, America/Merida
	&ZoneEurope_Chisinau, // 0xad58aa18, Europe/Chisinau
	&ZoneAmerica_Nassau, // 0xaedef011, America/Nassau -> America/Toronto
	&ZoneAmerica_Anguilla, // 0xafe31333, America/Anguilla -> America/Puerto_Rico
	&ZoneEurope_Uzhgorod, // 0xb066f5d6, Europe/Uzhgorod -> Europe/Kyiv
	&ZoneAustralia_Broken_Hill, // 0xb06eada3, Australia/Broken_Hill
	&ZoneAsia_Bishkek, // 0xb0728553, Asia/Bishkek
	&ZoneChile_EasterIsland, // 0xb0982af8, Chile/EasterIsland -> Pacific/Easter
	&ZonePacific_Johnston, // 0xb15d7b36, Pacific/Johnston -> Pacific/Honolulu
	&ZoneAfrica_Timbuktu, // 0xb164d56f, Africa/Timbuktu -> Africa/Abidjan
	&ZoneAmerica_St_Thomas, // 0xb1b3d778, America/St_Thomas -> America/Puerto_Rico
	&ZoneAmerica_Paramaribo, // 0xb319e4c4, America/Paramaribo
	&ZoneAmerica_Panama, // 0xb3863854, America/Panama
	&ZoneCanada_Newfoundland, // 0xb396e991, Canada/Newfoundland -> America/St_Johns
	&ZoneAsia_Harbin, // 0xb5af1186, Asia/Harbin -> Asia/Shanghai
	&ZoneAsia_Hebron, // 0xb5eef250, Asia/Hebron
	&ZoneAmerica_Goose_Bay, // 0xb649541e, America/Goose_Bay
	&ZoneAmerica_Los_Angeles, // 0xb7f7e8f2, America/Los_Angeles
	&ZoneAmerica_Recife, // 0xb8730494, America/Recife
	&ZoneAmerica_Regina, // 0xb875371c, America/Regina
	&ZoneAsia_Ashgabat, // 0xba87598d, Asia/Ashgabat
	&ZoneIsrael, // 0xba88c9e5, Israel -> Asia/Jerusalem
	&ZonePacific_Yap, // 0xbb40138d, Pacific/Yap -> Pacific/Port_Moresby
	&ZoneAmerica_Halifax, // 0xbc5b7183, America/Halifax
	&ZoneEurope_Ljubljana, // 0xbd98cdb7, Europe/Ljubljana -> Europe/Belgrade
	&ZoneAsia_Kuwait, // 0xbe1b2f27, Asia/Kuwait -> Asia/Riyadh
	&ZoneEurope_Tiraspol, // 0xbe704472, Europe/Tiraspol -> Europe/Chisinau
	&ZoneAsia_Srednekolymsk, // 0xbf8e337d, Asia/Srednekolymsk
	&ZoneAmerica_Argentina_Cordoba, // 0xbfccc308, America/Argentina/Cordoba
	&ZoneAmerica_Tegucigalpa, // 0xbfd6fd4c, America/Tegucigalpa
	&ZoneAmerica_Antigua, // 0xc067a32f, America/Antigua -> America/Puerto_Rico
	&ZoneEurope_Busingen, // 0xc06d2cdf, Europe/Busingen -> Europe/Zurich
	&ZoneAsia_Manila, // 0xc156c944, Asia/Manila
	&ZoneAfrica_Kampala, // 0xc1d30e31, Africa/Kampala -> Africa/Nairobi
	&ZoneAmerica_Knox_IN, // 0xc1db9a1c, America/Knox_IN -> America/Indiana/Knox
	&ZoneAfrica_Abidjan, // 0xc21305a3, Africa/Abidjan
	&ZoneAmerica_Virgin, // 0xc2183ab5, America/Virgin -> America/Puerto_Rico
	&ZoneAsia_Phnom_Penh, // 0xc224945e, Asia/Phnom_Penh -> Asia/Bangkok
	&ZoneAsia_Muscat, // 0xc2c3565f, Asia/Muscat -> Asia/Dubai
	&ZoneAmerica_Punta_Arenas, // 0xc2c3bce7, America/Punta_Arenas
	&ZonePortugal, // 0xc3274593, Portugal -> Europe/Lisbon
	&ZoneNavajo, // 0xc4ef0e24, Navajo -> America/Denver
	&ZoneAfrica_Casablanca, // 0xc59f1b33, Africa/Casablanca
	&ZoneAmerica_Argentina_Rio_Gallegos, // 0xc5b0f565, America/Argentina/Rio_Gallegos
	&ZoneAsia_Jayapura, // 0xc6833c2f, Asia/Jayapura
	&ZoneAmerica_Resolute, // 0xc7093459, America/Resolute
	&ZoneAsia_Chungking, // 0xc7121dd0, Asia/Chungking -> Asia/Shanghai
	&ZoneGreenwich, // 0xc84d4221, Greenwich -> Etc/GMT
	&ZoneAmerica_Rankin_Inlet, // 0xc8de4984, America/Rankin_Inlet
	&ZonePoland, // 0xca913b23, Poland -> Europe/Warsaw
	&ZoneUS_Central, // 0xcabdcb25, US/Central -> America/Chicago
	&ZoneEurope_Vatican, // 0xcb485dca, Europe/Vatican -> Europe/Rome
	&ZoneAmerica_Barbados, // 0xcbbc3b04, America/Barbados
	&ZoneAmerica_Porto_Acre, // 0xcce5bf54, America/Porto_Acre -> America/Rio_Branco
	&ZoneAmerica_Guadeloupe, // 0xcd1f8a31, America/Guadeloupe -> America/Puerto_Rico
	&ZoneAntarctica_South_Pole, // 0xcd96b290, Antarctica/South_Pole -> Pacific/Auckland
	&ZoneAsia_Riyadh, // 0xcd973d93, Asia/Riyadh
	&ZoneAmerica_Dominica, // 0xcecb4c4a, America/Dominica -> America/Puerto_Rico
	&ZoneEurope_San_Marino, // 0xcef7724b, Europe/San_Marino -> Europe/Rome
	&ZoneAsia_Saigon, // 0xcf52f713, Asia/Saigon -> Asia/Ho_Chi_Minh
	&ZonePacific_Easter, // 0xcf54f7e7, Pacific/Easter
	&ZoneAsia_Singapore, // 0xcf8581fa, Asia/Singapore
	&ZoneAsia_Krasnoyarsk, // 0xd0376c6a, Asia/Krasnoyarsk
	&ZoneEurope_Belfast, // 0xd07dd1e5, Europe/Belfast -> Europe/London
	&ZoneAmerica_Mexico_City, // 0xd0d93f43, America/Mexico_City
	&ZoneUniversal, // 0xd0ff523e, Universal -> Etc/UTC
	&ZoneAsia_Taipei, // 0xd1a844ae, Asia/Taipei
	&ZoneAsia_Tehran, // 0xd1f02254, Asia/Tehran
	&ZoneAsia_Thimbu, // 0xd226e31b, Asia/Thimbu -> Asia/Thimphu
	&ZoneArctic_Longyearbyen, // 0xd23e7859, Arctic/Longyearbyen -> Europe/Berlin
	&ZoneAustralia_Queensland, // 0xd326ed0a, Australia/Queensland -> Australia/Brisbane
	&ZoneEurope_Kaliningrad, // 0xd33b2f28, Europe/Kaliningrad
	&ZoneAmerica_Argentina_Buenos_Aires, // 0xd43b4c0d, America/Argentina/Buenos_Aires
	&ZoneTurkey, // 0xd455e469, Turkey -> Europe/Istanbul
	&ZoneAfrica_Juba, // 0xd51b395c, Africa/Juba
	&ZoneAfrica_Lome, // 0xd51c3a07, Africa/Lome -> Africa/Abidjan
	&ZoneAsia_Urumqi, // 0xd5379735, Asia/Urumqi
	&ZoneAmerica_Cambridge_Bay, // 0xd5a44aff, America/Cambridge_Bay
	&ZoneAfrica_Johannesburg, // 0xd5d157a0, Africa/Johannesburg
	&ZoneAmerica_Port_of_Spain, // 0xd8b28d59, America/Port_of_Spain -> America/Puerto_Rico
	&ZoneEtc_GMT, // 0xd8e2de58, Etc/GMT
	&ZoneEtc_UCT, // 0xd8e3189c, Etc/UCT -> Etc/UTC
	&ZoneEtc_UTC, // 0xd8e31abc, Etc/UTC
	&ZoneAmerica_Yakutat, // 0xd8ee31e9, America/Yakutat
	&ZoneAfrica_Algiers, // 0xd94515c1, Africa/Algiers
	&ZonePST8PDT, // 0xd99ee2dc, PST8PDT
	&ZoneEurope_Bratislava, // 0xda493bed, Europe/Bratislava -> Europe/Prague
	&ZoneEurope_Simferopol, // 0xda9eb724, Europe/Simferopol
	&ZonePacific_Funafuti, // 0xdb402d65, Pacific/Funafuti -> Pacific/Tarawa
	&ZoneAmerica_Matamoros, // 0xdd1b0259, America/Matamoros
	&ZonePacific_Kanton, // 0xdd512f0e, Pacific/Kanton
	&ZoneAsia_Yangon, // 0xdd54a8be, Asia/Yangon
	&ZoneEurope_Vilnius, // 0xdd63b8ce, Europe/Vilnius
	&ZoneAustralia_West, // 0xdd858a5d, Australia/West -> Australia/Perth
	&ZonePacific_Kosrae, // 0xde5139a8, Pacific/Kosrae
	&ZoneAmerica_Kentucky_Monticello, // 0xde71c439, America/Kentucky/Monticello
	&ZoneEurope_Brussels, // 0xdee07337, Europe/Brussels
	&ZoneAmerica_Swift_Current, // 0xdef98e55, America/Swift_Current
	&ZoneAmerica_Rosario, // 0xdf448665, America/Rosario -> America/Argentina/Cordoba
	&ZoneAsia_Irkutsk, // 0xdfbf213f, Asia/Irkutsk
	&ZoneEurope_Ulyanovsk, // 0xe03783d0, Europe/Ulyanovsk
	&ZoneAustralia_Lindeman, // 0xe05029e2, Australia/Lindeman
	&ZoneEurope_Belgrade, // 0xe0532b3a, Europe/Belgrade
	&ZoneAfrica_Blantyre, // 0xe08d813b, Africa/Blantyre -> Africa/Maputo
	&ZoneAmerica_Menominee, // 0xe0e9c583, America/Menominee
	&ZoneEurope_Copenhagen, // 0xe0ed30bc, Europe/Copenhagen -> Europe/Berlin
	&ZoneAtlantic_Faroe, // 0xe110a971, Atlantic/Faroe
	&ZonePacific_Majuro, // 0xe1f95371, Pacific/Majuro -> Pacific/Tarawa
	&ZoneAntarctica_Casey, // 0xe2022583, Antarctica/Casey
	&ZoneAntarctica_Davis, // 0xe2144b45, Antarctica/Davis
	&ZoneEurope_Astrakhan, // 0xe22256e1, Europe/Astrakhan
	&ZonePacific_Midway, // 0xe286d38e, Pacific/Midway -> Pacific/Pago_Pago
	&ZoneAntarctica_Syowa, // 0xe330c7e1, Antarctica/Syowa -> Asia/Riyadh
	&ZoneAntarctica_Troll, // 0xe33f085b, Antarctica/Troll
	&ZoneEurope_Saratov, // 0xe4315da4, Europe/Saratov
	&ZonePacific_Noumea, // 0xe551b788, Pacific/Noumea
	&ZoneIceland, // 0xe56a35b5, Iceland -> Africa/Abidjan
	&ZoneIndian_Mayotte, // 0xe6880bca, Indian/Mayotte -> Africa/Nairobi
	&ZoneAustralia_Tasmania, // 0xe6d76648, Australia/Tasmania -> Australia/Hobart
	&ZonePacific_Honolulu, // 0xe6e70af9, Pacific/Honolulu
	&ZoneAmerica_Kralendijk, // 0xe7c456c5, America/Kralendijk -> America/Puerto_Rico
	&ZoneAmerica_Argentina_Tucuman, // 0xe96399eb, America/Argentina/Tucuman
	&ZonePacific_Ponape, // 0xe9f80086, Pacific/Ponape -> Pacific/Guadalcanal
	&ZoneEurope_Zaporozhye, // 0xeab9767f, Europe/Zaporozhye -> Europe/Kyiv
	&ZoneEurope_Isle_of_Man, // 0xeaf84580, Europe/Isle_of_Man -> Europe/London
	&ZoneAsia_Magadan, // 0xebacc19b, Asia/Magadan
	&ZoneAmerica_Ojinaga, // 0xebfde83f, America/Ojinaga
	&ZonePacific_Saipan, // 0xeff7a35f, Pacific/Saipan -> Pacific/Guam
	&ZoneCST6CDT, // 0xf0e87d00, CST6CDT
	&ZonePacific_Tahiti, // 0xf24c2446, Pacific/Tahiti
	&ZonePacific_Tarawa, // 0xf2517e63, Pacific/Tarawa
	&ZoneMST7MDT, // 0xf2af9375, MST7MDT
	&ZoneCanada_Eastern, // 0xf3612d5e, Canada/Eastern -> America/Toronto
	&ZoneAsia_Tashkent, // 0xf3924254, Asia/Tashkent
	&ZoneAsia_Sakhalin, // 0xf4a1c9bd, Asia/Sakhalin
	&ZonePacific_Guadalcanal, // 0xf4dd25f0, Pacific/Guadalcanal
	&ZoneEtc_GMT0, // 0xf53ea988, Etc/GMT0 -> Etc/GMT
	&ZoneEtc_Zulu, // 0xf549c240, Etc/Zulu -> Etc/UTC
	&ZoneAmerica_Danmarkshavn, // 0xf554d204, America/Danmarkshavn
	&ZoneAsia_Shanghai, // 0xf895a7f5, Asia/Shanghai
	&ZoneEurope_Gibraltar, // 0xf8e325fc, Europe/Gibraltar
	&ZoneAsia_Chongqing, // 0xf937fb90, Asia/Chongqing -> Asia/Shanghai
	&ZoneAtlantic_Azores, // 0xf93ed918, Atlantic/Azores
	&ZonePacific_Wallis, // 0xf94ddb0f, Pacific/Wallis -> Pacific/Tarawa
	&ZoneAmerica_Thunder_Bay, // 0xf962e71b, America/Thunder_Bay -> America/Toronto
	&ZoneAmerica_Eirunepe, // 0xf9b29683, America/Eirunepe
	&ZoneAmerica_Montevideo, // 0xfa214780, America/Montevideo
	&ZoneUS_Alaska, // 0xfa300bc9, US/Alaska -> America/Anchorage
	&ZoneGB_Eire, // 0xfa70e300, GB-Eire -> Europe/London
	&ZoneEurope_Kirov, // 0xfaf5abef, Europe/Kirov
	&ZoneEurope_Malta, // 0xfb1560f3, Europe/Malta
	&ZoneEurope_Minsk, // 0xfb19cc66, Europe/Minsk
	&ZoneEurope_Bucharest, // 0xfb349ec5, Europe/Bucharest
	&ZoneAfrica_Khartoum, // 0xfb3d4205, Africa/Khartoum
	&ZoneEurope_Paris, // 0xfb4bc2a3, Europe/Paris
	&ZoneAsia_Yekaterinburg, // 0xfb544c6e, Asia/Yekaterinburg
	&ZoneEurope_Sofia, // 0xfb898656, Europe/Sofia
	&ZoneCanada_Mountain, // 0xfb8a8217, Canada/Mountain -> America/Edmonton
	&ZoneEurope_Vaduz, // 0xfbb81bae, Europe/Vaduz -> Europe/Zurich
	&ZoneAtlantic_Canary, // 0xfc23f2c2, Atlantic/Canary
	&ZoneMexico_BajaNorte, // 0xfcf7150f, Mexico/BajaNorte -> America/Tijuana
	&ZoneAmerica_Santa_Isabel, // 0xfd18a56c, America/Santa_Isabel -> America/Tijuana
	&ZoneAmerica_Campo_Grande, // 0xfec3e7a6, America/Campo_Grande

}
