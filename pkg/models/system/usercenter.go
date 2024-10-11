package system

import (
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
)

type (
	User struct {
		models.BaseModel
		Mobile    string            `json:"mobile"`
		Password  string            `json:"password"`
		NickName  string            `json:"nick_name"`
		Rule      int               `json:"rule"`
		Labels    string            `json:"-"`
		LabelsMap map[string]string `json:"labels" gorm:"-"`
	}

	PasswordLoginFrom struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
)

func (User) TableName() string {
	return "seeker_system_user"
}

// Create 新建数据
func (u *User) Create() error {
	return global.DBCli.Create(&u).Error
}

// Update 更新数据
func (u *User) UpdateById() error {
	return global.DBCli.Where("id=?", u.Id).Updates(&u).Error
}

// DeleteByIds 以ID批量删除数据
func (u *User) DeleteByIds(ids []int) error {
	return global.DBCli.Where("id in ?", ids).Delete(&User{}).Error
}

// FindPage 分页查找所有数据
func (u *User) FindPage(index, size int) ([]User, int64, error) {
	users := make([]User, 0)
	var total int64
	if err := global.DBCli.Model(&User{}).Count(&total).
		Offset((index - 1) * size).Limit(size).Find(&users).Error; err != nil {
		return users, total, err
	}

	return users, total, nil
}

// FindByMobile 用mobile检查一条数据
func (u *User) FindByMobile(mobile string) (User, error) {
	var user User
	if err := global.DBCli.Where(&User{Mobile: mobile}).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// ExistByMobile 使用mobile检查数据是否存在
func (u *User) ExistByMobile(mobile string) (bool, error) {
	var total int64
	if err := global.DBCli.Model(&User{}).
		Where("mobile=?", mobile).Count(&total).Error; err != nil {
		return false, err
	}

	return total != 0, nil
}
