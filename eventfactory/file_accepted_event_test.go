package eventfactory

import (
	"fmt"
	"testing"

	"github.com/jmontesinos91/oevents"

	"github.com/stretchr/testify/assert"
)

func TestFileAcceptedEvent(t *testing.T) {
	p := FileAcceptedPayload{
		Id: "a5389c8d-fed3-4f32-86b0-8ab5be2da555",
		Path: []string{
			"http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180934.jpeg",
			"http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180935.jpeg",
		},
		EventDate:        "2024-10-03T14:15:22Z",
		DetectionEventID: "11e39e8a-eb08-4c3d-bce7-7c3b31364692",
	}

	ev, err := NewFileAcceptedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.NotEmpty(t, ev.Data, "Data should not be empty")
	assert.Equal(t, FileAcceptedEvent, ev.EventType, "it should be a file accepted event")
}

func TestFileAcceptedMarshallAndUnMarshallEvent(t *testing.T) {
	p := FileAcceptedPayload{
		Id: "a5389c8d-fed3-4f32-86b0-8ab5be2da555",
		Path: []string{
			"http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180934.jpeg",
			"http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180935.jpeg",
		},
		EventDate:        "2024-10-03T14:15:22Z",
		DetectionEventID: "11e39e8a-eb08-4c3d-bce7-7c3b31364692",
	}

	ev, err := NewFileAcceptedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	serialized := ev.ToJSON()

	fmt.Printf("JSON ALARM: %s \n", serialized)

	parsed, err := oevents.ParseEvent([]byte(serialized))
	assert.NoError(t, err, "No error should be thrown")

	payload, err := ToFileAcceptedPayload(parsed.Data)
	assert.NoError(t, err, "No error should be thrown")

	assert.Equal(t, "a5389c8d-fed3-4f32-86b0-8ab5be2da555", payload.Id)
	assert.Equal(t, "http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180934.jpeg", payload.Path[0])
	assert.Equal(t, "http://localhost:8081/v1/files/view?idTenant=111&fileName=TVSS_Beta/detection_event/10025/180935.jpeg", payload.Path[1])
	assert.Equal(t, "11e39e8a-eb08-4c3d-bce7-7c3b31364692", payload.DetectionEventID)
}
