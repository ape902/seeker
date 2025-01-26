package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/handler"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/models/system"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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
	hostinfo_pb.RegisterHostInfoServer(server, &cmdb.HostPB{})
	user_center_pb.RegisterUserServer(server, &system.UserCenterPB{})
	agent_pb.RegisterAgentServer(server, &handler.RemoteHostControllerPB{})
	minio_pb.RegisterMinioServer(server, &handler.MinioServerPB{})

	addr := fmt.Sprintf("%s:%d", ip, port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logx.Fatal(err)
	}
	go func() {
		logx.Infof("GRPC服务启动成功: %s", addr)
		if err = server.Serve(listen); err != nil {
			logx.Errorf("GRPC服务运行错误: %v", err)
			return
		}
	}()

	// 优雅关闭服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logx.Info("正在关闭GRPC服务...")
	// 优雅停止GRPC服务
	server.GracefulStop()
	logx.Info("GRPC服务已关闭")

	// 关闭监听
	if err := listen.Close(); err != nil {
		logx.Errorf("关闭监听失败: %v", err)
	}

	// 清理GRPC连接池
	grpc_cli.CloseAllConnections()
}
