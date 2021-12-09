package secrets

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/secrets"
	"github.com/urfave/cli/v2"
)

var CmdSecretsWrite = cli.Command{
	Name:        "write",
	Aliases:     []string{"w"},
	Usage:       "write secrets",
	Description: `Write secrets to vault`,
	Action:      RunSecretsWrite,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "addr",
			Aliases:  []string{"a"},
			Usage:    "vault address",
			Required: true,
			EnvVars:  []string{"VAULT_ADDR"},
		},
		&cli.StringFlag{
			Name:     "token",
			Aliases:  []string{"t"},
			Usage:    "vault token",
			Required: true,
			EnvVars:  []string{"VAULT_TOKEN"},
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Debug output",
			Value:   false,
		},
		&cli.PathFlag{
			Name:     "input",
			Aliases:  []string{"i"},
			Usage:    "read env from file",
			Required: false,
		},
	},
}

func RunSecretsWrite(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	secrets.SecretWrite(ctx, cmd)
	return nil
}
