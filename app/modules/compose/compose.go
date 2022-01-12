package compose

import (
	"bytes"
	"errors"
	"github.com/Masterminds/semver"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

const DockerComposeVersion = ">= 1.28.0"

func CheckDockerComposeVersion() {
	// cmd := exec.Command("docker-compose", "--version")
	cmd := exec.Command("docker-compose", "--version")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()

	if err != nil {
		logger.Critical("No docker-compose installed. Please install docker-compose ie. via `pip3 install docker-compose`", err)
	}

	r := regexp.MustCompile("(.*){1}(version\\ ){1}(([0-9]+)\\.([0-9]+)\\.([0-9]+))")
	result := r.FindStringSubmatch(output.String())
	composeVersion := result[3]

	verConstraint, _ := semver.NewConstraint(DockerComposeVersion)
	composeSemVer, _ := semver.NewVersion(composeVersion)

	if !verConstraint.Check(composeSemVer) {
		logger.Critical("Wrong docker-compose version, please update to " + DockerComposeVersion + " or higher.", nil)
	}
}

func MergeComposerFiles(filenames ...string) (string, error) {
	var resultValues map[string]interface{}

	if len(filenames) <= 0 {
		return "", errors.New("You must provide at least one filename for reading Values")
	}

	for _, filename := range filenames {

		var override map[string]interface{}
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			logger.Error("Merge compose error", err)
			continue
		}
		if err := yaml.Unmarshal(bs, &override); err != nil {
			logger.Error("Merge compose error", err)
			continue
		}

		if resultValues == nil {
			resultValues = override
		} else {
			for k, v := range override {
				resultValues[k] = v
			}
		}
	}

	bs, err := yaml.Marshal(resultValues)
	if err != nil {
		logger.Critical("Compose error", err)
		return "", err
	}

	return string(bs), nil
}

func PrintCurrentMode(ctx *context.LedoContext) {
	pterm.Success.Prefix = pterm.Prefix{
		Text:  "MODE",
		Style: pterm.NewStyle(pterm.BgGreen, pterm.FgBlack),
	}
	pterm.Success.Printf("%v\n", ctx.Mode.CurrentMode)
}

func ExecComposerUp(ctx *context.LedoContext, noDetach bool) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs

	args = append(args, "up")
	if noDetach == false {
		args = append(args, "-d")
	}
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerPull(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "pull")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerStop(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "stop")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerBuild(ctx *context.LedoContext, command cli.Context) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "build", "--pull")
	if command.Bool("no-cache") == true {
		args = append(args, "--no-cache")
	}
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerDown(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "down")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerStart(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "start")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerRestart(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "restart")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerLogs(ctx *context.LedoContext, command cli.Args) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "logs", "--follow", "--tail", "100")
	args = append(args, command.Slice()...)
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerPs(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "ps")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerRm(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "rm", "-f")
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerShell(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "exec", strings.ToLower(ctx.Config.Docker.MainService), ctx.Config.Docker.Shell)
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerDebug(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "run", "--entrypoint=", strings.ToLower(ctx.Config.Docker.MainService), ctx.Config.Docker.Shell)
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerRun(ctx *context.LedoContext, command cli.Args) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "run", strings.ToLower(ctx.Config.Docker.MainService))
	if ctx.Config.Docker.Username != "" {
		args = append(args, "sudo", "-E", "-u", ctx.Config.Docker.Username)
	}
	args = append(args, command.Slice()...)
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerExec(ctx *context.LedoContext, command cli.Args) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "exec", strings.ToLower(ctx.Config.Docker.MainService))
	if ctx.Config.Docker.Username != "" {
		args = append(args, "sudo", "-E", "-u", ctx.Config.Docker.Username)
	}
	args = append(args, command.Slice()...)
	ctx.ExecCmd("docker-compose", args[0:])
}

func ExecComposerUpOnce(ctx *context.LedoContext) {
	PrintCurrentMode(ctx)
	args := ctx.ComposeArgs
	args = append(args, "up", "--force-recreate", "--renew-anon-volumes", "--abort-on-container-exit", "--exit-code-from", ctx.Config.Docker.MainService)
	ctx.ExecCmd("docker-compose", args[0:])
}
