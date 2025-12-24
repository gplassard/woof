package organizations

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UploadIdPMetadataCmd = &cobra.Command{
	Use:   "uploadidpmetadata",
	Short: "Upload IdP metadata",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/saml_configurations/idp_metadata")
		fmt.Println("OperationID: UploadIdPMetadata")
	},
}

func init() {
	Cmd.AddCommand(UploadIdPMetadataCmd)
}
