package image

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/context"
	"ledo/app/modules/docker"
)

var CmdDockerRetag = cli.Command{
	Name:        "retag",
	Aliases:     []string{"r"},
	Usage:       "retag docker image",
	Description: `Change docker image tag`,
	ArgsUsage:   "fromTag toTag",
	Action:      RunDockerRetag,
}

func RunDockerRetag(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	docker.ExecDockerRetag(ctx, cmd.Args())
	return nil
}
