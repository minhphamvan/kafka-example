package userrepo

import (
	"kafka-example/internal/model/entity"
)

type IRepository interface {
	GetUserByID(id int) (entity.User, error)
}
