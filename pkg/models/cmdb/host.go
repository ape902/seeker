package cmdb

import (
	"context"
	"errors"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	Authentication struct {
		Username string `json:"username"`
		AuthMode int    `json:"auth_mode"`
		Auth     string `json:"auth"`
	}
	Host struct {
		models.BaseModel
		IP    string `json:"ip"`
		Port  int    `json:"port"`
		OS    string `json:"os"`
		Label string `json:"label"`
	}
	HostInfo struct {
		Host
		Authentication
	}
)

// HostByHostInfo 通过HostInfo转换数据为host
func (h *HostInfo) HostByHostInfo() *Host {
	return &Host{
		IP:    h.IP,
		Port:  h.Port,
		OS:    h.OS,
		Label: h.Label,
	}
}

func (HostInfo) TableName() string {
	return "seeker_cmdb_hosts"
}

// FindAllHostInfo 查询所有主机信息数据
//func FindAllHostInfo() ([]HostInfo, error) {
//	data := make([]HostInfo, 0)
//	if err := global.DBCli.Model(&HostInfo{}).Find(&data).Error; err != nil {
//		return data, err
//	}
//
//	return data, nil
//}

func (h *HostInfo) FindAll(ctx context.Context, emp *emptypb.Empty) (*hostinfo_pb.HostInfoResp, error) {
	pb := &hostinfo_pb.HostInfoResp{}
	if err := global.DBCli.Model(&HostInfo{}).Count(&pb.Total).Find(&pb.Data).Error; err != nil {
		return pb, err
	}

	return pb, nil
}

// FindPage 主机分页查询
func (h *HostInfo) FindPage(ctx context.Context, page *hostinfo_pb.HostInfoPageInfo) (*hostinfo_pb.HostInfoResp, error) {
	pbHost := &hostinfo_pb.HostInfoResp{}
	if err := global.DBCli.Model(&HostInfo{}).Count(&pbHost.Total).
		Offset((int(page.Index) - 1) * int(page.Size)).Limit(int(page.Size)).
		Find(&pbHost.Data).Error; err != nil {
		return pbHost, err
	}

	return pbHost, nil
}

// Create 创建主机
func (h *HostInfo) Create(ctx context.Context, hostInfo *hostinfo_pb.HostAndAuthentication) (*emptypb.Empty, error) {
	//当密码不是空的时候，做密码加密。
	if hostInfo.Auth != "" {
		newPassword, err := encryptions.Base64AESCBCEncrypt([]byte(hostInfo.Auth), []byte(global.ENCRYPTKEY))
		if err != nil {
			return nil, err
		}
		h.Auth = newPassword
	}

	h.IP = hostInfo.Ip
	h.Port = int(hostInfo.Port)
	h.OS = hostInfo.OS
	h.Label = hostInfo.Label
	h.Username = hostInfo.Username
	h.AuthMode = int(hostInfo.AuthMode)
	h.Auth = hostInfo.Auth

	return nil, global.DBCli.Create(&h).Error
}

// Delete 删除主机
func (h *HostInfo) Delete(ctx context.Context, ids *hostinfo_pb.HostInfoIdsRequest) (*emptypb.Empty, error) {
	return nil, global.DBCli.Where("id in ?", ids.Ids).Delete(&HostInfo{}).Error
}

// UpdateHost 更新主机
func (h *HostInfo) UpdateHost(ctx context.Context, host *hostinfo_pb.Host) (*emptypb.Empty, error) {
	hostInfo := &Host{
		IP:    host.Ip,
		Port:  int(host.Port),
		OS:    host.OS,
		Label: host.Label,
	}

	return nil, global.DBCli.Model(&HostInfo{}).Where("id=?", host.Id).Updates(hostInfo).Error
}

// UpdateAuthentication 更新主机认证信息
func (h *HostInfo) UpdateAuthentication(ctx context.Context, auth *hostinfo_pb.Authentication) (*emptypb.Empty, error) {
	var count int64
	if err := global.DBCli.Model(&HostInfo{}).Where("id=?", auth.ID).Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("data not exist")
	}

	newPassword, err := encryptions.Base64AESCBCEncrypt([]byte(auth.Auth), []byte(global.ENCRYPTKEY))
	if err != nil {
		return nil, err
	}
	var authInfo Authentication
	authInfo.AuthMode = int(auth.AuthMode)
	authInfo.Username = auth.Username
	authInfo.Auth = newPassword

	return nil, global.DBCli.Model(&HostInfo{}).Where("id = ?", auth.ID).
		Updates(&authInfo).Error
}

// IsExistByIp 以IP检查主机是否存在
func (h *HostInfo) IsExistByIp(ctx context.Context, ip *hostinfo_pb.HostInfoIpRequest) (*hostinfo_pb.HostInfoIsExists, error) {
	var count int64
	pb := &hostinfo_pb.HostInfoIsExists{}
	if err := global.DBCli.Model(&HostInfo{}).Where("ip = ?", ip.Ip).Count(&count).Error; err != nil {
		pb.IsExist = false
		return pb, err
	}
	pb.IsExist = count != 0

	return pb, nil
}
