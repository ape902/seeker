package main

import (
	"flag"
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/handler"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/models/system"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "监听地址")
	port := flag.Int("port", 50050, "端口")
	config := flag.String("config", "config/seeker.yaml", "配置文件")
	flag.Parse()

	// 初始化日志
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig(*config)

	//初始化数据库
	initialize.InitGorm()

	// 初始化Minio
	initialize.InitMinio()

	useGrpcListen(*ip, *port)
}
func useGrpcListen(ip string, port int) {
	server := grpc.NewServer()
	hostinfo_pb.RegisterHostInfoServer(server, &cmdb.HostInfo{})
	user_center_pb.RegisterUserServer(server, &system.User{})
	command_pb.RegisterCommandServer(server, &handler.RemoteHostController{})
	minio_pb.RegisterMinioServer(server, &handler.MinioServer{})

	addr := fmt.Sprintf("%s:%d", ip, port)
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
