package interact

import "github.com/AlecAivazis/survey/v2"

func GetVaultEnvsForMode() bool {
	runAdv := false
	prompt := &survey.Confirm{Message: "File .env not found, do You want get envs from valut server?"}
	survey.AskOne(prompt, &runAdv)
	return runAdv
}
