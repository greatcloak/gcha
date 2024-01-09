// Package cmd contains our sub commands.
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Version: "0.0.1",

		Use:   "gcha",
		Short: "CLI tool for managing Great Cloak Hosted Apps",
		Long:  `gcha is a CLI tool for deploying and managing applications hosted on the Great Cloak Hosted Apps platform.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
