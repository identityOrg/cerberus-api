package impl

import (
	"github.com/identityOrg/cerberus-api/backend/api/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) ListUser(ctx echo.Context, params user.ListUserParams) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (u UserServiceImpl) UserDetail(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
