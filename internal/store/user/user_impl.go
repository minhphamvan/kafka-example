package userstore

import (
	"context"

	"kafka-example/internal/cerror"
	"kafka-example/internal/model/entity"
)

type UserImpl struct {
	data []entity.User
}

func New() *UserImpl {
	users := []entity.User{
		{ID: 1, Name: "First User"},
		{ID: 2, Name: "Second User"},
		{ID: 3, Name: "Third User"},
		{ID: 4, Name: "Fourth User"},
		{ID: 5, Name: "Fifth User"},
	}

	return &UserImpl{
		data: users,
	}
}

func (u *UserImpl) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	for _, user := range u.data {
		if user.ID == id {
			return user, nil
		}
	}

	return entity.User{}, cerror.ErrUserNotFound
}
