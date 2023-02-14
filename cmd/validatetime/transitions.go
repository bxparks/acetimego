package main

import (
	"strconv"
	"time"
)

// A tuple that represents the before and after time of a timezone transition.
type TransitionTimes struct {
	before time.Time
	after  time.Time
}

// findTransitions() finds the timezone transitions and returns an array of
// tuples of (before, after).
func findTransitions(
	startYear int,
	untilYear int,
	samplingInterval int, // hours
	tz *time.Location) []TransitionTimes {

	dt := time.Date(startYear, 1, 1, 0, 0, 0, 0, time.UTC)
	dtLocal := dt.In(tz)

	transitions := make([]TransitionTimes, 0, 500)
	samplingIntervalNanos := samplingInterval * 3600 * 1000000000
	intervalDuration := time.Duration(samplingIntervalNanos)
	for {
		dtNext := dt.Add(intervalDuration)
		dtNextLocal := dtNext.In(tz)
		if dtNextLocal.Year() >= untilYear {
			break
		}

		// Look for utc offset transition
		if isTransition(dtLocal, dtNextLocal) {
			dtLeft, dtRight := binarySearchTransition(tz, dt, dtNext)
			dtLeftLocal := dtLeft.In(tz)
			dtRightLocal := dtRight.In(tz)
			transitions = append(
				transitions, TransitionTimes{dtLeftLocal, dtRightLocal})
		}

		dt = dtNext
		dtLocal = dtNextLocal
	}

	return transitions
}

func isTransition(before time.Time, after time.Time) bool {
	return utcOffsetMinutes(before) != utcOffsetMinutes(after)
}

func binarySearchTransition(
	tz *time.Location,
	dtLeft time.Time,
	dtRight time.Time) (time.Time, time.Time) {

	dtLeftLocal := dtLeft.In(tz)
	for {
		duration := dtRight.Sub(dtLeft)
		durationMinutes := int(duration.Minutes())
		deltaMinutes := durationMinutes / 2
		if deltaMinutes == 0 {
			break
		}

		dtMid := dtLeft.Add(time.Duration(deltaMinutes * 60 * 1000000000))
		dtMidLocal := dtMid.In(tz)
		if isTransition(dtLeftLocal, dtMidLocal) {
			dtRight = dtMid
		} else {
			dtLeft = dtMid
			dtLeftLocal = dtMidLocal
		}
	}
	return dtLeft, dtRight
}

// utcOffsetMinutes() returns the UTC offset of the given time as minutes.
func utcOffsetMinutes(t time.Time) int {
	utcOffsetString := t.Format("-07:00")
	return convertOffsetStringToMinutes(utcOffsetString)
}

func convertOffsetStringToMinutes(offset string) int {
	hourString := offset[0:3]
	minuteString := offset[4:6]
	hour, _ := strconv.Atoi(hourString)
	minute, _ := strconv.Atoi(minuteString)
	return convertHourMinuteToMinutes(hour, minute)
}

func convertHourMinuteToMinutes(hour int, minute int) int {
	sign := 1
	if hour < 0 {
		sign = -1
		hour = -hour
	}
	minutes := hour*60 + minute
	return sign * minutes
}
