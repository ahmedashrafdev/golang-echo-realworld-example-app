package handler

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/user"
)

type Handler struct {
	userStore user.Store
}

func NewHandler(us user.Store) *Handler {
	return &Handler{
		userStore: us,
	}
}
