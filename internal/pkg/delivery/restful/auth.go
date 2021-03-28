package restful

import (
	"net/http"
	"pokemon/internal/pkg/delivery/restful/view"
	"pokemon/internal/pkg/model"

	"github.com/labstack/echo/v4"
)

// Register ...
func (h *handler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var (
		req      view.RegisterReq
		iAccount model.IdentityAccount
	)

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	iAccount = req.ConvertToIdentityAccount()
	err := h.authSvc.CreateIdentityAccount(ctx, &iAccount)
	if err != nil {
		return err
	}

	token, err := iAccount.CreateToken(h.appCfg)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, view.RegisterResp{
		Token: token,
	})
}

func (h *handler) Login(ctx echo.Context) error {
	return nil
}
