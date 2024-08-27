package handler

import (
	"bytes"
	"context"
	"github.com/ape902/seeker/pkg/contoller/pb/command_pb"
	"os/exec"
	"syscall"
)

type (
	Commands struct {
		command_pb.UnsafeCommandServer
	}
)

func (c *Commands) Command(ctx context.Context, in *command_pb.Info) (*command_pb.Response, error) {
	cmd := exec.CommandContext(ctx, "sh", "-c", in.Command)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	resp := &command_pb.Response{}
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
