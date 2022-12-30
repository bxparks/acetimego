package acetime

// These are the data structures used on the 'zonedb' files which capture
// the information contained in the IANA TZ database.

const (
	// The minimum value of ZoneRule::fromYear and ZoneRule::toYear. Used
	// by synthetic entries for certain zones, to guarantee that all zones have at
	// least one transition.
	minZoneRuleYear = 0

	// The maximum value of ZoneRule::fromYear and ZoneRule::toYear,
	// representing the sentinel value "max" in the TO and FROM columns of the
	// TZDB files. Must be less than kAtcMaxZoneEraUntilYear.
	maxZoneRuleYear = 9999

	// The maximum value of ZoneEra::untilYear, representing the sentinel value
	// "-" in the UNTIL column of the TZDB files. Must be greater than
	// kAtcMaxZoneRuleYear.
	maxZoneEraUntilYear = maxZoneRuleYear + 1
)

//-----------------------------------------------------------------------------

type ZoneRule struct {

	/** FROM year. */
	fromYear int16

	/** TO year. */
	toYear int16

	/** Determined by the IN column. 1=Jan, 12=Dec. */
	inMonth uint8

	/**
	 * Determined by the ON column. Possible values are: 0, 1=Mon, 7=Sun.
	 * There are 4 combinations:
	 * @verbatim
	 * onDayOfWeek=0, onDayOfMonth=(1-31): exact match
	 * onDayOfWeek=1-7, onDayOfMonth=1-31: day_of_week>=onDayOfMonth
	 * onDayOfWeek=1-7, onDayOfMonth=-(1-31): day_of_week<=onDayOfMonth
	 * onDayOfWeek=1-7, onDayOfMonth=0: last{day_of_week}
	 * @endverbatim
	 */
	onDayOfWeek uint8

	/**
	 * Determined by the ON column. Used with onDayOfWeek. Possible values are:
	 * 0, 1-31, or its corresponding negative values.
	 */
	onDayOfMonth int8

	/**
	 * Determined by the AT column in units of 15-minutes from 00:00. The range
	 * is (0 - 100) corresponding to 00:00 to 25:00.
	 */
	atTimeCode uint8

	/**
	   * The atTimeModifier is a packed field containing 2 pieces of info:
	   *
	   * * The upper 4 bits represent the AT time suffix: 'w', 's' or 'u',
	   *   represented by suffixW, suffixS and suffixU.
	   * * The lower 4 bits represent the remaining 0-14 minutes of the AT field
		 *   after truncation into atTimeCode. In other words, the full AT field in
		 *   one-minute resolution is (15 * atTimeCode + (atTimeModifier & 0x0f)).
	*/
	atTimeModifier uint8

	/**
	 * Determined by the SAVE column, containing the offset from UTC, in 15-min
	 * increments.
	 *
	 * If the '--scope extended' flag is given to tzcompiler.py, this field
	 * should be interpreted as an uint8_t field, whose lower 4-bits hold a
	 * slightly modified value of offsetCode equal to (originalDeltaCode + 4).
	 * This allows the 4-bits to represent DST offsets from -1:00 to 2:45 in
	 * 15-minute increments. This is the same algorithm used by
	 * ZoneEra::deltaCode field for consistency. The
	 * extended::ZonePolicyBroker::deltaMinutes() method knows how to convert
	 * this field into minutes.
	 */
	deltaCode int8

	/**
	 * Determined by the LETTER column. Determines the substitution into the '%s'
	 * field (implemented here by just a '%') of the ZoneInfo::format field.
	 * Possible values are 'S', 'D', '-', or a number < 32 (i.e. a non-printable
	 * character). If the value is < 32, then this number is an index offset into
	 * the ZonePolicy.letters[] array which contains a (char*) of the
	 * actual multi-character letter.
	 *
	 * BasicZoneProcessor supports only a single LETTER value (i.e. >= 32), which
	 * also means that ZonePolicy.numLetters will always be 0 for a
	 * BasicZoneProcessor. ExtendedZoenProcessor supports a LETTER value of < 32,
	 * indicating a multi-character string.
	 *
	 * As of TZ DB version 2018i, there are 4 ZonePolicies which have ZoneRules
	 * with a LETTER field longer than 1 character:
	 *
	 *  - Belize ('CST'; used by America/Belize)
	 *  - Namibia ('WAT', 'CAT'; used by Africa/Windhoek)
	 *  - StJohns ('DD'; used by America/St_Johns and America/Goose_Bay)
	 *  - Troll ('+00' '+02'; used by Antarctica/Troll)
	 */
	letter uint8
}

//-----------------------------------------------------------------------------

/**
 * A collection of transition rules which describe the DST rules of a given
 * administrative region. A given time zone (ZoneInfo) can follow a different
 * ZonePolicy at different times. Conversely, multiple time zones (ZoneInfo)
 * can choose to follow the same ZonePolicy at different times.
 *
 * If numLetters is non-zero, then 'letters' will be a pointer to an array of
 * (char*) pointers. Any ZoneRule.letter < 32 (i.e. non-printable) will
 * be an offset into this array of pointers.
 */
type ZonePolicy struct {
	/** Pointer to array of rules. */
	rules []ZoneRule

	/** Pointer to an array of DST letters (e.g. "D", "S"). */
	letters []string
}

