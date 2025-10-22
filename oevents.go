package oevents

import (
	"encoding/json"
	"errors"
)

// OmniViewEvent Structure to represent a OmniView Event
type OmniViewEvent struct {
	ID        string                 `json:"id"`
	Source    string                 `json:"source"`
	EventType string                 `json:"type"`
	Timestamp string                 `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

// ParseEvent Parses a JSON string into a OmniView Event
func ParseEvent(ev []byte) (*OmniViewEvent, error) {
	parsed := &OmniViewEvent{}

	err := json.Unmarshal(ev, parsed)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}

// Validate Validates that all data required in a OmniView Event is present
func (e OmniViewEvent) Validate() error {
	if len(e.ID) == 0 {
		return errors.New("id is required")
	}

	if len(e.Source) == 0 {
		return errors.New("source is required")
	}

	if len(e.EventType) == 0 {
		return errors.New("type is required")
	}

	if len(e.Timestamp) == 0 {
		return errors.New("timestamp is required")
	}

	return nil
}

// ToJSON Converts a OmniView Event struct into a serialized JSON string
func (e OmniViewEvent) ToJSON() string {
	o, err := json.Marshal(e)
	if err != nil {
		return "{}"
	}

	return string(o)
}
