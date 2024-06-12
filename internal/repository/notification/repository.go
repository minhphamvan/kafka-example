package notificationrepo

import (
	"context"

	"kafka-example/internal/model/entity"
)

type IRepository interface {
	Add(
		ctx context.Context, userID string, notification entity.Notification,
	) error
	GetByUserID(
		ctx context.Context, userID string,
	) ([]entity.Notification, error)
}
