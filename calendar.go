package calendar

import (
	"github.com/jseely/bst"
)

type Calendar struct {
	Events EventsContext
}

func NewCalendar() *Calendar {
	return &Calendar{
		Events: EventsContext{events: bst.New(EventSortFunc)},
	}
}

type EventsContext struct {
	events *bst.BinarySearchTree
}

func (ec *EventsContext) Add(event *Event) error {
	return ec.events.Insert(event)
}

func (ec *EventsContext) First() *Event {
	firstEvent, wasEvent := ec.events.First().(*Event)
	if !wasEvent {
		return nil
	}
	return firstEvent
}
