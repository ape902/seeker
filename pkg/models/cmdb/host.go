package cmdb

import (
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
)

type (
	Authentication struct {
		Username string `json:"username"`
		AuthMode int    `json:"auth_mode"`
		Auth     string `json:"auth"`
	}
	Host struct {
		models.BaseModel
		IP       string            `json:"ip"`
		Port     int               `json:"port"`
		OS       string            `json:"os"`
		Label    string            `json:"-"`
		LabelMap map[string]string `json:"label" gorm:"-"`
	}
	HostInfo struct {
		Host
		Authentication
	}
)

func (HostInfo) TableName() string {
	return "seeker_cmdb_hosts"
}

// IsExistById 通过ID判断数据是否存在。不存在为false
func (h *HostInfo) IsExistById(id int) (bool, error) {
	var count int64
	if err := global.DBCli.Model(h).Where("id=?", id).Count(&count).Error; err != nil {
		return false, err
	}

	return count != 0, nil
}

// IsExistByIp 通过IP判断数据是否存在。不存在为false
func (h *HostInfo) IsExistByIp(ip string) (bool, error) {
	var count int64
	if err := global.DBCli.Model(h).Where("ip=?", ip).Count(&count).Error; err != nil {
		return false, err
	}

	return count != 0, nil
}

// Create 主机信息创建
func (h *HostInfo) Create() error {
	return global.DBCli.Create(h).Error
}

// FindPage 分页查询
func (h *HostInfo) FindPage(index, size int) ([]HostInfo, int64, error) {
	var total int64
	hosts := make([]HostInfo, 0)
	if err := global.DBCli.Model(&HostInfo{}).Count(&total).
		Offset((index - 1) * size).Limit(size).Find(&hosts).Error; err != nil {
		return hosts, total, err
	}

	return hosts, total, nil
}

// UpdateHost 更新主机
func (h *HostInfo) UpdateHost() error {
	return global.DBCli.Model(&HostInfo{}).
		Where("id=?", h.Id).Updates(h).Error
}

// UpdateAuthentication 更新主机认证信息
func (h *HostInfo) UpdateAuthentication(id int, auth Authentication) error {
	return global.DBCli.Model(&HostInfo{}).Where("id = ?", id).Updates(auth).Error
}

// DeleteById 删除
func (h *HostInfo) DeleteById(ids []int) error {
	return global.DBCli.Where("id in ?", ids).Delete(&HostInfo{}).Error
}

// FindAll 查看所有主机数据
func (h *HostInfo) FindAll() ([]HostInfo, int64, error) {
	hosts := make([]HostInfo, 0)
	var total int64
	if err := global.DBCli.Model(&HostInfo{}).Count(&total).
		Find(&hosts).Error; err != nil {
		return hosts, total, err
	}

	return hosts, total, nil
}

// FindByIp 通过IP查看主机信息
func (h *HostInfo) FindByIp(ip string) (HostInfo, error) {
	var hostInfo HostInfo
	if err := global.DBCli.Model(&HostInfo{}).Where("ip=?", ip).
		First(&hostInfo).Error; err != nil {
		return hostInfo, err
	}

	return hostInfo, nil
}
