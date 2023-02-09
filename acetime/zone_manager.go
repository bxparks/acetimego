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
		return TimeZoneError
	}
	return NewTimeZoneFromZoneInfo(zi)
}

func (zm *ZoneManager) NewTimeZoneFromName(name string) TimeZone {
	zi := zm.store.ZoneInfoByName(name)
	if zi == nil {
		return TimeZoneError
	}
	return NewTimeZoneFromZoneInfo(zi)
}

func (zm *ZoneManager) NewTimeZoneFromIndex(index uint16) TimeZone {
	if index >= zm.ZoneCount() {
		return TimeZoneError
	}
	return NewTimeZoneFromZoneInfo(zm.store.ZoneInfo(index))
}

// ZoneCount returns the number of zones (Zones and Links) in the database.
func (zm *ZoneManager) ZoneCount() uint16 {
	return zm.store.ZoneCount()
}

// ZoneNames returns the list of zone names in the database. The list will
// probably be *not* sorted.
func (zm *ZoneManager) ZoneNames() []string {
	return zm.store.ZoneNames()
}

// ZoneIds returns a list of ZoneIDsin the database. The list will likely be
// sorted but that is not guaranteed.
func (zm *ZoneManager) ZoneIDs() []uint32 {
	return zm.store.ZoneIDs()
}
