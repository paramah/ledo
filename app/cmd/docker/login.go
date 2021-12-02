package docker

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
)

var CmdDockerLogin = cli.Command{
	Name:        "login",
	Aliases:     []string{"l"},
	Usage:       "Docker Registry login",
	Description: `Login to docker registry`,
	Subcommands: []*cli.Command{
		&CmdDockerEcrLogin,
	},
}

var CmdDockerEcrLogin = cli.Command{
	Name:        "ecr",
	Aliases:     []string{"e"},
	Usage:       "AWS Elastic Docker Registry",
	Description: `Login to docker registry`,
	Action:      RunDockerEcrLogin,
}

func RunDockerEcrLogin(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	docker.DockerEcrLogin(ctx)
	return nil
}

