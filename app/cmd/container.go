package cmd

import (
	"github.com/paramah/ledo/app/cmd/container"
	"github.com/urfave/cli/v2"
)

var CmdContainer = cli.Command{
	Name:        "container",
	Aliases:     []string{"d"},
	Category:    catHelpers,
	Usage:       "container helper",
	Description: `Manage compose tools (docker or podman) in project`,
	Subcommands: []*cli.Command{
		&container.CmdDockerPs,
		&container.CmdDockerUp,
		&container.CmdComposeBuild,
		&container.CmdComposeDebug,
		&container.CmdComposeDown,
		&container.CmdComposeLogs,
		&container.CmdComposeRestart,
		&container.CmdComposeRun,
		&container.CmdComposeExec,
		&container.CmdDockerRm,
		&container.CmdComposeShell,
		&container.CmdComposeStart,
		&container.CmdComposeUpOnce,
		&container.CmdComposePull,
		&container.CmdComposeStop,
		&container.CmdDockerLogin,
		&container.CmdPrune,
	},
}
