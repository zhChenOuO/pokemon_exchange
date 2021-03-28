package model

import (
	"pokemon/configuration"
	"pokemon/pkg/claims"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/howmay/gopher/db"
	"gitlab.com/howmay/gopher/errors"
)

// IdentityAccount ...
type IdentityAccount struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  db.Crypto `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName ...
func (IdentityAccount) TableName() string {
	return "identity_accounts"
}

// CreateToken ...
func (identity *IdentityAccount) CreateToken(cfg *configuration.App) (string, error) {
	if cfg.JwtExpireSec == 0 {
		cfg.JwtExpireSec = 3600 * 24
	}

	claims := &claims.Claims{
		ID:    identity.ID,
		Name:  identity.Name,
		Email: identity.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Duration(cfg.JwtExpireSec) * time.Second).Unix(),
		},
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwtClaims.SignedString([]byte(cfg.JwtSecrets))
	if err != nil {
		return "", errors.WithMessagef(errors.ErrInternalError, "err: %s", err.Error())
	}
	return t, nil
}
