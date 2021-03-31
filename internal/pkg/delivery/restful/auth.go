package restful

import (
	"net/http"
	"pokemon/internal/pkg/delivery/restful/view"
	"pokemon/internal/pkg/model"

	"github.com/labstack/echo/v4"
)

// @Title Register
// @Description 註冊
// @Success  200  object  view.RegisterReq   "註冊需要欄位"
// @Failure  500  object  view.ErrorResp     "系統錯誤"
// @Resource 註冊
// @Router /apis/v1/auth/register [post]
func (h *handler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var (
		req  view.RegisterReq
		iAcc model.IdentityAccount
	)

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	iAcc = req.ConvertToIdentityAccount()

	if err := h.authSvc.CreateIdentityAccount(ctx, &iAcc); err != nil {
		return err
	}

	token, err := iAcc.CreateToken(h.appCfg)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, view.RegisterResp{
		Token: token,
	})
}

// @Title Login
// @Description 註冊
// @Success  200  object  view.LoginReq   "登陸"
// @Failure  400  object  view.ErrorResp  "信箱或密碼為空"
// @Failure  401  object  view.ErrorResp  "密碼錯誤"
// @Failure  500  object  view.ErrorResp  "系統錯誤"
// @Resource 註冊
// @Router /apis/v1/auth/login [post]
func (h *handler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var (
		req  view.LoginReq
		iAcc model.IdentityAccount
		err  error
	)

	if err := req.BindAndVerify(c); err != nil {
		return err
	}

	iAcc = req.ConvertToIdentityAccount()
	iAcc, err = h.authSvc.VerifyIdentityAccount(ctx, iAcc)
	if err != nil {
		return err
	}

	token, err := iAcc.CreateToken(h.appCfg)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, view.LoginResp{
		Token: token,
	})
}
