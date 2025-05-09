package main

import (
	"linux-tools/pkg/cmd"
	"os"
	"os/exec"
)

func main() {
	c := cmd.NewToolCommand()
	_ = exec.Command("/bin/bash", "-c", "ulimit -u 65535").Run()
	_ = exec.Command("/bin/bash", "-c", "ulimit -n 65535").Run()
	// Execute adds all child commands to the root command and sets flags appropriately.
	// This is called by main.main(). It only needs to happen once to the rootCmd.
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
