package view

import (
	"pokemon/pkg/model"

	"github.com/labstack/echo/v4"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/errors"
)

// RegisterType 註冊模式
type RegisterType int8

const (
	// RegisterByEmail 通過 email 註冊
	RegisterByEmail RegisterType = iota + 1
	// RegisterByPhone 通過 電話 註冊
	RegisterByPhone
)

// RegisterReq ...
type RegisterReq struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
	AcceptLanguage string `json:"-"`
}

var (
	ErrNameNotFilledIn            = errors.NewWithMessage(errors.ErrInvalidInput, "ErrNameNotFilledIn")
	ErrPasswordInvalidInput       = errors.NewWithMessage(errors.ErrInvalidInput, "ErrPasswordInvalidInput")
	ErrConfirmPasswordNotFilledIn = errors.NewWithMessage(errors.ErrInvalidInput, "ErrConfirmPasswordNotFilledIn")
	ErrConfirmPasswordIncorrect   = errors.NewWithMessage(errors.ErrInvalidInput, "ErrConfirmPasswordIncorrect")
)

// BindAndVerify ...
func (req *RegisterReq) BindAndVerify(c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return errors.WithStack(errors.ErrInvalidInput)
	}

	req.AcceptLanguage = c.Request().Header.Get("Accept-Language")

	if req.Name == "" {
		return errors.WithStack(ErrNameNotFilledIn)
	}

	if len(req.Password) < 8 {
		return errors.WithStack(ErrPasswordInvalidInput)
	}

	if req.VerifyPassword == "" {
		return errors.WithStack(ErrConfirmPasswordNotFilledIn)
	}

	if req.Password != req.VerifyPassword {
		return errors.WithStack(ErrConfirmPasswordIncorrect)
	}

	return nil
}

// ConvertToIdentityAccount ...
func (req *RegisterReq) ConvertToIdentityAccount() model.IdentityAccount {
	acc := model.IdentityAccount{
		Name:     req.Name,
		Password: db.Crypto(req.Password),
	}
	return acc
}

// RegisterResp ...
type RegisterResp struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (req *LoginReq) ConvertToIdentityAccount() model.IdentityAccount {
	return model.IdentityAccount{
		Name:     req.Name,
		Password: db.Crypto(req.Password),
	}
}

func (req *LoginReq) BindAndVerify(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return errors.WithStack(errors.ErrInvalidInput)
	}

	if req.Name == "" {
		return errors.WithStack(ErrNameNotFilledIn)
	} else if req.Password == "" {
		return errors.WithStack(ErrPasswordInvalidInput)
	}

	return nil
}

type LoginResp struct {
	Token string `json:"token"`
}
