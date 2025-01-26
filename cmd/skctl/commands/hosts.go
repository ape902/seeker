package commands

import (
	"github.com/ape902/seeker/pkg/api/seeker/skctl"
	"github.com/spf13/cobra"
)

var (
	HostStart = &cobra.Command{
		Use:               "host",
		Short:             "主机操作控制",
		Long:              "主机的同步、基础操作等",
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return hostCommand()
		},
	}
)

func hostCommand() error {
	return skctl.GetHostInfo()
}
