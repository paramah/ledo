package docker

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
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
