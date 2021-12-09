/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Terraform project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initializing your terraform project...")
		getVersion, _ := cmd.Flags().GetString("tf-version")
		getType, _ := cmd.Flags().GetString("type")
		err := Create(getVersion, getType)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	initCmd.PersistentFlags().StringP("type", "t", "generic", "Type of project, can be module or generic")
}
