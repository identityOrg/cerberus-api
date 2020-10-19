package impl

import (
	api "github.com/identityOrg/cerberus-api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetClaims(ctx echo.Context, params api.GetClaimsParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) CreateClaim(ctx echo.Context) error {
	createReq := &api.Claim{}
	_ = ctx.Bind(createReq)
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) DeleteClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) FindClaimByName(ctx echo.Context, params api.FindClaimByNameParams) error {
	panic("implement me")
}
