package eventfactory

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmontesinos91/oevents"
	"github.com/mitchellh/mapstructure"
)

const (
	LogCreatedEvent = "omni.view.log_created"
)

// TenantItem represents a single tenant's information
type TenantItem struct {
	ID   int    `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

// LogCreatedPayload holds the data for a created log event
type LogCreatedPayload struct {
	IpAddress   string       `mapstructure:"ip_address"`
	ClientHost  string       `mapstructure:"client_host"`
	Provider    string       `mapstructure:"provider"`
	Level       int          `mapstructure:"level"`
	Message     int          `mapstructure:"message"`
	Description string       `mapstructure:"description"`
	Resource    string       `mapstructure:"resource"`
	Path        string       `mapstructure:"path"`
	Action      string       `mapstructure:"action"`
	Data        string       `mapstructure:"data"`
	OldData     string       `mapstructure:"old_data"`
	UserID      string       `mapstructure:"user_id"`
	Target      string       `mapstructure:"target"`
	TenantCat   []TenantItem `mapstructure:"tenant_cat"`
	TenantID    string       `mapstructure:"tenant_id"`
}

// NewLogCreatedEvent creates a new log event from the provided source and payload.
func NewLogCreatedEvent(source string, payload LogCreatedPayload) (*oevents.OmniViewEvent, error) {
	var data = make(map[string]interface{})
	err := mapstructure.Decode(payload, &data)

	if err != nil {
		return nil, err
	}

	ev := &oevents.OmniViewEvent{
		ID:        uuid.NewString(),
		Source:    source,
		EventType: LogCreatedEvent,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	return ev, nil
}

// ToLogCreatedPayload converts a map into a LogCreatedPayload.
func ToLogCreatedPayload(data map[string]interface{}) (*LogCreatedPayload, error) {
	parsed := &LogCreatedPayload{}

	err := mapstructure.Decode(data, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

// ToTenantCatJson converts TenantItem slice to JSON string
func ToTenantCatJson(items []TenantItem) (string, error) {

	var rawItems []map[string]interface{}
	err := mapstructure.Decode(items, &rawItems)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(rawItems)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
