package seeker

import (
	"context"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Discovery(c *gin.Context) {
	ip := c.Query("ip")
	port := c.Query("port")

	procs, err := connCommandGrpc(ip, port).FindProcInfo(context.Background(), nil)
	if err != nil {
		logx.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": procs,
	})

}

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
	ip := c.Query("ip")
	port := c.Query("port")

	resp, err := connCommandGrpc(ip, port).Command(context.Background(), &command_pb.Info{
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
