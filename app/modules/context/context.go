package context

import (
	"fmt"
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/config"
	"github.com/paramah/ledo/app/modules/mode"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strings"
)

type LedoContext struct {
	*cli.Context
	Config      *config.LedoFile
	ComposeArgs []string
	Mode        mode.Mode
}

func InitCommand(ctx *cli.Context) *LedoContext {

	var (
		c   LedoContext
		cfg *config.LedoFile
	)

	configYml := ".ledo.yml"
	modeYml := ".ledo-mode"
	envFile := ".env"

	// Compat with jzcli (StreamSage and Jazzy deployment tool)
	if _, err := os.Stat(".jz-project.yml"); err == nil {
		configYml = ".jz-project.yml"
		modeYml = ".jz-mode"
	}

	if _, err := os.Stat(configYml); err != nil {
		logger.Critical("Config file not found. Please run ledo init", err)
	}

	ledoMode := mode.InitMode(modeYml, configYml)
	c.Mode = ledoMode

	//c.Output = ctx.String("output")
	//c.Verbosity = 1

	cfg, _ = config.NewLedoFile(configYml)
	c.Config = cfg

	args := []string{"--env-file", envFile}
	args = append(args, "--project-name", strings.ToLower(strings.Replace(c.Config.Docker.Namespace, "/", "-", -1)))
	
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
	logger.Execute(fmt.Sprintf("%v %v\n", cmd, strings.Join(cmdArgs, " ")))
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		logger.Critical("Execute critical error", err)
	}

	if err := command.Wait(); err != nil {
		logger.Critical("Execute critical error", err)
	}
	return nil
}

func (lx *LedoContext) ExecCmdSilent(cmd string, cmdArgs []string) error {
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = nil

	if err := command.Start(); err != nil {
		logger.Critical("Execute critical error", err)
	}

	if err := command.Wait(); err != nil {
		logger.Critical("Execute critical error", err)
	}
	return nil
}
