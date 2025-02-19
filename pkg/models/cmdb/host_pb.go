package cmdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"github.com/ape902/seeker/pkg/tools/format"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	HostPB struct {
		// 继承GRPC的服务端扩展
		hostinfo_pb.UnimplementedHostInfoServer

		// 继承主机结构体，主要进行数据库操作方法调用
		host HostInfo
	}
)

// FindPage 主机分页查询
func (h *HostPB) FindPage(ctx context.Context, page *hostinfo_pb.HostInfoPageInfo) (*hostinfo_pb.HostInfoResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoResp{}
	pb.Code = codex.Success

	data, total, err := h.host.FindPage(int(page.Index), int(page.Size))
	if err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return nil, err
	}

	pb.Total = total

	for i := 0; i < len(data); i++ {
		pb.Data = append(pb.Data, &hostinfo_pb.Host{
			Id:          int32(data[i].Id),
			Ip:          data[i].IP,
			Port:        int32(data[i].Port),
			Os:          data[i].OS,
			Label:       format.StringToMap(data[i].Label),
			AgentStatus: int32(data[i].AgentStatus),
		})
	}

	return pb, nil
}

// Create 创建主机
func (h *HostPB) Create(ctx context.Context, host *hostinfo_pb.HostAndAuthentication) (*hostinfo_pb.HostInfoDefResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoDefResp{}
	pb.Code = codex.Success

	exist, err := h.host.IsExistByIp(host.Ip)
	if err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}
	//数据存在同样IP在就不用添加了
	if exist {
		pb.Error = errors.New("data exist").Error()
		pb.Code = codex.AlreadyExists
		return pb, nil
	}

	//当密码不是空的时候，做密码加密。
	if host.Auth != "" {
		newPassword, err := encryptions.Base64AESCBCEncrypt([]byte(host.Auth), []byte(global.ENCRYPTKEY))
		if err != nil {
			pb.Error = err.Error()
			pb.Code = codex.EncryptCreateFailed
			return pb, err
		}
		h.host.Auth = newPassword
	}

	h.host.IP = host.Ip
	h.host.Port = int(host.Port)
	h.host.OS = host.Os
	h.host.Label = format.MapToString(host.Label)
	h.host.Username = host.Username
	h.host.AuthMode = int(host.AuthMode)

	if err := h.host.Create(); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

// UpdateHost 更新主机
func (h *HostPB) UpdateHost(ctx context.Context, host *hostinfo_pb.Host) (*hostinfo_pb.HostInfoDefResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoDefResp{}
	pb.Code = codex.Success

	exist, err := h.host.IsExistById(int(host.Id))
	if err != nil {
		pb.Code = codex.ExecutionFailed
		pb.Error = err.Error()
		return pb, err
	}

	//数据不存在则不做数据更新
	if !exist {
		pb.Code = codex.NotExists
		pb.Error = fmt.Sprintf("%s data not exists", host.Ip)
		return pb, nil
	}

	h.host.Id = int(host.Id)
	h.host.IP = host.Ip
	h.host.Port = int(host.Port)
	h.host.OS = host.Os
	h.host.Label = format.MapToString(host.Label)

	if err := h.host.UpdateHost(); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

// UpdateAuthentication 更新主机认证信息
func (h *HostPB) UpdateAuthentication(ctx context.Context, auth *hostinfo_pb.Authentication) (*hostinfo_pb.HostInfoDefResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoDefResp{}
	pb.Code = codex.Success

	exist, err := h.host.IsExistById(int(auth.Id))
	if err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}
	if !exist {
		pb.Error = fmt.Sprint("data not exist")
		pb.Code = codex.NotExists
		return pb, nil
	}

	newPassword, err := encryptions.Base64AESCBCEncrypt([]byte(auth.Auth), []byte(global.ENCRYPTKEY))
	if err != nil {
		pb.Error = err.Error()
		pb.Code = codex.EncryptCreateFailed
		return pb, err
	}

	if err := h.host.UpdateAuthentication(int(auth.Id), Authentication{
		AuthMode: int(auth.AuthMode),
		Username: auth.Username,
		Auth:     newPassword,
	}); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

// Delete 删除主机
func (h *HostPB) Delete(ctx context.Context, ids *hostinfo_pb.HostInfoIdsRequest) (*hostinfo_pb.HostInfoDefResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoDefResp{}
	pb.Code = codex.Success

	i := format.Int32ToIntArray(ids.Ids)
	if err := h.host.DeleteById(i); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

// FindAll 查询所有主机信息数据
func (h *HostPB) FindAll(ctx context.Context, emp *emptypb.Empty) (*hostinfo_pb.HostInfoResp, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostInfoResp{}
	pb.Code = codex.Success

	data, total, err := h.host.FindAll()
	if err != nil {
		logx.Error(err)
		pb.Code = codex.DatabaseExecutionFailed
		pb.Error = err.Error()
		return pb, err
	}

	pb.Total = total

	for i := 0; i < len(data); i++ {
		pb.Data = append(pb.Data, &hostinfo_pb.Host{
			Id:    int32(data[i].Id),
			Ip:    data[i].IP,
			Port:  int32(data[i].Port),
			Os:    data[i].OS,
			Label: format.StringToMap(data[i].Label),
		})
	}

	return pb, nil
}

func (h *HostPB) FindHostByIp(ctx context.Context, ip *hostinfo_pb.HostInfoIpRequest) (*hostinfo_pb.HostAndAuthentication, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostAndAuthentication{}

	host, err := h.host.FindByIp(ip.Ip)
	if err != nil {
		logx.Error(err)
		return pb, err
	}

	pb.Id = int32(host.Id)
	pb.Ip = host.IP
	pb.Port = int32(host.Port)
	pb.Os = host.OS
	pb.Label = format.StringToMap(host.Label)
	pb.Username = host.Username
	pb.AuthMode = int32(host.AuthMode)
	pb.Auth = host.Auth

	return pb, nil
}

func (h *HostPB) FindById(ctx context.Context, req *hostinfo_pb.HostInfoIdsRequest) (*hostinfo_pb.HostAndAuthentication, error) {
	// 初始化host字段
	h.host = HostInfo{}

	pb := &hostinfo_pb.HostAndAuthentication{}

	host, err := h.host.FindById(int(req.Ids[0]))
	if err != nil {
		logx.Error(err)
		return pb, err
	}

	pb.Id = int32(host.Id)
	pb.Ip = host.IP
	pb.Port = int32(host.Port)
	pb.Os = host.OS
	pb.Label = format.StringToMap(host.Label)
	pb.Username = host.Username
	pb.AuthMode = int32(host.AuthMode)
	pb.Auth = host.Auth

	return pb, nil
}
