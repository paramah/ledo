package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdDockerUp = cli.Command{
	Name:        "up",
	Aliases:     []string{"u"},
	Usage:       "up containers",
	Description: `Up all containers defined in docker-compose use in current mode`,
	Action:      RunComposeUp, Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:     "no-detach",
			Aliases:  []string{"n"},
			Usage:    "run in foreground",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "build",
			Aliases:  []string{"b"},
			Usage:    "build local images before run",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "logs",
			Aliases:  []string{"l"},
			Usage:    "show logs after up",
			Required: false,
		},
	},
}

func RunComposeUp(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	if cmd.Bool("build") {
		compose.ExecComposerBuild(ctx, *cmd)
	}
	compose.ExecComposerUp(ctx, cmd.Bool("no-detach"))
	if cmd.Bool("logs") {
		compose.ExecComposerLogs(ctx, cmd.Args())
	}
	return nil
}
