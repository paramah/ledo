package compose

import (
	"html/template"
	"ledo/app/helper"
	"ledo/app/modules/context"
	"ledo/app/templates"
	"log"
	"os"
)

func CreateComposeFile(ctx *context.LedoContext, dockerProject helper.DockerProjectCfg, composeMode string) error {
	if _, err := os.Stat("./docker"); os.IsNotExist(err) {
	    err := os.Mkdir("./docker", 0755)
    	if err != nil {
        	log.Fatal(err)
    	}
	}

	log.Printf("%v", dockerProject)

	templateName := templates.LedoDockerComposeBaseFileTemplate_base

	tpl, err := template.New("dockercompose").Parse(templateName)
	if err != nil {
		log.Fatalln(err)
	}

	composeFilename := "./docker/docker-compose.yml"

	if composeMode != "base" {
		composeFilename = "./docker/docker-compose."+composeMode+".yml"
	}

	f, err := os.Create(composeFilename)
	if err != nil {
		log.Println("create file: ", err)
	}
	err = tpl.Execute(f, ctx)

	return err
}