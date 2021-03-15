package handler

import (
	"github.com/ahmedashrafdev/reports/server"
	"github.com/ahmedashrafdev/reports/user"
)

type Handler struct {
	userStore   user.Store
	serverStore server.Store
}

func NewHandler(us user.Store, ss server.Store) *Handler {
	return &Handler{
		userStore:   us,
		serverStore: ss,
	}
}
