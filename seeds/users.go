package seeds

import (
	"fmt"

	"github.com/ahmedashrafdev/reports/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *gorm.DB) error {
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("unable to encrypt", err)
	}
	db.Create(&model.User{Email: "halal@elnozom.com", ServerID: 1, Password: string(password)})
	db.Create(&model.User{Email: "dental@elnozom.com", ServerID: 2, Password: string(password)})
	return db.Create(&model.User{Email: "local@elnozom.com", ServerID: 3, Password: string(password)}).Error
}
