package service_scorecards

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "service-scorecards",
	Short: "service_scorecards endpoints",
	Aliases: []string{
		"scorecards",
	},
}
