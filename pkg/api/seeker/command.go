package seeker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type commandInfo struct {
	Commands string `json:"commands"`
}

func RunCommand(c *gin.Context) {
	var cmd commandInfo
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "无效参数",
		})
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "192.168.119.82", 58899),
		grpc.WithInsecure())
	if err != nil {
		logx.Error(err)
		return
	}

	cc := command_pb.NewCommandClient(conn)

	resp, err := cc.Command(context.Background(), &command_pb.Info{
		Command: cmd.Commands,
	})

	if err != nil {
		logx.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": string(resp.Data),
		"msg":  resp.Msg,
		"err":  string(resp.Error),
	})

}
