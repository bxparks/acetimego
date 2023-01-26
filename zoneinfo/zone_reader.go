package zoneinfo

// Functions that read the hex encoded string into a ZoneXxxRecord object.

//-----------------------------------------------------------------------------

type ZoneRuleReader struct {
	f DataIO
}

func (r *ZoneRuleReader) Reset() {
	r.f.Reset()
}

func (r *ZoneRuleReader) Seek(i uint16) {
	r.f.Seek(i * 11)
}

func (r *ZoneRuleReader) Read() ZoneRuleRecord {
	var record ZoneRuleRecord
	record.FromYear = int16(r.f.ReadU16())
	record.ToYear = int16(r.f.ReadU16())
	record.InMonth = r.f.ReadU8()
	record.OnDayOfWeek = r.f.ReadU8()
	record.OnDayOfMonth = int8(r.f.ReadU8())
	record.AtTimeCode = r.f.ReadU8()
	record.AtTimeModifier = r.f.ReadU8()
	record.DeltaCode = int8(r.f.ReadU8())
	record.LetterIndex = r.f.ReadU8()
	// TODO: Check f.err for errors
	return record
}

//-----------------------------------------------------------------------------

type ZonePolicyReader struct {
	f DataIO
}

func (r *ZonePolicyReader) Reset() {
	r.f.Reset()
}

func (r *ZonePolicyReader) Seek(i uint16) {
	r.f.Seek(i * 4)
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
	f DataIO
}

func (r *ZoneEraReader) Reset() {
	r.f.Reset()
}

func (r *ZoneEraReader) Seek(i uint16) {
	r.f.Seek(i * 11)
}

func (r *ZoneEraReader) Read() ZoneEraRecord {
	var record ZoneEraRecord
	record.FormatIndex = r.f.ReadU16()
	record.PolicyIndex = r.f.ReadU8()
	record.OffsetCode = int8(r.f.ReadU8())
	record.DeltaCode = r.f.ReadU8()
	record.UntilYear = int16(r.f.ReadU16())
	record.UntilMonth = r.f.ReadU8()
	record.UntilDay = r.f.ReadU8()
	record.UntilTimeCode = r.f.ReadU8()
	record.UntilTimeModifier = r.f.ReadU8()
	// TODO: Check r.f.err for errors
	return record
}

//-----------------------------------------------------------------------------

type ZoneInfoReader struct {
	f DataIO
}

func (r *ZoneInfoReader) Reset() {
	r.f.Reset()
}

func (r *ZoneInfoReader) Seek(i uint16) {
	r.f.Seek(i * 12)
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
