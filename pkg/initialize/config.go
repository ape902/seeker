package initialize

import (
	"fmt"
	"github.com/ape902/seeker/pkg/config"
	"github.com/ape902/seeker/pkg/global"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func InitConfig(configFile string) {
	context, err := os.ReadFile(configFile)
	if os.IsNotExist(err) {
		config.NewServerConfigByEnv()
		fmt.Println("Config Not Found.")
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
