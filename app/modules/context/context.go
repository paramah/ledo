package context

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/paramah/ledo/app/modules/config"
	"github.com/paramah/ledo/app/modules/mode"
	"github.com/urfave/cli/v2"
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
	envFile := ".env"

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
	fmt.Printf("Execute: %v %v\n", cmd, strings.Join(cmdArgs, " "))
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
		os.Exit(1)
		return err
	}

	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
				os.Exit(status.ExitStatus())
				return err
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
			os.Exit(1)
			return err
		}
	}
	return nil
}

func (lx *LedoContext) ExecCmdSilent(cmd string, cmdArgs []string) error {
	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = nil

	if err := command.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
		os.Exit(1)
		return err
	}

	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
				os.Exit(status.ExitStatus())
				return err
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
			os.Exit(1)
			return err
		}
	}
	return nil
}
