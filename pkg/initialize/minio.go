package initialize

import (
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() {
	endpoint := "192.168.119.80:11000"
	accessKeyID := "admin"
	secretAccessKey := "123456789"

	var err error
	global.MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		logx.Panic(err)
	}
}
