package container

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/container"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/urfave/cli/v2"
)

var CmdPrune = cli.Command{
	Name:        "prune",
	Usage:       "clean and prune container ",
	Description: `Old and working container system prune version.`,
	Action:      RunPrune,
}

func RunPrune(cmd *cli.Context) error {
	var err error
	ctx := context.InitCommand(cmd)

	if ctx.Mode.CurrentMode != "dev" {
		logger.Exit("container prune is only available in dev mode!")
		os.Exit(255)
	}

	wantPrune := false
	prompt := &survey.Confirm{
		Message: "Do You want prune containers (all data will be irretrievably lost) ?",
	}
	err = survey.AskOne(prompt, &wantPrune)
	if err != nil {
		return err
	}
	if wantPrune {
		err = container.ExecPrune(ctx)
		if err != nil {
			return err
		}
	} else {
		logger.Info("Done!")
	}

	return nil
}
