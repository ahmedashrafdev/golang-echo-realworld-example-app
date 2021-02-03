package handler

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/utils"
)

type userResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}

type serverResponse struct {
	DbUser     string `json:"username"`
	DbPassword string `json:"email"`
	DbIP       string `json:"bio"`
	DbName     string `json:"image"`
	ServerName string `json:"token"`
}

func newServerResponse(s *model.Server) *serverResponse {
	r := new(serverResponse)
	r.DbUser = s.DbUser
	r.DbPassword = s.DbPassword
	r.DbIP = s.DbIP
	r.DbName = s.DbName
	r.ServerName = s.ServerName
	return r
}
