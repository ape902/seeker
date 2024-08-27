package seeker

import (
	"context"
	"fmt"
	"github.com/ape902/seeker/pkg/contoller/pb/hostinfo_pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"strings"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/models/cmdb"
	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/ape902/seeker/pkg/tools/ginx"
	"github.com/gin-gonic/gin"
)

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
