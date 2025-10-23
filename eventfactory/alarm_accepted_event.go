package eventfactory

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmontesinos91/oevents"
	"github.com/mitchellh/mapstructure"
)

const (
	AlarmAcceptedEvent = "omni.view.alarm_accepted"
)

type AlarmPayload struct {
	Id               string `mapstructure:"id"`
	IMEI             string `mapstructure:"imei"`
	Description      string `mapstructure:"description"`
	Latitude         string `mapstructure:"latitude"`
	Longitude        string `mapstructure:"longitude"`
	AlarmType        string `mapstructure:"alarmType"`
	Waiting          string `mapstructure:"waiting"`
	Attending        string `mapstructure:"attending"`
	IsNotification   bool   `mapstructure:"notification"`
	EventDate        string `mapstructure:"event_date"`
	DetectionEventID string `mapstructure:"detection_event_id"`
}

func NewAlarmAcceptedEvent(source string, payload AlarmPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	ev := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: AlarmAcceptedEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return ev, nil
}

func ToAlarmPayload(data map[string]interface{}) (*AlarmPayload, error) {
	parsed := &AlarmPayload{}

	err := mapstructure.Decode(data, &parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
