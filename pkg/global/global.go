package global

import (
	"github.com/ape902/seeker/pkg/config"
	"gorm.io/gorm"
)

var (
	// 数据库
	DBCli *gorm.DB

	// 配置信息
	ServerConfig = new(config.ServerConfig)
	JWTConfig    = new(config.JWTConfig)
	DBConfig     = new(config.DatabaseConfig)
)
