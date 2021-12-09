package secrets

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-envparse"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/hashicorp_vault"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func Connect(ctx *context.LedoContext, command *cli.Context) (hashicorp_vault.Vault, error) {
	envPath := fmt.Sprintf("/%s/%s", strings.ToLower(ctx.Config.Docker.Namespace), strings.ToLower(ctx.Config.Docker.Name))
	var vaultConfig = hashicorp_vault.VaultConfig{
		Token:   command.String("token"),
		Address: command.String("addr"),
		Path:    envPath,
		Debug:   command.Bool("debug"),
	}

	vault, err := hashicorp_vault.New(vaultConfig)
	if err != nil {
		log.Println("%v", err)
		return hashicorp_vault.Vault{}, err
	}
	return vault, nil
}

func SecretRead(ctx *context.LedoContext, command *cli.Context) map[string]interface{} {
	vault, err := Connect(ctx, command)
	if err != nil {
		fmt.Errorf("read: %w", err)
	}

	dta, err := vault.Read(ctx.Mode.CurrentMode)
	if err != nil {
		out := fmt.Errorf("%w", err)
		log.Println(out)
	}

	return dta
}

func SecretWrite(ctx *context.LedoContext, command *cli.Context) error {
	var err error
	var vaultEnvs map[string]interface{}

	vault, err := Connect(ctx, command)
	if err != nil {
		log.Println("%v", err)
		return err
	}

	envFile := command.Path("input")
	if len(envFile) > 0 {
		fDta, err := os.ReadFile(envFile)
		if err != nil {
			log.Printf("readfile: %v", err)
			return err
		}

		envs, err := envparse.Parse(bytes.NewReader(fDta))
		if err != nil {
			log.Printf("unmarshal: %v", err)
			return err
		}

		vaultEnvs = make(map[string]interface{}, len(envs))
		for i, v := range envs {
			vaultEnvs[i] = v
		}

	} else {
		if len(command.Args().Slice()) == 0 {
			log.Println("no variables are specified")
			return err
		}
		vaultEnvs = make(map[string]interface{}, len(command.Args().Slice()))
		for _, v := range command.Args().Slice() {
			e := strings.Split(v, "=")
			key := e[0]
			val := strings.Join(e[1:], "")
			vaultEnvs[key] = val
		}
	}

	err = vault.Write(ctx.Mode.CurrentMode, vaultEnvs)
	if err != nil {
		log.Println("%v", err)
		return err
	}

	return err
}

func ParseVaultOutput(envs map[string]interface{}) {
	for i, v := range envs {
		fmt.Printf("%v=%v\n", i, v)
	}
}
