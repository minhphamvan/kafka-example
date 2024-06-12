package servicenotification

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"kafka-example/internal/constants/kafka"
	"kafka-example/internal/model/entity"
	"kafka-example/internal/service-notification/request"

	"github.com/IBM/sarama"
)

func (s *service) SendNotification(
	ctx context.Context, req request.SendNotificationReq,
) error {
	// Get user by ID
	fromUser, err := s.userRepo.GetUserByID(req.FromUserID)
	if err != nil {
		return err
	}

	toUser, err := s.userRepo.GetUserByID(req.ToUserID)
	if err != nil {
		return err
	}

	// Publish notification to Kafka
	err = s.publishNotification(ctx, fromUser, toUser, req.Message)
	if err != nil {
		return fmt.Errorf("failed to publish notification: %w", err)
	}

	return nil
}

func (s *service) publishNotification(
	ctx context.Context, fromUser, toUser entity.User, message string,
) error {
	// Create notification
	notification := entity.Notification{
		From:    fromUser,
		To:      toUser,
		Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// Publish to Kafka
	msg := &sarama.ProducerMessage{
		Topic: kafkaconst.KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}
	_, _, err = s.producer.SendMessage(msg)

	return nil
}
