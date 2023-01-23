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
// Total: 4 (3 zones, 1 links)
// ---------------------------------------------------------------------------

var ZoneAndLinkRegistry = []*zoneinfo.ZoneInfo{
	&ZoneAmerica_New_York, // 0x1e2a7654, America/New_York
	&ZoneUS_Pacific, // 0xa950f6ab, US/Pacific -> America/Los_Angeles
	&ZoneAmerica_Los_Angeles, // 0xb7f7e8f2, America/Los_Angeles
	&ZoneEtc_UTC, // 0xd8e31abc, Etc/UTC

}

// ---------------------------------------------------------------------------
// Zone IDs
// Total: 4 (3 zones, 1 links)
// ---------------------------------------------------------------------------

const (
	ZoneIDAmerica_Los_Angeles uint32 = 0xb7f7e8f2 // America/Los_Angeles
	ZoneIDAmerica_New_York uint32 = 0x1e2a7654 // America/New_York
	ZoneIDEtc_UTC uint32 = 0xd8e31abc // Etc/UTC
	ZoneIDUS_Pacific uint32 = 0xa950f6ab // US/Pacific

)
