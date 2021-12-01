package interact

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/thoas/go-funk"
	"ledo/app/modules/context"
)

func SelectMode(context *context.LedoContext, selectedMode string) (string, error) {
	mode := context.Mode
	if selectedMode == "" {
		availableModes := mode.GetModes()
		// the questions to ask
		var qs = []*survey.Question{
			{
				Name: "providers",
				Prompt: &survey.Select{
					Message:  "Select run mode",
					PageSize: 10,
					Options:  funk.Keys(availableModes).([]string),
				},
			},
		}
		err := survey.Ask(qs, &selectedMode)
		if err != nil {
			return "", err
		}
	}
	_, err := mode.SetMode(selectedMode)

	return selectedMode, err
}
