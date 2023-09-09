package compose

import (
	"github.com/paramah/ledo/app/helper"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/templates"
	"log"
	"os"
)

func CreateComposeFile(ctx *context.LedoContext, dockerProject helper.DockerProjectCfg, composeMode string) error {
	if _, err := os.Stat("./container"); os.IsNotExist(err) {
		err := os.Mkdir("./container", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	composeFilename := "./container/container-compose.yml"
	templateName := templates.LedoDockerComposeBaseFileTemplate_base

	switch composeMode {
	case "base":
		templateName = templates.LedoDockerComposeBaseFileTemplate_base
		composeFilename = "./container/container-compose.yml"
		break
	case "dev":
		templateName = templates.LedoDockerComposeBaseFileTemplate_dev
		composeFilename = "./container/container-compose." + composeMode + ".yml"
		break
	case "test":
		templateName = templates.LedoDockerComposeBaseFileTemplate_test
		composeFilename = "./container/container-compose." + composeMode + ".yml"
		break
	case "traefik":
		templateName = templates.LedoDockerComposeBaseFileTemplate_traefik
		composeFilename = "./container/container-compose." + composeMode + ".yml"
		break
	}

	_, err := helper.CreateFile(ctx, composeFilename, templateName)
	if err != nil {
		logger.Critical("Create file error", err)
	}

	return err
}
