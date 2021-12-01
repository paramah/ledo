package docker

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
)

var CmdComposeLogs = cli.Command{
	Name:        "logs",
	Aliases:     []string{"l"},
	Usage:       "logs from containers",
	Description: `Get fqn docker image defined as main service in config file`,
	Action:      RunComposeLogs,
}

func RunComposeLogs(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerLogs(ctx, cmd.Args())
	return nil
}
