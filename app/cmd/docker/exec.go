package docker

import (
	"github.com/urfave/cli/v2"
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
)

var CmdComposeExec = cli.Command{
	Name:        "exec",
	Aliases:     []string{"r"},
	Usage:       "exec cmd in a main running container",
	Description: `Execute command in a main running container`,
	ArgsUsage:   "[<cmd>]",
	Action:      RunComposeExec,
}

func RunComposeExec(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	if cmd.Args().Len() >= 1 {
		compose.ExecComposerExec(ctx, cmd.Args())
		return nil
	}
	return nil
}
