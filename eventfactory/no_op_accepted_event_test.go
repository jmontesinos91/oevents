package eventfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoOpAcceptedEvent(t *testing.T) {
	p := NoOpPayload{
		Field1: "hola",
		Field2: "adios",
	}

	ev, err := NewNoOpAcceptedEvent("test", p)
	assert.NoError(t, err, "No error should be thrown")

	err = ev.Validate()
	assert.NoError(t, err, "No error should be thrown")

	assert.NotEmpty(t, ev.ID, "Id should not be empty")
	assert.Equal(t, NoOpAcceptedEvent, ev.EventType, "It should be a No Op event")
	assert.Equal(t, "adios", ev.Data["field2"], "It should be a Adios")
}
