package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/tools/versionx"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	configFile string

	seekerCommand = &cobra.Command{
		Use:               "run",
		Short:             "Seeker",
		DisableAutoGenTag: true,
		Version:           versionx.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			initServer()
		},
	}
)

func init() {
	flag.StringVar(&configFile, "config", "config/seeker.yaml", "配置文件路径")
	flag.Parse()
}

func initServer() {
	// 初始化日志
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig(configFile)

	// 初始化gin
	e := initialize.Engine(gin.Mode())

	addr := global.ServerConfig.Host + ":" + strconv.Itoa(global.ServerConfig.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logx.Infof("HTTP服务启动成功: %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Fatalf("HTTP服务启动失败: %v", err)
		}
	}()

	<-quit
	logx.Info("正在关闭服务...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		logx.Errorf("服务关闭出错: %v", err)
	}

	logx.Info("服务已成功关闭")
}

func main() {
	if err := seekerCommand.Execute(); err != nil {
		logx.Fatal(err)
	}
}
