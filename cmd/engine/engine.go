package main

import (
	"flag"
	"fmt"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/models/system"
	"google.golang.org/grpc"
	"net"
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

	server := grpc.NewServer()
	hostinfo_pb.RegisterHostInfoServer(server, &cmdb.HostInfo{})
	user_center_pb.RegisterUserServer(server, &system.User{})
	fmt.Println(*ip, *port)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := server.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}
}
