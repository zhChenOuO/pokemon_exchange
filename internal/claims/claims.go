package claims

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/howmay/gopher/errors"
)

// Claims ...
type Claims struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.StandardClaims
}

// GetClaims ...
func GetClaims(c echo.Context) (*Claims, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.WithMessage(errors.ErrAuthenticationFailed, "need token")
	}
	claims, ok := user.Claims.(*Claims)
	if !ok {
		return nil, errors.WithMessage(errors.ErrAuthenticationFailed, "check token is legal")
	}

	return claims, nil
}

func (c *Claims) GetID() uint64 {
	return c.ID
}
