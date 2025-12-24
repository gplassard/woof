package organizations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOrgConfigCmd = &cobra.Command{
	Use:   "updateorgconfig",
	Short: "Update a specific Org Config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/org_configs/{org_config_name}")
		fmt.Println("OperationID: UpdateOrgConfig")
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConfigCmd)
}
