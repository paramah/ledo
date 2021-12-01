package image

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"ledo/app/modules/docker"
	"ledo/app/modules/context"
)

var CmdDockerFqn = cli.Command{
	Name:        "fqn",
	Aliases:     []string{"f"},
	Usage:       "Docker image fqn",
	Description: `Get fqn docker image defined as main service in config file`,
	Action:      RunDockerFqn,
}

func RunDockerFqn(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	fmt.Printf("%v", docker.ShowDockerImageFQN(ctx))
	return nil
}
