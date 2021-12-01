package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"html/template"
	"ledo/app/helper"
	"ledo/app/modules/compose"
	"ledo/app/modules/context"
	"ledo/app/modules/dockerfile"
	"ledo/app/modules/interact"
	"ledo/app/templates"
	"log"
	"os"
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
		fmt.Printf("Ledo config file not found!")
	}

	data, err := interact.InitLedoProject(config.Docker)
	if err != nil {
		return err
	}

	tpl, err := template.New("tpl").Parse(templates.LedoConfigurationFileTemplate)
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("./.ledo.yml")
	if err != nil {
		return err
	}
	err = tpl.Execute(f, data)
	if err != nil {
		return err
	}
	advRun = false
	// advRun := interact.InitAdvancedConfigurationAsk("Run advanced docker mode configuration?")
	if advRun == true {
		ctx := context.InitCommand(cmd)
		dConf, _ := interact.InitDocker()
		err = dockerfile.CreateDockerFile(dConf)
		if err != nil {
			return err
		}

		projectComposeConfig := helper.DockerProjectCfg{}
		projectComposeConfig.DockerBaseImage = dConf.DockerBaseImage

		var dockerComposeServices []helper.DockerProjectAdditionalServiceCfg
		var dockerComposeModeConfig []helper.DockerComposeModeCfg


		for _, composeMode := range interact.PredefinedDockerComposeModes {
			var configMode bool

			if composeMode == "base" {
				configMode = true
			} else {
				configMode = interact.InitAdvancedConfigurationAsk("Configure "+composeMode+" stack?")
			}

			if configMode == true {
				for {
					serviceCfg, _ := interact.CreateDockerService()
					serviceCfg.DockerServiceMode = composeMode
					dockerComposeServices = append(dockerComposeServices, serviceCfg)
					addAnother := interact.InitAdvancedConfigurationAsk("Add another service do " + composeMode + " stack?")
					if addAnother == false {
						break
					}
				}
				composeFilename := "./docker/docker-compose.yml"
				if composeMode != "base" {
					composeFilename = "./docker/docker-compose."+composeMode+".yml"
				}
				mdCfg := helper.DockerComposeModeCfg{
					DockerComposeName:     composeFilename,
					DockerComposeServices: dockerComposeServices,
				}
				dockerComposeModeConfig = append(dockerComposeModeConfig, mdCfg)
			}
		}
		projectComposeConfig.DockerComposeModes = dockerComposeModeConfig

		err = compose.CreateComposeFile(ctx, projectComposeConfig, "base")
		if err != nil {
			return err
		}
	}

	return err
}
