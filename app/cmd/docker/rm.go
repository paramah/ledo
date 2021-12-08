package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdDockerRm = cli.Command{
	Name:        "rm",
	Usage:       "remove docker-compose containers",
	Description: `Remove all containers defined in a docker-compose `,
	Action:      RunComposeRm,
}

func RunComposeRm(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerRm(ctx)
	return nil
}
