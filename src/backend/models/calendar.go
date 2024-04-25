package models

import (
	"time"

	gcal "google.golang.org/api/calendar/v3"
)

type Event struct {
	Title string
	Start time.Time
	End   time.Time
}

func ToEvents(items []*gcal.Event) []Event {
	var events []Event
	for _, item := range items {
		events = append(events, ToEvent(item))
	}
	return events
}

func ToEvent(item *gcal.Event) Event {
	startTime, _ := time.Parse(time.RFC3339, item.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, item.End.DateTime)
	return Event{
		Title: item.Summary,
		Start: startTime,
		End:   endTime,
	}
}
