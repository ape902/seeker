package handler

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/ape902/corex/logx"
	"github.com/ape902/seeker/pkg/contoller/pb/agent_pb"
	"github.com/ape902/seeker/pkg/tools/systemx/portx"
	"github.com/prometheus/procfs"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	RemoteHostControllerPB struct {
		agent_pb.UnimplementedAgentServer
	}
)

func (r *RemoteHostControllerPB) FindProcInfo(ctx context.Context, empty *emptypb.Empty) (*agent_pb.RespProcInfo, error) {
	pb := &agent_pb.RespProcInfo{}

	ports, err := portx.NewPorts()
	if err != nil {
		pb.Error = err.Error()
		return pb, err
	}

	procs, err := procfs.AllProcs()
	if err != nil {
		pb.Error = err.Error()
		return pb, err
	}

	for i := 0; i < len(procs); i++ {
		procInfo := &agent_pb.ProcInfo{}
		//如果是本服务进程直接跳过
		if procs[i].PID == os.Getpid() {
			continue
		}

		//获取proc中所有pid的Status信息
		procStatus, err := procs[i].NewStatus()
		if err != nil {
			pb.Error = err.Error()
			return pb, err
		}
		//进程名字
		procInfo.Name = procStatus.Name

		//有可能存在多个地址监听端口
		if ar, ok := ports[procs[i].PID]; !ok {
			continue
		} else {
			for i := 0; i < len(ar); i++ {
				ipAndPort := strings.Split(ar[i], ";;")
				listen := &agent_pb.ListenInfo{}

				listen.Ip = ipAndPort[0]
				portInt, _ := strconv.Atoi(ipAndPort[1])
				listen.Port = int32(portInt)

				procInfo.Listen = append(procInfo.Listen, listen)
			}
		}

		//获取进程的启动命令
		cmdLine, err := procs[i].CmdLine()
		if err != nil {
			logx.Error(err)
			continue
		}
		procInfo.CmdLine = strings.Join(cmdLine, " ")

		// 获取proc中所有pid的Stat信息
		stat, err := procs[i].Stat()
		if err != nil {
			pb.Error = err.Error()
			return pb, err
		}
		procInfo.Comm = stat.Comm

		//CPU执行时间
		const userHZ = 100
		procInfo.CpuUserTime = int64(stat.UTime / userHZ)
		procInfo.CpuSystemTime = int64(stat.STime / userHZ)

		// 进程内存信息
		procInfo.ResidentMemoryBytes = int64(stat.ResidentMemory())
		procInfo.VirtualMemoryBytes = int64(stat.VirtualMemory())

		pb.Data = append(pb.Data, procInfo)
	}

	return pb, nil
}

func (r *RemoteHostControllerPB) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*agent_pb.HealthStatus, error) {
	response := &agent_pb.HealthStatus{
		IsAlive: true,
	}

	// 获取进程信息
	proc, err := procfs.Self()
	if err != nil {
		response.Error = err.Error()
		return response, err
	}

	// 获取启动时间
	stat, err := proc.Stat()
	if err != nil {
		response.Error = err.Error()
		return response, err
	}

	// 计算运行时间（秒）
	startTime, err := stat.StartTime()
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.Uptime = int64(startTime)

	// 获取CPU使用率
	cpuStat, err := proc.Stat()
	if err != nil {
		response.Error = err.Error()
		return response, err
	}
	response.CpuUsage = float64((cpuStat.UTime + cpuStat.STime) / 100)

	// 获取内存使用量
	response.MemoryUsage = int64(stat.ResidentMemory())

	// 设置版本信息
	response.Version = "1.0.0" // 这里可以根据实际情况设置版本号

	return response, nil
}

func (r *RemoteHostControllerPB) AgentComm(ctx context.Context, in *agent_pb.Info) (*agent_pb.Response, error) {
	cmd := exec.CommandContext(ctx, "sh", "-c", in.Command)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	resp := &agent_pb.Response{}
	resp.Error = stderr.Bytes()
	resp.Data = stdout.Bytes()
	resp.Msg = "成功"

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus := exitError.Sys().(syscall.WaitStatus)
			switch waitStatus.ExitStatus() {
			case 1:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "常规错误"
			case 2:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "误用shell命令"
			case 126:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "命令不可执行"
			case 127:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "命令未找到"
			case 130:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "命令通过SIGINT被终止"
			default:
				resp.Error = stderr.Bytes()
				resp.Data = stdout.Bytes()
				resp.Msg = "错误"
			}
		}
	}

	return resp, nil
}
