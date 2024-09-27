package proxy

import (
	"context"
	"errors"
	"github.com/ape902/seeker/pkg/global"
	"google.golang.org/grpc/metadata"
	"strings"
)

// GRPCProxy GRPC代理地址检查，确定本地执行时，状态返回true
func GRPCProxy(ctx context.Context) (proxy string, next string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", errors.New("未设置上下文件信息")
	}

	header, exists := md[global.GRPCProxyDefaultHeader]
	if !exists {
		return "", "", errors.New("未查到代理信息")
	}

	addrs := strings.Split(header[0], ",")

	if len(addrs) == 1 {
		return "", addrs[0], nil
	}

	newAddr := strings.Join(addrs[1:], ",")
	return newAddr, addrs[0], nil
}
