package system

import (
	"errors"

	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
)

type (
	User struct {
		models.BaseModel
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		NickName string `json:"nick_name"`
		Rule     int    `json:"rule"`
	}

	PasswordLoginFrom struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
)

func (User) TableName() string {
	return "system_user"
}

// UserFindByMobile 通过手机号查询一条数据
func UserFindByMobile(mobile string) (User, error) {
	var user User
	if err := global.DBCli.Where(&User{Mobile: mobile}).First(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

// NotExistByMobile 通过手机号检查数据是否存在
func UserNotExistByMobile(mobile string) (bool, error) {
	if mobile == "" {
		return false, errors.New("手机号不允许为空")
	}

	var count int64
	if err := global.DBCli.Model(&User{}).Where(&User{Mobile: mobile}).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}

// UserFindAllByPage 分页查询所有数据
func UserFindAllByPage(index, size int) ([]User, int64, error) {
	var users []User
	var total int64
	if err := global.DBCli.Model(&User{}).Count(&total).
		Offset((index - 1) * size).Limit(size).
		Find(&users).Error; err != nil {
		return users, total, err
	}

	return users, total, nil
}

// Create 创建用户
func (u *User) Create() error {
	return global.DBCli.Create(&u).Error
}

// Update 更新用户
func (u *User) Update() error {
	return global.DBCli.Where("id=?", u.Id).Updates(&u).Error
}

// DeleteById 使用ID删除数据
func DeleteById(ids []int) error {
	return global.DBCli.Where("id in ?", ids).Delete(&User{}).Error
}
