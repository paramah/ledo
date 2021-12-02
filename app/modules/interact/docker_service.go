package interact

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/paramah/ledo/app/helper"
)

var PredefinedDockerComposeModes = []string{
	"base",
	"dev",
	"test",
}

func CreateDockerService() (helper.DockerProjectAdditionalServiceCfg, error) {
	dockerServiceConfig := helper.DockerProjectAdditionalServiceCfg{}

	var qs = []*survey.Question{
		//{
		//	Name: "DockerServiceType",
		//	Prompt: &survey.Select{
		//		Message:  "Select type of service",
		//		PageSize: 10,
		//		Options:  []string{"database", "development", "tools"},
		//	},
		//},
		//{
		//	Name: "DockerServiceMode",
		//	Prompt: &survey.Select{
		//		Message:  "Select docker-compose mode (file)",
		//		PageSize: 10,
		//		Options:  PredefinedDockerComposeModes,
		//	},
		//},
		{
			Name:      "DockerServiceImage",
			Prompt:    &survey.Input{Message: "Enter image name: "},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},

	}

	err := survey.Ask(qs, &dockerServiceConfig)

	if err != nil {
		return dockerServiceConfig, err
	}

	return dockerServiceConfig, err
}

