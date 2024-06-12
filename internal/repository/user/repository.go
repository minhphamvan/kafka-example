package userrepo

import (
	"context"

	"kafka-example/internal/model/entity"
)

type IRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
}
