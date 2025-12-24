package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSecurityFilterCmd = &cobra.Command{
	Use:   "createsecurityfilter",
	Short: "Create a security filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security_monitoring/configuration/security_filters")
		fmt.Println("OperationID: CreateSecurityFilter")
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityFilterCmd)
}
