package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	api "github.com/identityOrg/cerberus-api"
	"github.com/identityOrg/cerberus-api/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.IPExtractor = echo.ExtractIPFromXFFHeader(echo.TrustLoopback(true), echo.TrustPrivateNet(true))

	registerAllAPIs(e)

	if err := e.Start("localhost:9090"); err != nil {
		e.Logger.Errorf("failed to start server %v", err)
	}
}

var swagger *openapi3.Swagger

func init() {
	var err error
	swagger, err = api.GetSwagger()
	if err != nil {
		panic(err)
	}
}

func registerAllAPIs(e *echo.Echo) {
	apiImpl := impl.NewCerberusAPI()
	api.RegisterHandlers(e, apiImpl)
	e.GET("/v1/spec", respondApiSpec)
}

func respondApiSpec(ctx echo.Context) error {
	return ctx.JSONPretty(http.StatusOK, swagger, "  ")
}
