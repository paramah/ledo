package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "build container image",
	Description: `Build all container images`,
	Action:      RunComposeBuild,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:     "no-cache",
			Aliases:  []string{"n"},
			Usage:    "build without cache",
			Required: false,
		},
	},
}

func RunComposeBuild(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerBuild(ctx, *cmd)
	return nil
}
