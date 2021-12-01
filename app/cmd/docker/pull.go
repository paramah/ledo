package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposePull = cli.Command{
	Name:        "pull",
	Usage:       "docker image pull",
	Description: `Pull docker image from registry server`,
	Action:      RunComposePull,
}

func RunComposePull(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerPull(ctx)
	return nil
}
