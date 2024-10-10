package seeker

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"strconv"
	"strings"
	"time"

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
	options = &password.Options{16, 100, 32, sha512.New}
)

func UserCenterCreate(c *gin.Context) {
	var user system.User
	if err := c.BindJSON(&user); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	salt, encodedPwd := password.Encode(user.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	resp, err := connUserCenterGrpc().Create(context.Background(), &user_center_pb.UserCenterUserInfo{
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Rule:     int32(user.Rule),
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
		if _, err := connUserCenterGrpc().Update(context.Background(), &user_center_pb.UserCenterUserInfo{
			Id:       int32(users[i].Id),
			Mobile:   users[i].Mobile,
			Password: users[i].Password,
			NickName: users[i].NickName,
			Rule:     int32(users[i].Rule),
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

	if _, err := connUserCenterGrpc().DeleteByIds(context.Background(), &user_center_pb.UserCenterIDS{
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

	rest, err := connUserCenterGrpc().FindPage(context.Background(), &user_center_pb.UserCenterPageInfo{
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

func Login(c *gin.Context) {
	var loginFrom system.PasswordLoginFrom
	if err := c.BindJSON(&loginFrom); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	rest, err := connUserCenterGrpc().FindByMobile(context.Background(), &user_center_pb.UserCenterMobile{
		Mobile: loginFrom.Mobile,
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	passwordInfo := strings.Split(rest.Password, "$")
	if !password.Verify(loginFrom.Password, passwordInfo[2], passwordInfo[3], options) {
		ginx.RESP(c, codex.UserOrPassError, nil)
		return
	}

	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID:       uint(rest.Id),
		NickName: rest.NickName,
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
		"id":        rest.Id,
		"nick_name": rest.NickName,
		"token":     token,
	})
}
