package system

import (
	"context"
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/format"
)

type (
	UserCenterPB struct {
		//继承GRPC的服务端扩展
		user_center_pb.UnimplementedUserServer

		//继承用户结构体，主要进行数据库操作方法调用
		user User
	}
)

func (u *UserCenterPB) Create(ctx context.Context, user *user_center_pb.UserCenterUserInfo) (*user_center_pb.UserCenterDefResp, error) {
	pb := &user_center_pb.UserCenterDefResp{}
	pb.Code = codex.Success

	exist, err := u.user.ExistByMobile(user.Mobile)
	if err != nil {
		logx.Error(err)
		pb.Code = codex.DatabaseExecutionFailed
		pb.Error = err.Error()
		return pb, err
	}

	if exist {
		pb.Code = codex.AlreadyExists
		pb.Error = fmt.Sprintf("%s Exist", user.Mobile)
		return pb, nil
	}

	u.user.Mobile = user.Mobile
	u.user.Password = user.Password
	u.user.NickName = user.NickName
	u.user.Rule = int(user.Rule)
	u.user.Labels = format.MapToString(user.Labels)

	if err := u.user.Create(); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

func (u *UserCenterPB) Update(ctx context.Context, user *user_center_pb.UserCenterUserInfo) (*user_center_pb.UserCenterDefResp, error) {
	pb := &user_center_pb.UserCenterDefResp{}
	pb.Code = codex.Success

	u.user.Id = int(user.Id)
	u.user.Mobile = user.Mobile
	u.user.Password = user.Password
	u.user.NickName = user.NickName
	u.user.Rule = int(user.Rule)
	u.user.Labels = format.MapToString(user.Labels)

	if err := u.user.UpdateById(); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}
func (u *UserCenterPB) DeleteByIds(ctx context.Context, ids *user_center_pb.UserCenterIDS) (*user_center_pb.UserCenterDefResp, error) {
	pb := &user_center_pb.UserCenterDefResp{}
	pb.Code = codex.Success

	idInt := format.Int32ToIntArray(ids.Ids)
	if err := u.user.DeleteByIds(idInt); err != nil {
		logx.Error(err)
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, err
	}

	return pb, nil
}

func (u *UserCenterPB) FindPage(ctx context.Context, page *user_center_pb.UserCenterPageInfo) (*user_center_pb.UserCenterUserAll, error) {
	pb := &user_center_pb.UserCenterUserAll{}
	pb.Code = codex.Success

	data, total, err := u.user.FindPage(int(page.Index), int(page.Size))
	if err != nil {
		pb.Error = err.Error()
		pb.Code = codex.DatabaseExecutionFailed
		return pb, nil
	}

	pb.Total = total
	for i := 0; i < len(data); i++ {
		pb.Data = append(pb.Data, &user_center_pb.UserCenterUserResp{
			Id:       int32(data[i].Id),
			Mobile:   data[i].Mobile,
			NickName: data[i].NickName,
			Rule:     int32(data[i].Rule),
			Labels:   format.StringToMap(data[i].Labels),
		})
	}

	return pb, nil
}

func (u *UserCenterPB) FindByMobile(ctx context.Context, mobile *user_center_pb.UserCenterMobile) (*user_center_pb.UserCenterUserInfo, error) {
	pb := &user_center_pb.UserCenterUserInfo{}

	user, err := u.user.FindByMobile(mobile.Mobile)
	if err != nil {
		logx.Error(err)
		return pb, err
	}

	pb.Id = int32(user.Id)
	pb.Mobile = user.Mobile
	pb.Password = user.Password
	pb.NickName = user.NickName
	pb.Rule = int32(user.Rule)
	pb.Labels = format.StringToMap(user.Labels)

	return pb, nil
}
