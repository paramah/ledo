package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeShell = cli.Command{
	Name:        "shell",
	Aliases:     []string{"sh"},
	Usage:       "run shell from main service",
	Description: `Execute shell cmd in main service`,
	Action:      RunComposeShell,
}

func RunComposeShell(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerShell(ctx)
	return nil
}
