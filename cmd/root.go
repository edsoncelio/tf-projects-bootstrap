/*
Copyright © 2021 Edson Celio edsoncelio2020@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tf-bootstrap",
	Short: "A tool to boostrap your terraform project",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("tf-version", "v", "1.0.0", "Terraform version")
	rootCmd.AddCommand(initCmd)
}
