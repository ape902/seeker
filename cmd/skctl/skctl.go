package main

import (
	"fmt"
	"github.com/ape902/seeker/cmd/skctl/commands"
	"github.com/ape902/seeker/pkg/initialize"
	"github.com/spf13/cobra"
)

var (
	skctlStart = &cobra.Command{
		Short: "seeker tools",
		Long:  `Seeker运维管理工具`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	helmCmd = &cobra.Command{
		Use:   "help",
		Short: "帮助信息",
		Long:  "查看具体使用方法",
	}
)

func Execute() {
	skctlStart.SetHelpCommand(helmCmd)
	//用户登录认证操作
	skctlStart.AddCommand(commands.LoginStart)
	skctlStart.AddCommand(commands.HostStart)

	// 初始化日志
	initialize.InitLogger()

	//if err := skctl.JWTVerify(); err != nil {
	//	logx.Error(err)
	//	return
	//}

	if err := skctlStart.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	Execute()
}
