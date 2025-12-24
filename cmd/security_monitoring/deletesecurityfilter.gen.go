package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteSecurityFilterCmd = &cobra.Command{
	Use:   "deletesecurityfilter",
	Short: "Delete a security filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security_monitoring/configuration/security_filters/{security_filter_id}")
		fmt.Println("OperationID: DeleteSecurityFilter")
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityFilterCmd)
}
