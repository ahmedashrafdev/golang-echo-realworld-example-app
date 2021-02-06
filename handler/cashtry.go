package handler

import (
	"net/http"

	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/db"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CashTryAnalysis(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
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

	return c.JSON(http.StatusOK, cashtries)
}

func (h *Handler) CashTryStores(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
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

	return c.JSON(http.StatusOK, stores)
}
