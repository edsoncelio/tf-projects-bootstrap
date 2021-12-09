package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/edsoncelio/tf-projects-boostrap/tpl"
)

func Create(versions string) error {

	// create versions.tf
	versionsFile, err := os.Create(fmt.Sprintf("%s", "versions.tf"))
	if err != nil {
		return err
	}

	defer versionsFile.Close()

	versionsTemplate := template.Must(template.New("versions").Parse(string(tpl.VersionsTemplate())))
	err = versionsTemplate.Execute(versionsFile, versions)
	if err != nil {
		return err
	}
	//create empty files (main.tf outputs.tf variables.tf)
	_, err = os.Create(fmt.Sprintf("%s", "main.tf"))
	if err != nil {
		return err
	}

	_, err = os.Create(fmt.Sprintf("%s", "variables.tf"))
	if err != nil {
		return err
	}

	_, err = os.Create(fmt.Sprintf("%s", "outputs.tf"))
	if err != nil {
		return err
	}
	return nil

}
