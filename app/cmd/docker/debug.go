package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeDebug = cli.Command{
	Name:        "debug",
	Usage:       "debug main container",
	Description: `Run shell on main container without entrypoint execute`,
	Action:      RunComposeDebug,
}

func RunComposeDebug(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerDebug(ctx)
	return nil
}
