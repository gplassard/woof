package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTagConfigurationsCmd = &cobra.Command{
	Use:   "listtagconfigurations",
	Short: "Get a list of metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics")
		fmt.Println("OperationID: ListTagConfigurations")
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationsCmd)
}
