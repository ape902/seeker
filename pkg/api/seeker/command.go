package seeker

import (
	"context"
	"fmt"
	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/ape902/seeker/pkg/tools/remote_host"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"google.golang.org/grpc"
	"io"
	"net/http"
)

func Discovery(c *gin.Context) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "192.168.119.82", 58899), grpc.WithInsecure())
	if err != nil {
		logx.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}

	cc := command_pb.NewCommandClient(conn)

	procs, err := cc.FindProcInfo(context.Background(), nil)
	if err != nil {
		logx.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": procs,
	})

}

type commandInfo struct {
	Commands string `json:"commands"`
}

func RunCommand(c *gin.Context) {
	var cmd commandInfo
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "无效参数",
		})
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "192.168.119.80", 58899),
		grpc.WithInsecure())
	if err != nil {
		logx.Error(err)
		return
	}

	cc := command_pb.NewCommandClient(conn)

	resp, err := cc.Command(context.Background(), &command_pb.Info{
		Command: cmd.Commands,
	})

	if err != nil {
		logx.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": string(resp.Data),
		"msg":  resp.Msg,
		"err":  string(resp.Error),
	})

}

// SftpPut 远程文件Copy
func SftpPut(c *gin.Context) {
	cwd := c.PostForm("cwd")
	ip := c.PostForm("ip")
	hostInfo, err := connHostInfoGrpc().FindHostByIp(
		context.Background(), &hostinfo_pb.HostInfoIpRequest{
			Ip: ip,
		})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	decryptPassword, err := encryptions.Base64AESCBCDecrypt(hostInfo.Auth, []byte(global.ENCRYPTKEY))
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.Failure, "主机密码解析失败", nil)
		return
	}
	sshCli, err := remote_host.NewSSHDial(
		fmt.Sprintf("%s:%d", hostInfo.Ip, hostInfo.Port),
		hostInfo.Username, string(decryptPassword), int8(hostInfo.AuthMode))
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	ftpCli, err := sftp.NewClient(sshCli.Client)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	data, err := c.FormFile("file")
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}
	remoteFile, err := ftpCli.Create(sftp.Join(cwd, data.Filename))
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	defer remoteFile.Close()

	file, err := data.Open()
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	defer file.Close()

	written, err := io.Copy(remoteFile, file)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	logx.Infof("Size %s", format.FileSize(written))

	ginx.RESP(c, codex.Success, nil)
}
