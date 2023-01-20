package zoneinfo

// These are the data structures used on the 'zonedb' files which capture
// the information contained in the IANA TZ database.

const (
	// The minimum value of ZoneRule::FromYear and ZoneRule::ToYear. Used
	// by synthetic entries for certain zones, to guarantee that all zones have at
	// least one transition.
	MinZoneRuleYear = 0

	// The maximum value of ZoneRule::FromYear and ZoneRule::ToYear,
	// representing the sentinel value "max" in the TO and FROM columns of the
	// TZDB files. Must be less than MaxZoneEraUntilYear.
	MaxZoneRuleYear = 9999

	// The maximum value of ZoneEra::UntilYear, representing the sentinel value
	// "-" in the UNTIL column of the TZDB files. Must be greater than
	// MaxZoneRuleYear.
	MaxZoneEraUntilYear = MaxZoneRuleYear + 1
)

//-----------------------------------------------------------------------------

type ZoneRule struct {

	/** FROM year. */
	FromYear int16

	/** TO year. */
	ToYear int16

	/** Determined by the IN column. 1=Jan, 12=Dec. */
	InMonth uint8

	/**
	 * Determined by the ON column. Possible values are: 0, 1=Mon, 7=Sun.
	 * There are 4 combinations:
	 * @verbatim
	 * OnDayOfWeek=0, OnDayOfMonth=(1-31): exact match
	 * OnDayOfWeek=1-7, OnDayOfMonth=1-31: day_of_week>=OnDayOfMonth
	 * OnDayOfWeek=1-7, OnDayOfMonth=-(1-31): day_of_week<=OnDayOfMonth
	 * OnDayOfWeek=1-7, OnDayOfMonth=0: last{day_of_week}
	 * @endverbatim
	 */
	OnDayOfWeek uint8

	/**
	 * Determined by the ON column. Used with OnDayOfWeek. Possible values are:
	 * 0, 1-31, or its corresponding negative values.
	 */
	OnDayOfMonth int8

	/**
	 * Determined by the AT column in units of 15-minutes from 00:00. The range
	 * is (0 - 100) corresponding to 00:00 to 25:00.
	 */
	AtTimeCode uint8

	/**
	 * The AtTimeModifier is a packed field containing 2 pieces of info:
	 *
	 * * The upper 4 bits represent the AT time suffix: 'w', 's' or 'u',
	 *   represented by SuffixW, SuffixS and SuffixU.
	 * * The lower 4 bits represent the remaining 0-14 minutes of the AT field
	 *   after truncation into AtTimeCode. In other words, the full AT field in
	 *   one-minute resolution is (15 * AtTimeCode + (AtTimeModifier & 0x0f)).
	 */
	AtTimeModifier uint8

	/**
	 * Determined by the SAVE column, containing the offset from UTC, in 15-min
	 * increments.
	 *
	 * If the '--scope extended' flag is given to tzcompiler.py, this field
	 * should be interpreted as an uint8_t field, whose lower 4-bits hold a
	 * slightly modified value of offsetCode equal to (originalDeltaCode + 4).
	 * This allows the 4-bits to represent DST offsets from -1:00 to 2:45 in
	 * 15-minute increments. This is the same algorithm used by
	 * ZoneEra::DeltaCode field for consistency. The DeltaOffsetMinutes() method
	 * knows how to convert this field into minutes.
	 */
	DeltaCode int8

	/**
	 * Determined by the LETTER column. Determines the substitution into the '%s'
	 * field (implemented here by just a '%') of the ZoneInfo::Format field.
	 * Most comment values in the raw TZDB files are "S", "D", and "-". The "-" is
	 * stored as "" (empty string) to save memory, and because that's what the "-"
	 * in the raw file actually means.
	 *
	 * As of TZ DB version 2018i, there are 4 ZonePolicies which have ZoneRules
	 * with a LETTER field longer than 1 character:
	 *
	 *  - Belize ('CST'; used by America/Belize)
	 *  - Namibia ('WAT', 'CAT'; used by Africa/Windhoek)
	 *  - StJohns ('DD'; used by America/St_Johns and America/Goose_Bay)
	 *  - Troll ('+00' '+02'; used by Antarctica/Troll)
	 */
	Letter string
}

