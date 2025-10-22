package eventfactory

import (
	"time"

	"github.com/jmontesinos91/oevents"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	NoOpAcceptedEvent = "omniview.generic.no_op_accepted"
)

type NoOpPayload struct {
	Field1 string `mapstructure:"field1"`
	Field2 string `mapstructure:"field2"`
}

func NewNoOpAcceptedEvent(source string, payload NoOpPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	evt := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: NoOpAcceptedEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return evt, nil
}

func ToNoOpPayload(data map[string]interface{}) (*NoOpPayload, error) {
	parsed := &NoOpPayload{}

	err := mapstructure.Decode(data, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
