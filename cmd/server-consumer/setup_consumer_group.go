package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"kafka-example/internal/constants/kafka"
	"kafka-example/internal/model/entity"
	"kafka-example/internal/repository/notification"

	"github.com/IBM/sarama"
)

type Consumer struct {
	notificationRepo notificationrepo.IRepository
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim,
) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)

		var notification entity.Notification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}

		ctx := context.Background()
		consumer.notificationRepo.Add(ctx, userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{kafkaconst.KafkaServerAddress}, kafkaconst.ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func setupConsumerGroup(
	ctx context.Context, notificationRepo notificationrepo.IRepository,
) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		notificationRepo: notificationRepo,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{kafkaconst.ConsumerTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}
