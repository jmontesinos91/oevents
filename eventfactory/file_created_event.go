package eventfactory

import (
	"time"

	"github.com/jmontesinos91/oevents"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	FileCreatedEvent = "omni.view.file_created"
)

type FileCreatedPayload struct {
	Id               string `mapstructure:"id"`
	TenantID         int    `mapstructure:"tenantId"`
	Directory        string `mapstructure:"directory"` // detection-event/{cameraProxyID}
	Data             []Data `mapstructure:"data"`
	EventDate        string `mapstructure:"eventDate"`
	DetectionEventID string `mapstructure:"detectionEventId"`
}

type Data struct {
	FileName string `mapstructure:"fileName"` // detectionEventID
	File     string `mapstructure:"file"`
}

func NewFileCreatedEvent(source string, payload FileCreatedPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	ev := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: FileCreatedEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return ev, nil
}

func ToFileCreatedPayload(data map[string]interface{}) (*FileCreatedPayload, error) {
	parsed := &FileCreatedPayload{}

	err := mapstructure.Decode(data, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
