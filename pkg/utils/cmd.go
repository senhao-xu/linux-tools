package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
)

func ExecCmd(command string) error {
	sysType := runtime.GOOS
	var cmd *exec.Cmd
	if "linux" == sysType {
		// LINUX系统
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	if "windows" == sysType {
		// windows系统
		log.Error("Non Linux environment, command exit")
		os.Exit(1)
		cmd = exec.Command("cmd", "/C", command)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info("command: " + command)
	}
	return nil
}
