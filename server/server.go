package article

import (
	"github.com/xesina/golang-echo-realworld-example-app/model"
)

type Store interface {
	CreateServer(*model.Server) error
	UpdateServer(*model.Server, []string) error
	DeleteServer(*model.Server) error
	List(offset, limit int) ([]model.Server, int, error)
}
