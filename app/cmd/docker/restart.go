package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeRestart = cli.Command{
	Name:        "restart",
	Usage:       "Restart containers",
	Description: `Restart all containers defined in docker-compose`,
	Action:      RunComposeRestart,
}

func RunComposeRestart(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerRestart(ctx)
	return nil
}
