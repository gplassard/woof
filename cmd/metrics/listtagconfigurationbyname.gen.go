package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTagConfigurationByNameCmd = &cobra.Command{
	Use:   "listtagconfigurationbyname",
	Short: "List tag configuration by name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/tags")
		fmt.Println("OperationID: ListTagConfigurationByName")
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationByNameCmd)
}
