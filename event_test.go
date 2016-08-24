package calendar

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestEventMarshalUnmarshalJSON(t *testing.T) {
	orig := NewEvent("test event", time.Now(), time.Now(), &Contact{}, &Location{}, "this is a \"description\"")
	origBytes, err := json.Marshal(orig)
	if err != nil {
		t.Errorf(err.Error())
	}
	copy := &Event{}
	err = json.Unmarshal(origBytes, copy)
	if err != nil {
		t.Errorf(err.Error())
	}
	copyBytes, err := json.Marshal(copy)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !bytes.Equal(origBytes, copyBytes) {
		t.Errorf("Original json does not match copy json\n%v\n%v", string(origBytes), string(copyBytes))
	}
}
