package calendar

import (
	"encoding/json"
)

type Location struct{}

func (l *Location) MarshalJSON() ([]byte, error) {
	return json.Marshal("")
}

func (l *Location) UnmarshalJSON(data []byte) error {
	return nil
}
