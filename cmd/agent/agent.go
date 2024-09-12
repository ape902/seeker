package main

import (
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/ape902/seeker/pkg/handler"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	agentCommand = &cobra.Command{
		Use:               "run",
		Short:             "Agent",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			initServer()
		},
	}
)

func initServer() {
	// 初始化日志
	initialize.InitLogger()

	//启用GRPC监听
	useGrpcListen()
}

func useGrpcListen() {
	server := grpc.NewServer()

	// Agent 远程执行命令GRPC
	//多个Server注册时下面添加即可
	command_pb.RegisterCommandServer(server, &handler.Commands{})

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", 58899)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logx.Fatal(err)
	}
	go func() {
		if err = server.Serve(listen); err != nil {
			logx.Error(err)
			return
		}
		//关闭Listen监听
		if err := listen.Close(); err != nil {
			logx.Error(err)
			return
		}
		//停止GRPC服务
		server.Stop()
	}()
	fmt.Println(fmt.Sprintf("GRPC: %s running...", addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

func main() {
	if err := agentCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
