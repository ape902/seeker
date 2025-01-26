package seeker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"github.com/gin-gonic/gin"
)

func getAgentCli(c *gin.Context) agent_pb.AgentClient {
	ip := c.Query("ip")
	port := c.Query("port")
	addr := fmt.Sprintf("%s:%s", ip, port)

	return grpc_cli.GetGrpcClient[agent_pb.AgentClient](grpc_cli.Agent, addr)
}

func Discovery(c *gin.Context) {
	agentGRPCCli := getAgentCli(c)
	procs, err := agentGRPCCli.FindProcInfo(context.Background(), nil)
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

	agentGRPCCli := getAgentCli(c)
	resp, err := agentGRPCCli.AgentComm(context.Background(), &agent_pb.Info{
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
