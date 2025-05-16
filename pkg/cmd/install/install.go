package install

import "github.com/spf13/cobra"

func NewCmdInstall() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install command.",
	}

	cmd.AddCommand(NewCmdDocker())
	cmd.AddCommand(NewCmdOpenssh())
	return cmd
}
