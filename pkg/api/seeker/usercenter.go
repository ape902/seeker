package seeker

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/models"
	"github.com/ape902/seeker/pkg/models/system"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	userCliByGrpc = grpc_cli.GetGrpcClient[user_center_pb.UserClient](grpc_cli.UserCenter, global.EngineGrpcServerAddr)
)

func UserCenterCreate(c *gin.Context) {
	var user system.User
	if err := c.BindJSON(&user); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	salt, encodedPwd := password.Encode(user.Password, global.PasswordOption)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	resp, err := userCliByGrpc.Create(context.Background(), &user_center_pb.UserCenterUserInfo{
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Rule:     int32(user.Rule),
		Labels:   user.LabelsMap,
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	if resp.Code != codex.Success {
		logx.Error(err)
		ginx.RESP(c, int(resp.Code), nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}

func UserCenterUpdate(c *gin.Context) {
	var users []system.User
	if err := c.BindJSON(&users); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Password != "" {
			salt, encodedPwd := password.Encode(users[i].Password, global.PasswordOption)
			users[i].Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
		}

		if _, err := userCliByGrpc.Update(context.Background(), &user_center_pb.UserCenterUserInfo{
			Id:       int32(users[i].Id),
			Mobile:   users[i].Mobile,
			Password: users[i].Password,
			NickName: users[i].NickName,
			Rule:     int32(users[i].Rule),
			Labels:   users[i].LabelsMap,
		}); err != nil {
			logx.Error(err)
			ginx.RESP(c, codex.ExecutionFailed, nil)
			return
		}
	}

	ginx.RESP(c, codex.Success, nil)
}

func UserCenterDeleteById(c *gin.Context) {
	var ids IDS
	if err := c.BindJSON(&ids); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	if _, err := userCliByGrpc.DeleteByIds(context.Background(), &user_center_pb.UserCenterIDS{
		Ids: ids.IDS,
	}); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}

func UserCenterFindPage(c *gin.Context) {
	index := c.Query("index")
	indexToInt, err := strconv.Atoi(index)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	size := c.Query("size")
	sizeToInt, err := strconv.Atoi(size)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	rest, err := userCliByGrpc.FindPage(context.Background(), &user_center_pb.UserCenterPageInfo{
		Index: int32(indexToInt),
		Size:  int32(sizeToInt),
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	if rest.Code != codex.Success {
		logx.Error(rest.Error)
		ginx.RESP(c, int(rest.Code), nil)
		return
	}

	ginx.RESP(c, codex.Success, ginx.Page(rest.Total, rest.Data))
}

func UserCenterFindById(c *gin.Context) {
	id := c.Query("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	resp, err := userCliByGrpc.FindById(context.Background(), &user_center_pb.UserCenterId{
		Id: int32(idToInt),
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, resp)
}

func Login(c *gin.Context) {
	var loginFrom system.PasswordLoginFrom
	if err := c.BindJSON(&loginFrom); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	resp, err := userCliByGrpc.FindByMobile(context.Background(), &user_center_pb.UserCenterMobile{
		Mobile: loginFrom.Mobile,
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	passwordInfo := strings.Split(resp.Password, "$")
	if !password.Verify(loginFrom.Password, passwordInfo[2], passwordInfo[3], global.PasswordOption) {
		ginx.RESP(c, codex.UserOrPassError, nil)
		return
	}

	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID:       uint(resp.Id),
		NickName: resp.NickName,
		Labels:   format.MapToString(resp.Labels),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
			Issuer:    "seeker",                        // 是哪个机构做的签名
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.TokenCreateFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, gin.H{
		"id":        resp.Id,
		"nick_name": resp.NickName,
		"token":     token,
	})
}
