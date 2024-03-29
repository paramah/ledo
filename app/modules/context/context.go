package context

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/config"
	"github.com/paramah/ledo/app/modules/git"
	"github.com/paramah/ledo/app/modules/mode"
	"github.com/urfave/cli/v2"
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

	configYml := ResolveRootPath(".ledo.yml")
	modeYml := ".ledo-mode"
	envFile := ResolveRootPath(".env")
	jzProjectYml := ResolveRootPath(".jz-project.yml")

	// Compat with jzcli (StreamSage and Jazzy deployment tool)
	if _, err := os.Stat(jzProjectYml); err == nil {
		configYml = jzProjectYml
		modeYml = ResolveRootPath(".jz-mode")
	}

	if _, err := os.Stat(configYml); err != nil {
		logger.Critical("Config file not found. Please run ledo init", err)
	}

	ledoMode := mode.InitMode(modeYml, configYml)
	c.Mode = ledoMode

	// c.Output = ctx.String("output")
	// c.Verbosity = 1

	cfg, _ = config.NewLedoFile(configYml)
	c.Config = cfg

	args := []string{"--env-file", envFile}
	args = append(args, "--project-name", strings.ToLower(strings.Replace(c.Config.Container.Namespace, "/", "-", -1)))

	composes, _ := ledoMode.GetModeConfig()
	for _, element := range composes {
		args = append(args, "-f")
		args = append(args, ResolveRootPath(element))
	}

	c.ComposeArgs = args

	return &c
}

func LoadConfigFile() (*config.LedoFile, error) {
	configYml := ResolveRootPath(".ledo.yml")

	if _, err := os.Stat(configYml); err != nil {
		nilCfg := &config.LedoFile{}
		return nilCfg, err
	}

	cfg, _ := config.NewLedoFile(configYml)
	return cfg, nil
}

func (lx *LedoContext) GetRuntimeCommand() string {
	runtime := lx.Config.Runtime
	return runtime.Command()
}

func (lx *LedoContext) GetRuntimeCompose() string {
	runtime := lx.Config.Runtime
	return runtime.Compose()
}

func (lx *LedoContext) ExecCmdOutput(cmd string, cmdArgs []string) ([]byte, error) {
	command, _ := exec.Command(cmd, cmdArgs...).Output()
	return command, nil
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
	command.Stdout = nil
	command.Stderr = nil

	if err := command.Start(); err != nil {
		logger.Critical("Execute critical error", err)
	}

	if err := command.Wait(); err != nil {
		logger.Critical("Execute critical error", err)
	}
	return nil
}

func ResolveRootPath(configFile string) string {
	gitRoot, err := git.GetRepositoryRootDir()
	if err != nil {
		return configFile
	}
	return fmt.Sprintf("%s/%s", gitRoot, configFile)
}
