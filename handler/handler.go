package handler

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/server"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/user"
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
