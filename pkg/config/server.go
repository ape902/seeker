package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
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
	// 设置默认值
	defaultHost := "0.0.0.0"
	defaultPort := 8000

	// 获取并验证host环境变量
	host := os.Getenv("host")
	if host == "" {
		host = defaultHost
	}

	// 获取并验证port环境变量
	portStr := os.Getenv("port")
	port := defaultPort
	if portStr != "" {
		portInt, err := strconv.Atoi(portStr)
		if err == nil && portInt > 0 {
			port = portInt
		}
	}

	return &ServerConfig{
		Host: host,
		Port: port,
	}
}
