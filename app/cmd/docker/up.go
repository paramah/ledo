package docker

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
	Action:      RunComposeUp,
}

func RunComposeUp(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerUp(ctx)
	compose.ExecComposerLogs(ctx, cmd.Args())
	return nil
}
