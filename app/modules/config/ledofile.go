package config

import (
	"github.com/paramah/ledo/app/logger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type SupportedRuntime string

const (
	Docker SupportedRuntime = "docker"
	Podman SupportedRuntime = "podman"
)

type LedoFile struct {
	Runtime    SupportedRuntime  `yaml:"runtime"`
	Container  ContainerMap      `yaml:"container"`
	Modes      map[string]string `yaml:"modes"`
	Project    string            `yaml:"project"`
	Deployment []Deployment      `yaml:"deployment,omitempty"`
}

type ContainerMap struct {
	Registry    string `yaml:"registry,omitempty"`
	Namespace   string `yaml:"namespace,omitempty"`
	Name        string `yaml:"name,omitempty"`
	MainService string `yaml:"main_service,omitempty" env:"MAIN_SERVICE"`
	Shell       string `yaml:"shell,omitempty" env:"MAIN_SHELL"`
	Username    string `yaml:"username,omitempty"`
}

type Deployment struct {
	Host         string `yaml:"host"`
	IsSecure     bool   `yaml:"is_secure"`
	TlsDirectory string `yaml:"tls_directory"`
	Mode         string `yaml:"mode"`
}

func NewLedoFile(s string) (*LedoFile, error) {
	yamlFile, err := ioutil.ReadFile(s)
	if err != nil {
		logger.Critical("Read yaml file error", err)
	}
	t := &LedoFile{}
	err = yaml.Unmarshal(yamlFile, t)

	//Replace with env variables
	mainService := os.Getenv("MAIN_SERVICE")
	if len(mainService) != 0 {
		t.Container.MainService = mainService
	}

	mainShell := os.Getenv("MAIN_SHELL")
	if len(mainShell) != 0 {
		t.Container.Shell = mainShell
	}

	if err != nil {
		logger.Critical("Parse yaml file error", err)
	}
	return t, nil
}
