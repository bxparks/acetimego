package zoneinfo

// Functions that convert a ZoneXxxRecord to a ZoneXxx instance.

// InvalidIndex is returned by the FindByID() function to indicate that a ZoneID
// was not found. The max number of entries that can be described by a `uint16`
// is `2^16-1`, so the index of a valid entry can never be `2^16-1`.
const InvalidIndex = 1<<16 - 1

// A ZoneStore knows how to materialize a zoneinfo object (ZoneRule, ZonePolicy,
// ZoneEra, ZoneInfo) from the underlying persistent data storage described by
// the ZoneDataContext. It also knows how to search the underlying data storage
// for a ZoneInfo object by its Name (e.g. "America/Los_Angeles" or by a uint32
// integer ID (e.g. ZoneIDAmerica_Los_Angeles = 0xb7f7e8f2). It performs the
// search operation without loading the entire database into memory because it
// knows the internal layout of the database.
type ZoneStore struct {
	context      *ZoneDataContext
	nameIO       StringIO16
	formatIO     StringIO16
	letterIO     StringIO8
	infoReader   ZoneInfoReader
	eraReader    ZoneEraReader
	policyReader ZonePolicyReader
	ruleReader   ZoneRuleReader
	isSorted     bool
}

func NewZoneStore(c *ZoneDataContext) *ZoneStore {
	store := ZoneStore{
		context:    c,
		nameIO:     StringIO16{c.NameData, c.NameOffsets},
		formatIO:   StringIO16{c.FormatData, c.FormatOffsets},
		letterIO:   StringIO8{c.LetterData, c.LetterOffsets},
		infoReader: ZoneInfoReader{NewDataIO(c.ZoneInfosData), c.ZoneInfoChunkSize},
		eraReader:  ZoneEraReader{NewDataIO(c.ZoneErasData), c.ZoneEraChunkSize},
		policyReader: ZonePolicyReader{
			NewDataIO(c.ZonePoliciesData),
			c.ZonePolicyChunkSize,
		},
		ruleReader: ZoneRuleReader{NewDataIO(c.ZoneRulesData), c.ZoneRuleChunkSize},
	}
	store.isSorted = store.IsSorted()
	return &store
}

// ZoneCount() returns the number of zones in this database
func (store *ZoneStore) ZoneCount() uint16 {
	return store.context.ZoneInfoCount
}

// ZoneNames() returns an array of all zone names in the database.
func (store *ZoneStore) ZoneNames() []string {
	return store.nameIO.Strings()
}

// ZoneIDs() returns an array of all ZoneIDs in the database.
func (store *ZoneStore) ZoneIDs() []uint32 {
	count := store.context.ZoneInfoCount
	ids := make([]uint32, count)
	for i := uint16(0); i < count; i++ {
		store.infoReader.Seek(i)
		record := store.infoReader.Read()
		ids[i] = record.ZoneID
	}
	return ids
}

// ZoneInfo retrieves the ZoneXxxRecords from the various ZoneXxxData strings,
// then converts them into a fully populated ZoneInfo object.
func (store *ZoneStore) ZoneInfo(i uint16) *ZoneInfo {
	// Retrieve the ZoneInfoRecord and follow the graph of objects indicated by
	// the various XxxIndex foreign keys.
	store.infoReader.Seek(i)
	record := store.infoReader.Read()

	var info ZoneInfo
	store.fillZoneInfo(&info, &record)
	return &info
}

func (store *ZoneStore) fillZoneInfo(info *ZoneInfo, record *ZoneInfoRecord) {
	name := store.nameIO.StringAt(record.NameIndex)
	var eras []ZoneEra
	var target *ZoneInfo

	if record.EraCount == 0 { // Link, so recursively resolve the target ZoneInfo
		eras = nil
		target = store.ZoneInfo(record.TargetIndex)
	} else { // Zone
		eras = store.ZoneEras(record.EraIndex, record.EraCount)
		target = nil
	}

	info.Name = name
	info.ZoneID = record.ZoneID
	info.StartYear = store.context.StartYear
	info.UntilYear = store.context.UntilYear
	info.Eras = eras
	info.Target = target
}

func (store *ZoneStore) ZoneEras(i uint16, count uint16) []ZoneEra {
	eras := make([]ZoneEra, count)
	store.eraReader.Seek(i)
	for j := uint16(0); j < count; j++ {
		record := store.eraReader.Read()
		store.fillZoneEra(&eras[j], &record)
	}
	return eras
}