func (rule *ZoneRule) AtMinutes() int16 {
	return int16(rule.AtTimeCode)*15 + int16(rule.AtTimeModifier&0x0f)
}

func (rule *ZoneRule) AtSuffix() uint8 {
	return rule.AtTimeModifier & 0xf0
}

func (rule *ZoneRule) DstOffsetMinutes() int16 {
	return (int16(uint8(rule.DeltaCode)&0x0f) - 4) * 15
}

//-----------------------------------------------------------------------------

/**
 * A collection of transition rules which describe the DST rules of a given
 * administrative region. A given time zone (ZoneInfo) can follow a different
 * ZonePolicy at different times. Conversely, multiple time zones (ZoneInfo)
 * can choose to follow the same ZonePolicy at different times.
 */
type ZonePolicy struct {
	/** Slice to array of rules. */
	Rules []ZoneRule
}

//---------------------------------------------------------------------------

const (
	/** Represents 'w' or wall time. */
	SuffixW uint8 = 0x00

	/** Represents 's' or standard time. */
	SuffixS uint8 = 0x10

	/** Represents 'u' or UTC time. */
	SuffixU uint8 = 0x20
)

/**
 * An entry in ZoneInfo which describes which ZonePolicy was being followed
 * during a particular time period. Corresponds to one line of the ZONE record
 * in the TZ Database file ending with an UNTIL field. The ZonePolicy is
 * determined by the RULES column in the TZ Database file.
 *
 * There are 2 types of ZoneEra:
 *    1) ZonePolicy == nil. Then DeltaCode determines the additional
 *    offset from offsetCode. A value of '-' in the TZ Database file is stored
 *    as 0.
 *    2) ZonePolicy != nil. Then the DeltaCode offset is given by the
 *    ZoneRule.DeltaCode of the ZoneRule which matches the time instant of
 *    interest.
 */
