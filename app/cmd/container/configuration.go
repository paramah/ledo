package container

import (
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdConfiguration = cli.Command{
	Name:        "configuration",
	Usage:       "how to configure container service",
	Description: `Display configuration hints for containers`,
	Action:      RunConfiguration,
}

func RunConfiguration(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	compose.ExecComposerBuild(ctx, *cmd)
	return nil
}
