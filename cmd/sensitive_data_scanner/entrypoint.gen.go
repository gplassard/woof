package sensitive_data_scanner

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "sensitive-data-scanner",
	Short: "sensitive_data_scanner endpoints",
	Aliases: []string{
		"sds",
	},
}
