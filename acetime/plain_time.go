package acetime

// PlainTimeToSeconds converts (hour, minute, second) to seconds.
func PlainTimeToSeconds(hour uint8, minute uint8, second uint8) int32 {
	return ((int32(hour)*60)+int32(minute))*60 + int32(second)
}

// PlainTimeFromSeconds extracts the (hour, minute, second) components from the
// seconds. The seconds is assumed to be positive.
func PlainTimeFromSeconds(seconds int32) (
	hour uint8, minute uint8, second uint8) {

	// The compiler will probably combine the mod (%) and division (/) operations
	// into a single (dividend, remainder) function call.
	second = uint8(seconds % 60)
	minutes := seconds / 60
	minute = uint8(minutes % 60)
	hour = uint8(minutes / 60)
	return
}
