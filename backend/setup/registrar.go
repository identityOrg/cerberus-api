package setup

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/identityOrg/cerberus-api/backend/api"
	"github.com/identityOrg/cerberus-api/backend/api/impl"
	"github.com/labstack/echo/v4"
	"net/http"
)

var swagger *openapi3.Swagger

func init() {
	var err error
	swagger, err = api.GetSwagger()
	if err != nil {
		panic(err)
	}
}

func RegisterAllAPIs(e *echo.Echo) {
	apiImpl := impl.CerberusAPI{}
	api.RegisterHandlers(e, apiImpl)
	e.GET("/v1/spec", respondApiSpec)
}

func respondApiSpec(ctx echo.Context) error {
	return ctx.JSONPretty(http.StatusOK, swagger, "  ")
}
