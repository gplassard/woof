package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAssetsSBOMsCmd = &cobra.Command{
	Use:   "listassetssboms",
	Short: "List assets SBOMs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/sboms")
		fmt.Println("OperationID: ListAssetsSBOMs")
	},
}

func init() {
	Cmd.AddCommand(ListAssetsSBOMsCmd)
}
