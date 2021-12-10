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

func Create(versions string, projectType string) error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if projectType == "generic" {

		versionsFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "versions.tf"))
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
		mainFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "main.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("main.tf created\n")

		defer mainFile.Close()

		variableFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "variables.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("variables.tf created\n")

		defer variableFile.Close()
		variablesTemplate := template.Must(template.New("variables").Parse(string(tpl.VariablesTemplate())))
		err = variablesTemplate.Execute(variableFile, "")
		if err != nil {
			return err
		}

		outputFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "outputs.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("outputs.tf created\n")

		defer outputFile.Close()

		outputsTemplate := template.Must(template.New("outputs").Parse(string(tpl.OutputsTemplate())))
		err = outputsTemplate.Execute(outputFile, "")
		if err != nil {
			return err
		}

		backendFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "backend.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("backend.tf created\n")

		defer backendFile.Close()

		backendTemplate := template.Must(template.New("backend").Parse(string(tpl.BackendTemplate())))
		err = backendTemplate.Execute(backendFile, "")
		if err != nil {
			return err
		}

	} else {

		versionsFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "versions.tf"))
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

		//create empty files (main.tf outputs.tf variables.tf backend.tf)
		mainFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "main.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("main.tf created\n")

		defer mainFile.Close()

		variableFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "variables.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("variables.tf created\n")

		defer variableFile.Close()
		variablesTemplate := template.Must(template.New("variables").Parse(string(tpl.VariablesTemplate())))
		err = variablesTemplate.Execute(variableFile, "")
		if err != nil {
			return err
		}

		outputFile, err := os.Create(fmt.Sprintf("%s/%s", wd, "outputs.tf"))
		if err != nil {
			return err
		}

		fmt.Printf("outputs.tf created\n")

		defer outputFile.Close()

		outputsTemplate := template.Must(template.New("outputs").Parse(string(tpl.OutputsTemplate())))
		err = outputsTemplate.Execute(outputFile, "")
		if err != nil {
			return err
		}

		editorCfgFile, err := os.Create(fmt.Sprintf("%s/%s", wd, ".editorconfig"))
		if err != nil {
			return err
		}
		fmt.Printf(".editorconfig created\n")

		defer editorCfgFile.Close()
		editorCfgTemplate := template.Must(template.New("editorconfig").Parse(string(tpl.EditorCfgTemplate())))
		err = editorCfgTemplate.Execute(editorCfgFile, "")
		if err != nil {
			return err
		}
	}

	return nil

}
