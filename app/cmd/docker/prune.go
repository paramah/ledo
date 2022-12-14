package docker

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/paramah/ledo/app/modules/docker"
	"github.com/urfave/cli/v2"
	"os"
)

var CmdPrune = cli.Command{
	Name:        "prune",
	Usage:       "clean and prune docker ",
	Description: `Old and working docker system prune version.`,
	Action:      RunPrune,
}

func RunPrune(cmd *cli.Context) error {
	var err error
	ctx := context.InitCommand(cmd)

	if ctx.Mode.CurrentMode != "dev" {
		logger.Exit("docker prune is only available in dev mode!")
		os.Exit(255)
	}

	wantPrune := false
	prompt := &survey.Confirm{
    	Message: "Do You want prune docker (all data will be irretrievably lost) ?",
    }
	err = survey.AskOne(prompt, &wantPrune)
	if err != nil {
		return err
	}
	if wantPrune {
		err = docker.ExecDockerPrune(ctx)
		if err != nil {
			return err
		}
	} else {
		logger.Info("Done!")
	}

	return nil
}
