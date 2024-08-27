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
	clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", "0.0.0.0", 50050), grpc.WithInsecure())
	if err != nil {
		logx.Error(err)
		return nil
	}

	return clientConn
}