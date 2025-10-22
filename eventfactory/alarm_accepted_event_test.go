package eventfactory

import (
	"fmt"
	"testing"

	"github.com/jmontesinos91/oevents"
	"github.com/stretchr/testify/assert"
)

func TestAlarmEvent(t *testing.T) {
	p := AlarmPayload{
		Id:          "EV_sds23424",
		IMEI:        "12334223",
		Description: "test description",
		Longitude:   "1.1234567",
		Latitude:    "1.1234568",
		Tenant:      "1",
		Metadata: Event{
			Provider: "Lumeo",
		},
	}

	ev, err := NewAlarmAcceptedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.NotEmpty(t, ev.Data, "Data should not be empty")
	assert.Equal(t, AlarmAcceptedEvent, ev.EventType, "it should be a alarm event")
}

func TestAlarmMarshallAndUnMarshallEvent(t *testing.T) {
	p := AlarmPayload{
		Id:          "EV_sds23424",
		IMEI:        "12334223",
		Description: "test description",
		Longitude:   "1.1234567",
		Latitude:    "1.1234568",
		Tenant:      "1",
		Metadata: Event{
			Provider: "Lumeo",
		},
	}

	ev, err := NewAlarmAcceptedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	serialized := ev.ToJSON()

	fmt.Printf("JSON ALARM: %s \n", serialized)

	parsed, err := oevents.ParseEvent([]byte(serialized))
	assert.NoError(t, err, "No error should be thrown")

	payload, err := ToAlarmPayload(parsed.Data)
	assert.NoError(t, err, "No error should be thrown")

	assert.Equal(t, "EV_sds23424", payload.Id)
	assert.Equal(t, "12334223", payload.IMEI)
	assert.Equal(t, "test description", payload.Description)
	assert.Equal(t, "1.1234567", payload.Longitude)
	assert.Equal(t, "1.1234568", payload.Latitude)
	assert.Equal(t, "1", payload.Tenant)
	assert.Equal(t, "Lumeo", payload.Metadata.Provider)
}
