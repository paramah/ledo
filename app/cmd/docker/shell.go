package docker

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeShell = cli.Command{
	Name:        "shell",
	Aliases:     []string{"sh"},
	Usage:       "run shell from main service",
	Description: `Execute shell cmd in main service`,
	Action:      RunComposeShell,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Usage:    "Username or UID (format: <name|uid>)",
			Required: false,
		},
	},
}

func RunComposeShell(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerShell(ctx, *cmd)
	return nil
}
