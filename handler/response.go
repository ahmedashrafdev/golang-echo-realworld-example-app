package handler

import (
	"github.com/ahmedashrafdev/reports/model"
	"github.com/ahmedashrafdev/reports/server"
	"github.com/ahmedashrafdev/reports/user"
	"github.com/ahmedashrafdev/reports/utils"
)

type userResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
type userListResponse struct {
	Users      []*userResponse `json:"users"`
	UsersCount int             `json:"usersCount"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}

func newUserListResponse(us user.Store, users []model.User, count int) *userListResponse {
	r := new(userListResponse)
	r.Users = make([]*userResponse, 0)
	for _, a := range users {
		ar := new(userResponse)
		ar.Email = a.Email
		r.Users = append(r.Users, ar)
	}
	r.UsersCount = count
	return r
}

type serverResponse struct {
	DbUser     string `json:"DbUser"`
	DbPassword string `json:"DbPassword"`
	DbIP       string `json:"DbIP"`
	DbName     string `json:"DbName"`
	ServerName string `json:"ServerName"`
}
type serverListResponse struct {
	Servers      []*serverResponse `json:"users"`
	ServersCount int               `json:"usersCount"`
}

func newServerListResponse(ss server.Store, servers []model.Server, count int) *serverListResponse {
	r := new(serverListResponse)
	r.Servers = make([]*serverResponse, 0)
	for _, a := range servers {
		sr := new(serverResponse)
		sr.DbUser = a.DbUser
		sr.DbPassword = a.DbPassword
		sr.DbIP = a.DbIP
		sr.DbName = a.DbName
		sr.ServerName = a.ServerName
		r.Servers = append(r.Servers, sr)
	}
	r.ServersCount = count
	return r
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
