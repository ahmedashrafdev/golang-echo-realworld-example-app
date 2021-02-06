package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB) error {
	return db.Create(&model.User{Email: "test@test.com", ServerID: 1, Password: "123456"}).Error
}
