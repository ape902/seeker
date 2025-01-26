package commands

import (
	"github.com/ape902/seeker/pkg/api/seeker/skctl"
	"github.com/spf13/cobra"
)

var (
	LoginStart = &cobra.Command{
		Use:               "login",
		Short:             "初始化认证信息配置文件",
		Long:              "生成认证信息，主要用于连接seeker引擎服务",
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return loginCommand()
		},
	}
	helpCmd = &cobra.Command{
		Use:   "help",
		Short: "帮助",
		Long:  "初始化认证配置文件帮助",
	}
)

var (
	mobile    string
	pass      string
	expiresAt int64
)

func init() {
	LoginStart.SetHelpCommand(helpCmd)

	LoginStart.Flags().StringVarP(&mobile, "mobile", "m", "", "手机号")
	LoginStart.Flags().StringVarP(&pass, "pass", "p", "", "密码")
	LoginStart.Flags().Int64VarP(&expiresAt, "ExpiresAt", "e", 0, "过期时间，默认30天")
}

func loginCommand() error {

	return skctl.ClientLogin(mobile, pass, expiresAt)
}
