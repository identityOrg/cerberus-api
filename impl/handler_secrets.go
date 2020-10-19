package impl

import (
	api "github.com/identityOrg/cerberus-api"
	"github.com/labstack/echo/v4"
)

func (c *CerberusAPI) GetSecretChannels(ctx echo.Context, params api.GetSecretChannelsParams) error {
	panic("implement me")
}

func (c *CerberusAPI) CreateSecretChannel(ctx echo.Context) error {
	panic("implement me")
}

func (c *CerberusAPI) FindSecretChannelByAlgouse(ctx echo.Context, params api.FindSecretChannelByAlgouseParams) error {
	panic("implement me")
}

func (c *CerberusAPI) FindSecretChannelByName(ctx echo.Context, params api.FindSecretChannelByNameParams) error {
	panic("implement me")
}

func (c *CerberusAPI) DeleteSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) GetSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) RenewSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) UpdateSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}