//---------------------------------------------------------------------------

const (
	/** Represents 'w' or wall time. */
	suffixW uint8 = 0x00

	/** Represents 's' or standard time. */
	suffixS uint8 = 0x10

	/** Represents 'u' or UTC time. */
	suffixU uint8 = 0x20
)

/**
 * An entry in ZoneInfo which describes which ZonePolicy was being followed
 * during a particular time period. Corresponds to one line of the ZONE record
 * in the TZ Database file ending with an UNTIL field. The ZonePolicy is
 * determined by the RULES column in the TZ Database file.
 *
 * There are 2 types of ZoneEra:
 *    1) zonePolicy == nullptr. Then deltaCode determines the additional
 *    offset from offsetCode. A value of '-' in the TZ Database file is stored
 *    as 0.
 *    2) zonePolicy != nullptr. Then the deltaCode offset is given by the
 *    ZoneRule.deltaCode of the ZoneRule which matches the time instant of
 *    interest.
 */
type ZoneEra struct {
	/**
	 * Zone policy, determined by the RULES column. Set to nullptr if the RULES
	 * column is '-' or an explicit DST shift in the form of 'hh:mm'.
	 */
	zonePolicy *ZonePolicy

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
	 * This field will never be a 'nullptr' if it was derived from an actual
	 * entry from the TZ database. There is an internal object named
	 * `ExtendedZoneProcessor::kAnchorEra` which does set this field to nullptr.
	 * Maybe it should be set to ""?
	 */
	format string

	/** UTC offset in 15 min increments. Determined by the STDOFF column. */
	offsetCode int8

	/**
	 * If zonePolicy is nullptr, then this indicates the DST offset in 15 minute
	 * increments as defined by the RULES column in 'hh:mm' format. If the
	 * 'RULES' column is '-', then the deltaCode is 0.
	 *
	 * If the '--scope extended' flag is given to tzcompiler.py, the 'deltaCode`
	 * should be interpreted as a uint8_t field, composed of two 4-bit fields:
	 *
	 *    * The upper 4-bits is an unsigned integer from 0 to 14 that represents
	 *    the one-minute remainder from the offsetCode. This allows us to capture
	 *    STDOFF offsets in 1-minute resolution.
	 *    * The lower 4-bits is an unsigned integer that holds (originalDeltaCode
	 *    + 4). This allows us to represent DST offsets from -1:00 to +2:45, in
	 *    15-minute increments.
	 *
	 * The extended::ZoneEraBroker::deltaMinutes() and offsetMinutes() know how
	 * to convert offsetCode and deltaCode into the appropriate minutes.
	 */
	deltaCode int8

	/**
	 * Era is valid until currentTime < untilYear. Comes from the UNTIL column.
	 */
	untilYear int16

	/** The month field in UNTIL (1-12). Will never be 0. */
	untilMonth uint8

	/**
	 * The day field in UNTIL (1-31). Will never be 0. Also, there's no need for
	 * untilDayOfWeek, because the database generator will resolve the exact day
	 * of month based on the known year and month.
	 */
	untilDay uint8

	/**
	 * The time field of UNTIL field in 15-minute increments. A range of 00:00 to
	 * 25:00 corresponds to 0-100.
	 */
	untilTimeCode uint8

	/**
	 * The untilTimeModifier is a packed field containing 2 pieces of info:
	 *
	 *    * The upper 4 bits represent the UNTIL time suffix: 'w', 's' or 'u',
	 *    represented by suffixW, suffixS and suffixU.
	 *    * The lower 4 bits represent the remaining 0-14 minutes of the UNTIL
	 *    field after truncation into untilTimeCode. In other words, the full
	 *    UNTIL field in one-minute resolution is (15 * untilTimeCode +
	 *    (untilTimeModifier & 0x0f)).
	 */
	untilTimeModifier uint8
}

func (era *ZoneEra) StdOffsetMinutes() int16 {
  return int16(era.offsetCode) * 15 +
		int16((uint8(era.deltaCode) & 0xf0) >> 4)
}

//-----------------------------------------------------------------------------

/**
 * Representation of a given time zone, implemented as an array of ZoneEra
 * records.
 */
type ZoneInfo struct {
	/** Full name of zone (e.g. "America/Los_Angeles"). */
	name string

	/**
	 * Unique, stable ID of the zone name, created from a hash of the name.
	 * This ID will never change once assigned. This can be used for presistence
	 * and serialization.
	 */
	zoneID uint32

	/** Start year of the zone files. */
	startYear int16

	/** Until year of the zone files. */
	untilYear int16

	/**
	 * A `ZoneEras*` pointer or a `const ZoneInfo*` pointer. For a normal
	 * Zone, numEras is greater than 0, and this field is a pointer to the
	 * ZoneEra entries in increasing order of UNTIL time. For a Link entry,
	 * numEras == 0, and this field provides a level of indirection to a (const
	 * ZoneInfo*) pointer to the target Zone. We have to follow the indirect
	 * pointer, and resolve the target numEras and eras to obtain the actual
	 * ZoneEra entries.
	 */
	eras []ZoneEra

	/** If not nil, this entry is a Link to the target. */
	target *ZoneInfo
}
