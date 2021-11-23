module pokemon

go 1.16

require (
	github.com/cenk/backoff v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/labstack/echo/v4 v4.1.17
	github.com/ory/dockertest/v3 v3.8.0
	github.com/pressly/goose v2.7.0+incompatible
	github.com/rs/xid v1.2.1
	github.com/rs/zerolog v1.20.0
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	gitlab.com/howmay/gopher v0.0.47
	go.etcd.io/etcd/client/v3 v3.5.1
	go.uber.org/fx v1.13.1
	gorm.io/gorm v1.20.2

)

replace gitlab.com/howmay/gopher => ../../gitlab/gopher
