package calendar

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/satori/go.uuid"
)

type eventDataStoreRecorder struct {
	lastCallEvent  *Event
	lastCallMethod string
}

func TestEventGettersAndSetters(t *testing.T) {
	ds := &eventDataStoreRecorder{}
	now := time.Now().UTC()
	event := NewEvent("Original name", now, now, nil, nil, "original description")

	if event.SetDataStore(ds); event.dataStore == nil {
		t.Errorf("Failed to set dataStore")
	}

	if event.Id() == uuid.Nil {
		t.Errorf("Failed to get event Id")
	}

	ds.CheckDataStoreRecord(t, nil, "")

	event.SetName("New Name")
	if event.Name() != "New Name" {
		t.Errorf("Event name (%v) did not match expected (%v)", event.Name(), "New Name")
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")

	newTime := time.Now().UTC().Add(15 * time.Minute)
	event.SetStart(newTime)
	if event.Start() != newTime {
		t.Errorf("Event start (%v) did not match expected (%v)", event.Start(), newTime)
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")

	newTime = time.Now().UTC().Add(20 * time.Hour)
	event.SetEnd(newTime)
	if event.End() != newTime {
		t.Errorf("Event end (%v) did not match expected (%v)", event.End(), newTime)
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")

	organizers := []*Contact{&Contact{}, &Contact{}}
	event.SetOrganizers(organizers)
	eventsOrganizers := event.Organizers()
	if len(organizers) != len(eventsOrganizers) {
		t.Errorf("Event organizers length (%v) did not match expected length (%v)", len(organizers), len(eventsOrganizers))
	} else {
		for i := range organizers {
			if organizers[i] != eventsOrganizers[i] {
				t.Errorf("Event organizer at index %v (%v) did not match expected (%v)", i, eventsOrganizers[i], organizers[i])
			}
		}
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")

	newLocation := &Location{}
	event.SetLocation(newLocation)
	if event.Location() != newLocation {
		t.Errorf("Event location (%v) did not match expected (%v)", event.Location(), newLocation)
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")

	newDescription := "new description"
	event.SetDescription(newDescription)
	if event.Description() != newDescription {
		t.Errorf("Event description (%v) did not match expected (%v)", event.Description(), newDescription)
	}
	ds.CheckDataStoreRecord(t, event, "OnEventUpdated")
}

func TestEventMarshalUnmarshalJSON(t *testing.T) {
	orig := NewEvent("test event", time.Now(), time.Now(), []*Contact{&Contact{}}, &Location{}, "this is a \"description\"")
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

func (ds *eventDataStoreRecorder) CheckDataStoreRecord(t *testing.T, e *Event, expectedCall string) {
	if ds.lastCallEvent != e {
		t.Errorf("Last call event reference (%v) does not match expected (%v)", ds.lastCallEvent, e)
	}
	if ds.lastCallMethod != expectedCall {
		t.Errorf("Last call method (%v) does not match expected (%v)", ds.lastCallMethod, expectedCall)
	}
	ds.lastCallEvent = nil
	ds.lastCallMethod = ""
}

func (r *eventDataStoreRecorder) OnEventCreated(event *Event) {
	r.lastCallEvent = event
	r.lastCallMethod = "OnEventCreated"
}

func (r *eventDataStoreRecorder) OnEventUpdated(e *Event) {
	r.lastCallEvent = e
	r.lastCallMethod = "OnEventUpdated"
}

func (r *eventDataStoreRecorder) OnEventDeleted(e *Event) {
	r.lastCallEvent = e
	r.lastCallMethod = "OnEventDeleted"
}

func (r *eventDataStoreRecorder) PopulateEvent(e *Event) {
	r.lastCallEvent = e
	r.lastCallMethod = "PopulateEvent"
}
