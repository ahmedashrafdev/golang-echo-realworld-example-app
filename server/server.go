package server

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
)

type Store interface {
	GetByID(uint) (*model.Server, error)
	Create(*model.Server) error
	Update(*model.Server) error
}
