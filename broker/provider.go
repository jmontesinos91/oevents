package broker

import (
	"context"
	"sync"

	"github.com/jmontesinos91/oevents"
)

// OBrokerConfig Represents a configuration object for connecting to a broker
type OBrokerConfig struct {
	Servers           string
	User              string
	Password          string
	ClientName        string
	ConsumerEnabled   bool
	ConsumerGroupName string
	ConsumeFromTopics []string
}

// OmniViewMessage Represents a OmniView Event and a WaitGroup used for acknowledgements
type OmniViewMessage struct {
	Event oevents.OmniViewEvent
	Ack   *sync.WaitGroup
}

// MessagingBrokerProvider Our main interface to allow access to brokers
type MessagingBrokerProvider interface {
	Publish(ctx context.Context, topic string, events ...oevents.OmniViewEvent) bool
	Subscribe(ctx context.Context, maxRecords int, workerChannel chan<- OmniViewMessage)
	Close()
}
