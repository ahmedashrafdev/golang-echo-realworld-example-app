package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmedashrafdev/reports/db"
	"github.com/ahmedashrafdev/reports/model"
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
func (h *Handler) GetDocNo(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocNo []model.Doc
	rows, err := db.Raw("EXEC GetSdDocNo @DevNo = ?, @TrSerial = ?,@StoreCode = ?;", req.DevNo, req.TrSerial, req.StoreCode).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var doc model.Doc
		err = rows.Scan(
			&doc.DocNo,
		)
		print(rows)
		if err != nil {
			return c.JSON(http.StatusOK, 1)
		}
		DocNo = append(DocNo, doc)
	}

	return c.JSON(http.StatusOK, DocNo[0].DocNo+1)
}

func (h *Handler) GetOpenDocs(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.OpenDocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var OpenDocs []model.OpenDoc
	rows, err := db.Raw("EXEC GetOpenSdDocNo @DevNo = ?, @TrSerial = ?;", req.DevNo, req.TrSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var openDoc model.OpenDoc
		err = rows.Scan(
			&openDoc.DocNo,
			&openDoc.StoreCode,
			&openDoc.AccontSerial,
			&openDoc.TransSerial,
			&openDoc.AccountName,
			&openDoc.AccountCode,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		OpenDocs = append(OpenDocs, openDoc)
	}

	return c.JSON(http.StatusOK, OpenDocs)
}
func (h *Handler) GetCashFlow(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.CashFlowReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var CashFlows []model.CashFlow
	rows, err := db.Raw("EXEC cashFlow @DateFrom = ?, @DateTo = ?,@AccSerial = ?;", req.DateFrom, req.DateTo, req.AccSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var CashFlow model.CashFlow
		err = rows.Scan(
			&CashFlow.DocDate,
			&CashFlow.Income,
			&CashFlow.Supplier,
			&CashFlow.Expensis,
			&CashFlow.Others,
			&CashFlow.Bankin,
			&CashFlow.Cheqout,
			&CashFlow.Cheqin,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		CashFlows = append(CashFlows, CashFlow)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, CashFlows)

}

func (h *Handler) GetSupplierBalance(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	var Suppliers []model.Supplier
	rows, err := db.Raw("EXEC GetSupplierBalance").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var Supplier model.Supplier
		err = rows.Scan(
			&Supplier.AccountCode,
			&Supplier.AccountName,
			&Supplier.DBT,
			&Supplier.CRDT,
			&Supplier.ReturnBuy,
			&Supplier.Buy,
			&Supplier.Paid,
			&Supplier.CHEQUE,
			&Supplier.CHQUnderCollec,
			&Supplier.Discount,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		Suppliers = append(Suppliers, Supplier)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, Suppliers)

}
func (h *Handler) GetCashFlowYear(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.CashFlowYearReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var CashFlows []model.CashFlow
	rows, err := db.Raw("EXEC cashFlowYear @Year = ? ,@AccSerial = ?;", req.Year, req.AccSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var CashFlow model.CashFlow
		err = rows.Scan(
			&CashFlow.DocDate,
			&CashFlow.Income,
			&CashFlow.Supplier,
			&CashFlow.Expensis,
			&CashFlow.Others,
			&CashFlow.Bankin,
			&CashFlow.Cheqout,
			&CashFlow.Cheqin,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		CashFlows = append(CashFlows, CashFlow)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, CashFlows)

}

func (h *Handler) GetDocItems(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DocItemsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocItems []model.DocItem
	rows, err := db.Raw("EXEC GetSdItems @DevNo = ?, @TrSerial = ?,@StoreCode = ? , @DocNo = ?;", req.DevNo, req.TrSerial, req.StoreCode, req.DocNo).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var docItem model.DocItem
		err = rows.Scan(
			&docItem.Serial,
			&docItem.Qnt,
			&docItem.Item_BarCode,
			&docItem.ItemName,
			&docItem.MinorPerMajor,
			&docItem.ByWeight,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		DocItems = append(DocItems, docItem)
	}

	return c.JSON(http.StatusOK, DocItems)
}

func (h *Handler) DeleteItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DeleteItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := db.Raw("EXEC DeleteSdItem  @Serial = ?; ", req.Serial).Rows()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) InsertItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	print("asddd")

	db := db.DBConn
	req := new(model.InsertItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := db.Raw(
		"EXEC InsertSdDocNo  @DNo = ? ,@TrS = ? ,@AccS = ? ,@ItmS =?  ,@Qnt = ? ,@StCode = ? ,@InvNo = ? ,@ItmBarCode = ? ,@DevNo = ?,@StCode2 = ?; ", req.DNo, req.TrS, req.AccS, req.ItmS, req.Qnt, req.StCode, req.InvNo, req.ItmBarCode, req.DevNo, req.StCode2).Rows()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) OpenCashTry(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.OpenCashtryReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var openCashtries []model.OpenCashtry
	rows, err := db.Raw("EXEC GetOpenCashTry").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.OpenCashtry
		err = rows.Scan(
			&cashtry.EmpCode,
			&cashtry.OpenDate,
			&cashtry.StartCash,
			&cashtry.TotalCash,
			&cashtry.CompouterName,
			&cashtry.TotalOrder,
			&cashtry.TotalVisa,
			&cashtry.StoreName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")

		}
		openCashtries = append(openCashtries, cashtry)
	}

	return c.JSON(http.StatusOK, openCashtries)
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

func (h *Handler) GetBranchesProfit(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.BranchesProfitReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var branchesProfit []model.BranchesProfit
	rows, err := db.Raw("EXEC GetBranchesProfit @Year = ?, @Month = ?", req.Year, req.Month).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var branchSale model.BranchesProfit
		rows.Scan(&branchSale.StoreCode, &branchSale.StoreName, &branchSale.Profit)
		branchesProfit = append(branchesProfit, branchSale)
	}

	return c.JSON(http.StatusOK, branchesProfit)
}

func (h *Handler) GetAccount(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var accounts []model.Account
	rows, err := db.Raw("EXEC GetAccount @Code = ?, @Name = ? , @Type = ?", req.Code, req.Name, req.Type).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		rows.Scan(&account.Serial, &account.AccountCode, &account.AccountName)
		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.GetItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var items []model.Item
	rows, err := db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ?", req.BCode, req.StoreCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.Serial, &item.ItemName, &item.MinorPerMajor, &item.POSPP, &item.POSTP, &item.ByWeight)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
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

func (h *Handler) GetDailySales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DailySalesReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var dailySales []model.DailtlySales
	rows, err := db.Raw("EXEC GetDailySales @Month = ? , @Year = ?", req.Month, req.Year).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure")
	}

	defer rows.Close()
	for rows.Next() {
		var monthlySale model.DailtlySales
		if err := rows.Scan(&monthlySale.DocDay, &monthlySale.Totalamount); err != nil {
			panic(err)
		}
		fmt.Printf("is %x", monthlySale.DocDay)
		// println(topsale)
		dailySales = append(dailySales, monthlySale)
	}

	return c.JSON(http.StatusOK, dailySales)
}
