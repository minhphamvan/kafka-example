package servicenotification

import (
	"context"
	"fmt"

	"kafka-example/internal/model/entity"
	"kafka-example/internal/service-notification/request"
)

func (s *service) GetByUserID(
	ctx context.Context, req request.GetByUserIDReq,
) ([]entity.Notification, error) {
	// Get notifications by user ID
	notifications, err := s.notificationRepo.GetByUserID(ctx, fmt.Sprint(req.UserID))
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
