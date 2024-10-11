package seeker

import (
	"fmt"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
)

type (
	IDS struct {
		IDS []int32 `json:"ids"`
	}
)

// connUserCenterGrpc
func connUserCenterGrpc() user_center_pb.UserClient {
	dial := grpc_cli.NewGrpcDial("127.0.0.1:50050").Dial()

	return user_center_pb.NewUserClient(dial)
}

// connHostInfoGrpc
func connHostInfoGrpc() hostinfo_pb.HostInfoClient {
	dial := grpc_cli.NewGrpcDial("127.0.0.1:50050").Dial()

	return hostinfo_pb.NewHostInfoClient(dial)
}

// connStorageGrpc
func connStorageGrpc() minio_pb.MinioClient {
	dial := grpc_cli.NewGrpcDial("127.0.0.1:50050").Dial()

	return minio_pb.NewMinioClient(dial)
}

func connCommandGrpc(ip, port string) command_pb.CommandClient {
	dial := grpc_cli.NewGrpcDial(fmt.Sprintf("%s:%s", ip, port)).Dial()

	return command_pb.NewCommandClient(dial)
}
