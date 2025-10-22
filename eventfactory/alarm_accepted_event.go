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
	Id          string `mapstructure:"id"`
	IMEI        string `mapstructure:"imei"`
	Description string `mapstructure:"description"`
	Latitude    string `mapstructure:"latitude"`
	Longitude   string `mapstructure:"longitude"`
	Tenant      string `mapstructure:"tenant"`
	Metadata    Event  `mapstructure:"metadata"`
}

type Event struct {
	DetectionEventsID string    `mapstructure:"detection_events_id,omitempty"`
	EventType         string    `mapstructure:"event_type"`
	Provider          string    `mapstructure:"provider"`
	Type              string    `mapstructure:"type"`
	Confidence        float64   `mapstructure:"confidence"`
	Camera            int       `mapstructure:"camera"`
	Timestamp         time.Time `mapstructure:"timestamp"`
	Data              Metadata  `mapstructure:"data"`
	ExtraData         ExtraData `mapstructure:"extra_data"`
}

type Metadata struct {
	Plate                string      `mapstructure:"plate"`
	Color                string      `mapstructure:"color"`
	Brand                string      `mapstructure:"brand"`
	Model                interface{} `mapstructure:"model"`
	Class                interface{} `mapstructure:"class"`
	Speed                int         `mapstructure:"speed"`
	Confidence           float64     `mapstructure:"confidence"`
	Label                string      `mapstructure:"label"`
	UpperBodyColor       string      `mapstructure:"upper_body_color"`
	LowerBodyColor       string      `mapstructure:"lower_body_color"`
	Gender               string      `mapstructure:"gender"`
	Age                  int         `mapstructure:"age"`
	Mask                 string      `mapstructure:"mask"`
	QualityBlurriness    float64     `mapstructure:"quality_blurriness"`
	QualityDark          float64     `mapstructure:"quality_dark"`
	QualityLight         float64     `mapstructure:"quality_light"`
	QualitySaturation    float64     `mapstructure:"quality_saturation"`
	PredominantEmotion   string      `mapstructure:"predominant_emotion"`
	PredominantEthnicity string      `mapstructure:"predominant_ethnicity"`
	Occlusion            string      `mapstructure:"occlusion"`
	Name                 string      `mapstructure:"name"`
	Hair                 string      `mapstructure:"hair"`
	Glasses              string      `mapstructure:"glasses"`
	Accessory            string      `mapstructure:"accessory"`
	Vehicle              string      `mapstructure:"vehicle"`
	Country              string      `mapstructure:"country"`
}

type ExtraData struct {
	Priority   string `mapstructure:"priority"`
	MatchRules string `mapstructure:"matchedRules"`
	Timestamp  int64  `mapstructure:"timestamp"`
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
