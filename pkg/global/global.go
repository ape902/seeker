package global

import (
	"github.com/ape902/seeker/pkg/config"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type (
	//AuthMode 主机访问认证类型
	AuthMode int8
)

const (
	//PASSWORD 主机密码
	PASSWORD AuthMode = iota + 1
	//PUBLICKEY 主机公钥
	PUBLICKEY
)

const (
	// 密钥长度必须为16、24或32个字节
	ENCRYPTKEY             = "*c*dTwJ%!JaGM7zL"
	GRPCProxyDefaultHeader = "grpc_proxy_name"
)

var (
	//DBCli 数据库
	DBCli *gorm.DB

	//MinioClient Minio客户端
	MinioClient *minio.Client

	// 配置信息
	ServerConfig = new(config.ServerConfig)
	JWTConfig    = new(config.JWTConfig)
	DBConfig     = new(config.DatabaseConfig)
)
