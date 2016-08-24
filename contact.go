package calendar

import (
	"encoding/json"
)

type Contact struct{}

func (c *Contact) MarshalJSON() ([]byte, error) {
	return json.Marshal("")
}

func (c *Contact) UnmarshalJSON(data []byte) error {
	return nil
}
