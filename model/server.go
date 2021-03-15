package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	gorm.Model
	DbUser     string `gorm:"not null;" json:"DbUser"`
	DbPassword string `gorm:"not null" json:"DbPassword"`
	DbIP       string `gorm:"not null;unique" json:"DbIP"`
	DbName     string `gorm:"not null;unique" json:"DbName"`
	ServerName string `gorm:"not null;unique" json:"ServerName"`
	Users      []User
}

func (s *Server) HashServerPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}
