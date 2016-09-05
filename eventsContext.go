package calendar

import (
	"encoding/json"
	"fmt"

	"github.com/jseely/bst"
)

type EventsContext struct {
	events *bst.BinarySearchTree
}

func NewEventsContext() *EventsContext {
	return &EventsContext{
		events: bst.New(EventSortFunc),
	}
}

func (ec *EventsContext) Count() int {
	return ec.events.Count()
}

func (ec *EventsContext) Add(event *Event) error {
	return ec.events.Insert(event)
}

func (ec *EventsContext) First() (*Event, error) {
	firstEvent, wasEvent := ec.events.First().(*Event)
	if !wasEvent {
		return nil, fmt.Errorf("Bad State: Item was not of type *Event")
	}
	return firstEvent, nil
}

func (ec *EventsContext) MarshalJSON() ([]byte, error) {
	eInterface := ec.events.SortedSlice()
	events := make([]*Event, len(eInterface))
	var ok bool
	for i := range eInterface {
		if events[i], ok = eInterface[i].(*Event); !ok {
			return nil, fmt.Errorf("Failed to casts element (%v) to *Event", i)
		}
	}
	return json.Marshal(events)
}

func (ec *EventsContext) UnmarshalJSON(data []byte) error {
	var events []*Event
	if err := json.Unmarshal(data, &events); err != nil {
		return err
	}
	for i := range events {
		if err := ec.events.Insert(events[i]); err != nil {
			return err
		}
	}
	return nil
}
