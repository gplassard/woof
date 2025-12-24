package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTagConfigurationCmd = &cobra.Command{
	Use:   "updatetagconfiguration",
	Short: "Update a tag configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/metrics/{metric_name}/tags")
		fmt.Println("OperationID: UpdateTagConfiguration")
	},
}

func init() {
	Cmd.AddCommand(UpdateTagConfigurationCmd)
}
