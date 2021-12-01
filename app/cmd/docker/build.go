package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "build docker image",
	Description: `Build all docker images`,
	Action:      RunComposeBuild,
}

func RunComposeBuild(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerBuild(ctx)
	return nil
}
