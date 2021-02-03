package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	gorm.Model
	DbUser     string `gorm:"size:50;not null;unique" json:"db_user"`
	DbPassword string `gorm:"size:50;not null;unique" json:"db_password"`
	DbIP       string `gorm:"size:50;not null;unique" json:"db_ip"`
	DbName     string `gorm:"size:50;not null;unique" json:"db_name"`
	ServerName string `gorm:"size:50;not null;unique" json:"server_name"`
	Users      []User
}

func (s *Server) HashServerPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}
