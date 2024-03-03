package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeExec = cli.Command{
	Name:        "exec",
	Aliases:     []string{"e"},
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
