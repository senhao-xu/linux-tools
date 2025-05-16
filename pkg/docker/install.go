package docker

import (
	"github.com/sirupsen/logrus"
	"linux-tools/pkg/utils"
)

func Install(args []string) {
	logrus.Infof("Install Docker.")

	utils.ExecCmd("curl -fsSL https://get.docker.com | sudo bash -s docker --mirror Aliyun && sudo systemctl enable --now docker")
}
