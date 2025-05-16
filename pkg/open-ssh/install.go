package open_ssh

import (
	"github.com/sirupsen/logrus"
	"linux-tools/pkg/utils"
)

func Install() {
	logrus.Infof("Install Open-ssh.")

	utils.ExecCmd("sudo apt update && sudo apt upgrade")
	utils.ExecCmd("sudo apt install openssh-server")
	utils.ExecCmd("sudo systemctl enable ssh")

}
