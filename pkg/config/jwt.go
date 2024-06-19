package config

import "github.com/spf13/viper"

type (
	JWTConfig struct {
		Key string
	}
)

func NewJWTConfigByViper(cfg *viper.Viper) *JWTConfig {
	return &JWTConfig{
		Key: cfg.GetString("key"),
	}
}
