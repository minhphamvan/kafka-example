package servicenotification

import (
	notificationrepo "kafka-example/internal/repository/notification"
	"kafka-example/internal/repository/user"

	"github.com/IBM/sarama"
)

type service struct {
	userRepo         userrepo.IRepository
	notificationRepo notificationrepo.IRepository

	producer sarama.SyncProducer
}

const traceTag = "serviceNotification"

func NewService(
	userRepo userrepo.IRepository,
	notificationRepo notificationrepo.IRepository,
	producer sarama.SyncProducer,
) IService {
	return &service{
		userRepo:         userRepo,
		notificationRepo: notificationRepo,
		producer:         producer,
	}
}