type ZoneEra struct {
	/**
	 * Zone policy, determined by the RULES column. Set to nil if the RULES
	 * column is '-' or an explicit DST shift in the form of 'hh:mm'.
	 */
	ZonePolicy *ZonePolicy

	/**
	 * Zone abbreviations (e.g. PST, EST) determined by the FORMAT column. It has
	 * 3 encodings in the TZ DB files:
	 *
	 *  1) A fixed string, e.g. "GMT".
	 *  2) Two strings separated by a '/', e.g. "-03/-02" indicating
	 *     "{std}/{dst}" options.
	 *  3) A single string with a substitution, e.g. "E%sT", where the "%s" is
	 *  replaced by the LETTER value from the ZoneRule.
	 *
	 * BasicZoneProcessor supports only a single letter subsitution from LETTER,
	 * but ExtendedZoneProcessor supports substituting multi-character strings
	 * (e.g. "CAT", "DD", "+00").
	 *
	 * The TZ DB files use '%s' to indicate the substitution, but for simplicity,
	 * AceTime replaces the "%s" with just a '%' character with no loss of
	 * functionality. This also makes the string-replacement code a little
	 * simpler. For example, 'E%sT' is stored as 'E%T', and the LETTER
	 * substitution is performed on the '%' character.
	 *
	 * This field will never be a 'nil' if it was derived from an actual
	 * entry from the TZ database. There is an internal object named
	 * `ExtendedZoneProcessor::kAnchorEra` which does set this field to nil.
	 * Maybe it should be set to ""?
	 */
	Format string

	/** UTC offset in 15 min increments. Determined by the STDOFF column. */
	OffsetCode int8

	/**
	 * If ZonePolicy is nil, then this indicates the DST offset in 15 minute
	 * increments as defined by the RULES column in 'hh:mm' format. If the
	 * 'RULES' column is '-', then the DeltaCode is 0.
	 *
	 * If the '--scope extended' flag is given to tzcompiler.py, the 'DeltaCode`
	 * should be interpreted as a uint8_t field, composed of two 4-bit fields:
	 *
	 *    * The upper 4-bits is an unsigned integer from 0 to 14 that represents
	 *    the one-minute remainder from the OffsetCode. This allows us to capture
	 *    STDOFF offsets in 1-minute resolution.
	 *    * The lower 4-bits is an unsigned integer that holds (originalDeltaCode
	 *    + 4). This allows us to represent DST offsets from -1:00 to +2:45, in
	 *    15-minute increments.
	 *
	 * The StdOffsetMinutes() and DstOffsetMinutes() functions know how to convert
	 * OffsetCode and DeltaCode into the appropriate minutes.
	 */
	DeltaCode uint8

	/**
	 * Era is valid until currentTime < UntilYear. Comes from the UNTIL column.
	 */
	UntilYear int16

	/** The month field in UNTIL (1-12). Will never be 0. */
	UntilMonth uint8

	/**
	 * The day field in UNTIL (1-31). Will never be 0. Also, there's no need for
	 * UntilDayOfWeek, because the database generator will resolve the exact day
	 * of month based on the known year and month.
	 */
	UntilDay uint8

	/**
	 * The time field of UNTIL field in 15-minute increments. A range of 00:00 to
	 * 25:00 corresponds to 0-100.
	 */
	UntilTimeCode uint8

	/**
	 * The UntilTimeModifier is a packed field containing 2 pieces of info:
	 *
	 *    * The upper 4 bits represent the UNTIL time suffix: 'w', 's' or 'u',
	 *    represented by SuffixW, SuffixS and SuffixU.
	 *    * The lower 4 bits represent the remaining 0-14 minutes of the UNTIL
	 *    field after truncation into UntilTimeCode. In other words, the full
	 *    UNTIL field in one-minute resolution is (15 * UntilTimeCode +
	 *    (UntilTimeModifier & 0x0f)).
	 */
	UntilTimeModifier uint8
}

func (era *ZoneEra) StdOffsetMinutes() int16 {
	return int16(era.OffsetCode)*15 + int16((era.DeltaCode&0xf0)>>4)
}

func (era *ZoneEra) DstOffsetMinutes() int16 {
	return int16((int8(era.DeltaCode&0x0f) - 4) * 15)
}

func (era *ZoneEra) UntilMinutes() int16 {
	return int16(era.UntilTimeCode)*15 + int16(era.UntilTimeModifier&0x0f)
}

func (era *ZoneEra) UntilSuffix() uint8 {
	return era.UntilTimeModifier & 0xf0
}

//-----------------------------------------------------------------------------

/**
 * Representation of a given time zone, implemented as an array of ZoneEra
 * records.
 */
type ZoneInfo struct {
	/** Full name of zone (e.g. "America/Los_Angeles"). */
	Name string

	/**
	 * Unique, stable ID of the zone name, created from a hash of the name.
	 * This ID will never change once assigned. This can be used for presistence
	 * and serialization.
	 */
	ZoneID uint32

	/** Start year of the zone files. */
	StartYear int16

	/** Until year of the zone files. */
	UntilYear int16

	/**
	 * A slice of ZoneEra instances. For a normal Zone, num(Eras) is greater than
	 * 0, and the ZoneEra entries are arranged in increasing order
	 * of UNTIL time. For a Link entry, num(Eras) == 0, and instead 'target' is
	 * non-nil and points to the target Zone. We have to follow the indirect
	 * pointer, and resolve the target.Eras to obtain the actual ZoneEra entries.
	 */
	Eras []ZoneEra

	/** If not nil, this entry is a Link to the target. */
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
