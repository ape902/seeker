package grpc_cli

import (
	"context"
	"sync"
	"time"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
)

const (
	defaultTimeout       = 10 * time.Second
	defaultRetries       = 3
	defaultCheckInterval = 5 * time.Second
)

// connPool 连接池管理
type connPool struct {
	mu    sync.RWMutex
	conns map[string]*grpc.ClientConn
}

var (
	pool     = &connPool{conns: make(map[string]*grpc.ClientConn)}
	cleanup  sync.Once
	shutdown = make(chan struct{})
)

// getDefaultGrpcOptions 获取默认的GRPC连接选项
func getDefaultGrpcOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  100 * time.Millisecond,
				Multiplier: 1.6,
				Jitter:     0.2,
				MaxDelay:   3 * time.Second,
			},
			MinConnectTimeout: defaultTimeout,
		}),
	}
}

// connectGRPC 通用的GRPC连接函数
func connectGRPC(addr string, serviceName string) (*grpc.ClientConn, error) {
	pool.mu.RLock()
	if conn, ok := pool.conns[addr]; ok && conn.GetState() != connectivity.Shutdown {
		pool.mu.RUnlock()
		return conn, nil
	}
	pool.mu.RUnlock()

	pool.mu.Lock()
	defer pool.mu.Unlock()

	// 双重检查，避免在获取写锁期间其他goroutine已经创建了连接
	if conn, ok := pool.conns[addr]; ok && conn.GetState() != connectivity.Shutdown {
		return conn, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	dial, err := grpc.DialContext(ctx, addr, getDefaultGrpcOptions()...)
	if err != nil {
		logx.Errorf("连接%sGRPC服务失败: %v", serviceName, err)
		return nil, err
	}

	// 启动连接状态监控
	go monitorConnection(addr, serviceName, dial)

	pool.conns[addr] = dial
	return dial, nil
}

// monitorConnection 监控连接状态
func monitorConnection(addr, serviceName string, conn *grpc.ClientConn) {
	ticker := time.NewTicker(defaultCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-shutdown:
			return
		case <-ticker.C:
			state := conn.GetState()
			if state == connectivity.TransientFailure || state == connectivity.Shutdown {
				logx.Warnf("%s连接状态异常: %v，尝试重新连接", serviceName, state)
				if err := conn.Close(); err != nil {
					logx.Errorf("关闭%s连接失败: %v", serviceName, err)
				}

				pool.mu.Lock()
				delete(pool.conns, addr)
				pool.mu.Unlock()
				return
			}
		}
	}
}

// CloseAllConnections 关闭所有连接
func CloseAllConnections() {
	cleanup.Do(func() {
		close(shutdown)
		pool.mu.Lock()
		defer pool.mu.Unlock()

		for addr, conn := range pool.conns {
			if err := conn.Close(); err != nil {
				logx.Errorf("关闭连接失败 %s: %v", addr, err)
			}
		}
		pool.conns = make(map[string]*grpc.ClientConn)
	})
}

// ServiceType GRPC服务类型
type ServiceType int

const (
	UserCenter ServiceType = iota
	HostInfo
	Storage
	Agent
)

// GetGrpcClient 获取GRPC客户端
// 使用示例:
// user_center_pb.UserClient client = GetGrpcClient[user_center_pb.UserClient](UserCenter, addr)
// hostinfo_pb.HostInfoClient client = GetGrpcClient[hostinfo_pb.HostInfoClient](HostInfo, addr)
// minio_pb.MinioClient client = GetGrpcClient[minio_pb.MinioClient](Storage, addr)
// command_pb.CommandClient client = GetGrpcClient[command_pb.CommandClient](Command, addr)
func GetGrpcClient[T any](serviceType ServiceType, addr string) T {
	if addr == "" {
		addr = "127.0.0.1:50050"
	}

	var serviceName string
	switch serviceType {
	case UserCenter:
		serviceName = "用户中心"
	case HostInfo:
		serviceName = "主机信息"
	case Storage:
		serviceName = "存储"
	case Agent:
		serviceName = "Agent"
	}

	dial, err := connectGRPC(addr, serviceName)
	if err != nil {
		var zero T
		return zero
	}

	switch serviceType {
	case UserCenter:
		return any(user_center_pb.NewUserClient(dial)).(T)
	case HostInfo:
		return any(hostinfo_pb.NewHostInfoClient(dial)).(T)
	case Storage:
		return any(minio_pb.NewMinioClient(dial)).(T)
	case Agent:
		return any(agent_pb.NewAgentClient(dial)).(T)
	default:
		var zero T
		return zero
	}
}
