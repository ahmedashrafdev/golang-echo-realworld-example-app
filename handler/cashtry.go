package handler

import (
	"net/http"

	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/db"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/labstack/echo/v4"
)

func CashTryAnalysis(c echo.Context) error {
	db := db.DBConn
	req := new(model.CashtryReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var cashtries []model.Cashtry
	rows, err := db.Raw("EXEC CashtryAnalysis @StoreCode = ?, @Year = ?;", req.Store, req.Year).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.Cashtry
		rows.Scan(&cashtry.TotalCash, &cashtry.TotalOrder, &cashtry.TVisa, &cashtry.TVoid, &cashtry.MonthNo, &cashtry.AverageCash, &cashtry.NoOfCashTry, &cashtry.AvgBasket)
		cashtries = append(cashtries, cashtry)
	}

	c.Set("Content-Type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, cashtries)
}

func CashTryStores(c echo.Context) error {
	db := db.DBConn
	var stores []model.CashtryStores
	rows, err := db.Raw("EXEC GetStoreName").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var store model.CashtryStores
		rows.Scan(&store.StoreCode, &store.StoreName)
		stores = append(stores, store)
	}

	c.Set("Content-Type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, stores)
}
