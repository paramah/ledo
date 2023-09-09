package image

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
)

var CmdDockerBuild = cli.Command{
	Name:        "build",
	Aliases:     []string{"b"},
	Usage:       "build container image",
	Description: `Build container image`,
	ArgsUsage:   "version",
	Action:      RunDockerBuild,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "stage",
			Aliases:  []string{"s"},
			Value:    "",
			Usage:    "select stage to build (multi-stage dockerfile)",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "dockerfile",
			Aliases:  []string{"f"},
			Value:    "./Dockerfile",
			Usage:    "select dockerfile",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "opts",
			Aliases:  []string{"o"},
			Value:    "--compress",
			Usage:    "Additional build options",
			Required: false,
		},
	},
}

func RunDockerBuild(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	docker.ExecDockerBuild(ctx, cmd.Args(), *cmd)
	return nil
}
