package db

import (
	"database/sql/driver"
	"encoding/base64"
)

// EncryptoKey 加密key
var EncryptoKey []byte

// Crypto 支援加解密
type Crypto string

// SetEncrypKey ...
func SetEncrypKey(cfg *Config) {
	EncryptoKey = []byte(cfg.Secrets)
}

// String 方便轉成string
// 轉成Crypto 請使用 db.Crypto("要轉的字串")
func (t Crypto) String() string {
	return string(t)
}

// Scan from db
func (t *Crypto) Scan(src interface{}) error {
	switch tmp := src.(type) {
	case string:
		b, err := base64.StdEncoding.DecodeString(tmp)
		if err != nil {
			*t = Crypto(tmp)
			return nil
		}
		decrypto, err := AESDecrypt(b, EncryptoKey)
		if err != nil {
			*t = Crypto(tmp)
			return nil
		}
		*t = Crypto(decrypto)
	case []byte:
		b, err := base64.StdEncoding.DecodeString(string(tmp))
		if err != nil {
			*t = Crypto(tmp)
			return nil
		}
		decrypto, err := AESDecrypt(b, EncryptoKey)
		if err != nil {
			*t = Crypto(tmp)
			return nil
		}
		*t = Crypto(decrypto)
	default:
		*t = ""
	}
	return nil
}

// Value to db
func (t Crypto) Value() (driver.Value, error) {
	encrypto := AESEncrypt([]byte(t), EncryptoKey)
	return base64.StdEncoding.EncodeToString(encrypto), nil
}
