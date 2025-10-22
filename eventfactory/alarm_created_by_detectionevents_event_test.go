package eventfactory

import (
	"fmt"
	"testing"

	"github.com/jmontesinos91/oevents"

	"github.com/stretchr/testify/assert"
)

func TestAlarmCreatedByDetectionEventsEvent(t *testing.T) {
	p := AlarmCreatedPayload{
		Id:               "EV_sds23424",
		AlarmID:          "a_id_sds23424",
		DetectionEventID: "EV_id_sds23424",
		EventDate:        "2024-08-24T14:15:22Z",
	}

	ev, err := NewAlarmCreatedPayload("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.NotEmpty(t, ev.Data, "Data should not be empty")
	assert.Equal(t, AlarmCreatedByDetectionEventsEvent, ev.EventType, "it should be a alarm created by detectionevent event")
}

func TestAlarmCreatedByDetectionEventsMarshallAndUnMarshallEvent(t *testing.T) {
	p := &AlarmCreatedPayload{
		Id:               "EV_sds23424",
		AlarmID:          "A_id_sds23424",
		DetectionEventID: "EV_id_sds23424",
		EventDate:        "2024-08-24T14:15:22Z",
	}

	ev, err := NewAlarmCreatedPayload("test", *p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	serialized := ev.ToJSON()

	fmt.Printf("JSON ALARM: %s \n", serialized)

	parsed, err := oevents.ParseEvent([]byte(serialized))
	assert.NoError(t, err, "No error should be thrown")

	payload, err := ToAlarmCreatedPayload(parsed.Data)
	assert.NoError(t, err, "No error should be thrown")

	assert.Equal(t, "EV_sds23424", payload.Id)
	assert.Equal(t, "A_id_sds23424", payload.AlarmID)
	assert.Equal(t, "EV_id_sds23424", payload.DetectionEventID)
	assert.Equal(t, "2024-08-24T14:15:22Z", payload.EventDate)
}
