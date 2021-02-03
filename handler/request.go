package handler

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/labstack/echo/v4"
)

type userUpdateRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func newUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

func (r *userUpdateRequest) populate(u *model.User) {
	r.Email = u.Email
	r.Password = u.Password
}

func (r *userUpdateRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.Email
	if r.Password != u.Password {
		h, err := u.HashUserPassword(r.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	return nil
}

type userRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	ServerId uint   `json:"server_id" validate:"required"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.Email
	u.ServerID = r.ServerId
	h, err := u.HashUserPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type ServerRequest struct {
	DbUser     string `validate:"required" json:"db_user"`
	DbPassword string `validate:"required" json:"db_password"`
	DbIP       string `validate:"required" json:"db_ip"`
	DbName     string `validate:"required" json:"db_name"`
	ServerName string `validate:"required" json:"server_name"`
}

func (r *ServerRequest) bind(c echo.Context, s *model.Server) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	s.ServerName = r.ServerName
	s.DbUser = r.DbUser
	s.DbIP = r.DbIP
	h, err := s.HashServerPassword(r.DbPassword)
	if err != nil {
		return err
	}
	s.DbPassword = h
	return nil
}
