package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeStop = cli.Command{
	Name:        "stop",
	Usage:       "stop containers",
	Description: `Stop all containers defined in docker-compose stack mode`,
	Action:      RunComposeStop,
}

func RunComposeStop(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerStop(ctx)
	return nil
}
