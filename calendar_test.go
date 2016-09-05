package calendar

import (
	"bytes"
	"testing"
	"time"
)

func TestNewCalendar(t *testing.T) {
	cal := NewCalendar()
	if cal.Events == nil {
		t.Errorf("NewCalendar() did not instantiate an EventsContext")
	}
}

func TestCalendarSaveAndLoad(t *testing.T) {
	cal1 := NewCalendar()
	cal1.Events.Add(NewEvent("EventName", time.Now(), time.Now(), nil, nil, ""))

	buffer := new(bytes.Buffer)
	if err := cal1.Save(buffer); err != nil {
		t.Errorf("Error saving calendar: %v", err.Error())
	}

	cal2 := NewCalendar()
	if err := cal2.Load(buffer); err != nil {
		t.Errorf("Error loading calendar: %v", err.Error())
	}

	if cal2.Events.Count() != 1 {
		t.Errorf("Recovered %v events, expected %v", cal2.Events.Count(), 1)
	}
}
