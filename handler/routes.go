package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xesina/golang-echo-realworld-example-app/router/middleware"
	"github.com/xesina/golang-echo-realworld-example-app/utils"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

}
