package user

import (
	"github.com/ahmedashrafdev/reports/model"
)

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByIDSec(uint64) (*model.User, error)
	GetServer(uint) (model.Server, error)
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	ConnectDb(uint) error
	ListUsers() ([]model.User, int, error)
	DeleteUser(u *model.User) error
}
