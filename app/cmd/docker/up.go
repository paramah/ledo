package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdDockerUp = cli.Command{
	Name:        "up",
	Aliases:     []string{"u"},
	Usage:       "up containers",
	Description: `Up all containers defined in docker-compose use in current mode`,
	Action:      RunComposeUp,
}

func RunComposeUp(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerUp(ctx)
	compose.ExecComposerLogs(ctx, cmd.Args())
	return nil
}
