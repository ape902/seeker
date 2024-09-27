package seeker

import (
	"context"
	"fmt"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/remote_host"
	"github.com/pkg/sftp"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/gin-gonic/gin"
)

// SftpPut 远程文件Copy
// 前端传输的文件采用SFTP进行远程发送。
// cwd: 目标主机绝对路径
// IP: 目标主机
// file: 传输文件
func SftpPut(c *gin.Context) {
	cwd := c.PostForm("cwd")
	ip := c.PostForm("ip")
	// 通过IP 从engine中获取该主机信息
	hostInfo, err := connHostInfoGrpc().FindHostByIp(
		context.Background(), &hostinfo_pb.HostInfoIpRequest{
			Ip: ip,
		})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	//加密密码操作解析
	decryptPassword, err := encryptions.Base64AESCBCDecrypt(hostInfo.Auth, []byte(global.ENCRYPTKEY))
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.Failure, "主机密码解析失败", nil)
		return
	}

	// 初始化主机SSH客户端
	sshCli, err := remote_host.NewSSHDial(
		fmt.Sprintf("%s:%d", hostInfo.Ip, hostInfo.Port),
		hostInfo.Username, string(decryptPassword), int8(hostInfo.AuthMode))
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	// 使用SSH客户端进行SFTP Client初始化
	ftpCli, err := sftp.NewClient(sshCli.Client)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	// 获取前端传输文件
	data, err := c.FormFile("file")
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	// 在远程主机创建空文件（绝对路径）
	remoteFile, err := ftpCli.Create(sftp.Join(cwd, data.Filename))
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	defer remoteFile.Close()

	// 读取前端传输的文件内容
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

type (
	promsContent struct {
		Targets []string          `json:"targets"`
		Labels  map[string]string `json:"labels"`
	}
)

// HttpSDConfig 采用Prometheus的 HttpSDConfig模块
// 由Prometheus自动请求，从该接口获取数据源
func HttpSDConfig(c *gin.Context) {
	pc := make([]promsContent, 0)

	data, err := connHostInfoGrpc().FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	for i := 0; i < len(data.Data); i++ {
		pc = append(pc, promsContent{
			Targets: []string{fmt.Sprintf("%s:%d", data.Data[i].Ip, data.Data[i].Port)},
			Labels:  stringToMap(data.Data[i].Label),
		})
	}

	c.JSON(http.StatusOK, pc)
}

// stringToMap 字符串转map
// 例：aa=aa,bb=bb转换成{"aa":"aa", "bb":"bb"}
func stringToMap(str string) map[string]string {
	m := make(map[string]string)
	if str == "" {
		return m
	}

	strS := strings.Split(str, ",")
	for i := 0; i < len(strS); i++ {
		kv := strings.Split(strS[i], "=")
		if _, ok := m[kv[0]]; ok {
			continue
		}

		m[kv[0]] = kv[1]
	}

	return m
}

// HostInfoFindPage 主机信息分页查询
func HostInfoFindPage(c *gin.Context) {
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

	resp, err := connHostInfoGrpc().FindPage(context.Background(), &hostinfo_pb.HostInfoPageInfo{
		Index: int32(indexToInt),
		Size:  int32(sizeToInt),
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}
	ginx.RESP(c, codex.Success, ginx.Page(resp.Total, resp.Data))
}

func HostInfoCreate(c *gin.Context) {
	hosts := make([]cmdb.HostInfo, 0)
	if err := c.BindJSON(&hosts); err != nil {
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	for i := 0; i < len(hosts); i++ {
		exist, err := connHostInfoGrpc().IsExistByIp(context.Background(), &hostinfo_pb.HostInfoIpRequest{
			Ip: hosts[i].IP,
		})
		if err != nil {
			ginx.RESP(c, codex.ExecutionFailed, nil)
			return
		}

		if exist.IsExist {
			ginx.RESP(c, codex.AlreadyExists, nil)
			return
		}

		if _, err := connHostInfoGrpc().Create(context.Background(), &hostinfo_pb.HostAndAuthentication{
			Ip:       hosts[i].IP,
			Port:     int32(hosts[i].Port),
			OS:       hosts[i].OS,
			Label:    hosts[i].Label,
			Username: hosts[i].Username,
			AuthMode: int32(hosts[i].AuthMode),
			Auth:     hosts[i].Auth,
		}); err != nil {
			msg := fmt.Sprintf("%s 创建失败", hosts[i].IP)
			ginx.RESPCustomMsg(c, codex.ExecutionFailed, msg, nil)
			return
		}
	}

	ginx.RESP(c, codex.Success, nil)
}

func HostInfoDelete(c *gin.Context) {
	var ids IDS
	if err := c.BindJSON(&ids); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	if _, err := connHostInfoGrpc().Delete(context.Background(), &hostinfo_pb.HostInfoIdsRequest{
		Ids: ids.IDS,
	}); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}

func HostInfoUpdateHost(c *gin.Context) {
	var hosts []cmdb.HostInfo
	if err := c.BindJSON(&hosts); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	for i := 0; i < len(hosts); i++ {
		if _, err := connHostInfoGrpc().UpdateHost(context.Background(), &hostinfo_pb.Host{
			Id:    int32(hosts[i].Id),
			Ip:    hosts[i].IP,
			Port:  int32(hosts[i].Port),
			OS:    hosts[i].OS,
			Label: hosts[i].Label,
		}); err != nil {
			msg := fmt.Sprintf("%s 更新失败", hosts[i].IP)
			logx.Error(err)
			ginx.RESPCustomMsg(c, codex.ExecutionFailed, msg, nil)
			return
		}
	}

	ginx.RESP(c, codex.Success, nil)
}

func HostInfoUpdateAuthentication(c *gin.Context) {
	id := c.Query("id")
	var auth cmdb.Authentication
	if err := c.BindJSON(&auth); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	if _, err := connHostInfoGrpc().UpdateAuthentication(context.Background(), &hostinfo_pb.Authentication{
		ID:       id,
		Username: auth.Username,
		AuthMode: int32(auth.AuthMode),
		Auth:     auth.Auth,
	}); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.ExecutionFailed, nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}
