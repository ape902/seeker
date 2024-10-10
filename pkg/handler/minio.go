package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/minio_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/remote_host"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/sftp"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"strings"
)

type (
	MinioServerPB struct {
		minio_pb.UnimplementedMinioServer
	}
)

func (m *MinioServerPB) GetObject(ctx context.Context, info *minio_pb.GetObjectInfo) (*emptypb.Empty, error) {
	object, err := global.MinioClient.GetObject(ctx, info.BucketName, info.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	decryptPassword, err := encryptions.Base64AESCBCDecrypt(info.Auth, []byte(global.ENCRYPTKEY))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	sshCli, err := remote_host.NewSSHDial(info.Addr, info.Username, string(decryptPassword), int8(info.AuthMode))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	ftpCli, err := sftp.NewClient(sshCli.Client)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	objectSplit := strings.Split(info.ObjectName, "/")
	remoteFile, err := ftpCli.Create(sftp.Join(info.Prefix, objectSplit[len(objectSplit)-1]))
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	defer remoteFile.Close()

	written, err := io.Copy(remoteFile, object)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	fmt.Println(format.FileSize(written))

	return nil, nil
}

func (m *MinioServerPB) PutObject(ctx context.Context, info *minio_pb.PutRest) (*minio_pb.PutResp, error) {
	pb := &minio_pb.PutResp{}
	uploadInfo, err := global.MinioClient.PutObject(ctx, info.BucketName, info.Name, bytes.NewReader(info.Data), info.Size, minio.PutObjectOptions{})
	if err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		return pb, err
	}
	pb.Error = uploadInfo.VersionID

	return pb, nil
}

func (m *MinioServerPB) RemoveObject(ctx context.Context, info *minio_pb.RemoveObjectRest) (*emptypb.Empty, error) {
	if err := global.MinioClient.RemoveObject(ctx, info.BucketName, info.ObjectName, minio.RemoveObjectOptions{}); err != nil {
		logx.Error(err)
		return nil, err
	}

	return nil, nil
}

func (m *MinioServerPB) ListBucket(ctx context.Context, info *emptypb.Empty) (*minio_pb.BucketListResp, error) {
	pb := &minio_pb.BucketListResp{}

	bucketInfo, err := global.MinioClient.ListBuckets(ctx)
	if err != nil {
		logx.Error(err)
		return pb, err
	}

	for i := 0; i < len(bucketInfo); i++ {
		bucket := &minio_pb.BucketInfo{}
		bucket.Name = bucketInfo[i].Name
		bucketInfo[i].CreationDate.Unix()
		bucket.Second = int64(bucketInfo[i].CreationDate.Unix())

		pb.Data = append(pb.Data, bucket)
	}

	return pb, nil
}

func (m *MinioServerPB) ListObject(ctx context.Context, info *minio_pb.ListObjectRest) (*minio_pb.ListObjectResp, error) {
	pb := &minio_pb.ListObjectResp{}

	objectCh := global.MinioClient.ListObjects(ctx, info.Bucket, minio.ListObjectsOptions{
		Prefix:    info.Prefix,
		Recursive: info.Recursive,
	})

	for object := range objectCh {
		if object.Err != nil {
			logx.Error(object.Err)
			return pb, object.Err
		}

		objectInfo := &minio_pb.ObjectInfo{}
		keySplice := strings.Split(object.Key, "/")
		objectInfo.Name = keySplice[len(keySplice)-1]
		objectInfo.LastModified = int64(object.LastModified.Second())
		objectInfo.Size = object.Size
		objectInfo.Etag = object.ETag

		pb.Data = append(pb.Data, objectInfo)
	}

	return pb, nil
}
