package zoneinfo

// Functions that read the hex encoded string into a ZoneXxxRecord object.

//-----------------------------------------------------------------------------

type ZoneRuleReader struct {
	f         DataIO
	chunkSize uint8
}

func NewZoneRuleReader(f DataIO, chunkSize uint8) ZoneRuleReader {
	return ZoneRuleReader{f, chunkSize}
}

func (r *ZoneRuleReader) Reset() {
	r.f.Reset()
}

func (r *ZoneRuleReader) Seek(i uint16) {
	r.f.Seek(i * uint16(r.chunkSize))
}

func (r *ZoneRuleReader) Read() ZoneRuleRecord {
	var record ZoneRuleRecord
	record.FromYear = int16(r.f.ReadU16())
	record.ToYear = int16(r.f.ReadU16())
	record.InMonth = r.f.ReadU8()
	record.OnDayOfWeek = r.f.ReadU8()
	record.OnDayOfMonth = int8(r.f.ReadU8())
	record.AtSecondsModifier = r.f.ReadU8()
	record.AtSecondsCode = r.f.ReadU16()
	record.DeltaMinutes = int8(r.f.ReadU8())
	record.LetterIndex = r.f.ReadU8()
	// TODO: Check f.err for errors
	return record
}

//-----------------------------------------------------------------------------

type ZonePolicyReader struct {
	f         DataIO
	chunkSize uint8
}

func NewZonePolicyReader(f DataIO, chunkSize uint8) ZonePolicyReader {
	return ZonePolicyReader{f, chunkSize}
}

func (r *ZonePolicyReader) Reset() {
	r.f.Reset()
}

func (r *ZonePolicyReader) Seek(i uint16) {
	r.f.Seek(i * uint16(r.chunkSize))
}

func (r *ZonePolicyReader) Read() ZonePolicyRecord {
	var record ZonePolicyRecord
	record.RuleIndex = r.f.ReadU16()
	record.RuleCount = r.f.ReadU16()
	// TODO: Check r.f.err for errors
	return record
}

//-----------------------------------------------------------------------------

type ZoneEraReader struct {
	f         DataIO
	chunkSize uint8
}

func NewZoneEraReader(f DataIO, chunkSize uint8) ZoneEraReader {
	return ZoneEraReader{f, chunkSize}
}

func (r *ZoneEraReader) Reset() {
	r.f.Reset()
}

func (r *ZoneEraReader) Seek(i uint16) {
	r.f.Seek(i * uint16(r.chunkSize))
}

func (r *ZoneEraReader) Read() ZoneEraRecord {
	var record ZoneEraRecord
	record.FormatIndex = r.f.ReadU16()
	record.PolicyIndex = r.f.ReadU8()
	record.OffsetSecondsRemainder = r.f.ReadU8()
	record.OffsetSecondsCode = int16(r.f.ReadU16())
	record.UntilYear = int16(r.f.ReadU16())
	record.DeltaMinutes = int8(r.f.ReadU8())
	record.UntilMonth = r.f.ReadU8()
	record.UntilDay = r.f.ReadU8()
	record.UntilSecondsModifier = r.f.ReadU8()
	record.UntilSecondsCode = r.f.ReadU16()
	// TODO: Check r.f.err for errors
	return record
}

//-----------------------------------------------------------------------------

type ZoneInfoReader struct {
	f         DataIO
	chunkSize uint8
}

func NewZoneInfoReader(f DataIO, chunkSize uint8) ZoneInfoReader {
	return ZoneInfoReader{f, chunkSize}
}

func (r *ZoneInfoReader) Reset() {
	r.f.Reset()
}

func (r *ZoneInfoReader) Seek(i uint16) {
	r.f.Seek(i * uint16(r.chunkSize))
}

func (r *ZoneInfoReader) Read() ZoneInfoRecord {
	var record ZoneInfoRecord
	record.ZoneID = r.f.ReadU32()
	record.NameIndex = r.f.ReadU16()
	record.EraIndex = r.f.ReadU16()
	record.EraCount = r.f.ReadU16()
	record.TargetIndex = r.f.ReadU16()
	// TODO: Check r.f.err for errors
	return record
}
