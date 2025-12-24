package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTagConfigurationCmd = &cobra.Command{
	Use:   "createtagconfiguration",
	Short: "Create a tag configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/metrics/{metric_name}/tags")
		fmt.Println("OperationID: CreateTagConfiguration")
	},
}

func init() {
	Cmd.AddCommand(CreateTagConfigurationCmd)
}
