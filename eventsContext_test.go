package calendar

import (
	"testing"
	"time"
)

func TestNewEventsContext(t *testing.T) {
	ec := NewEventsContext()
	if ec.events == nil {
		t.Errorf("Constructor did not initialize 'events'")
	}
}

func TestAddEventAndCount(t *testing.T) {
	ec := NewEventsContext()
	// NewEvent(Name, Start, End, Organizer, Location, Description)
	event := NewEvent("Event Name", time.Now(), time.Now(), nil, nil, "")
	ec.Add(event)

	firstEvent, err := ec.First()
	if err != nil {
		t.Errorf("This is really bad. You seriously fucked up! (%v)", err.Error())
	}
	if event != firstEvent {
		t.Errorf("First event in calendar was not a reference to the given event")
	}

	if ec.Count() != 1 {
		t.Errorf("Count() returned %v, expected %v", ec.Count(), 1)
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	ec := NewEventsContext()
	ec.Add(NewEvent("Event1", time.Now(), time.Now(), nil, nil, ""))
	ec.Add(NewEvent("Event2", time.Now(), time.Now(), nil, nil, ""))

	ecBytes, err := ec.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshaling eventsContext to JSON: %v", err.Error())
	}

	ec2 := NewEventsContext()
	if err = ec2.UnmarshalJSON(ecBytes); err != nil {
		t.Errorf("Error unmarshaling eventsContext from JSON: %v", err.Error())
	}

	if ec2.Count() != 2 {
		t.Errorf("Recovered %v events, expected %v", ec2.Count(), 2)
	}
}

/*
func TestCalendarAfter(t *testing.T) {
	now := time.Now().UTC()
	cal := NewCalendar()
	cal.Events.Add(NewEvent(
		"Event1",
		now.Add(-time.Hour),
		now.Add(-time.Hour+time.Minute),
		nil,
		nil,
		""))
	cal.Events.Add(NewEvent(
		"Event2",
		now,
		now.Add(time.Minute),
		nil,
		nil,
		""))
	cal.Events.Add(NewEvent(
		"Event3",
		now.Add(time.Hour),
		now.Add(time.Hour+time.Minute),
		nil,
		nil,
		""))
	cal.Events.Add(NewEvent(
		"Event4",
		now.Add(time.Hour*2),
		now.Add(time.Hour*2+time.Minute),
		nil,
		nil,
		""))
	events := cal.Events.AtOrAfter(now)
	if len(events) != 3 {
		t.Errorf("Expected 3 events got (%v)", len(events))
	}

	eventOrder := "Event2 => Event3 => Event4"
	var actualOrder string
	for i := range events {
		event, ok := events.(*Event)
		if !ok {
			t.Errorf("Could not cast element at (%v) to *Event", i)
		}
		actualOrder += event.Name
		if i+1 != len(events) {
			actualOrder += " => "
		}
	}
	if eventOrder != actualOrder {
		t.Errorf("Event order (%v) did not match expected event order (%v)", actualOrder, eventOrder)
	}
}
*/
