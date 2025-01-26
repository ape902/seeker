package seeker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"github.com/ape902/seeker/pkg/global"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/encryptions"
	"github.com/ape902/seeker/pkg/tools/format"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/ape902/seeker/pkg/tools/grpc_cli"
	"github.com/ape902/seeker/pkg/tools/remote_host"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	hostInfoCliByGRPC = grpc_cli.GetGrpcClient[hostinfo_pb.HostInfoClient](grpc_cli.HostInfo, global.EngineGrpcServerAddr)
)

// SftpPut 远程文件Copy
// 前端传输的文件采用SFTP进行远程发送。
// cwd: 目标主机绝对路径
// IP: 目标主机
// file: 传输文件
func SftpPut(c *gin.Context) {
	// 参数验证
	cwd := c.PostForm("cwd")
	if cwd == "" {
		ginx.RESPCustomMsg(c, codex.InvalidParameter, "目标路径不能为空", nil)
		return
	}

	ip := c.PostForm("ip")
	if ip == "" {
		ginx.RESPCustomMsg(c, codex.InvalidParameter, "目标主机IP不能为空", nil)
		return
	}

	// 获取文件
	data, err := c.FormFile("file")
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.InvalidParameter, "文件上传失败", nil)
		return
	}

	// 检查文件大小（这里设置最大100MB）
	if data.Size > 100*1024*1024 {
		ginx.RESPCustomMsg(c, codex.InvalidParameter, "文件大小超过限制", nil)
		return
	}

	// 通过IP 从engine中获取该主机信息
	resp, err := hostInfoCliByGRPC.FindHostByIp(
		context.Background(), &hostinfo_pb.HostInfoIpRequest{
			Ip: ip,
		})
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.GRPCConnectionFailed, "获取主机信息失败", nil)
		return
	}

	//加密密码操作解析
	decryptPassword, err := encryptions.Base64AESCBCDecrypt(resp.Auth, []byte(global.ENCRYPTKEY))
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.Failure, "主机密码解析失败", nil)
		return
	}

	// 初始化主机SSH客户端
	sshCli, err := remote_host.NewSSHDial(
		fmt.Sprintf("%s:%d", resp.Ip, resp.Port),
		resp.Username, string(decryptPassword), int8(resp.AuthMode))
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.SSHConnectionFailed, "SSH连接失败", nil)
		return
	}
	defer sshCli.Client.Close()

	// 使用SSH客户端进行SFTP Client初始化
	ftpCli, err := sftp.NewClient(sshCli.Client)
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.ExecutionFailed, "SFTP客户端初始化失败", nil)
		return
	}
	defer ftpCli.Close()

	// 在远程主机创建空文件（绝对路径）
	remoteFile, err := ftpCli.Create(sftp.Join(cwd, data.Filename))
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.ExecutionFailed, "远程文件创建失败", nil)
		return
	}
	defer remoteFile.Close()

	// 读取前端传输的文件内容
	file, err := data.Open()
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.ExecutionFailed, "本地文件读取失败", nil)
		return
	}
	defer file.Close()

	// 复制文件内容
	written, err := io.Copy(remoteFile, file)
	if err != nil {
		logx.Error(err)
		ginx.RESPCustomMsg(c, codex.ExecutionFailed, "文件传输失败", nil)
		return
	}

	logx.Infof("文件传输完成，大小: %s", format.FileSize(written))
	ginx.RESPCustomMsg(c, codex.Success, "文件上传成功", nil)
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

	resp, err := hostInfoCliByGRPC.FindAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.GRPCConnectionFailed, nil)
		return
	}
	if resp.Code != codex.Success {
		logx.Error(resp.Error)
		ginx.RESP(c, int(resp.Code), nil)
		return
	}

	for i := 0; i < len(resp.Data); i++ {
		pc = append(pc, promsContent{
			Targets: []string{fmt.Sprintf("%s:%d", resp.Data[i].Ip, resp.Data[i].Port)},
			Labels:  resp.Data[i].Label,
		})
	}

	c.JSON(http.StatusOK, pc)
}

// HostInfoFindPage 主机信息分页查询
func HostInfoFindPage(c *gin.Context) {
	index := c.Query("index")
	indexToInt, err := strconv.Atoi(index)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}
	size := c.Query("size")
	sizeToInt, err := strconv.Atoi(size)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	resp, err := hostInfoCliByGRPC.FindPage(context.Background(), &hostinfo_pb.HostInfoPageInfo{
		Index: int32(indexToInt),
		Size:  int32(sizeToInt),
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.GRPCConnectionFailed, nil)
		return
	}

	if resp.Code != codex.Success {
		logx.Error(resp.Error)
		ginx.RESP(c, int(resp.Code), nil)
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
		resp, err := hostInfoCliByGRPC.Create(context.Background(), &hostinfo_pb.HostAndAuthentication{
			Ip:       hosts[i].IP,
			Port:     int32(hosts[i].Port),
			Os:       hosts[i].OS,
			Label:    hosts[i].LabelMap,
			Username: hosts[i].Username,
			AuthMode: int32(hosts[i].AuthMode),
			Auth:     hosts[i].Auth,
		})
		if err != nil {
			msg := fmt.Sprintf("%s 创建失败", hosts[i].IP)
			ginx.RESPCustomMsg(c, codex.GRPCConnectionFailed, msg, nil)
			return
		}

		if resp.Code != codex.Success {
			logx.Error(resp.Error)
			ginx.RESP(c, int(resp.Code), nil)
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

	resp, err := hostInfoCliByGRPC.Delete(context.Background(), &hostinfo_pb.HostInfoIdsRequest{
		Ids: ids.IDS,
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.GRPCConnectionFailed, nil)
		return
	}

	if resp.Code != codex.Success {
		logx.Error(resp.Error)
		ginx.RESP(c, int(resp.Code), nil)
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
		resp, err := hostInfoCliByGRPC.UpdateHost(context.Background(), &hostinfo_pb.Host{
			Id:    int32(hosts[i].Id),
			Ip:    hosts[i].IP,
			Port:  int32(hosts[i].Port),
			Os:    hosts[i].OS,
			Label: hosts[i].LabelMap,
		})
		if err != nil {
			logx.Error(err)
			ginx.RESP(c, codex.GRPCConnectionFailed, nil)
			return
		}
		if resp.Code != codex.Success {
			msg := fmt.Sprintf("%s 更新失败", hosts[i].IP)
			logx.Error(err)
			ginx.RESPCustomMsg(c, int(resp.Code), msg, nil)
			return
		}
	}

	ginx.RESP(c, codex.Success, nil)
}

func HostInfoUpdateAuthentication(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}
	var auth cmdb.Authentication
	if err := c.BindJSON(&auth); err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.InvalidParameter, nil)
		return
	}

	resp, err := hostInfoCliByGRPC.UpdateAuthentication(context.Background(), &hostinfo_pb.Authentication{
		Id:       int32(idInt),
		Username: auth.Username,
		AuthMode: int32(auth.AuthMode),
		Auth:     auth.Auth,
	})
	if err != nil {
		logx.Error(err)
		ginx.RESP(c, codex.GRPCConnectionFailed, nil)
		return
	}
	if resp.Code != codex.Success {
		logx.Error(err)
		ginx.RESP(c, int(resp.Code), nil)
		return
	}

	ginx.RESP(c, codex.Success, nil)
}
