package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateSecurityFilterCmd = &cobra.Command{
	Use:   "updatesecurityfilter",
	Short: "Update a security filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security_monitoring/configuration/security_filters/{security_filter_id}")
		fmt.Println("OperationID: UpdateSecurityFilter")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityFilterCmd)
}