func (store *ZoneStore) fillZoneEra(era *ZoneEra, record *ZoneEraRecord) {
	if record.PolicyIndex == 0 {
		era.Policy = nil
	} else {
		era.Policy = store.ZonePolicy(uint16(record.PolicyIndex))
	}
	era.Format = store.formatIO.StringAt(record.FormatIndex)
	era.OffsetSecondsRemainder = record.OffsetSecondsRemainder
	era.OffsetSecondsCode = record.OffsetSecondsCode
	era.DeltaMinutes = record.DeltaMinutes
	era.UntilYear = record.UntilYear
	era.UntilMonth = record.UntilMonth
	era.UntilDay = record.UntilDay
	era.UntilSecondsCode = record.UntilSecondsCode
	era.UntilSecondsModifier = record.UntilSecondsModifier
}

func (store *ZoneStore) ZonePolicy(i uint16) *ZonePolicy {
	store.policyReader.Seek(i)
	record := store.policyReader.Read()
	var policy ZonePolicy
	store.fillZonePolicy(&policy, &record)
	return &policy
}

func (store *ZoneStore) fillZonePolicy(
	policy *ZonePolicy, record *ZonePolicyRecord) {

	policy.Rules = store.ZoneRules(record.RuleIndex, record.RuleCount)
}

func (store *ZoneStore) ZoneRules(i uint16, count uint16) []ZoneRule {
	rules := make([]ZoneRule, count)
	store.ruleReader.Seek(i)
	for j := uint16(0); j < count; j++ {
		record := store.ruleReader.Read()
		store.fillZoneRule(&rules[j], &record)
	}
	return rules
}

func (store *ZoneStore) fillZoneRule(rule *ZoneRule, record *ZoneRuleRecord) {
	rule.FromYear = record.FromYear
	rule.ToYear = record.ToYear
	rule.InMonth = record.InMonth
	rule.OnDayOfWeek = record.OnDayOfWeek
	rule.OnDayOfMonth = record.OnDayOfMonth
	rule.AtSecondsCode = record.AtSecondsCode
	rule.AtSecondsModifier = record.AtSecondsModifier
	rule.DeltaMinutes = record.DeltaMinutes
	rule.Letter = store.letterIO.StringAt(record.LetterIndex)
}

func (store *ZoneStore) ZoneInfoByID(id uint32) *ZoneInfo {
	i := store.FindByID(id)
	if i == InvalidIndex {
		return nil
	}
	return store.ZoneInfo(i)
}

func (store *ZoneStore) ZoneInfoByName(name string) *ZoneInfo {
	id := ZoneNameHash(name)
	i := store.FindByID(id)
	if i == InvalidIndex {
		return nil
	}

	// Check for hash collision
	store.infoReader.Seek(i)
	record := store.infoReader.Read()
	recordName := store.nameIO.StringAt(record.NameIndex)
	if recordName != name {
		return nil
	}

	return store.ZoneInfo(i)
}

func (store *ZoneStore) FindByID(id uint32) uint16 {
	if store.isSorted {
		return store.FindByIDBinary(id)
	} else {
		return store.FindByIDLinear(id)
	}
}

func (store *ZoneStore) FindByIDLinear(id uint32) uint16 {
	store.infoReader.Reset()
	for i := uint16(0); i < store.context.ZoneInfoCount; i++ {
		record := store.infoReader.Read()
		if record.ZoneID == id {
			return i
		}
	}
	return InvalidIndex
}

func (store *ZoneStore) FindByIDBinary(id uint32) uint16 {
	store.infoReader.Reset()
	var a uint16 = 0
	var b uint16 = store.context.ZoneInfoCount
	for {
		diff := b - a
		if diff == 0 {
			break
		}

		c := a + diff/2 // avoids overflow of uint16
		store.infoReader.Seek(c)
		zi := store.infoReader.Read()
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

func (store *ZoneStore) IsSorted() bool {
	var prevID uint32 = 0
	store.infoReader.Reset()
	for i := uint16(0); i < store.context.ZoneInfoCount; i++ {
		record := store.infoReader.Read()
		id := record.ZoneID
		if id < prevID {
			return false
		}
		prevID = id
	}
	return true
}

// ZoneNameHash is the hash function that converts the zone name (e.g.
// "America/Los_Angeles" into a uint32 integer. This is intended to be stable
// and unique for all future versions of acetimego. If a new zone is added in
// the future that causes a hash collision, a modified form of this function
// will be created so that previous hashes remain stable, while resolving the
// hash collision.
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
