package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
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
