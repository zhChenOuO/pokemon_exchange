package configuration

// App ..
type App struct {
	JwtSecrets   string `json:"jwt_secrets" mapstructure:"jwt_secrets"`
	JwtExpireSec int    `json:"jwt_expire_sec" mapstructure:"jwt_expire_sec"`
}
