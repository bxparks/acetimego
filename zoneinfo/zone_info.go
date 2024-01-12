package zoneinfo

// These are the data structures used on the 'zonedb' files which capture
// the information contained in the IANA TZ database.

const (
	// The minimum value of ZoneRule.FromYear and ZoneRule.ToYear. Used by
	// synthetic rule entries for certain zones, to guarantee that all zones have
	// at least one transition.
	MinYear int16 = -32767

	// The maximum value of ZoneRule.ToYear, representing the sentinel value "max"
	// in the TO columns of the TZDB files. Must be less than MaxUntilYear.
	MaxToYear int16 = 32766

	// The maximum value of ZoneEra.UntilYear, representing the sentinel value
	// "-" in the UNTIL column of the TZDB files. Must be greater than MaxYear.
	MaxUntilYear int16 = MaxToYear + 1
)

//-----------------------------------------------------------------------------

// ZoneContext contains references to variables which describe the current time
// zone database. This version contains the ZoneRules, ZonePolicies, ZoneEras,
// and ZoneInfos arrays which are the expanded versions of the data structures.
type ZoneRecordContext struct {
	TzDatabaseVersion string
	StartYear         int16
	UntilYear         int16
	StartYearAccurate int16
	UntilYearAccurate int16
	MaxTransitions    int16
	LetterData        string
	LetterOffsets     []uint8
	FormatData        string
	FormatOffsets     []uint16
	NameData          string
	NameOffsets       []uint16
	ZoneRuleRecords   []ZoneRuleRecord
	ZonePolicyRecords []ZonePolicyRecord
	ZoneEraRecords    []ZoneEraRecord
	ZoneInfoRecords   []ZoneInfoRecord
}

// ZoneDataContext is an alternate version of ZoneRecordContext that encodes the
// various record arrays (ZoneRuleRecords, ZonePolicyRecords, ZoneEraRecords,
// ZoneInfoRecords) as hex-encoded string variables (ZoneRulesData,
// ZonePoliciesData, ZoneErasData, and ZoneInfosData) so that they are placed
// into flash memory instead of consuming RAM. This is important in
// microcontroller environments using the TinyGo compiler.
type ZoneDataContext struct {
	TzDatabaseVersion   string
	StartYear           int16
	UntilYear           int16
	StartYearAccurate   int16
	UntilYearAccurate   int16
	MaxTransitions      int16
	LetterData          string
	LetterOffsets       []uint8
	FormatData          string
	FormatOffsets       []uint16
	NameData            string
	NameOffsets         []uint16
	ZoneRuleChunkSize   uint8
	ZonePolicyChunkSize uint8
	ZoneEraChunkSize    uint8
	ZoneInfoChunkSize   uint8
	ZoneRuleCount       uint16
	ZonePolicyCount     uint16
	ZoneEraCount        uint16
	ZoneInfoCount       uint16
	ZoneRulesData       string
	ZonePoliciesData    string
	ZoneErasData        string
	ZoneInfosData       string
}

//-----------------------------------------------------------------------------

// ZoneRule is the in-memory hydrated version of ZoneRuleRecord which describes
// one RULE line from the raw TZDB file.
type ZoneRule struct {

	// FROM year.
	FromYear int16

	// TO year.
	ToYear int16

	// Determined by the IN column. 1=Jan, 12=Dec.
	InMonth uint8

	// Determined by the ON column. Possible values are: 0, 1=Mon, 7=Sun.
	// There are 4 combinations:
	//
	//  * OnDayOfWeek=0, OnDayOfMonth=(1-31): exact match
	//  * OnDayOfWeek=1-7, OnDayOfMonth=1-31: day_of_week>=OnDayOfMonth
	//  * OnDayOfWeek=1-7, OnDayOfMonth=-(1-31): day_of_week<=OnDayOfMonth
	//  * OnDayOfWeek=1-7, OnDayOfMonth=0: last{day_of_week}
	OnDayOfWeek uint8

	// Determined by the ON column. Used with OnDayOfWeek. Possible values are:
	// 0, 1-31, or its corresponding negative values.
	OnDayOfMonth int8

	// The AtSecondsModifier is a packed field containing 2 pieces of info:
	//
	//  * The upper 4 bits represent the AT time suffix: 'w', 's' or 'u',
	//    represented by SuffixW, SuffixS and SuffixU.
	//  * The lower 4 bits represent the remaining 0-14 seconds of the AT field
	//    after truncation into AtSecondsCode. In other words, the full AT field
	//    is (15 * AtSecondsCode + (AtSecondsModifier & 0x0f)).
	AtSecondsModifier uint8

	// Determined by the AT column in units of 15-seconds, [00:00:00,25:00:00]
	// which corresponds to [0,6000] in 15-second units.
	AtSecondsCode uint16

	// Determined by the SAVE column, containing the offset from UTC in minutes,
	// in the range of [-128,+127].
	DeltaMinutes int8

	// Determined by the LETTER column. Determines the substitution into the '%s'
	// field (implemented here by just a '%') of the ZoneInfo.Format field.
	// Most comment values in the raw TZDB files are "S", "D", and "-". The "-" is
	// stored as "" (empty string).
	//
	// As of TZ DB version 2018i, there are 4 ZonePolicies which have ZoneRules
	// with a LETTER field longer than 1 character:
	//
	//  * Belize ('CST'; used by America/Belize)
	//  * Namibia ('WAT', 'CAT'; used by Africa/Windhoek)
	//  * StJohns ('DD'; used by America/St_Johns and America/Goose_Bay)
	//  * Troll ('+00' '+02'; used by Antarctica/Troll)
	Letter string
}

