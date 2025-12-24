package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListScannedAssetsMetadataCmd = &cobra.Command{
	Use:   "listscannedassetsmetadata",
	Short: "List scanned assets metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/scanned-assets-metadata")
		fmt.Println("OperationID: ListScannedAssetsMetadata")
	},
}

func init() {
	Cmd.AddCommand(ListScannedAssetsMetadataCmd)
}
