package handler

import (
	"fmt"
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
		return c.JSON(http.StatusBadRequest, err)
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

func (h *Handler) GetTopSalesItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.TopsaleReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var topsales []model.Topsale
	rows, err := db.Raw("EXEC GetTopSalesItem @Year = ?, @Month = ?,@StoreCode = ?;", req.Year, req.Month, req.Store).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var topsale model.Topsale
		rows.Scan(&topsale.ItemName, &topsale.TotalAmount, &topsale.TotalQnt, &topsale.Profit, &topsale.Disc)
		// println(topsale)
		topsales = append(topsales, topsale)
	}

	return c.JSON(http.StatusOK, topsales)
}

func (h *Handler) GetBranchesSales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.BranchesSaleReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var branchesSales []model.BranchesSale
	rows, err := db.Raw("EXEC GetBranchesSales @Year = ?, @Month = ?", req.Year, req.Month).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var branchSale model.BranchesSale
		rows.Scan(&branchSale.StoreCode, &branchSale.StoreName, &branchSale.Totalamount)
		// println(topsale)
		branchesSales = append(branchesSales, branchSale)
	}

	return c.JSON(http.StatusOK, branchesSales)
}
func (h *Handler) GetMonthlySales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.MonthlySalesReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var monthlySales []model.MonthlySales
	rows, err := db.Raw("EXEC GetMonthlySales @Year = ?", req.Year).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure")
	}

	defer rows.Close()
	for rows.Next() {
		var monthlySale model.MonthlySales
		if err := rows.Scan(&monthlySale.DocMonth, &monthlySale.Totalamount); err != nil {
			panic(err)
		}
		fmt.Printf("is %x", monthlySale.DocMonth)
		// println(topsale)
		monthlySales = append(monthlySales, monthlySale)
	}

	return c.JSON(http.StatusOK, monthlySales)
}