func (rule *ZoneRule) AtSeconds() int32 {
	return int32(rule.AtSecondsCode)*15 + int32(rule.AtSecondsModifier&0x0f)
}

func (rule *ZoneRule) AtSuffix() uint8 {
	return rule.AtSecondsModifier & 0xf0
}

func (rule *ZoneRule) DstOffsetSeconds() int32 {
	return int32(rule.DeltaMinutes) * 60
}

//-----------------------------------------------------------------------------

// ZoneRuleRecord is the distilled version of ZoneRule which can be persisted
// in a file or a hex encoded string.
type ZoneRuleRecord struct {
	FromYear          int16
	ToYear            int16
	InMonth           uint8
	OnDayOfWeek       uint8
	OnDayOfMonth      int8
	AtSecondsModifier uint8  // AT time, remainder seconds [0,14]
	AtSecondsCode     uint16 // AT time, units of 15 seconds, [0h,25h]
	DeltaMinutes      int8   // the DST SAVE offset in minutes
	LetterIndex       uint8  // index into LetterData
}

//-----------------------------------------------------------------------------

// A collection of transition rules which describe the DST rules of a given
// administrative region. A given time zone (ZoneInfo) can follow a different
// ZonePolicy at different times. Conversely, multiple time zones (ZoneInfo)
// can choose to follow the same ZonePolicy at different times.
type ZonePolicy struct {
	Rules []ZoneRule
}

//-----------------------------------------------------------------------------

// A ZonePolicyRecord describes a collection ZoneRuleRecords through the
// RuleIndex into the ZoneRuleData, and the number of ZoneRuleData. The
// collection of rules for the policy at index RuleIndex is
// `ZoneRuleRecords[RuleIndex:RuleIndex+RuleCount]`.
//
// The transition rules describe the DST rules of a given administrative region.
// A given time zone (ZoneInfo) can follow a different ZonePolicy at different
// times. Conversely, multiple time zones (ZoneInfo) can choose to follow the
// same ZonePolicy at different times.
//
// A synthetic record entry of {0, 0} exists to represent a ZoneEraRecord that
// has no policy. (This is needed because a PolicyIndex is an integer which
// cannot be nil, so we use an index of 0 to represent nil.)
//
// TODO: Maybe add a PolicyNameIndex, to reveal the name of the policy (e.g.
// "US" or "WS". It isn't used anywhere in the code, but could be useful in
// debugging.
type ZonePolicyRecord struct {
	RuleIndex uint16 // index into the ZoneRulesData
	RuleCount uint16 // always > 0, every policy has at least one Rule
}

//---------------------------------------------------------------------------

// The Suffix* constants are stored in the upper nibble of various uint8 fields
// named XxxModifier below.
const (
	SuffixW uint8 = 0x00 // Represents 'w' or wall time.
	SuffixS uint8 = 0x10 // Represents 's' or standard time.
	SuffixU uint8 = 0x20 // Represents 'u' or UTC time.
)

