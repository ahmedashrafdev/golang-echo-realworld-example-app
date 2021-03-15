package db

import (
	"fmt"

	"github.com/ahmedashrafdev/reports/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	DBConn *gorm.DB
)

func InitDatabase(server model.Server) error {
	var err error
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:1433?database=%s", server.DbUser, server.DbPassword, server.DbIP, server.DbName)
	fmt.Println(connectionString)
	// DBConn, err = gorm.Open("mssql", "sqlserver://mcs:123@41.38.87.59:1433?database=stock_main")
	DBConn, err = gorm.Open("mssql", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to external database")
		return err
	}
	fmt.Println("Connection Opened to External Database")
	return nil

}
