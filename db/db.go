package db

import (
	"fmt"

	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/config"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func New() *gorm.DB {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER_TEST"), config.Config("DB_PASSWORD_TEST"), config.Config("DB_HOST_TEST"), config.Config("DB_PORT_TEST"), config.Config("DB_NAME_TEST"))
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.LogMode(false)
	return db
}

func DropTestDB() error {

	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	//ALTER TABLE users ADD CONSTRAINT server_id_foreign FOREIGN KEY (server_id) REFERENCES servers(id);
	// INSERT INTO `servers` (`id`, `created_at`, `updated_at`, `deleted_at`, `db_user`, `db_password`, `db_ip`, `db_name`, `server_name`) VALUES (NULL, NULL, NULL, NULL, 'a', 'a', 'a', 'a', 'a');
	db.AutoMigrate(
		&model.Server{},
		&model.User{},
	)
}
