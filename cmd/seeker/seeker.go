package main

import (
	"fmt"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/tools/httpx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var (
	seekerCommand = &cobra.Command{
		Use:               "run",
		Short:             "Seeker",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			initServer()
		},
	}
)

func initServer() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig("config/seeker.yaml")

	// 初始化数据库
	initialize.InitGorm()

	// 初始化gin
	e := initialize.Engine(gin.Mode())

	addr := fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	fmt.Println(fmt.Sprintf("%s running...", addr))
	httpx.RunHttp(srv)
}

func main() {
	if err := seekerCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
