package docker

// install docker and setting daemon

import (
	"github.com/sirupsen/logrus"
	"linux-tools/pkg/utils"
)

const (
	daemon = `sudo tee /etc/docker/daemon.json <<-'EOF'
{
    "registry-mirrors": [ 
      "https://hub.senhao.xyz",
      "https://docker.1ms.run"
    ],
    "log-opts": {
      "max-size": "100m",
      "max-file": "5"
    }
}
EOF`
)

func Install(args []string) {
	logrus.Infof("Install Docker.")

	utils.ExecCmd("curl -fsSL https://get.docker.com | sudo bash -s docker --mirror Aliyun && sudo systemctl enable --now docker")

	logrus.Infof("Start config image source.")
	utils.ExecCmd(daemon)

	utils.ExecCmd("sudo systemctl daemon-reload && sudo systemctl restart docker")

}
