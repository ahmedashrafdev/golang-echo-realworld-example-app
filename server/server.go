package server

import (
	"github.com/ahmedashrafdev/reports/model"
)

type Store interface {
	GetByID(uint64) (*model.Server, error)
	Create(*model.Server) error
	ListServers() ([]model.Server, int, error)
	DeleteServer(s *model.Server) error
	Update(*model.Server) error
}
