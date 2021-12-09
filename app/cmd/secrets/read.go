package secrets

import (
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/secrets"
	"github.com/urfave/cli/v2"
)

var CmdSecretsRead = cli.Command{
	Name:        "read",
	Aliases:     []string{"r"},
	Usage:       "read secrets",
	Description: `Read secrets from vault`,
	Action:      RunSecretsRead,
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
	},
}

func RunSecretsRead(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	envs := secrets.SecretRead(ctx, cmd)
	secrets.ParseVaultOutput(envs)
	return nil
}
