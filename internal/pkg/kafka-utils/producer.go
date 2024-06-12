package kafkautils

import (
	"fmt"

	"github.com/IBM/sarama"
)

func SetupProducer(kafkaServerAddress string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafkaServerAddress},
		config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	
	return producer, nil
}
