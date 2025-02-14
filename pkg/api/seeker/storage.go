package seeker

import (
	"context"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	storageCliByGrpc = grpc_cli.GetGrpcClient[minio_pb.MinioClient](grpc_cli.Storage, global.EngineGrpcServerAddr)
)

func MinioBucketList(c *gin.Context) {
	bucketList, err := storageCliByGrpc.ListBucket(context.Background(), &emptypb.Empty{})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, bucketList)
}
