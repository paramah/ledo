package container

import (
	"github.com/paramah/ledo/app/modules/container"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdDockerLogin = cli.Command{
	Name:        "login",
	Aliases:     []string{"l"},
	Usage:       "Docker Registry login",
	Description: `Login to container registry`,
	Subcommands: []*cli.Command{
		&CmdDockerEcrLogin,
	},
}

var CmdDockerEcrLogin = cli.Command{
	Name:    "ecr",
	Aliases: []string{"e"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "region",
			Aliases:  []string{"r"},
			Usage:    "aws-region",
			Required: true,
			EnvVars:  []string{"AWS_REGION"},
		},
		&cli.StringFlag{
			Name:     "key",
			Aliases:  []string{"k"},
			Usage:    "AWS access key",
			Required: true,
			EnvVars:  []string{"AWS_ACCESS_KEY_ID"},
		},
		&cli.StringFlag{
			Name:     "secret",
			Aliases:  []string{"s"},
			Usage:    "AWS secret key",
			Required: true,
			EnvVars:  []string{"AWS_SECRET_ACCESS_KEY"},
		},
	},
	Usage:       "AWS Elastic Docker Registry",
	Description: `Login to AWS Elastic Container Registry`,
	Action:      RunDockerEcrLogin,
}

func RunDockerEcrLogin(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	container.DockerEcrLogin(ctx)
	return nil
}
