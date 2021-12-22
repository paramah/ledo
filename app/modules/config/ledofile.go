package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type LedoFile struct {
	Docker  DockerMap         `yaml:"docker,omitempty`
	Modes   map[string]string `yaml:"modes,omitempty"`
	Project string            `yaml:"project,omitempty"`
}

type DockerMap struct {
	Registry    string `yaml:"registry,omitempty"`
	Namespace   string `yaml:"namespace,omitempty"`
	Name        string `yaml:"name,omitempty"`
	MainService string `yaml:"main_service,omitempty" env:"MAIN_SERVICE"`
	Shell       string `yaml:"shell,omitempty" env:"MAIN_SHELL"`
	Username    string `yaml:"username,omitempty"`
}

func NewLedoFile(s string) (*LedoFile, error) {
	yamlFile, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	t := &LedoFile{}
	err = yaml.Unmarshal(yamlFile, t)

	//Replace with env variables
	mainService := os.Getenv("MAIN_SERVICE")
	if len(mainService) != 0 {
		t.Docker.MainService = mainService
	}

	mainShell := os.Getenv("MAIN_SHELL")
	if len(mainShell) != 0 {
		t.Docker.Shell = mainShell
	}


	if err != nil {
		return nil, err
	}
	return t, nil
}
