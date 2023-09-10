package image

import (
	"github.com/paramah/ledo/app/modules/container"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdDockerPush = cli.Command{
	Name:        "push",
	Aliases:     []string{"p"},
	Usage:       "push container to registry",
	Description: `Push container image to container registry`,
	ArgsUsage:   "version",
	Action:      RunDockerPush,
}

func RunDockerPush(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	container.ExecPush(ctx, cmd.Args())
	return nil
}
