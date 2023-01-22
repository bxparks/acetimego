package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
	"math"
)

//-----------------------------------------------------------------------------
// ZoneRegistrar provides lookup functions for a given zone registry. The
// internal implementation of the zone registry is encapsulation by the
// registrar. The initial implementation is a map[uint32]*ZoneInfo. However, we
// may be able to save significant memory by using an array of []*ZoneInfo, and
// using a binary search on zoneId, just like the AceTime and AceTimeC
// libraries.
//-----------------------------------------------------------------------------

const (
	InvalidRegistryIndex = math.MaxUint16
)

type ZoneRegistry = map[uint32]*zoneinfo.ZoneInfo

type ZoneRegistry2 = []*zoneinfo.ZoneInfo

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

func IsZoneRegistrySorted(zis ZoneRegistry2) bool {
	var prevID uint32 = 0
	for _, zi := range zis {
		id := zi.ZoneID
		if id < prevID {
			return false
		}
		prevID = id
	}
	return true
}

func FindByIdLinear(zis ZoneRegistry2, id uint32) uint16 {
	for i, zi := range zis {
		if zi.ZoneID == id {
			return uint16(i)
		}
	}
	return InvalidRegistryIndex
}

func FindByIdBinary(zis ZoneRegistry2, id uint32) uint16 {
  var a uint16 = 0
  var b uint16 = uint16(len(zis))
  for {
    diff := b - a
    if (diff == 0) {
			break
		}

    c := a + diff / 2
    current := zis[c].ZoneID
    if id == current {
			return c
		}
    if (id < current) {
      b = c
    } else {
      a = c + 1
    }
  }
  return InvalidRegistryIndex
}
