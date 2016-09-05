package calendar

import (
	"encoding/json"
	"io"
)

type Calendar struct {
	Events *EventsContext `json:"events"`
}

func NewCalendar() *Calendar {
	return &Calendar{
		Events: NewEventsContext(),
	}
}

func (c *Calendar) Save(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(c)
}

func (c *Calendar) Load(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(c)
}
