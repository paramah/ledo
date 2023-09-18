package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdComposeStart = cli.Command{
	Name:        "start",
	Usage:       "start containers",
	Description: `Start all containers defined in docker-compose stack run mode`,
	Action:      RunComposeStart,
}

func RunComposeStart(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerStart(ctx)
	return nil
}
