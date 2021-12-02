package image

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
)

var CmdDockerPush = cli.Command{
	Name:        "push",
	Aliases:     []string{"p"},
	Usage:       "push docker to registry",
	Description: `Push docker image to docker registry`,
	ArgsUsage:   "version",
	Action:      RunDockerPush,
}

func RunDockerPush(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	docker.ExecDockerPush(ctx, cmd.Args())
	return nil
}
