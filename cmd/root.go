package cmd

import (
	"os"
	"ouaf/cmd/roles"
	"ouaf/cmd/users"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ouaf",
	Short: "ouaf is a Datadog CLI",
	Long:  `ouaf is a Datadog CLI built with Cobra.`,
}

func Execute() {
	rootCmd.AddCommand(
		roles.Cmd,
		users.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
