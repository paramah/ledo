package image

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
)

var CmdDockerRetag = cli.Command{
	Name:        "retag",
	Aliases:     []string{"r"},
	Usage:       "retag container image",
	Description: `Change container image tag`,
	ArgsUsage:   "fromTag toTag",
	Action:      RunDockerRetag,
}

func RunDockerRetag(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	docker.ExecDockerRetag(ctx, cmd.Args())
	return nil
}
