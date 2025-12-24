package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSBOMCmd = &cobra.Command{
	Use:   "getsbom",
	Short: "Get SBOM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/sboms/{asset_type}")
		fmt.Println("OperationID: GetSBOM")
	},
}

func init() {
	Cmd.AddCommand(GetSBOMCmd)
}
