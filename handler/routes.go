package handler

import (
	"github.com/ahmedashrafdev/reports/router/middleware"
	"github.com/ahmedashrafdev/reports/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("/list", h.ListUsers)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUser)

	server := v1.Group("/server", jwtMiddleware)
	server.POST("", h.CreateServer)
	server.GET("/list", h.ListServers)
	server.DELETE("/:id", h.DeleteServer)
	server.PUT("/:id", h.UpdateServer)

	cashtry := v1.Group("/cashtry", jwtMiddleware)
	cashtry.GET("", h.CashTryAnalysis)
	cashtry.GET("/open", h.OpenCashTry)
	cashtry.GET("/stores", h.CashTryStores)

	v1.GET("/top", h.GetTopSalesItem, jwtMiddleware)
	v1.GET("/cash-flow", h.GetCashFlow, jwtMiddleware)
	v1.GET("/cash-flow-year", h.GetCashFlowYear, jwtMiddleware)
	v1.GET("/supplier-balance", h.GetSupplierBalance, jwtMiddleware)
	v1.GET("/branches-sales", h.GetBranchesSales, jwtMiddleware)
	v1.GET("/branches-profit", h.GetBranchesProfit, jwtMiddleware)
	v1.GET("/monthly-sales", h.GetMonthlySales, jwtMiddleware)
	v1.GET("/daily-sales", h.GetDailySales, jwtMiddleware)
	v1.POST("/get-account", h.GetAccount, jwtMiddleware)
	v1.POST("/get-item", h.GetItem, jwtMiddleware)
	v1.POST("/get-doc", h.GetDocNo, jwtMiddleware)
	v1.POST("/get-doc-items", h.GetDocItems, jwtMiddleware)
	v1.POST("/insert-item", h.InsertItem, jwtMiddleware)
	v1.POST("/delete-item", h.DeleteItem, jwtMiddleware)
	v1.POST("/get-docs", h.GetOpenDocs, jwtMiddleware)

}
