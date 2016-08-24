package calendar

import (
	"testing"
	"time"
)

func TestNewCalendarAddEvent(t *testing.T) {
	cal := NewCalendar()
	// NewEvent(Name, Start, End, Organizer, Location, Description)
	event := NewEvent("Event Name", time.Now(), time.Now(), nil, nil, "")
	cal.Events.Add(event)
	if event != cal.Events.First() {
		t.Errorf("First event in calendar was not a reference to the given event")
	}
}
