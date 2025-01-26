package initialize

import (
	"log"
	"os"
	"strings"

	"github.com/ape902/seeker/pkg/config"
	"github.com/ape902/seeker/pkg/global"
	"github.com/spf13/viper"
)

func InitConfig(configFile string) {
	context, err := os.ReadFile(configFile)
	if os.IsNotExist(err) {
		log.Printf("警告: %s 配置文件未找到，将使用环境变量或默认地址作为配置", configFile)
		global.ServerConfig = config.NewServerConfigByEnv()
		return
	} else {
		viper.SetConfigFile(configFile)
		if err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(context)))); err != nil {
			log.Fatal(err)
		}

		// 服务配置
		server := viper.Sub("server")
		if server == nil {
			log.Fatal(err)
		}
		global.ServerConfig = config.NewServerConfigByViper(server)

		// 数据库配置
		database := viper.Sub("database")
		if server == nil {
			log.Fatal(err)
		}
		global.DBConfig = config.NewDatabaseConfig(database)

		// JWT配置
		jwt := viper.Sub("jwt")
		if jwt == nil {
			log.Fatal(err)
		}
		global.JWTConfig = config.NewJWTConfigByViper(jwt)
	}
}
