package mode

import (
	"github.com/urfave/cli/v2"
	"ledo/app/modules/context"
)

var CmdModeList = cli.Command{
	Name:        "list",
	Usage:       "list run modes",
	Description: `List modes defined in project config file`,
	Action:      RunModeList,
}

func RunModeList(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Mode.PrintListModes()
	return nil
}