// An entry in ZoneInfo which corresponds to a single RULE line from the raw
// TZDB file. It describes which ZonePolicy is being followed during a
// particular time period. The ZonePolicy is determined by the RULES column in
// the TZ Database file.
//
// There are 2 types of ZoneEra:
//
//  1. ZonePolicy == nil. Then DeltaCode determines the additional
//     offset from offsetCode. A value of '-' in the TZ Database file is
//     stored as 0.
//  2. ZonePolicy != nil. Then the DeltaCode offset is given by the
//     ZoneRule.DeltaCode of the ZoneRule which matches the time instant of
//     interest.
type ZoneEra struct {
	// Zone policy, determined by the RULES column. Set to nil if the RULES
	// column is '-' or an explicit DST shift in the form of 'hh:mm'.
	Policy *ZonePolicy

	// Zone abbreviations (e.g. PST, EST) determined by the FORMAT column. It has
	// 3 encodings in the TZ DB files:
	//
	//  1) A fixed string, e.g. "GMT".
	//  2) Two strings separated by a '/', e.g. "-03/-02" indicating
	//     "{std}/{dst}" options.
	//  3) A single string with a substitution, e.g. "E%sT", where the "%s" is
	//  replaced by the LETTER value from the ZoneRule.
	//
	// BasicZoneProcessor supports only a single letter subsitution from LETTER,
	// but ExtendedZoneProcessor supports substituting multi-character strings
	// (e.g. "CAT", "DD", "+00").
	//
	// The TZ DB files use '%s' to indicate the substitution, but for simplicity,
	// AceTime replaces the "%s" with just a '%' character with no loss of
	// functionality. This also makes the string-replacement code a little
	// simpler. For example, 'E%sT' is stored as 'E%T', and the LETTER
	// substitution is performed on the '%' character.
	Format string

	// The remainder seconds [0-14] from OffsetSecondsCode.
	OffsetSecondsRemainder uint8

	// UTC offset for the zone in standard time, as determined by the STDOFF
	// column in the TZ database. This field is in units of 15-seconds. Any
	// remaining seconds are encoded into the DeltaCode field.
	//
	// All zones after about 1974 uses STDOFF which are in units of 15 minute
	// increments. But some zones before 1974 use STDOFF in increments of one
	// second. We need a range of about -10h to +14h in one second increments,
	// 90000 seconds, which corresponds to about 17 bits of seconds.
	//
	// Instead of wasting an entire int32, we can use 16-bits in this field,
	// representing the offset in 15 *second* increments, then store any remainder
	// 15 seconds in the 8-bit OffsetSecondsRemainder field.
	OffsetSecondsCode int16

	// Era is valid until currentTime < UntilYear. Comes from the UNTIL column.
	UntilYear int16

	// The DSTOFF during DST phase. If ZonePolicy is nil, then the DST offsets are
	// defined by the RULES column in 'hh:mm' format, so this is set to 0.
	// Otherwise, the hh:mm is converted into minutes and stored here. Range is
	// [-128,127], so this handles [-02:00,+02:00].
	DeltaMinutes int8

	// The month field in UNTIL (1-12). Will never be 0.
	UntilMonth uint8

	// The day field in UNTIL (1-31). Will never be 0. Also, there's no need for
	// UntilDayOfWeek, because the database generator will resolve the exact day
	// of month based on the known year and month.
	UntilDay uint8

	// The time field of UNTIL field in 15-second increments. A range of 00:00:00
	// 25:00:00, which is [0,90000] seconds, which means we need 17-bits to store
	// this information. The remainder secondes [0-14] are stored in
	// UntilSecondsModifier.
	//
	// A number of zones before 1970 have UNTIL fields in one-second resolution:
	//
	//	* America/Adak {UNTIL time '12:44:35'}
	// 	* America/Anchorage {UNTIL time '14:31:37'}
	// 	* America/Juneau {UNTIL time '15:33:32'}
	// 	* America/Metlakatla {UNTIL time '15:44:55'}
	// 	* America/Nome {UNTIL time '13:29:35'}
	// 	* America/Yakutat {UNTIL time '15:12:18'}
	// 	* Europe/Brussels {UNTIL time '00:17:30'}
	UntilSecondsCode uint16

	// The UntilSecondsModifier is a packed field containing 2 pieces of info:
	//
	//  * The upper 4 bits represent the UNTIL time suffix: 'w', 's' or 'u',
	//    represented by SuffixW, SuffixS and SuffixU.
	//  * The lower 4 bits represent the remaining 0-14 seconds of the UNTIL
	//    field after truncation into UntilSecondsCode. In other words, the full
	//    UNTIL field is (15 * UntilSecondsCode + (UntilSecondsModifier & 0x0f)).
	UntilSecondsModifier uint8
}

