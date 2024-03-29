package image

import (
	"fmt"
	"github.com/paramah/ledo/app/modules/container"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdDockerFqn = cli.Command{
	Name:        "fqn",
	Aliases:     []string{"f"},
	Usage:       "container image fqn",
	Description: `Get fqn container image defined as main service in config file`,
	Action:      RunDockerFqn,
}

func RunDockerFqn(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	fmt.Printf("%v", container.ShowImageFQN(ctx))
	return nil
}
