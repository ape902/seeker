package seeker

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
	"github.com/ape902/seeker/pkg/models/system"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	options = &password.Options{16, 100, 32, sha512.New}
)

func CreateUser(c *gin.Context) {
	var user system.User
	if err := c.BindJSON(&user); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	var total int64
	if err := global.DBCli.Model(&system.User{}).Where(&system.User{Mobile: user.Mobile}).Count(&total).Error; err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	if total > 0 {
		ginx.RESP(c, codex.AlreadyExists, nil)
		return
	}

	salt, encodedPwd := password.Encode(user.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	if err := global.DBCli.Create(&user).Error; err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}

func Login(c *gin.Context) {
	var loginFrom system.PasswordLoginFrom
	if err := c.BindJSON(&loginFrom); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	var user system.User
	if err := global.DBCli.Where(&system.User{Mobile: loginFrom.Mobile}).First(&user).Error; err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	passwordInfo := strings.Split(user.Password, "$")
	if !password.Verify(loginFrom.Password, passwordInfo[2], passwordInfo[3], options) {
		ginx.RESP(c, codex.UserOrPassError, nil)
		return
	}

	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID:       uint(user.Id),
		NickName: user.NickName,
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
		"id":        user.Id,
		"nick_name": user.NickName,
		"token":     token,
	})
}
