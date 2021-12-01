package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
	"github.com/urfave/cli/v2"
)

// CmdAutocomplete manages autocompletion
var CmdAutocomplete = cli.Command{
	Name:        "shellcompletion",
	Aliases:     []string{"autocomplete"},
	Category:    catSetup,
	Usage:       "Install shell completion for ledo",
	Description: "Install shell completion for ledo",
	ArgsUsage:   "<shell type> (bash, zsh, powershell, fish)",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "install",
			Usage: "Persist in shell config instead of printing commands",
		},
	},
	Action: runAutocompleteAdd,
}

func runAutocompleteAdd(ctx *cli.Context) error {
	var remoteFile, localFile, cmds string
	shell := ctx.Args().First()

	switch shell {
	case "zsh":
		remoteFile = "contrib/autocomplete.zsh"
		localFile = "autocomplete.zsh"
		cmds = "echo 'PROG=ledo _CLI_ZSH_AUTOCOMPLETE_HACK=1 source \"%s\"' >> ~/.zshrc && source ~/.zshrc"

	case "bash":
		remoteFile = "contrib/autocomplete.sh"
		localFile = "autocomplete.sh"
		cmds = "echo 'PROG=ledo source \"%s\"' >> ~/.bashrc && source ~/.bashrc"

	case "powershell":
		remoteFile = "contrib/autocomplete.ps1"
		localFile = "tea.ps1"
		cmds = "\"& %s\" >> $profile"

	case "fish":
		return writeFishAutoCompleteFile(ctx)

	default:
		return fmt.Errorf("Must specify valid %s", ctx.Command.ArgsUsage)
	}

	localPath, err := xdg.ConfigFile("ledo/" + localFile)
	if err != nil {
		return err
	}

	cmds = fmt.Sprintf(cmds, localPath)
	if err = writeRemoteAutoCompleteFile(remoteFile, localPath); err != nil {
		return err
	}

	if ctx.Bool("install") {
		fmt.Println("Installing in your shellrc")
		installer := exec.Command(shell, "-c", cmds)
		if shell == "powershell" {
			installer = exec.Command("powershell.exe", "-Command", cmds)
		}
		out, err := installer.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Couldn't run the commands: %s %s", err, out)
		}
	} else {
		fmt.Println("\n# Run the following commands to install autocompletion (or use --install)")
		fmt.Println(cmds)
	}

	return nil
}

func writeRemoteAutoCompleteFile(file, destPath string) error {
	url := fmt.Sprintf("https://git.cynarski.pl/LeDo/ledo/raw/branch/master/%s", file)
	fmt.Println("Fetching " + url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	writer, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, res.Body)
	return err
}

func writeFishAutoCompleteFile(ctx *cli.Context) error {
	script, err := ctx.App.ToFishCompletion()
	if err != nil {
		return err
	}

	localPath, err := xdg.ConfigFile("fish/conf.d/tea_completion.fish")
	if err != nil {
		return err
	}
	writer, err := os.Create(localPath)
	if err != nil {
		return err
	}
	if _, err = io.WriteString(writer, script); err != nil {
		return err
	}
	fmt.Printf("Installed tab completion to %s\n", localPath)
	return nil
}
