package system

import (
	"context"
	"errors"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/global"
	"google.golang.org/protobuf/types/known/emptypb"

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
	return "seeker_system_user"
}

// Create 新建数据
func (u *User) Create(ctx context.Context, user *user_center_pb.UserCenterUserInfo) (*emptypb.Empty, error) {
	u.Mobile = user.Mobile
	u.Password = user.Password
	u.NickName = user.NickName
	u.Rule = int(user.Rule)
	return nil, global.DBCli.Create(&u).Error
}

// Update 更新数据
func (u *User) Update(ctx context.Context, user *user_center_pb.UserCenterUserInfo) (*emptypb.Empty, error) {
	u.Id = int(user.ID)
	u.Mobile = user.Mobile
	u.Password = user.Password
	u.NickName = user.NickName
	u.Rule = int(user.Rule)
	return nil, global.DBCli.Where("id=?", u.Id).Updates(&u).Error
}

// DeleteByIds 以ID批量删除数据
func (u *User) DeleteByIds(ctx context.Context, ids *user_center_pb.UserCenterIDS) (*emptypb.Empty, error) {
	return nil, global.DBCli.Where("id in ?", ids.Ids).Delete(&User{}).Error
}

// FindPage 分页查找所有数据
func (u *User) FindPage(ctx context.Context, user *user_center_pb.UserCenterPageInfo) (*user_center_pb.UserCenterUserAll, error) {
	userPB := &user_center_pb.UserCenterUserAll{}
	if err := global.DBCli.Model(&User{}).Count(&userPB.Total).
		Offset((int(user.Index) - 1) * int(user.Size)).Limit(int(user.Size)).
		Find(&userPB.Data).Error; err != nil {
		return userPB, err
	}
	return userPB, nil
}

// FindByMobile 用mobile检查一条数据
func (u *User) FindByMobile(ctx context.Context, user *user_center_pb.UserCenterMobile) (*user_center_pb.UserCenterUserInfo, error) {
	userResp := &user_center_pb.UserCenterUserInfo{}
	if err := global.DBCli.Where(&User{Mobile: user.Mobile}).
		First(&userResp).Error; err != nil {
		return userResp, err
	}
	return userResp, nil
}

// IsExistByMobile 使用mobile检查数据是否存在
func (u *User) IsExistByMobile(ctx context.Context, mobile *user_center_pb.UserCenterMobile) (*user_center_pb.UserCenterIsExists, error) {
	pb := &user_center_pb.UserCenterIsExists{
		IsExist: false,
	}
	if mobile.Mobile == "" {
		return pb, errors.New("手机号不允许为空")
	}

	var count int64
	if err := global.DBCli.Model(&User{}).Where(&User{Mobile: mobile.Mobile}).Count(&count).Error; err != nil {
		return pb, err
	}
	pb.IsExist = count == 0
	return pb, nil
}
