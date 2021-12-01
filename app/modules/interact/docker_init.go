package interact

import (
	"ledo/app/helper"
)

func InitDocker() (helper.DockerProjectCfg, error) {
	dockerConfig := helper.DockerProjectCfg{}

	image, err := EnterDockerImage()
	if err != nil {
		return dockerConfig, err
	}

	tag, err := SelectDockerHubTag(image)
	if err != nil {
		return dockerConfig, err
	}

	dockerConfig.DockerBaseImage = image
	dockerConfig.DockerBaseTag = tag

	return dockerConfig, err
}
