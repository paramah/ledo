package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeUpOnce = cli.Command{
	Name:        "uponce",
	Usage:       "up one container",
	Description: `Up one container from docker compose stack`,
	Action:      RunComposeUpOnce,
}

func RunComposeUpOnce(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerUpOnce(ctx)
	return nil
}
