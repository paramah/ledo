package docker

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"ledo/app/modules/aws_ledo"
	"ledo/app/modules/context"
	"net/url"
	"os"
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
		fmt.Println("Ecr login error: %s", err)
		os.Exit(1)
	}
	password := *ecr.AuthorizationData[0].AuthorizationToken
	ecrUrl := *ecr.AuthorizationData[0].ProxyEndpoint
	sDec, _ := b64.StdEncoding.DecodeString(password)
	registryAddr, err := url.Parse(ecrUrl)
	if err != nil {
		fmt.Printf("Ecr endpoint addr parse error: %s", err)
		os.Exit(1)
	}
	ctx.ExecCmdSilent("docker", []string{"login", "-u", "AWS", "-p", string(trimLeftChars(string(sDec), 4)), registryAddr.Host})

	return nil
}

func ShowDockerImageFQN(ctx *context.LedoContext) string {
	fqn := fmt.Sprintf("%s/%s/%s", ctx.Config.Docker.Registry, ctx.Config.Docker.Namespace, ctx.Config.Docker.Name)
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
	args = append(args, image + ":" + version)
	ctx.ExecCmd("docker", args[0:])
}

func ExecDockerRetag(ctx *context.LedoContext, command cli.Args) {
	var args []string

	from := command.Get(0)
	to := command.Get(1)
	image := ShowDockerImageFQN(ctx)
	args = append(args, "tag")
	args = append(args, image + ":" + from)
	args = append(args, image + ":" + to)
	ctx.ExecCmd("docker", args[0:])
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
	ctx.ExecCmd("docker", args[0:])
}