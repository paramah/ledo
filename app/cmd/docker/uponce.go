package docker

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
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
