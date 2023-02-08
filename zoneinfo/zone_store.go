package zoneinfo

// Functions that convert a ZoneXxxRecord to a ZoneXxx instance.

const InvalidIndex = 1<<16 - 1

type ZoneStore struct {
	context      *ZoneDataContext
	nameIO       StringIO16
	formatIO     StringIO16
	letterIO     StringIO8
	infoReader   ZoneInfoReader
	eraReader    ZoneEraReader
	policyReader ZonePolicyReader
	ruleReader   ZoneRuleReader
}

func NewZoneStore(c *ZoneDataContext) *ZoneStore {
	return &ZoneStore{
		context:    c,
		nameIO:     StringIO16{c.NameData, c.NameOffsets},
		formatIO:   StringIO16{c.FormatData, c.FormatOffsets},
		letterIO:   StringIO8{c.LetterData, c.LetterOffsets},
		infoReader: ZoneInfoReader{NewDataIO(c.ZoneInfosData), c.ZoneInfoChunkSize},
		eraReader:  ZoneEraReader{NewDataIO(c.ZoneErasData), c.ZoneEraChunkSize},
		policyReader: ZonePolicyReader{
			NewDataIO(c.ZonePoliciesData), c.ZonePolicyChunkSize},
		ruleReader: ZoneRuleReader{NewDataIO(c.ZoneRulesData), c.ZoneRuleChunkSize},
	}
}

// ZoneNames() returns an array of all zone names in the database.
func (zs *ZoneStore) ZoneNames() []string {
	return zs.nameIO.Strings()
}

// ZoneIDs() returns an array of all ZoneIDs in the database.
func (zs *ZoneStore) ZoneIDs() []uint32 {
	count := zs.context.ZoneInfoCount
	ids := make([]uint32, count)
	for i := uint16(0); i < count; i++ {
		zs.infoReader.Seek(i)
		record := zs.infoReader.Read()
		ids[i] = record.ZoneID
	}
	return ids
}

// ZoneInfo retrieves the ZoneXxxRecords from the various ZoneXxxData strings,
// then converts them into a fully populated ZoneInfo object.
func (zs *ZoneStore) ZoneInfo(i uint16) *ZoneInfo {
	// Retrieve the ZoneInfoRecord and follow the graph of objects indicated by
	// the various XxxIndex foreign keys.
	zs.infoReader.Seek(i)
	record := zs.infoReader.Read()

	var info ZoneInfo
	zs.fillZoneInfo(&info, &record)
	return &info
}

func (zs *ZoneStore) fillZoneInfo(info *ZoneInfo, record *ZoneInfoRecord) {
	name := zs.nameIO.StringAt(record.NameIndex)
	var eras []ZoneEra
	var target *ZoneInfo

	if record.EraCount == 0 { // Link, so recursively resolve the target ZoneInfo
		eras = nil
		target = zs.ZoneInfo(record.TargetIndex)
	} else { // Zone
		eras = zs.ZoneEras(record.EraIndex, record.EraCount)
		target = nil
	}

	info.Name = name
	info.ZoneID = record.ZoneID
	info.StartYear = zs.context.StartYear
	info.UntilYear = zs.context.UntilYear
	info.Eras = eras
	info.Target = target
}

func (zs *ZoneStore) ZoneEras(i uint16, count uint16) []ZoneEra {
	eras := make([]ZoneEra, count)
	zs.eraReader.Seek(i)
	for j := uint16(0); j < count; j++ {
		record := zs.eraReader.Read()
		zs.fillZoneEra(&eras[j], &record)
	}
	return eras
}

