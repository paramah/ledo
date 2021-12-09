package hashicorp_vault

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

const base = "/environment/data"

type VaultConfig struct {
	Token   string
	Address string
	Path    string
	Debug   bool
}

type Vault struct {
	client *api.Client
	config VaultConfig
}

func New(config VaultConfig) (Vault, error) {
	client, err := api.NewClient(&api.Config{Address: config.Address})
	if err != nil {
		return Vault{}, err
	}
	client.SetToken(config.Token)
	return Vault{client: client, config: config}, nil
}

func (v Vault) Write(key string, value map[string]interface{}) error {
	scr, err := v.client.Logical().Write(
		fmt.Sprintf("%s%s/%s", base, v.config.Path, key),
		map[string]interface{}{"data": value},
	)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	if v.config.Debug {
		dat, err := json.Marshal(scr)
		if err != nil {
			return fmt.Errorf("debug: %w", err)
		}
		log.Println(string(dat))
	}

	return nil
}

func (v Vault) Read(key string) (map[string]interface{}, error) {
	scr, err := v.client.Logical().Read(fmt.Sprintf("%s%s/%s", base, v.config.Path, key))
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	if scr == nil {
		return nil, fmt.Errorf("path: not found")
	}

	if v.config.Debug {
		dat, err := json.Marshal(scr)
		if err != nil {
			return nil, fmt.Errorf("debug: %w", err)
		}
		log.Println(string(dat))
	}

	dat, ok := scr.Data["data"]
	if !ok {
		return nil, fmt.Errorf("secret path %s/%s: not found", v.config.Path, key)
	}

	return dat.(map[string]interface{}), nil
}
