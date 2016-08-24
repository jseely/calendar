package calendar

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

type Event struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Start       time.Time  `json:"start"`
	End         time.Time  `json:"end"`
	Modified    time.Time  `json:"modified"`
	Organizers  []*Contact `json:"organizers"`
	Location    *Location  `json:"location"`
	Description string     `json:"description"`
}

func NewEvent(name string, start, end time.Time, organizer *Contact, location *Location, description string) *Event {
	event := &Event{
		ID:          uuid.NewV4(),
		Name:        name,
		Start:       start,
		End:         end,
		Modified:    time.Now(),
		Organizers:  make([]*Contact, 0),
		Location:    location,
		Description: description,
	}
	if organizer != nil {
		event.Organizers = append(event.Organizers, organizer)
	}
	return event
}

func (e *Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		*Alias
		Start    int64 `json:"start"`
		End      int64 `json:"end"`
		Modified int64 `json:"modified"`
	}{
		Start:    e.Start.UnixNano(),
		End:      e.End.UnixNano(),
		Modified: e.Modified.UnixNano(),
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
	e.Start = time.Unix(0, aux.Start)
	e.End = time.Unix(0, aux.End)
	e.Modified = time.Unix(0, aux.Modified)
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
	if ea.Start.Before(eb.Start) {
		return -1, nil
	} else {
		return 1, nil
	}
}
