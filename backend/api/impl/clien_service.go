package impl

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClientServiceImpl struct {
}

func (c ClientServiceImpl) ListAll(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (c ClientServiceImpl) ClientDetail(ctx echo.Context, clientId string) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (c ClientServiceImpl) CreateClient(ctx echo.Context, clientId string) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
