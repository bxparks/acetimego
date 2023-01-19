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

// Supported Zones: 3
var ZoneRegistry = map[string]*zoneinfo.ZoneInfo{
	"America/Los_Angeles": &ZoneAmerica_Los_Angeles,
	"America/New_York": &ZoneAmerica_New_York,
	"Etc/UTC": &ZoneEtc_UTC,

}

// Supported Zones and Links: 4
var ZoneAndLinkRegistry = map[string]*zoneinfo.ZoneInfo{
	"America/Los_Angeles": &ZoneAmerica_Los_Angeles,
	"America/New_York": &ZoneAmerica_New_York,
	"Etc/UTC": &ZoneEtc_UTC,
	"US/Pacific": &ZoneUS_Pacific,

}