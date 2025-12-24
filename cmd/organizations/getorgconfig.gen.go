package organizations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOrgConfigCmd = &cobra.Command{
	Use:   "getorgconfig",
	Short: "Get a specific Org Config value",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/org_configs/{org_config_name}")
		fmt.Println("OperationID: GetOrgConfig")
	},
}

func init() {
	Cmd.AddCommand(GetOrgConfigCmd)
}
