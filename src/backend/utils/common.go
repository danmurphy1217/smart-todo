package utils

import "time"

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	// get start of tomorrow
	date = StartOfDay(date.AddDate(0, 0, 1))
	return date.Add(-1 * time.Nanosecond)
}
