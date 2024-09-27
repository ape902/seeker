package grpc_cli

import (
	"fmt"
	"github.com/ape902/corex/logx"
	"google.golang.org/grpc"
)

type (
	GrpcDial interface {
		Dial() *grpc.ClientConn
	}
	GrpcConfig struct {
		addr string
	}
)

func NewGrpcDial(addr string) GrpcDial {
	return &GrpcConfig{
		addr: addr,
	}
}

func (g *GrpcConfig) Dial() *grpc.ClientConn {
	if g.addr == "" {
		g.addr = fmt.Sprintf("%s:%d", "127.0.0.1", 50050)
	}
	clientConn, err := grpc.NewClient(g.addr, grpc.WithInsecure())
	if err != nil {
		logx.Error(err)
		return nil
	}

	return clientConn
}
