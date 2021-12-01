package cmd

import (
	"github.com/urfave/cli/v2"
	"ledo/app/cmd/mode"
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
