package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/edsoncelio/tf-projects-boostrap/tpl"
)

type Version struct {
	Version string
}

func Create(versions string) error {

	// create versions.tf
	versionsFile, err := os.Create(fmt.Sprintf("%s", "versions.tf"))
	if err != nil {
		return err
	}
	fmt.Printf("versions.tf created\n")

	defer versionsFile.Close()
	version := Version{versions}
	versionsTemplate := template.Must(template.New("version").Parse(string(tpl.VersionsTemplate())))
	err = versionsTemplate.Execute(versionsFile, version)
	if err != nil {
		return err
	}

	//create empty files (main.tf outputs.tf variables.tf)
	mainFile, err := os.Create(fmt.Sprintf("%s", "main.tf"))
	if err != nil {
		return err
	}

	fmt.Printf("main.tf created\n")

	defer mainFile.Close()

	variableFile, err := os.Create(fmt.Sprintf("%s", "variables.tf"))
	if err != nil {
		return err
	}

	defer variableFile.Close()

	outputFile, err := os.Create(fmt.Sprintf("%s", "outputs.tf"))
	if err != nil {
		return err
	}

	defer outputFile.Close()

	backendFile, err := os.Create(fmt.Sprintf("%s", "backend.tf"))
	if err != nil {
		return err
	}

	defer backendFile.Close()

	return nil

}
