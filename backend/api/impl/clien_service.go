package impl

import (
	"github.com/identityOrg/cerberus-api/backend/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c CerberusAPI) GetClients(ctx echo.Context, params api.GetClientsParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) CreateClient(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) DeleteClient(ctx echo.Context, id string) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GetClient(ctx echo.Context, id string) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) UpdateClient(ctx echo.Context, id string) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c CerberusAPI) GenerateSecret(ctx echo.Context, id string) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}
