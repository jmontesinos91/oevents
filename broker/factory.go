package broker

import (
	"github.com/jmontesinos91/ologs/logger"
)

// ConnectKafka Creates a new instance of a Kafka provider
func ConnectKafka(config OBrokerConfig, logger *logger.ContextLogger) (MessagingBrokerProvider, error) {
	return Connect(config, logger)
}

// ConnectKafkaInsecure Creates a new instance of a Kafka provider (only use it for development purposes)
func ConnectKafkaInsecure(config OBrokerConfig, logger *logger.ContextLogger) (MessagingBrokerProvider, error) {
	return ConnectInsecure(config, logger)
}
