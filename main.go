package main

import (
	"fmt"
	"os"

	"github.com/paramah/ledo/app/cmd"
	"github.com/urfave/cli/v2"
)

var version = "dev"

func GetCurrentVersion() string {
	return version
}

func main() {
	app := cli.NewApp()
	app.Name = "ledo"
	app.Usage = "docker-compose and container workflow improvement tool"
	app.Description = appDescription
	app.CustomAppHelpTemplate = helpTemplate
	app.Version = GetCurrentVersion()
	app.Commands = []*cli.Command{
		&cmd.CmdInit,
		&cmd.CmdContainer,
		&cmd.CmdImage,
		&cmd.CmdSecrets,
		&cmd.CmdMode,
		&cmd.CmdAutocomplete,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}
}

var appDescription = `LeadDocker (ledo) is a simple tool for improve container and docker-compose workflow in your project.
What you can do with this tool:
=> create and manage docker-compose workflow in a project
=> build container image for project (automatic fqn and container registry)
=> login to the container registry (with AWS ECR support)

LeadDocker is helpful in a continuous methodologies. 
If you want use it as a container service, try dind image: https://hub.docker.com/r/paramah/dind

Enjoy (-_-)
`

var helpTemplate = bold(`
 _                _ ___          _
| |   ___ __ _ __| |   \ ___  __| |_____ _ _
| |__/ -_) _' / _' | |) / _ \/ _| / / -_) '_|
|____\___\__,_\__,_|___/\___/\__|_\_\___|_|    {{.Version}} 

{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}`) + `

 USAGE
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .Commands}} command [subcommand] [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Description}}

 DESCRIPTION
   {{.Description | nindent 3 | trim}}{{end}}{{if .VisibleCommands}}

 COMMANDS{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

 OPTIONS
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}

 EXAMPLES
   ledo init                       # init ledo in your project
   ledo container ps               # print list of container containers

  
 ABOUT
   Written & maintained by Aleksander "paramah" Cynarski

   More info about ledo on https://leaddocker.tech
   ` + bold(`
   Thanks for:
    StreamSage Team        https://streamsage.io
    Jazzy Innovations Team https://jazzy.pro

	Contributors:
		-  muchzill4 (Bartek Mucha)
		-  vertisan (Paweł Farys) 
`) + "\n"

func bold(t string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", t)
}
