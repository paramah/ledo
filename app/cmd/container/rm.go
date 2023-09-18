package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
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
