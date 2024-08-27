package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/ape902/seeker/pkg/handler"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

	server := grpc.NewServer()
	// Agent 远程执行命令GRPC
	command_pb.RegisterCommandServer(server, &handler.Commands{})

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", 58899)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logx.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s running...", addr))

	if err = server.Serve(listen); err != nil {
		logx.Fatal(err)
	}
}

func main() {
	if err := agentCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
