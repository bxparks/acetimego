package acetime

import (
	"github.com/bxparks/AceTimeGo/zoneinfo"
)

//-----------------------------------------------------------------------------
// ZoneManager creates TimeZone objects from its ZoneID or its Zone Name.
//-----------------------------------------------------------------------------

type ZoneManager struct {
	zoneContext   *zoneinfo.ZoneContext
	zoneRegistrar *ZoneRegistrar
}

func NewZoneManager(context *zoneinfo.ZoneContext) ZoneManager {
	registrar := NewZoneRegistrar(context.ZoneRegistry)
	manager := ZoneManager{
		zoneContext:   context,
		zoneRegistrar: &registrar,
	}

	return manager
}

func (zm *ZoneManager) NewTimeZoneFromID(zoneID uint32) TimeZone {
	zi := zm.zoneRegistrar.FindZoneInfoByID(zoneID)
	if zi == nil {
		return NewTimeZoneError()
	}
	return NewTimeZoneFromZoneInfo(zm.zoneContext, zi)
}

func (zm *ZoneManager) NewTimeZoneFromName(name string) TimeZone {
	zi := zm.zoneRegistrar.FindZoneInfoByName(name)
	if zi == nil {
		return NewTimeZoneError()
	}
	return NewTimeZoneFromZoneInfo(zm.zoneContext, zi)
}
