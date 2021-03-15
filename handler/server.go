package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ahmedashrafdev/reports/model"
	"github.com/ahmedashrafdev/reports/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateServer(c echo.Context) error {
	var s model.Server
	req := &ServerRequest{}
	fmt.Println(req)
	fmt.Println(s)
	if err := req.bind(c, &s); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.serverStore.Create(&s); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, "server created successfully")
}

func (h *Handler) ListServers(c echo.Context) error {
	var (
		servers []model.Server
		err     error
		count   int
	)

	servers, count, err = h.serverStore.ListServers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, newServerListResponse(h.serverStore, servers, count))
}

func (h *Handler) DeleteServer(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error parsing id from body")

	}
	u, err := h.serverStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	err = h.serverStore.DeleteServer(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (h *Handler) UpdateServer(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error parsing id from body")

	}
	s, err := h.serverStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if s == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := newServerUpdateRequest()
	req.populate(s)
	if err := req.bind(c, s); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.serverStore.Update(s); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newServerResponse(s))
}
