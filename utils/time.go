package utils

import "time"

// DayStart returns the start of the current day
func DayStart(now time.Time) time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// DayEnd returns the end of the current day
// This is technically just the start of the following day, because it's easier
func DayEnd(now time.Time) time.Time {
	return DayStart(now.Add(time.Hour * 24))
}

// HourStart returns the start of the current hour
func HourStart(now time.Time) time.Time {
	y, m, d := now.Date()
	h := now.Hour()
	return time.Date(y, m, d, h, 0, 0, 0, time.Local)
}

// Returns the end of the current hour
// This works akin to DayEnd in that it's just the start of the following hour
func HourEnd(now time.Time) time.Time {
	return HourStart(now.Add(time.Hour))
}
