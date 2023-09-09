package docker

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/aws_ledo"
	"github.com/paramah/ledo/app/modules/context"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"net/url"
	"strings"
)

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func DockerEcrLogin(ctx *context.LedoContext) error {
	ecr, err := aws_ledo.EcrLogin()
	if err != nil {
		logger.Critical("Ecr login error", err)
	}
	password := *ecr.AuthorizationData[0].AuthorizationToken
	ecrUrl := *ecr.AuthorizationData[0].ProxyEndpoint
	sDec, _ := b64.StdEncoding.DecodeString(password)
	registryAddr, err := url.Parse(ecrUrl)
	if err != nil {
		logger.Critical("ECR endpoint address parse error", err)
		return err
	}
	err = ctx.ExecCmdSilent("container", []string{"login", "-u", "AWS", "-p", string(trimLeftChars(string(sDec), 4)), registryAddr.Host})
	if err != nil {
		return err
	}

	return nil
}

func ShowDockerImageFQN(ctx *context.LedoContext) string {
	fqn := fmt.Sprintf("%s/%s/%s", ctx.Config.Container.Registry, ctx.Config.Container.Namespace, ctx.Config.Container.Name)
	return strings.ToLower(fqn)
}

func ExecDockerPush(ctx *context.LedoContext, command cli.Args) {

	var version string
	var args []string
	if command.First() == "" {
		version = "latest"
	} else {
		version = command.First()
	}
	image := ShowDockerImageFQN(ctx)
	args = append(args, "push")
	args = append(args, image+":"+version)
	err := ctx.ExecCmd("container", args[0:])
	if err != nil {
		return
	}
}

func ExecDockerRetag(ctx *context.LedoContext, command cli.Args) {
	var args []string

	from := command.Get(0)
	to := command.Get(1)
	image := ShowDockerImageFQN(ctx)
	args = append(args, "tag")
	args = append(args, image+":"+from)
	args = append(args, image+":"+to)
	err := ctx.ExecCmd("container", args[0:])
	if err != nil {
		return
	}
}

func ExecDockerBuild(ctx *context.LedoContext, command cli.Args, cmdCtx cli.Context) {
	var version string
	var args []string
	if command.First() == "" {
		version = "latest"
	} else {
		version = command.First()
	}

	opts := strings.Split(strings.Trim(cmdCtx.String("opts"), " "), " ")
	image := ShowDockerImageFQN(ctx)
	args = append(args, "build")
	args = append(args, "-t", image+":"+version)
	args = append(args, "-f", cmdCtx.String("dockerfile"))
	args = append(args, opts...)
	if cmdCtx.String("stage") != "" {
		args = append(args, "--target", cmdCtx.String("stage"))
	}
	args = append(args, ".")
	err := ctx.ExecCmd("container", args[0:])
	if err != nil {
		return
	}
}

func ExecDockerPruneContainers(ctx *context.LedoContext) error {
	var formatArgs []string
	formatArgs = append(formatArgs, "--format")
	formatArgs = append(formatArgs, "{{.ID}}")

	var containerArgs []string
	containerArgs = append(containerArgs, "ps")
	containerArgs = append(containerArgs, formatArgs...)

	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Getting containers to prune...")
	output, _ := ctx.ExecCmdOutput("container", containerArgs[0:])
	spinnerLiveText.Success("Getting containers to prune... Done!")

	lines := strings.Split(string(output[:]), "\n")
	progressbar := pterm.DefaultProgressbar.WithTotal(len(lines) - 1).WithShowElapsedTime(false)
	progressbar.Title = "Prune containers"

	for _, container := range lines {
		if container == "" {
			continue
		}
		var rmargs []string
		rmargs = append(rmargs, "rm")
		rmargs = append(rmargs, container)
		rmargs = append(rmargs, "--force")
		_, err := ctx.ExecCmdOutput("container", rmargs[0:])
		if err != nil {
			return err
		}

		progressbar.Increment()
	}

	_, err := progressbar.Stop()
	if err != nil {
		return err
	}

	return nil
}

