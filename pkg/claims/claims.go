package claims

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gitlab.com/howmay/gopher/errors"
)

// Claims ...
type Claims struct {
	ID        int       `json:"id"`
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
		return nil, errors.WithMessage(errors.ErrAuthenticationFailed, "fail to convert to jwt token")
	}
	claims, ok := user.Claims.(*Claims)
	if !ok {
		return nil, errors.WithMessage(errors.ErrAuthenticationFailed, "fail to convert to claims")
	}

	return claims, nil
}
