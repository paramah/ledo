package cmd

import (
	"html/template"
	"os"

	"github.com/paramah/ledo/app/helper"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/dockerfile"
	"github.com/paramah/ledo/app/modules/interact"
	"github.com/paramah/ledo/app/templates"
	"github.com/urfave/cli/v2"
)

var CmdInit = cli.Command{
	Name:        "init",
	Category:    catSetup,
	Usage:       "init ledo in project",
	Description: `Initialize LeadDocker in current project`,
	Action:      runInitLedo,
}

func runInitLedo(cmd *cli.Context) error {
	var advRun bool

	config, err := context.LoadConfigFile()
	if err != nil {
		logger.Error("Ledo config file not found!", err)
	}

	data, err := interact.InitLedoProject(config.Container)
	if err != nil {
		logger.Critical("Initialize ledo critical error", err)
	}

	tpl, err := template.New("tpl").Parse(templates.LedoConfigurationFileTemplate)
	if err != nil {
		logger.Critical("Template parse error", err)
	}

	f, err := os.Create("./.ledo.yml")
	if err != nil {
		logger.Critical(".ledo.yml create file error", err)
	}
	err = tpl.Execute(f, data)
	if err != nil {
		logger.Critical(".ledo.yml render template error", err)
	}
	// advRun = false
	advRun = interact.InitAdvancedConfigurationAsk("Run advanced container mode configuration?")
	if advRun == true {
		ctx := context.InitCommand(cmd)
		dConf, _ := interact.InitDocker()
		err = dockerfile.CreateDockerFile(dConf)
		if err != nil {
			logger.Critical("Unable to create a Dockerfile", err)
			return err
		}

		projectComposeConfig := helper.DockerProjectCfg{}
		projectComposeConfig.DockerBaseImage = dConf.DockerBaseImage

		for _, composeMode := range interact.PredefinedDockerComposeModes {
			err = compose.CreateComposeFile(ctx, projectComposeConfig, composeMode)
			if err != nil {
				logger.Critical("Create docker-compose file error", err)
			}
		}

		_, err = helper.CreateFile(ctx, "./docker/docker-entrypoint.sh", templates.DockerEntrypointTemplate_bash, true)
		if err != nil {
			logger.Critical("Create docker-entrypoint.sh file error", err)
		}

		_, err = helper.CreateFile(ctx, "./docker/test-entrypoint.sh", templates.TestEntrypointTemplate_bash, true)
		if err != nil {
			logger.Critical("Create test-entrypoint.sh file error", err)
		}
	}

	return err
}
