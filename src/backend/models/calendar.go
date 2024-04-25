package integrations

import "time"

type Event struct {
	Title string
	Start time.Time
	End   time.Time
}
