package eventfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogEvent(t *testing.T) {
	p := LogCreatedPayload{
		IpAddress:   "192.105.31.1",
		ClientHost:  "localhost",
		Provider:    "PSIM",
		Level:       0,
		Message:     1200,
		Description: "test",
		Path:        "test",
		Resource:    "test",
		Action:      "test",
		Data:        "[]",
		OldData:     "[]",
		TenantCat:   []TenantItem{},
		TenantID:    "21",
		UserID:      "21",
		Target:      "LOG TARGET",
	}

	ev, err := NewLogCreatedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.NotEmpty(t, ev.Data, "Data should not be empty")
	assert.Equal(t, LogCreatedEvent, ev.EventType, "it should be a log created event")
}

func TestToLogCreatedPayload(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		expected    *LogCreatedPayload
		expectError bool
	}{
		{
			name: "Happy path - valid data",
			input: map[string]interface{}{
				"ip_address":  "192.168.0.1",
				"client_host": "localhost",
				"provider":    "aws",
				"level":       1,
				"message":     1234,
				"description": "Test description",
				"resource":    "test_resource",
				"path":        "/api/v1",
				"action":      "create",
				"data":        "test data",
				"old_data":    "old data",
				"user_id":     "user123",
				"target":      "LOG TARGET",
				"tenant_cat": []map[string]interface{}{
					{"id": 1, "name": "Tenant A"},
					{"id": 2, "name": "Tenant B"},
				},
				"tenant_id": "tenant123",
			},
			expected: &LogCreatedPayload{
				IpAddress:   "192.168.0.1",
				ClientHost:  "localhost",
				Provider:    "aws",
				Level:       1,
				Message:     1234,
				Description: "Test description",
				Resource:    "test_resource",
				Path:        "/api/v1",
				Action:      "create",
				Data:        "test data",
				OldData:     "old data",
				UserID:      "user123",
				Target:      "LOG TARGET",
				TenantCat: []TenantItem{
					{ID: 1, Name: "Tenant A"},
					{ID: 2, Name: "Tenant B"},
				},
				TenantID: "tenant123",
			},
			expectError: false,
		},
		{
			name: "Missing fields -  do not expect an error",
			input: map[string]interface{}{
				"ip_address":  "192.168.0.1",
				"client_host": "localhost",
			},
			expected: &LogCreatedPayload{
				IpAddress:  "192.168.0.1",
				ClientHost: "localhost",
			},
			expectError: false,
		},
		{
			name: "Invalid field types - expect error",
			input: map[string]interface{}{
				"ip_address": 1234,
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToLogCreatedPayload(tt.input)

			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
				assert.Nil(t, result, "Expected result to be nil on error")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.NotNil(t, result, "Expected result to be non-nil")
				assert.Equal(t, tt.expected, result, "Result does not match expected")
			}
		})
	}
}

func TestSerializeTenantCatWithMapstructure(t *testing.T) {
	tests := []struct {
		name        string
		input       []TenantItem
		expected    string
		expectError bool
	}{
		{
			name: "Happy path - valid TenantCat",
			input: []TenantItem{
				{ID: 111, Name: "VSS Beta"},
				{ID: 113, Name: "Kapa8"},
			},
			expected:    `[{"id":111,"name":"VSS Beta"},{"id":113,"name":"Kapa8"}]`,
			expectError: false,
		},
		{
			name:        "Empty slice",
			input:       []TenantItem{},
			expected:    `[]`,
			expectError: false,
		},
		{
			name: "Invalid data - should not fail with mapstructure",
			input: []TenantItem{
				{ID: 0, Name: ""},
			},
			expected:    `[{"id":0,"name":""}]`,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToTenantCatJson(tt.input)

			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.JSONEq(t, tt.expected, result, "Serialized JSON does not match expected")
			}
		})
	}
}
