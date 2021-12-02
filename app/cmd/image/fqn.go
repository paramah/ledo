package image

import (
	"fmt"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
)

var CmdDockerFqn = cli.Command{
	Name:        "fqn",
	Aliases:     []string{"f"},
	Usage:       "docker image fqn",
	Description: `Get fqn docker image defined as main service in config file`,
	Action:      RunDockerFqn,
}

func RunDockerFqn(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	fmt.Printf("%v", docker.ShowDockerImageFQN(ctx))
	return nil
}
