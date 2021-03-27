package http

import (
	"encoding/json"

	"gitlab.com/howmay/gopher/errors"

	"github.com/labstack/echo/v4"
)

// Token ...
type Token struct {
	AccountID        int64             `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Namespace        string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ExpiresIn        int64             `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	TokenString      string            `protobuf:"bytes,4,opt,name=token_string,json=tokenString,proto3" json:"token_string,omitempty"`
	Claims           map[string]string `protobuf:"bytes,5,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Username         string            `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	AccountType      int64             `protobuf:"varint,7,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	RefreshExpiresIn int64             `protobuf:"varint,8,opt,name=refresh_expires_in,json=refreshExpiresIn,proto3" json:"refresh_expires_in,omitempty"`
}

// Claims 提供 token 認證的資訊
type Claims struct {
	AccountID   int64  `json:"accountID" description:"帳號id"`
	TokenString string `json:"tokenString" description:"token"`
	Namespace   string `json:"namespace" description:"命名空間"`
	Username    string `json:"username" description:"使用者名稱"`
	AccountType int32  `json:"accountType" description:"帳號類別,1:operator,2:merchant,3:agent,4:customerService,5:divideCommission,6:openAccount,7:divideCommissionAndOpenAccount"`
	// Roles       []Role            `json:"roles"`
	Others map[string]string `json:"others"`
}

// GetXSecret ...
func GetXSecret(c echo.Context) (*Claims, error) {

	r := c.Request()
	claimsStr := r.Header.Get("X-Secret")
	if claimsStr == "" {
		return nil, errors.WithMessage(errors.ErrMissingRequiredHeader, "token was not found")
	}

	token := Token{}
	err := json.Unmarshal([]byte(claimsStr), &token)
	if err != nil {
		return nil, errors.WithMessage(errors.ErrInternalError, "unmarshal token error")
	}

	return ConvertProtoToClaims(&token)
}

//ConvertProtoToClaims ...
func ConvertProtoToClaims(token *Token) (claims *Claims, err error) {

	claims = new(Claims)
	claims.AccountID = token.AccountID
	claims.TokenString = token.TokenString
	claims.Namespace = token.Namespace
	claims.Others = token.Claims
	claims.Username = token.Username
	claims.AccountType = int32(token.AccountType)

	// var roles []Role
	// roleStr, exist := token.Claims["Roles"]
	// if exist {
	// 	err := json.Unmarshal([]byte(roleStr), &roles)
	// 	if err != nil {
	// 		return nil, errors.WithMessage(errors.ErrInternalError, "unmarshal Roles error")
	// 	}
	// }

	// for _, role := range roles {
	// 	claims.Roles = append(claims.Roles, role)
	// }

	return claims, nil
}
