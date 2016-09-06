package calendar

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

type EventDataStore interface {
	OnEventCreated(event *Event)
	OnEventUpdated(event *Event)
	OnEventDeleted(event *Event)

	PopulateEvent(event *Event)
}

type Event struct {
	id          uuid.UUID  `json:"id"`
	name        string     `json:"name"`
	start       time.Time  `json:"start"`
	end         time.Time  `json:"end"`
	modified    time.Time  `json:"modified"`
	organizers  []*Contact `json:"organizers"`
	location    *Location  `json:"location"`
	description string     `json:"description"`

	dataStore EventDataStore `json:"-"`
}

func NewEvent(name string, start, end time.Time, organizers []*Contact, location *Location, description string) *Event {
	event := &Event{
		id:          uuid.NewV4(),
		name:        name,
		start:       start,
		end:         end,
		modified:    time.Now(),
		location:    location,
		description: description,
	}
	event.organizers = make([]*Contact, len(organizers))
	copy(event.organizers, organizers)
	return event
}

func (e Event) Id() uuid.UUID {
	return e.id
}

func (e Event) Name() string {
	return e.name
}

func (e *Event) SetName(name string) {
	e.name = name
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e Event) Start() time.Time {
	return e.start
}

func (e *Event) SetStart(start time.Time) {
	e.start = start
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e Event) End() time.Time {
	return e.end
}

func (e *Event) SetEnd(end time.Time) {
	e.end = end
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e Event) Modified() time.Time {
	return e.modified
}

func (e Event) Organizers() []*Contact {
	organizers := make([]*Contact, len(e.organizers))
	copy(organizers, e.organizers)
	return organizers
}

func (e *Event) SetOrganizers(organizers []*Contact) {
	e.organizers = make([]*Contact, len(organizers))
	copy(e.organizers, organizers)
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e Event) Location() *Location {
	return e.location
}

func (e *Event) SetLocation(location *Location) {
	e.location = location
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e Event) Description() string {
	return e.description
}

func (e *Event) SetDescription(description string) {
	e.description = description
	e.modified = time.Now().UTC()
	if e.dataStore != nil {
		e.dataStore.OnEventUpdated(e)
	}
}

func (e *Event) SetDataStore(dataStore EventDataStore) {
	e.dataStore = dataStore
}

func (e1 *Event) Equals(e2 *Event) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	return e1.id == e2.id && e1.modified == e2.modified
}

func (e *Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		start    int64 `json:"start"`
		end      int64 `json:"end"`
		modified int64 `json:"modified"`
		*Alias
	}{
		start:    e.start.UnixNano(),
		end:      e.end.UnixNano(),
		modified: e.modified.UnixNano(),
		Alias:    (*Alias)(e),
	})
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		*Alias
		Start    int64 `json:"start"`
		End      int64 `json:"end"`
		Modified int64 `json:"modified"`
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	e.start = time.Unix(0, aux.Start)
	e.end = time.Unix(0, aux.End)
	e.modified = time.Unix(0, aux.Modified)
	return nil
}

func EventSortFunc(a, b interface{}) (int, error) {
	ea, oka := a.(*Event)
	eb, okb := b.(*Event)
	if !oka {
		return 0, errors.New("a is not an Event")
	}
	if !okb {
		return 0, errors.New("b is not an Event")
	}
	if ea.start.Before(eb.start) {
		return -1, nil
	} else {
		return 1, nil
	}
}
