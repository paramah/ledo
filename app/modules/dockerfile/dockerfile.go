package dockerfile

import (
	"html/template"
	"ledo/app/helper"
	"ledo/app/templates"
	"log"
	"os"
)

func CreateDockerFile(cfg helper.DockerProjectCfg) error {
	if _, err := os.Stat("./Dockerfile"); err == nil {
		log.Printf("Dockerfile exists!")
		os.Exit(1)
	}

	templateName := templates.DockerFileTemplate_default
	if cfg.DockerBaseImage == "paramah/php" {
		templateName = templates.DockerFileTemplate_php
	}
	tpl, err := template.New("dockerfile").Parse(templateName)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("./Dockerfile")
	if err != nil {
		log.Println("create file: ", err)
	}
	err = tpl.Execute(f, cfg)

	return err
}