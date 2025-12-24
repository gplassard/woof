package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCustomFrameworkCmd = &cobra.Command{
	Use:   "createcustomframework",
	Short: "Create a custom framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cloud_security_management/custom_frameworks")
		fmt.Println("OperationID: CreateCustomFramework")
	},
}

func init() {
	Cmd.AddCommand(CreateCustomFrameworkCmd)
}
