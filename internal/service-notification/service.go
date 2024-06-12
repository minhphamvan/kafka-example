package servicenotification

import (
	"context"

	"kafka-example/internal/model/entity"
	"kafka-example/internal/service-notification/request"
)

type IService interface {
	SendNotification(ctx context.Context, req request.SendNotificationReq) error
	GetByUserID(ctx context.Context, req request.GetByUserIDReq) ([]entity.Notification, error)
}
