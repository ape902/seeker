package seeker

import (
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
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
	dial := grpc_cli.NewGrpcDial("0.0.0.0:50050").Dial()

	return user_center_pb.NewUserClient(dial)
}

// connHostInfoGrpc
func connHostInfoGrpc() hostinfo_pb.HostInfoClient {
	dial := grpc_cli.NewGrpcDial("0.0.0.0:50050").Dial()

	return hostinfo_pb.NewHostInfoClient(dial)
}
