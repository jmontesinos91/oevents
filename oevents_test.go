package oevents_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jmontesinos91/oevents"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	te := &oevents.OmniViewEvent{}

	err := te.Validate()

	assert.Error(t, err, "An error should have been thrown")
}

func TestJsonMarshalling(t *testing.T) {
	data := map[string]interface{}{
		"customerId":   "3c320faf-4057-4994-a943-56a7e607da8b",
		"cardNumber":   "2245",
		"cardCurrency": "mxn",
		"cardAmount":   254350,
	}

	te := &oevents.OmniViewEvent{
		ID:        "xyz",
		Source:    "test",
		EventType: "my_event",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	err := te.Validate()

	assert.NoError(t, err, "No error should have been thrown")

	json := te.ToJSON()

	fmt.Printf("Json: %s", json)

	assert.NotEmpty(t, json, "Json should not be empty")
	assert.NotEqual(t, "{}", json, "Json did not serialize properly")
}

func TestEventUnmarshalling(t *testing.T) {
	event := `{
		"id": "1f421b96-c00e-4c14-bd5d-e24668218500",
		"type": "omniview.cc.auth_finalized",
		"source": "talos",
		"timestamp": "2021-11-19T18:30:00Z",
		"data": {
			"customerId": "3c320faf-4057-4994-a943-56a7e607da8b",
			"cardNumber": "2245",
			"cardCurrency": "mxn",
			"cardAmount": 254350
		}
	}`

	parsedEvent, err := oevents.ParseEvent([]byte(event))

	assert.NoError(t, err, "No error should have been thrown")
	assert.NotEmpty(t, parsedEvent, "no empty")
	assert.Equal(t, "talos", parsedEvent.Source)
	assert.Equal(t, "2245", parsedEvent.Data["cardNumber"])
}
