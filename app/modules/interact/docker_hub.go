package interact

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/paramah/ledo/app/modules/docker_hub"
)

func SelectDockerHubTag(dockerImage string) (string, error) {
	dockerImageTags := docker_hub.GetImageTags(dockerImage)
	selectedTag := "latest"

	var qs = []*survey.Question{
		{
			Name: "tags",
			Prompt: &survey.Select{
				Message:  "Select available container image tag",
				PageSize: 10,
				Options:  docker_hub.ImageTagsToArray(dockerImageTags),
			},
		},
	}

	err := survey.Ask(qs, &selectedTag)

	if err != nil {
		return "", err
	}

	return selectedTag, err
}

func EnterDockerImage() (string, error) {
	dockerImage := ""
	prompt := &survey.Input{
		Message: "Enter container base image: ",
		Help:    "This is base Dockerfile image (FROM image)",
	}
	survey.AskOne(prompt, &dockerImage)
	return dockerImage, nil
}

func SearchDockerImage(image string) (string, error) {
	dockerImages := docker_hub.GetImage(image)
	selectedImage := ""

	var qs = []*survey.Question{
		{
			Name: "image",
			Prompt: &survey.Select{
				Message:  "Select available container image",
				PageSize: 100,
				Options:  docker_hub.DockerImageToArray(dockerImages),
			},
		},
	}

	err := survey.Ask(qs, &selectedImage)

	if err != nil {
		return "", err
	}

	return selectedImage, err
}
