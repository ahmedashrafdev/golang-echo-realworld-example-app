package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

func CreateServer(db *gorm.DB) error {
	db.Create(&model.Server{
		DbUser:     "mcs",
		DbPassword: "123",
		DbIP:       "41.38.87.59",
		DbName:     "stock_main",
		ServerName: "halal",
	})
	db.Create(&model.Server{
		DbUser:     "web",
		DbPassword: "123",
		DbIP:       "dental.myfirewall.co",
		DbName:     "STOCK2021",
		ServerName: "dental",
	})
	return db.Create(&model.Server{
		DbUser:     "mcs",
		DbPassword: "123",
		DbIP:       "192.168.1.167",
		DbName:     "khedawy2020",
		ServerName: "local",
	}).Error
}
