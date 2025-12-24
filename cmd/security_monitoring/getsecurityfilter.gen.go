package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSecurityFilterCmd = &cobra.Command{
	Use:   "getsecurityfilter",
	Short: "Get a security filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/security_filters/{security_filter_id}")
		fmt.Println("OperationID: GetSecurityFilter")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityFilterCmd)
}
