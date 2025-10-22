package eventfactory

import (
	"time"

	"github.com/jmontesinos91/oevents"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	FileAcceptedEvent = "omni.view.file_accepted"
)

type FileAcceptedPayload struct {
	Id               string   `mapstructure:"id"`
	Path             []string `mapstructure:"path"`
	EventDate        string   `mapstructure:"eventDate"`
	DetectionEventID string   `mapstructure:"detectionEventId"`
}

func NewFileAcceptedEvent(source string, payload FileAcceptedPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	ev := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: FileAcceptedEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return ev, nil
}

func ToFileAcceptedPayload(data map[string]interface{}) (*FileAcceptedPayload, error) {
	parsed := &FileAcceptedPayload{}

	err := mapstructure.Decode(data, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