func (zs *ZoneStore) fillZoneEra(era *ZoneEra, record *ZoneEraRecord) {
	era.Format = zs.formatIO.StringAt(record.FormatIndex)
	era.OffsetSecondsCode = record.OffsetSecondsCode
	era.DeltaCode = record.DeltaCode
	era.UntilYear = record.UntilYear
	era.UntilMonth = record.UntilMonth
	era.UntilDay = record.UntilDay
	era.UntilTimeCode = record.UntilTimeCode
	era.UntilTimeModifier = record.UntilTimeModifier
	if record.PolicyIndex == 0 {
		era.Policy = nil
	} else {
		era.Policy = zs.ZonePolicy(uint16(record.PolicyIndex))
	}
}

func (zs *ZoneStore) ZonePolicy(i uint16) *ZonePolicy {
	zs.policyReader.Seek(i)
	record := zs.policyReader.Read()
	var policy ZonePolicy
	zs.fillZonePolicy(&policy, &record)
	return &policy
}

func (zs *ZoneStore) fillZonePolicy(
	policy *ZonePolicy, record *ZonePolicyRecord) {

	policy.Rules = zs.ZoneRules(record.RuleIndex, record.RuleCount)
}

func (zs *ZoneStore) ZoneRules(i uint16, count uint16) []ZoneRule {
	rules := make([]ZoneRule, count)
	zs.ruleReader.Seek(i)
	for j := uint16(0); j < count; j++ {
		record := zs.ruleReader.Read()
		zs.fillZoneRule(&rules[j], &record)
	}
	return rules
}

func (zs *ZoneStore) fillZoneRule(rule *ZoneRule, record *ZoneRuleRecord) {
	rule.FromYear = record.FromYear
	rule.ToYear = record.ToYear
	rule.InMonth = record.InMonth
	rule.OnDayOfWeek = record.OnDayOfWeek
	rule.OnDayOfMonth = record.OnDayOfMonth
	rule.AtTimeCode = record.AtTimeCode
	rule.AtTimeModifier = record.AtTimeModifier
	rule.DeltaCode = record.DeltaCode
	rule.Letter = zs.letterIO.StringAt(record.LetterIndex)
}

func (zs *ZoneStore) ZoneInfoByID(id uint32) *ZoneInfo {
	// TODO: Incorporate binary search.
	i := zs.FindByIDLinear(id)
	if i == InvalidIndex {
		return nil
	}
	return zs.ZoneInfo(i)
}

func (zs *ZoneStore) ZoneInfoByName(name string) *ZoneInfo {
	id := ZoneNameHash(name)
	// TODO: Incorporate binary search.
	i := zs.FindByIDLinear(id)
	if i == InvalidIndex {
		return nil
	}

	// Check for hash collision
	zs.infoReader.Seek(i)
	record := zs.infoReader.Read()
	recordName := zs.nameIO.StringAt(record.NameIndex)
	if recordName != name {
		return nil
	}

	return zs.ZoneInfo(i)
}

func (zs *ZoneStore) FindByIDLinear(id uint32) uint16 {
	zs.infoReader.Reset()
	for i := uint16(0); i < zs.context.ZoneInfoCount; i++ {
		record := zs.infoReader.Read()
		if record.ZoneID == id {
			return i
		}
	}
	return InvalidIndex
}

func (zs *ZoneStore) FindByIDBinary(id uint32) uint16 {
	zs.infoReader.Reset()
	var a uint16 = 0
	var b uint16 = zs.context.ZoneInfoCount
	for {
		diff := b - a
		if diff == 0 {
			break
		}

		c := a + diff/2
		zs.infoReader.Seek(c)
		zi := zs.infoReader.Read()
		current := zi.ZoneID
		if id == current {
			return c
		}
		if id < current {
			b = c
		} else {
			a = c + 1
		}
	}
	return InvalidIndex
}

func (zs *ZoneStore) IsSorted() bool {
	var prevID uint32 = 0
	zs.infoReader.Reset()
	for i := uint16(0); i < zs.context.ZoneInfoCount; i++ {
		record := zs.infoReader.Read()
		id := record.ZoneID
		if id < prevID {
			return false
		}
		prevID = id
	}
	return true
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
