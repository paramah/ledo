package cmd

import (
	"github.com/paramah/ledo/app/cmd/secrets"
	"github.com/urfave/cli/v2"
)

var CmdSecrets = cli.Command{
	Name:        "secrets",
	Aliases:     []string{"s"},
	Category:    catHelpers,
	Usage:       "secrets helper",
	Description: `Managing secrets with hashicorp vault.

Requires a vault server with a KV2 resource prefixed /environment to function properly.

The vault path is created from project namespace, project name and selected mode. 

`,
	Subcommands: []*cli.Command{
		&secrets.CmdSecretsRead,
		&secrets.CmdSecretsWrite,
	},
}
