package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdDockerPs = cli.Command{
	Name:        "ps",
	Aliases:     []string{"p"},
	Usage:       "list running containers",
	Description: `List all containers defined in docker-compose use in current mode`,
	Action:      RunComposePs,
}

func RunComposePs(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerPs(ctx)
	return nil
}
