package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"github.com/ape902/seeker/pkg/tools/worker_pool"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AgentChecker struct {
	pool *worker_pool.Pool
}

func NewAgentChecker() *AgentChecker {
	return &AgentChecker{
		pool: worker_pool.NewPool(10), // 创建一个大小为10的worker pool
	}
}

// StartCheck 启动agent状态检查
func (ac *AgentChecker) StartCheck() {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			ac.checkAllAgents()
		}
	}()
}

// checkAllAgents 检查所有agent的状态
func (ac *AgentChecker) checkAllAgents() {
	// 获取所有主机信息
	hostModel := &cmdb.HostInfo{}
	hosts, _, err := hostModel.FindAll()
	if err != nil {
		logx.Errorf("获取主机列表失败: %v", err)
		return
	}

	// 遍历所有主机，将检查任务提交到worker pool
	for _, host := range hosts {
		host := host // 创建副本以避免闭包问题
		ac.pool.Submit(func() {
			ac.checkAgent(&host)
		})
	}
}

// checkAgent 检查单个agent的状态
func (ac *AgentChecker) checkAgent(host *cmdb.HostInfo) {
	// 创建agent的gRPC客户端
	addr := fmt.Sprintf("%s:%d", host.IP, host.Port)
	agentCli := grpc_cli.GetGrpcClient[agent_pb.AgentClient](grpc_cli.Agent, addr)

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用HealthCheck接口
	resp, err := agentCli.HealthCheck(ctx, &emptypb.Empty{})

	// 更新agent状态
	var status int
	if err != nil || !resp.IsAlive {
		status = 2
		logx.Warnf("Agent检查失败 [%s]: %v", host.IP, err)
	} else {
		status = 1
		logx.Infof("Agent正常 [%s]: uptime=%d, cpu=%.2f%%, memory=%d",
			host.IP, resp.Uptime, resp.CpuUsage, resp.MemoryUsage)
	}

	// 更新数据库中的状态
	host.AgentStatus = status
	if err := host.UpdateHost(); err != nil {
		logx.Errorf("更新主机状态失败 [%s]: %v", host.IP, err)
	}
}
