package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCustomFrameworkCmd = &cobra.Command{
	Use:   "getcustomframework",
	Short: "Get a custom framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cloud_security_management/custom_frameworks/{handle}/{version}")
		fmt.Println("OperationID: GetCustomFramework")
	},
}

func init() {
	Cmd.AddCommand(GetCustomFrameworkCmd)
}
