package handler

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/router/middleware"
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/utils"
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
	cashtry.GET("/stores", h.CashTryStores)

	v1.GET("/top", h.GetTopSalesItem, jwtMiddleware)
	v1.GET("/branches-sales", h.GetBranchesSales, jwtMiddleware)
	v1.GET("/monthly-sales", h.GetMonthlySales, jwtMiddleware)
}
