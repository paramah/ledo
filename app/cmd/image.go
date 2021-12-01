package cmd

import (
	"ledo/app/cmd/image"

	"github.com/urfave/cli/v2"
)

var CmdImage = cli.Command{
	Name:        "image",
	Aliases:     []string{"i"},
	Category:    catHelpers,
	Usage:       "docker container helper",
	Description: `docker container helper`,
	Subcommands: []*cli.Command{
		&image.CmdDockerFqn,
		&image.CmdDockerPush,
		&image.CmdDockerRetag,
		&image.CmdDockerBuild,
	},
}


