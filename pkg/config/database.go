package config

import "github.com/spf13/viper"

type (
	DatabaseConfig struct {
		Type               string
		Username           string
		Password           string
		Host               string
		Port               int
		DBName             string
		MaxOpenConn        int
		MaxIdleConn        int
		ConnMaxLifeSecond  int
		SlowLogMillisecond int
		PrepareStmt        bool
		SingularTable      bool
	}
)

func NewDatabaseConfig(cfg *viper.Viper) *DatabaseConfig {
	return &DatabaseConfig{
		Type:               cfg.GetString("type"),
		Username:           cfg.GetString("username"),
		Password:           cfg.GetString("password"),
		Host:               cfg.GetString("host"),
		Port:               cfg.GetInt("port"),
		DBName:             cfg.GetString("dbname"),
		MaxOpenConn:        cfg.GetInt("MaxOpenConn"),
		MaxIdleConn:        cfg.GetInt("MaxIdleConn"),
		ConnMaxLifeSecond:  cfg.GetInt("ConnMaxLifeSecond"),
		SlowLogMillisecond: cfg.GetInt("SlowLogMillisecond"),
		PrepareStmt:        cfg.GetBool("PrepareStmt"),
		SingularTable:      cfg.GetBool("SingularTable"),
	}
}
