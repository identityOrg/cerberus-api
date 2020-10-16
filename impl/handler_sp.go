package impl

import (
	"github.com/identityOrg/cerberus-api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c CerberusAPI) GetServiceProviders(ctx echo.Context, params api.GetServiceProvidersParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) CreateServiceProvider(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) DeleteServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) UpdateServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetCredentials(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GenerateCredentials(ctx echo.Context, id int, params api.GenerateCredentialsParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}
