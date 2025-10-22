package eventfactory

import (
	"time"

	"github.com/jmontesinos91/oevents"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	AlarmCreatedByDetectionEventsEvent = "omni.view.alarm_created_by_detectionevents_event"
)

type AlarmCreatedPayload struct {
	Id               string `mapstructure:"id"`
	AlarmID          string `mapstructure:"alarm_id"`
	DetectionEventID string `mapstructure:"detection_event_id"`
	EventDate        string `mapstructure:"event_date"`
}

func NewAlarmCreatedPayload(source string, payload AlarmCreatedPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	ev := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: AlarmCreatedByDetectionEventsEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return ev, nil
}

func ToAlarmCreatedPayload(data map[string]interface{}) (*AlarmCreatedPayload, error) {
	parsed := &AlarmCreatedPayload{}

	err := mapstructure.Decode(data, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
