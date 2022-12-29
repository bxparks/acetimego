package acetime

type DateTuple struct {
	/** [0,10000] */
	year int16

	/** [1-12] */
	month uint8

	/** [1-31] */
	day uint8

	/** negative values allowed */
	minutes int16

	/** suffixS, suffixW, suffixU */
	suffix uint8
}

func dateTupleCompare(a *DateTuple, b *DateTuple) int8 {
	if a.year < b.year {
		return -1
	}
	if a.year > b.year {
		return 1
	}
	if a.month < b.month {
		return -1
	}
	if a.month > b.month {
		return 1
	}
	if a.day < b.day {
		return -1
	}
	if a.day > b.day {
		return 1
	}
	if a.minutes < b.minutes {
		return -1
	}
	if a.minutes > b.minutes {
		return 1
	}
	return 0
}

//-----------------------------------------------------------------------------

type MatchingEra struct {
	/**
	 * The effective start time of the matching ZoneEra, which uses the
	 * UTC offsets of the previous matching era.
	 */
	startDt DateTuple

	/** The effective until time of the matching ZoneEra. */
	untilDt DateTuple

	/** The ZoneEra that matched the given year. NonNullable. */
	era *ZoneEra

	/** The previous MatchingEra, needed to interpret start_dt.  */
	prevMatch *MatchingEra

	/** The STD offset of the last Transition in this MatchingEra. */
	lastOffsetMinutes int16

	/** The DST offset of the last Transition in this MatchingEra. */
	lastDeltaMinutes int16
}

type Transition struct {
	/** The matching_era which generated this Transition. */
	match *MatchingEra

	/**
	 * The Zone transition rule that matched for the the given year. Set to
	 * nullptr if the RULES column is '-', indicating that the MatchingEra was
	 * a "simple" ZoneEra.
	 */
	rule *ZoneRule

	/**
	 * The original transition time, usually 'w' but sometimes 's' or 'u'. After
	 * expandDateTuple() is called, this field will definitely be a 'w'. We must
	 * remember that the transition_time* fields are expressed using the UTC
	 * offset of the *previous* Transition.
	 */
	transitionTime DateTuple

	//union {
	/**
	* Version of transition_time in 's' mode, using the UTC offset of the
	* *previous* Transition. Valid before
	* ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	transitionTimeS DateTuple

	/**
	* Start time expressed using the UTC offset of the current Transition.
	* Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	startDt DateTuple
	//}

	//union {
	/**
	* Version of transition_time in 'u' mode, using the UTC offset of the
	* *previous* transition. Valid before
	* ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	transitionTimeU DateTuple

	/**
	* Until time expressed using the UTC offset of the current Transition.
	* Valid after ExtendedZoneProcessor::generateStartUntilTimes() is called.
	 */
	untilDt DateTuple
	//}

	/** The calculated transition time of the given rule. */
	startEpochSeconds int32

	/**
	 * The base offset minutes, not the total effective UTC offset. Note that
	 * this is different than basic::Transition::offsetMinutes used by
	 * BasicZoneProcessor which is the total effective offsetMinutes. (It may be
	 * possible to make this into an effective offsetMinutes (i.e. offsetMinutes
	 * + deltaMinutes) but it does not seem worth making that change right now.)
	 */
	offsetMinutes int16

	/** The DST delta minutes. */
	deltaMinutes int16

	/** The calculated effective time zone abbreviation, e.g. "PST" or "PDT". */
	abbrev string

	/** Storage for the single letter 'letter' field if 'rule' is not null. */
	letter string

	/**
	 * During findCandidateTransitions(), this flag indicates whether the
	 * current transition is a valid "prior" transition that occurs before other
	 * transitions. It is set by setFreeAgentAsPriorIfValid() if the transition
	 * is a prior transition.
	 */
	isValidPrior bool

	/**
	 * During processTransitionMatchStatus(), this flag indicates how the
	 * transition falls within the time interval of the MatchingEra.
	 */
	matchStatus bool
}

type TransitionStorage struct {
	/** Pointers into the pool of Transition objects. */
	transitions []Transition
	/** Index of the most recent prior transition [0,kTransitionStorageSize) */
	indexPrior uint8
	/** Index of the candidate pool [0,kTransitionStorageSize) */
	indexCandidate uint8
	/** Index of the free agent transition [0, kTransitionStorageSize) */
	indexFree uint8

	/** Number of allocated transitions. */
	allocSize uint8
}

//-----------------------------------------------------------------------------

type TimeZone struct {
	zoneInfo          *ZoneInfo
	year              int16
	isFilled          bool
	matches           []MatchingEra
	transitionStorage TransitionStorage
}

// TimeZoneFromZoneInfo creates a new TimeZone from the given ZoneInfo instance.
func TimeZoneFromZoneInfo(zoneInfo *ZoneInfo) *TimeZone {
	return &TimeZone{zoneInfo: zoneInfo}
}
