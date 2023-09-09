package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeStop = cli.Command{
	Name:        "stop",
	Usage:       "stop containers",
	Description: `Stop all containers defined in container-compose stack mode`,
	Action:      RunComposeStop,
}

func RunComposeStop(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerStop(ctx)
	return nil
}
