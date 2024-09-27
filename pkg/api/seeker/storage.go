package seeker

import (
	"context"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func MinioBucketList(c *gin.Context) {
	bucketList, err := connStorageGrpc().ListBucket(context.Background(), &emptypb.Empty{})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, bucketList)
}
