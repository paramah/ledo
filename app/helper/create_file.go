package helper

import (
	"github.com/paramah/ledo/app/logger"
	"github.com/paramah/ledo/app/modules/context"
	"os"
	"text/template"
)

func CreateFile(ctx *context.LedoContext, path string, templateBody string, is_executable ...bool) (bool, error) {
	executable := false

	if len(is_executable) > 0 {
		executable = is_executable[0]
	}

	tpl, err := template.New("file").Parse(templateBody)
	if err != nil {
		logger.Error("Template parse error", err)
		return false, err
	}

	f, err := os.Create(path)
	if err != nil {
		logger.Error("Create file error", err)

		return false, err
	}

	err = tpl.Execute(f, ctx)
	if err != nil {
		logger.Error("Render template error", err)

		return false, err
	}

	if executable {
		err = os.Chmod(path, 0755)
		if err != nil {
			logger.Error("Chmod error", err)
			return false, err
		}
	}

	return true, nil
}
