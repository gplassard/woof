package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCustomFrameworkCmd = &cobra.Command{
	Use:   "updatecustomframework",
	Short: "Update a custom framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/cloud_security_management/custom_frameworks/{handle}/{version}")
		fmt.Println("OperationID: UpdateCustomFramework")
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomFrameworkCmd)
}
