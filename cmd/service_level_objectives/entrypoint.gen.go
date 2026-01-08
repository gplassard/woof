package service_level_objectives

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "service-level-objectives",
	Short: "service_level_objectives endpoints",
	Aliases: []string{
		"slo",
	},
}
