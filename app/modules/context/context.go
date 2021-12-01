package context

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"ledo/app/modules/config"
	"ledo/app/modules/mode"
	"os"
	"os/exec"
	"strings"
)

type LedoContext struct {
	*cli.Context
	Config      *config.LedoFile
	ComposeArgs []string
	Mode        mode.Mode
	Output      string
}

func InitCommand(ctx *cli.Context) *LedoContext {

	var (
		c   LedoContext
		cfg *config.LedoFile
	)

	configYml := ".ledo.yml"
	modeYml := ".ledo-mode"

	// Compat with jzcli (StreamSage and Jazzy deployment tool)
	if _, err := os.Stat(".jz-project.yml"); err == nil {
		configYml = ".jz-project.yml"
		modeYml = ".jz-mode"
	}

	if _, err := os.Stat(configYml); err != nil {
		fmt.Printf("Config file not found. Please run ledo init\n")
		os.Exit(1)
	}

	ledoMode := mode.InitMode(modeYml, configYml)
	c.Mode = ledoMode

	c.Output = ctx.String("output")

	cfg, _ = config.NewLedoFile(configYml)
	c.Config = cfg

	args := []string{"--env-file", ".env"}
	args = append(args, "--project-name", strings.ToLower(strings.Replace(c.Config.Docker.Name, "/", "-", -1)))

	composes, _ := ledoMode.GetModeConfig()
	for _, element := range composes {
		args = append(args, "-f")
		args = append(args, element)
	}

	c.ComposeArgs = args

	return &c
}

func LoadConfigFile() (*config.LedoFile, error) {
	configYml := ".ledo.yml"

	if _, err := os.Stat(configYml); err != nil {
		nilCfg := &config.LedoFile{}
		return nilCfg, err
	}

	cfg, _ := config.NewLedoFile(configYml)
	return cfg, nil
}

func (lx *LedoContext) ExecCmd(cmd string, cmdArgs []string) error {
	fmt.Printf("Execute: %v %v\n", cmd, strings.Join(cmdArgs, " "))
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func (lx *LedoContext) ExecCmdSilent(cmd string, cmdArgs []string) error {
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = nil
	return command.Run()
}