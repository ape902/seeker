package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type ServerConfig struct {
	Host string
	Port int
}

// NewServerConfigByViper 通过Viper获取配置信息
func NewServerConfigByViper(cfg *viper.Viper) *ServerConfig {
	return &ServerConfig{
		Host: cfg.GetString("host"),
		Port: cfg.GetInt("port"),
	}
}

// NewServerConfigByEnv 通过环境变量获取配置信息
func NewServerConfigByEnv() *ServerConfig {
	port := os.Getenv("port")
	portInt, _ := strconv.Atoi(port)

	return &ServerConfig{
		Host: os.Getenv("host"),
		Port: portInt,
	}
}
