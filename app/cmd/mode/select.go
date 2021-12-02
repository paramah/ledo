package mode

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/interact"
	"github.com/urfave/cli/v2"
)

var CmdModeSelect = cli.Command{
	Name:        "select",
	Usage:       "select run mode",
	Description: `Select mode defined in project config file`,
	ArgsUsage:   "[<mode>]",
	Action:      RunSelectMode,
}

func RunSelectMode(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	if cmd.Args().Len() == 1 {
		ctx.Mode.SetMode(cmd.Args().First())
		return nil
	}
	// interact.SelectDockerHubTag("paramah/php")
	interact.SelectMode(ctx, "")
	return nil
}
