package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

func CreateServer(db *gorm.DB) error {
	return db.Create(&model.Server{
		DbUser:     "mcs",
		DbPassword: "123",
		DbIP:       "41.38.87.59",
		DbName:     "stock_main",
		ServerName: "Halal",
	}).Error
}
