package cmd

import (
	"github.com/paramah/ledo/app/cmd/docker"
	"github.com/urfave/cli/v2"
)

var CmdDocker = cli.Command{
	Name:        "docker",
	Aliases:     []string{"d"},
	Category:    catHelpers,
	Usage:       "docker helper",
	Description: `Manage docker-compose in project`,
	Subcommands: []*cli.Command{
		&docker.CmdDockerPs,
		&docker.CmdDockerUp,
		&docker.CmdComposeBuild,
		&docker.CmdComposeDebug,
		&docker.CmdComposeDown,
		&docker.CmdComposeLogs,
		&docker.CmdComposeRestart,
		&docker.CmdComposeRun,
		&docker.CmdComposeExec,
		&docker.CmdDockerRm,
		&docker.CmdComposeShell,
		&docker.CmdComposeStart,
		&docker.CmdComposeUpOnce,
		&docker.CmdComposePull,
		&docker.CmdComposeStop,
		&docker.CmdDockerLogin,
	},
}
