package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmontesinos91/oevents"
	"github.com/jmontesinos91/oevents/broker"
	"github.com/jmontesinos91/oevents/eventfactory"
	"github.com/jmontesinos91/ologs/logger"
	"github.com/sirupsen/logrus"
)

func main() {

	// Clean context
	ctx := context.Background()

	numArgs := len(os.Args)

	if numArgs < 2 {
		fmt.Println(">> You need to use 'producer' or 'consumer' as the first arg")
		return
	}

	// Initialize the logger
	log := logger.NewContextLogger("EVENTS", "debug", logger.TextFormat)

	mode := os.Args[1]

	// Kafka config
	var streamConfig broker.OBrokerConfig

	if mode == "producer" {
		streamConfig = broker.OBrokerConfig{
			Servers:         "10.26.23.5:9092",
			User:            "the_user",
			Password:        "the_password",
			ClientName:      "demo",
			ConsumerEnabled: false,
		}
	} else {
		streamConfig = broker.OBrokerConfig{
			Servers:           "10.26.23.5:9092",
			User:              "the_user",
			Password:          "the_password",
			ClientName:        "demo",
			ConsumerEnabled:   true,
			ConsumerGroupName: "demo-group",
			ConsumeFromTopics: []string{"omniview.alarms.all"},
		}
	}

	// Connect to Kafka (insecure only for development)
	stream, err := broker.ConnectKafkaInsecure(streamConfig, log)

	if err != nil {
		log.Log(
			logrus.ErrorLevel,
			"Main",
			"Not able to connect to the broker")
		return
	}

	// Make sure to close the connection at the end
	defer stream.Close()

	if mode == "producer" {
		log.Log(logrus.InfoLevel, "main", "Starting up in producer mode...")
		producer(ctx, stream, log)
	} else {
		log.Log(logrus.InfoLevel, "main", "Starting up in consumer mode...")
		consumer(ctx, stream, log)
	}
}

func producer(ctx context.Context, stream broker.MessagingBrokerProvider, log *logger.ContextLogger) {
	start := time.Now()
	var seen int
	// Publish a bunch of events
	for i := 0; i < 10; i++ {

		// Creates an event of type NoOp
		payload := &eventfactory.NoOpPayload{
			Field1: "field1",
			Field2: "field2",
		}

		evt, err := eventfactory.NewNoOpAcceptedEvent("my_service", *payload)
		if err != nil {
			log.WithContext(
				logrus.ErrorLevel,
				"producer",
				"Ouch, not able to construct event.",
				logger.Context{}, err)
			return
		}

		seen = seen + 1

		// Publishes the batch
		ok := stream.Publish(ctx, oevents.OmniViewTopic, *evt)

		if !ok {
			log.WithContext(
				logrus.ErrorLevel,
				"producer",
				"Ouch, message batch publication failed",
				logger.Context{}, err)
			return
		}
	}

	message := fmt.Sprintf("Produced %v messages with speed %.2f/s", seen, float64(seen)/time.Since(start).Seconds())
	log.Log(logrus.InfoLevel, "producer", message)
}

func consumer(ctx context.Context, stream broker.MessagingBrokerProvider, log *logger.ContextLogger) {
	// Creates a channel for listening for ctrl-c keystroke
	signalChannel := make(chan os.Signal, 1)

	// The workers channel, must be a bounded channel to avoid running out of memory
	workerChannel := make(chan broker.OmniViewMessage, 100)

	// Spawn some workers
	for w := 1; w <= 10; w++ {
		go myHandler(workerChannel, log)
	}

	// Subscribe to the topic
	stream.Subscribe(ctx, 100, workerChannel)

	// Block until ctrl-c is issued
	<-signalChannel
}

func myHandler(workerChannel <-chan broker.OmniViewMessage, log *logger.ContextLogger) {
	for msg := range workerChannel {
		log.WithContext(
			logrus.InfoLevel,
			"myHandler",
			"Processing event",
			logger.Context{
				"ID": msg.Event.ID,
			},
			nil)

		// Simulate some processing
		time.Sleep(300 * time.Millisecond)

		// Ack the message
		msg.Ack.Done()
	}
}