func (era *ZoneEra) HasPolicy() bool {
	return era.Policy != nil
}

func (era *ZoneEra) StdOffsetSeconds() int32 {
	return int32(era.OffsetSecondsCode)*15 + int32(era.OffsetSecondsRemainder)
}

func (era *ZoneEra) DstOffsetSeconds() int32 {
	return int32(era.DeltaMinutes) * 60
}

func (era *ZoneEra) UntilSeconds() int32 {
	return int32(era.UntilSecondsCode)*15 + int32(era.UntilSecondsModifier&0x0f)
}

func (era *ZoneEra) UntilSuffix() uint8 {
	return era.UntilSecondsModifier & 0xf0
}

//-----------------------------------------------------------------------------

// ZoneEraRecord is a version of ZoneEra suitable for persisted in a file or hex
// encoded string.
type ZoneEraRecord struct {
	FormatIndex            uint16 // index into FormatData
	PolicyIndex            uint8  // index into ZonePoliciesData
	OffsetSecondsRemainder uint8  // remainder OffsetSeconds [0,14]
	OffsetSecondsCode      int16  // STDOFF units of 15 seconds, [-12h,+14h]
	UntilYear              int16  // year of UNTIL column
	DeltaMinutes           int8   // DSTOFF in units of 1 minute [-128,+127]
	UntilMonth             uint8  // month of UNTIL column
	UntilDay               uint8  // day of UNTIL column
	UntilSecondsModifier   uint8  // remainder UNTIL seconds, and w, s, u suffix
	UntilSecondsCode       uint16 // UNTIL time in units of 15 seconds, [0h,25h]
}

//-----------------------------------------------------------------------------

// The ZONE entries of a given time zone from the TZDB files, implemented as an
// array of ZoneEra records.
type ZoneInfo struct {
	// Full name of zone (e.g. "America/Los_Angeles").
	Name string

	// Unique, stable ID of the zone name, created from a hash of the name.
	// This ID will never change once assigned. This can be used for presistence
	// and serialization.
	ZoneID uint32

	// TODO: Maybe replace the StartYear, UntilYear, StartYearAccurate,
	// UntilYearAccurate with reference to ZoneContext, now that 4 fields are
	// copied from ZoneContext.

	// Start year of the zone files as requested.
	StartYear int16

	// Until year of the zone files as requested.
	UntilYear int16

	// Start year of accurate transitions. MinYear indicates -Infinity.
	StartYearAccurate int16

	// Until year of accurate transitions. MaxUntilYear indicates -Infinity.
	UntilYearAccurate int16

	// A slice of ZoneEra instances. For a normal Zone, num(Eras) is greater than
	// 0, and the ZoneEra entries are arranged in increasing order
	// of UNTIL time. For a Link entry, num(Eras) == 0, and instead 'target' is
	// non-nil and points to the target Zone. We have to follow the indirect
	// pointer, and resolve the target.Eras to obtain the actual ZoneEra entries.
	Eras []ZoneEra

	// If not nil, this entry is a Link to the target.
	Target *ZoneInfo
}

// IsLink returns true if the current zone is a Link.
func (zi *ZoneInfo) IsLink() bool {
	return zi.Target != nil
}

// ErasActive returns the Eras of the current zone, or the Eras of the target
// zone if the current zone is a Link.
func (zi *ZoneInfo) ErasActive() []ZoneEra {
	if zi.Target != nil {
		return zi.Target.Eras
	} else {
		return zi.Eras
	}
}

//-----------------------------------------------------------------------------

// ZoneInfoRecord is a version of ZoneInfo suitable for persisting in a file or
// a hex encoded string.
type ZoneInfoRecord struct {
	ZoneID    uint32
	NameIndex uint16 // index into NameData
	EraIndex  uint16 // index into ZoneErasData

	// EraCount should always be > 0 for a Zone. If set to 0, indicates that this
	// is a Link and the TargetIndex should be used to retrieve the target
	// ZoneInfo.
	EraCount uint16

	TargetIndex uint16 // index into ZoneInfosData if IsLink() is true
}
