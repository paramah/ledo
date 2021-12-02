package cmd

import (
	"github.com/paramah/ledo/app/cmd/mode"
	"github.com/urfave/cli/v2"
)

var CmdMode = cli.Command{
	Name:        "mode",
	Aliases:     []string{"m"},
	Category:    catSetup,
	Usage:       "run mode management",
	Description: `Manage run mode`,
	Action:      runModeDefault,
	Subcommands: []*cli.Command{
		&mode.CmdModeSelect,
		&mode.CmdModeList,
	},
}

func runModeDefault(cmd *cli.Context) error {
	return mode.RunSelectMode(cmd)
}
