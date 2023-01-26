package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

//-----------------------------------------------------------------------------
// ZoneManager creates TimeZone objects from its ZoneID or its Zone Name.
//-----------------------------------------------------------------------------

type ZoneManager struct {
	store *zoneinfo.ZoneStore
}

func NewZoneManager(context *zoneinfo.ZoneDataContext) ZoneManager {
	return ZoneManager{zoneinfo.NewZoneStore(context)}
}

func (zm *ZoneManager) NewTimeZoneFromID(zoneID uint32) TimeZone {
	zi := zm.store.ZoneInfoByID(zoneID)
	if zi == nil {
		return NewTimeZoneError()
	}
	return NewTimeZoneFromZoneInfo(zi)
}

func (zm *ZoneManager) NewTimeZoneFromName(name string) TimeZone {
	zi := zm.store.ZoneInfoByName(name)
	if zi == nil {
		return NewTimeZoneError()
	}
	return NewTimeZoneFromZoneInfo(zi)
}
