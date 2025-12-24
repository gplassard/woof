package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCustomFrameworkCmd = &cobra.Command{
	Use:   "deletecustomframework",
	Short: "Delete a custom framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cloud_security_management/custom_frameworks/{handle}/{version}")
		fmt.Println("OperationID: DeleteCustomFramework")
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomFrameworkCmd)
}
