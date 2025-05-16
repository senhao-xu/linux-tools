package install

import (
	"github.com/spf13/cobra"
	open_ssh "linux-tools/pkg/open-ssh"
)

func NewCmdOpenssh() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "Openssh install command.",
		Run: func(cmd *cobra.Command, args []string) {
			open_ssh.Install()
		},
	}
	return cmd
}
