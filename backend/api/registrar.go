package api

import (
	"github.com/identityOrg/cerberus-api/backend/api/client"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAllAPI(e *echo.Echo) {
	apiImpl := ClientImpl{}
	client.RegisterHandlers(e, apiImpl)
	e.GET("/v1/api/client-spec", respondApiSpec)
}

func respondApiSpec(ctx echo.Context) error {
	swagger, err := client.GetSwagger()
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, swagger, "  ")
}

type ClientImpl struct {
}

func (c ClientImpl) ListAll(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Not implemented")
}

func (c ClientImpl) ClientDetail(ctx echo.Context, clientId string) error {
	return ctx.String(http.StatusOK, "Not implemented: "+clientId)
}

func (c ClientImpl) CreateClient(ctx echo.Context, clientId string) error {
	return ctx.String(http.StatusOK, "Not implemented: "+clientId)
}
