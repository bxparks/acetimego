package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

type ZoneRegistry = map[uint32]*zoneinfo.ZoneInfo

//-----------------------------------------------------------------------------
// ZoneRegistrar provides lookup functions for a given zone registry.
//-----------------------------------------------------------------------------

type ZoneRegistrar struct {
	Registry ZoneRegistry
}

func (zr *ZoneRegistrar) FindZoneInfoByID(id uint32) *zoneinfo.ZoneInfo {
	return zr.Registry[id]
}

func (zr *ZoneRegistrar) FindZoneInfoByName(name string) *zoneinfo.ZoneInfo {
	return nil
}
