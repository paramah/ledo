package docker

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeRun = cli.Command{
	Name:        "run",
	Aliases:     []string{"r"},
	Usage:       "run cmd in main container",
	Description: `Run command in main container`,
	ArgsUsage:   "[<cmd>]",
	Action:      RunComposeRun,
}

func RunComposeRun(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	if cmd.Args().Len() >= 1 {
		compose.ExecComposerRun(ctx, cmd.Args())
		return nil
	}
	return nil
}
