package user

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
)

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
