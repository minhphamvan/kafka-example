package userstore

import (
	"kafka-example/internal/cerror"
	"kafka-example/internal/model/entity"
)

var Users = []entity.User{
	{ID: 1, Name: "First User"},
	{ID: 2, Name: "Second User"},
	{ID: 3, Name: "Third User"},
	{ID: 4, Name: "Fourth User"},
}

type UserImpl struct{}

func New() *UserImpl {
	return &UserImpl{}
}

func (u *UserImpl) GetUserByID(id int) (entity.User, error) {
	for _, user := range Users {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, cerror.ErrUserNotFound
}
