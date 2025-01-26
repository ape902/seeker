package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/handler"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/tools/versionx"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	ip   string
	port int

	agentCommand = &cobra.Command{
		Use:               "run",
		Short:             "Agent",
		DisableAutoGenTag: true,
		Version:           versionx.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			initServer()
		},
	}
)

func init() {
	flag.StringVar(&ip, "ip", "0.0.0.0", "监听地址")
	flag.IntVar(&port, "port", 58899, "监听端口")
	flag.Parse()
}

func initServer() {
	// 初始化日志
	initialize.InitLogger()

	//启用GRPC监听
	useGrpcListen()
}

func useGrpcListen() {
	// GRPC服务器配置
	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
			MaxConnectionAge:  10 * time.Minute,
			Time:              20 * time.Second,
			Timeout:           5 * time.Second,
		}),
		grpc.MaxConcurrentStreams(100),
	)

	// 注册Agent服务
	agent_pb.RegisterAgentServer(server, &handler.RemoteHostControllerPB{})

	// 监听地址
	addr := fmt.Sprintf("%s:%d", ip, port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logx.Fatalf("启动GRPC监听失败: %v", err)
	}

	// 启动服务
	go func() {
		logx.Infof("GRPC服务启动成功: %s", addr)
		if err = server.Serve(listen); err != nil {
			logx.Errorf("GRPC服务运行错误: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logx.Info("正在关闭GRPC服务...")
	server.GracefulStop()
	logx.Info("GRPC服务已关闭")

	if err := listen.Close(); err != nil {
		logx.Errorf("关闭监听失败: %v", err)
	}
}

func main() {
	if err := agentCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
