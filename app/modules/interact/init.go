package interact

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/paramah/ledo/app/modules/config"
)

func InitLedoProject(dockerConfig config.ContainerMap) (config.ContainerMap, error) {
	if dockerConfig.Registry == "" {
		dockerConfig.Registry = ""
	}

	if dockerConfig.Shell == "" {
		dockerConfig.Shell = "/bin/bash"
	}

	if dockerConfig.Username == "" {
		dockerConfig.Username = "root"
	}

	var qs = []*survey.Question{
		{
			Name:      "Registry",
			Prompt:    &survey.Input{Message: "Enter container registry address: ", Default: dockerConfig.Registry, Help: "Container registry for main service image"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "Namespace",
			Prompt:    &survey.Input{Message: "Enter project namespace: ", Default: dockerConfig.Namespace, Help: "Project namespace (eq. GitLab project group)"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "Name",
			Prompt:    &survey.Input{Message: "Enter project name: ", Default: dockerConfig.Name},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "MainService",
			Prompt:    &survey.Input{Message: "Enter docker-compose main service name: ", Default: dockerConfig.MainService, Help: "Main service, important if you want use ledo shell command or ledo run command"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "Shell",
			Prompt:    &survey.Input{Message: "Enter default shell: ", Default: dockerConfig.Shell},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "Username",
			Prompt:    &survey.Input{Message: "Enter container main service username: ", Default: dockerConfig.Username, Help: "Default user, if set ledo run command was execute with sudo user"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
	}

	err := survey.Ask(qs, &dockerConfig)

	if err != nil {
		return config.ContainerMap{}, err
	}

	return dockerConfig, err
}

func InitAdvancedConfigurationAsk(question string) bool {
	runAdv := false
	prompt := &survey.Confirm{Message: question}
	survey.AskOne(prompt, &runAdv)
	return runAdv
}
