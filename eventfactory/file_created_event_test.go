package eventfactory

import (
	"fmt"
	"testing"

	"github.com/jmontesinos91/oevents"

	"github.com/stretchr/testify/assert"
)

func TestFileCreatedEvent(t *testing.T) {
	p := FileCreatedPayload{
		Id:        "a5389c8d-fed3-4f32-86b0-8ab5be2da555",
		TenantID:  1,
		Directory: "detection-event/2158",
		Data: []Data{
			{
				FileName: "test",
				File:     "1234567890",
			},
			{
				FileName: "test1",
				File:     "1234567891",
			},
		},
		EventDate:        "2024-10-03T14:15:22Z",
		DetectionEventID: "11e39e8a-eb08-4c3d-bce7-7c3b31364692",
	}

	ev, err := NewFileCreatedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.NotEmpty(t, ev.Data, "Data should not be empty")
	assert.Equal(t, FileCreatedEvent, ev.EventType, "it should be a file accepted event")
}

func TestFileCreatedMarshallAndUnMarshallEvent(t *testing.T) {
	p := FileCreatedPayload{
		Id:        "a5389c8d-fed3-4f32-86b0-8ab5be2da555",
		TenantID:  1,
		Directory: "detection-event/2158",
		Data: []Data{
			{
				FileName: "test",
				File:     "1234567890",
			},
			{
				FileName: "test1",
				File:     "1234567891",
			},
		},
		EventDate:        "2024-10-03T14:15:22Z",
		DetectionEventID: "11e39e8a-eb08-4c3d-bce7-7c3b31364692",
	}

	ev, err := NewFileCreatedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	serialized := ev.ToJSON()

	fmt.Printf("JSON ALARM: %s \n", serialized)

	parsed, err := oevents.ParseEvent([]byte(serialized))
	assert.NoError(t, err, "No error should be thrown")

	payload, err := ToFileCreatedPayload(parsed.Data)
	assert.NoError(t, err, "No error should be thrown")

	assert.Equal(t, "a5389c8d-fed3-4f32-86b0-8ab5be2da555", payload.Id)
	assert.Equal(t, "detection-event/2158", payload.Directory)
	assert.Equal(t, "11e39e8a-eb08-4c3d-bce7-7c3b31364692", payload.DetectionEventID)
	assert.Equal(t, "test", payload.Data[0].FileName)
	assert.Equal(t, "1234567890", payload.Data[0].File)
	assert.Equal(t, "test1", payload.Data[1].FileName)
	assert.Equal(t, "1234567891", payload.Data[1].File)
}
