package skctl

import (
	"context"
	"fmt"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	hostInfoCliByGRPC = grpc_cli.GetGrpcClient[hostinfo_pb.HostInfoClient](grpc_cli.HostInfo, global.EngineGrpcServerAddr)
)

func GetHostInfo() error {
	hosts, err := hostInfoCliByGRPC.FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		logx.Error(err)
		return err
	}
	fmt.Printf("%s %-23s %-10s %s\n", "ID", "IP", "PORT", "OS")
	for i := 0; i < len(hosts.Data); i++ {
		fmt.Printf("%d %-25s %-10d %s\n", hosts.Data[i].Id, hosts.Data[i].Ip, hosts.Data[i].Port, hosts.Data[i].Os)
	}

	return nil
}
