package acetime

import (
	"github.com/bxparks/acetimego/zoneinfo"
)

// ZoneManager creates TimeZone objects from its ZoneID or its Zone Name.
type ZoneManager struct {
	store *zoneinfo.ZoneStore
}

func NewZoneManager(context *zoneinfo.ZoneDataContext) ZoneManager {
	return ZoneManager{zoneinfo.NewZoneStore(context)}
}

func (zm *ZoneManager) TimeZoneFromZoneID(zoneID uint32) TimeZone {
	info := zm.store.ZoneInfoByID(zoneID)
	if info == nil {
		return TimeZoneError
	}
	return newTimeZoneFromZoneInfo(info)
}

func (zm *ZoneManager) TimeZoneFromName(name string) TimeZone {
	info := zm.store.ZoneInfoByName(name)
	if info == nil {
		return TimeZoneError
	}
	return newTimeZoneFromZoneInfo(info)
}

// TODO: The "index" of a ZoneInfo is currently not exported, so this function
// is not useful to the end-user. Maybe remove it altogether.
func (zm *ZoneManager) timeZoneFromIndex(index uint16) TimeZone {
	if index >= zm.ZoneCount() {
		return TimeZoneError
	}
	return newTimeZoneFromZoneInfo(zm.store.ZoneInfo(index))
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
