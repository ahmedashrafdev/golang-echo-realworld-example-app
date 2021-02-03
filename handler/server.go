package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateServer(c echo.Context) error {
	var s model.Server
	req := &ServerRequest{}
	fmt.Println(req)
	fmt.Println(s)
	fmt.Println(c.Request().GetBody())
	return c.JSON(http.StatusOK, newServerResponse(&s))

	// if err := req.bind(c, &s); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	// }
	// if err := h.serverStore.Create(&s); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	// }
	// return c.JSON(http.StatusCreated, newServerResponse(&s))
}
