package api

import (
	"github.com/identityOrg/cerberus-api/backend/api/client"
	"github.com/identityOrg/cerberus-api/backend/api/impl"
	"github.com/identityOrg/cerberus-api/backend/api/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAllAPIs(e *echo.Echo) {
	registerClientAPIs(e)
	registerUserAPIs(e)
}

func registerClientAPIs(e *echo.Echo) {
	apiImpl := impl.ClientServiceImpl{}
	client.RegisterHandlers(e, apiImpl)
	e.GET("/v1/api/client-spec", respondClientApiSpec)
}

func respondClientApiSpec(ctx echo.Context) error {
	swagger, err := client.GetSwagger()
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, swagger, "  ")
}

func registerUserAPIs(e *echo.Echo) {
	apiImpl := impl.UserServiceImpl{}
	user.RegisterHandlers(e, apiImpl)
	e.GET("/v1/api/user-spec", respondUserApiSpec)
}

func respondUserApiSpec(ctx echo.Context) error {
	swagger, err := user.GetSwagger()
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, swagger, "  ")
}
