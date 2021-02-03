package handler

// import (
// 	"log"
// 	"os"
// 	"testing"

// 	"encoding/json"

// 	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/db"
// 	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
// 	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/router"
// 	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/store"
// 	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/user"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// 	"github.com/labstack/echo/v4"
// )

// var (
// 	d  *gorm.DB
// 	us user.Store
// 	ss server.Store
// 	h  *Handler
// 	e  *echo.Echo
// )

// func TestMain(m *testing.M) {
// 	setup()
// 	code := m.Run()
// 	tearDown()
// 	os.Exit(code)
// }

// func authHeader(token string) string {
// 	return "Token " + token
// }

// func setup() {
// 	d = db.TestDB()
// 	db.AutoMigrate(d)
// 	us = store.NewUserStore(d)
// 	ss = store.NewServerStore(d)
// 	h = NewHandler(us , ss)
// 	e = router.New()
// 	loadFixtures()
// }

// func tearDown() {
// 	_ = d.Close()
// 	if err := db.DropTestDB(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func responseMap(b []byte, key string) map[string]interface{} {
// 	var m map[string]interface{}
// 	json.Unmarshal(b, &m)
// 	return m[key].(map[string]interface{})
// }

// func loadFixtures() error {
// 	u1 := model.User{
// 		Email: "user1@realworld.io",
// 	}
// 	u1.Password, _ = u1.HashUserPassword("secret")
// 	if err := us.Create(&u1); err != nil {
// 		return err
// 	}

// 	u2 := model.User{
// 		Email: "user2@realworld.io",
// 	}
// 	u2.Password, _ = u2.HashUserPassword("secret")
// 	if err := us.Create(&u2); err != nil {
// 		return err
// 	}

// 	return nil
// }
