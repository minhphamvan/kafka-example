package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"kafka-example/env/kafka"
	"kafka-example/internal/model/entity"
	"kafka-example/internal/repository/notification"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct {
	notificationRepo notificationrepo.IRepository
}

func (*ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumerGroupHandler *ConsumerGroupHandler) ConsumeClaim(
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
		consumerGroupHandler.notificationRepo.Add(ctx, userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{kafkaconst.KafkaServerAddress}, kafkaconst.ConsumerGroupNotifications, config)
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

	consumerGroupHandler := &ConsumerGroupHandler{
		notificationRepo: notificationRepo,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{kafkaconst.TopicNotifications}, consumerGroupHandler)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}
