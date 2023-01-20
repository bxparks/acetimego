package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

//-----------------------------------------------------------------------------
// ZoneRegistrar provides lookup functions for a given zone registry. The
// internal implementation of the zone registry is encapsulation by the
// registrar. The initial implementation is a map[uint32]*ZoneInfo. However, we
// may be able to save significant memory by using an array of []*ZoneInfo, and
// using a binary search on zoneId, just like the AceTime and AceTimeC
// libraries.
//-----------------------------------------------------------------------------

type ZoneRegistry = map[uint32]*zoneinfo.ZoneInfo

type ZoneRegistrar struct {
	Registry ZoneRegistry
}

func (zr *ZoneRegistrar) FindZoneInfoByID(id uint32) *zoneinfo.ZoneInfo {
	return zr.Registry[id]
}

func (zr *ZoneRegistrar) FindZoneInfoByName(name string) *zoneinfo.ZoneInfo {
	id := ZoneNameHash(name)
	zi := zr.FindZoneInfoByID(id)
	if zi == nil {
		return nil
	}
	if zi.Name != name {
		return nil
	}
	return zi
}

func ZoneNameHash(s string) uint32 {
	return djb2(s)
}

func djb2(s string) uint32 {
	var hash uint32 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + uint32(c) /* hash * 33 + c */
	}

	return hash
}