func ExecDockerPruneImages(ctx *context.LedoContext) error {
	var formatArgs []string
	formatArgs = append(formatArgs, "--format")
	formatArgs = append(formatArgs, "{{.ID}}")

	var containerArgs []string
	containerArgs = append(containerArgs, "images")
	containerArgs = append(containerArgs, "-a")
	containerArgs = append(containerArgs, formatArgs...)

	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Getting images to prune...")
	output, _ := ctx.ExecCmdOutput("container", containerArgs[0:])
	spinnerLiveText.Success("Getting images to prune... Done!")

	lines := strings.Split(string(output[:]), "\n")
	progressbar := pterm.DefaultProgressbar.WithTotal(len(lines) - 1).WithShowElapsedTime(false)
	progressbar.Title = "Prune images"

	for _, image := range lines {
		if image == "" {
			continue
		}
		var rmargs []string
		rmargs = append(rmargs, "rmi")
		rmargs = append(rmargs, image)
		rmargs = append(rmargs, "--force")
		_, err := ctx.ExecCmdOutput("container", rmargs[0:])
		if err != nil {
			return err
		}

		progressbar.Increment()
	}

	_, err := progressbar.Stop()
	if err != nil {
		return err
	}

	return nil
}

func ExecDockerPruneVolumes(ctx *context.LedoContext) error {
	var formatArgs []string
	formatArgs = append(formatArgs, "--format")
	formatArgs = append(formatArgs, "{{.ID}}")

	var containerArgs []string
	containerArgs = append(containerArgs, "volume")
	containerArgs = append(containerArgs, "ls")
	containerArgs = append(containerArgs, formatArgs...)

	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Getting volumes to prune...")
	output, _ := ctx.ExecCmdOutput("container", containerArgs[0:])
	spinnerLiveText.Success("Getting volumes to prune... Done!")

	lines := strings.Split(string(output[:]), "\n")
	progressbar := pterm.DefaultProgressbar.WithTotal(len(lines) - 1).WithShowElapsedTime(false)
	progressbar.Title = "Prune volumes"

	for _, image := range lines {
		if image == "" {
			continue
		}
		var rmargs []string
		rmargs = append(rmargs, "volume")
		rmargs = append(rmargs, "rm")
		rmargs = append(rmargs, image)
		rmargs = append(rmargs, "--force")
		_, err := ctx.ExecCmdOutput("container", rmargs[0:])
		if err != nil {
			return err
		}

		progressbar.Increment()
	}

	_, err := progressbar.Stop()
	if err != nil {
		return err
	}

	return nil
}

func ExecDockerPruneNetworks(ctx *context.LedoContext) error {
	var formatArgs []string
	formatArgs = append(formatArgs, "--format")
	formatArgs = append(formatArgs, "{{.ID}}")

	var containerArgs []string
	containerArgs = append(containerArgs, "network")
	containerArgs = append(containerArgs, "ls")
	containerArgs = append(containerArgs, formatArgs...)

	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Getting networks to prune...")
	output, _ := ctx.ExecCmdOutput("container", containerArgs[0:])
	spinnerLiveText.Success("Getting networks to prune... Done!")

	lines := strings.Split(string(output[:]), "\n")
	progressbar := pterm.DefaultProgressbar.WithTotal(len(lines) - 1).WithShowElapsedTime(false)
	progressbar.Title = "Prune networks"

	for _, network := range lines {
		if network == "" {
			continue
		}
		var rmargs []string
		rmargs = append(rmargs, "network")
		rmargs = append(rmargs, "rm")
		rmargs = append(rmargs, network)
		_, err := ctx.ExecCmdOutput("container", rmargs[0:])
		if err != nil {
			return err
		}

		progressbar.Increment()
	}

	_, err := progressbar.Stop()
	if err != nil {
		return err
	}

	return nil
}

func ExecDockerSystemPrune(ctx *context.LedoContext) error {
	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Container system prune...")

	var containerArgs []string
	containerArgs = append(containerArgs, "system")
	containerArgs = append(containerArgs, "prune")
	containerArgs = append(containerArgs, "--all")
	containerArgs = append(containerArgs, "--volumes")
	containerArgs = append(containerArgs, "--force")

	err := ctx.ExecCmd("container", containerArgs[0:])
	if err != nil {
		spinnerLiveText.Fail("Container system prune... Error!")
		return err
	}

	spinnerLiveText.Success("Container system prune... Done!")

	return nil
}

func ExecDockerPrune(ctx *context.LedoContext) error {
	var err error

	err = ExecDockerPruneContainers(ctx)
	if err != nil {
		return err
	}

	err = ExecDockerPruneImages(ctx)
	if err != nil {
		return err
	}

	err = ExecDockerPruneVolumes(ctx)
	if err != nil {
		return err
	}

	err = ExecDockerPruneNetworks(ctx)
	if err != nil {
		return err
	}

	err = ExecDockerSystemPrune(ctx)
	if err != nil {
		return err
	}

	return nil
}
