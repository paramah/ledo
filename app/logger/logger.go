package logger

import (
	"github.com/pterm/pterm"
	"os"
)

const (
	levelDebugName    = "DEBUG"
	levelExecuteName  = "EXECUTE"
	levelInfoName     = "INFO"
	levelSuccessName  = "SUCCESS"
	levelErrorName    = "ERROR"
	levelCriticalName = "CRITICAL"
)

func Execute(msg string) {
	pterm.Success.Prefix = pterm.Prefix{
		Text:  levelExecuteName,
		Style: pterm.NewStyle(pterm.BgGreen, pterm.FgBlack),
	}
	pterm.Success.Printf("%v\n", msg)
}

func Info(msg string) {
	pterm.Success.Prefix = pterm.Prefix{
		Text:  levelInfoName,
		Style: pterm.NewStyle(pterm.BgGreen, pterm.FgBlack),
	}
	pterm.Success.Printf("%v\n", msg)
}

func Debug(msg string) {
	pterm.Debug.Prefix = pterm.Prefix{
		Text:  levelDebugName,
		Style: pterm.NewStyle(pterm.BgLightRed, pterm.FgBlack),
	}
	pterm.Debug.Printf("%v\n", msg)
}

func Error(msg string, err error) {
	pterm.Error.Prefix = pterm.Prefix{
		Text:  levelErrorName,
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgLightWhite),
	}
	pterm.Error.Printf("%v\n", msg)
	pterm.Error.Printf("%v\n", err.Error())
}

func Critical(msg string, err error) {
	pterm.Error.Prefix = pterm.Prefix{
		Text:  levelCriticalName,
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgLightWhite),
	}
	pterm.Error.Printf("%v\n", msg)
	pterm.Error.Printf("%v\n", err.Error())
	os.Exit(1)
}

func Exit(msg string) {
	pterm.Error.Prefix = pterm.Prefix{
		Text:  levelCriticalName,
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgLightWhite),
	}
	pterm.Error.Printf("%v\n", msg)
	os.Exit(1)
}
