package install

import (
	"github.com/spf13/cobra"
	"linux-tools/pkg/docker"
)

func NewCmdDocker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker",
		Short: "Docker install command.",
		Run: func(cmd *cobra.Command, args []string) {
			docker.Install(args)
		},
	}
	return cmd
}
