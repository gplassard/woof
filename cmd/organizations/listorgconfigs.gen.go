package organizations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListOrgConfigsCmd = &cobra.Command{
	Use:   "listorgconfigs",
	Short: "List Org Configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/org_configs")
		fmt.Println("OperationID: ListOrgConfigs")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConfigsCmd)
}
