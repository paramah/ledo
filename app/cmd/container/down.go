package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
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
