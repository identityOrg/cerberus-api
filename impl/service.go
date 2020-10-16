package impl

import (
	 "github.com/identityOrg/cerberus-api"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CerberusAPI struct {
}

func NewCerberusAPI() *CerberusAPI {
	return &CerberusAPI{}
}

func (c CerberusAPI) GetClaims(ctx echo.Context, params api.GetClaimsParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) CreateClaim(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) DeleteClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) UpdateClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetScopes(ctx echo.Context, params api.GetScopesParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) CreateScope(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) DeleteScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) UpdateScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}
