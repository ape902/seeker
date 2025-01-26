package skctl

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/system_pb/user_center_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"github.com/dgrijalva/jwt-go"
)

var (
	userCliByGrpc = grpc_cli.GetGrpcClient[user_center_pb.UserClient](grpc_cli.UserCenter, global.EngineGrpcServerAddr)
)

// JWTVerify 做为JWT验证
func JWTVerify() error {
	ctx, err := os.ReadFile(global.ClientAuthFilePath)
	if os.IsNotExist(err) {
		logx.Error(err)
		return errors.New(fmt.Sprintf("未找到[%s]配置文件，请重新登录生成.", global.ClientAuthFilePath))
	}

	j := middleware.NewJWT()
	claims, err := j.ParseToken(string(ctx))
	if err != nil {
		if err == middleware.TokenExpired {
			return errors.New("授权已过期,请重新登录")
		}

		logx.Error(err)
		return err
	}

	fmt.Println(claims)
	return nil
}

// ClientLogin 客户端登录
func ClientLogin(mobile, pass string, expiresAt int64) error {
	// 参数验证
	if mobile == "" || pass == "" {
		return errors.New(codex.CodeText(codex.InvalidParameter))
	}

	// 设置GRPC调用超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 查询用户信息
	resp, err := userCliByGrpc.FindByMobile(ctx, &user_center_pb.UserCenterMobile{
		Mobile: mobile,
	})

	if err != nil {
		logx.Errorf("查询用户信息失败: %v", err)
		return fmt.Errorf("查询用户信息失败: %v", err)
	}

	// 密码验证
	passwordInfo := strings.Split(resp.Password, "$")
	if len(passwordInfo) != 4 {
		logx.Error("密码格式错误")
		return errors.New(codex.CodeText(codex.UserOrPassError))
	}

	if !password.Verify(pass, passwordInfo[2], passwordInfo[3], global.PasswordOption) {
		logx.Error("密码验证失败")
		return errors.New(codex.CodeText(codex.UserOrPassError))
	}

	// 设置Token过期时间
	if expiresAt == 0 {
		expiresAt = 30
	}

	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID:       uint(resp.Id),
		NickName: resp.NickName,
		Labels:   format.MapToString(resp.Labels),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                      // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*expiresAt, // 30天过期
			Issuer:    "seeker",                               // 是哪个机构做的签名
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return errors.New(codex.CodeText(codex.TokenCreateFailed))
	}

	// 确保目录存在
	dir := filepath.Dir(global.ClientAuthFilePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 写入token文件，使用0600权限以提高安全性
	return ioutil.WriteFile(global.ClientAuthFilePath, []byte(token), 0600)
}
