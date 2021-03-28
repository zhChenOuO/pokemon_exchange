package view

import (
	"pokemon/internal/pkg/model"

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
	RegisterType   RegisterType `json:"register_type"`
	Name           string       `json:"name"`
	Email          string       `json:"email"`
	Phone          string       `json:"phone"`
	PhoneAreaCode  string       `json:"phone_area_code"`
	Password       db.Crypto    `json:"password"`
	VerifyPassword db.Crypto    `json:"verify_password"`
	AcceptLanguage string       `json:"-"`
}

// BindAndVerify ...
func (req *RegisterReq) BindAndVerify(c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	req.AcceptLanguage = c.Request().Header.Get("Accept-Language")

	switch req.RegisterType {
	case RegisterByEmail:
		if req.Email == "" {
			return errors.WithStack(errors.ErrEmailNotFilledIn)
		}
	case RegisterByPhone:
		if req.PhoneAreaCode == "" {
			return errors.WithStack(errors.ErrPhoneAreaCodeNotFilledIn)
		}

		if req.Phone == "" {
			return errors.WithStack(errors.ErrPhoneNumberNotFilledIn)
		}
	default:
		return errors.WithStack(errors.ErrRegistrationTypeInvalidInput)
	}

	if req.Name == "" {
		return errors.WithStack(errors.ErrNameNotFilledIn)
	}

	if len(req.Password) < 8 {
		return errors.WithStack(errors.ErrPasswordInvalidInput)
	}

	if req.VerifyPassword == "" {
		return errors.WithStack(errors.ErrConfirmPasswordNotFilledIn)
	}

	if req.Password != req.VerifyPassword {
		return errors.WithStack(errors.ErrConfirmPasswordIncorrect)
	}

	return nil
}

// ConvertToIdentityAccount ...
func (req *RegisterReq) ConvertToIdentityAccount() model.IdentityAccount {
	acc := model.IdentityAccount{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	return acc
}

// RegisterResp ...
type RegisterResp struct {
	Token string `json:"token"`
}
