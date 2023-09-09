package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdDockerPs = cli.Command{
	Name:        "ps",
	Aliases:     []string{"p"},
	Usage:       "list running containers",
	Description: `List all containers defined in container-compose use in current mode`,
	Action:      RunComposePs,
}

func RunComposePs(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerPs(ctx)
	return nil
}
