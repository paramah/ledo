package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeDown = cli.Command{
	Name:        "down",
	Usage:       "down all containers",
	Description: `Down all containers defined in docker-compose`,
	Action:      RunComposeDown,
}

func RunComposeDown(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerDown(ctx)
	return nil
}
