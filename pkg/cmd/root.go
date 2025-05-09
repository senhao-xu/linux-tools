package cmd

import (
	"github.com/spf13/cobra"
	"linux-tools/pkg/cmd/install"
)

func NewToolCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "xsh",
		Short: "linux Aggregation tool",
		Long:  `Install commonly used software and configuration under Linux.`,
	}
	cmds.AddCommand(install.NewCmdInstall())
	return cmds
}
