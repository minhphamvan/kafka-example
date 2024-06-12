package notificationstore

import (
	"context"
	"sync"

	"kafka-example/internal/model/entity"
	"kafka-example/internal/repository/notification"
)

type NotificationImpl struct {
	data map[string][]entity.Notification // map[userID]entity.Notification
	mu   sync.RWMutex
}

func New() notificationrepo.IRepository {
	return &NotificationImpl{
		data: make(map[string][]entity.Notification),
		mu:   sync.RWMutex{},
	}
}

func (r *NotificationImpl) Add(
	ctx context.Context, userID string, notification entity.Notification,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[userID] = append(r.data[userID], notification)
	return nil
}

func (r *NotificationImpl) GetByUserID(
	ctx context.Context, userID string,
) ([]entity.Notification, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.data[userID], nil
}
