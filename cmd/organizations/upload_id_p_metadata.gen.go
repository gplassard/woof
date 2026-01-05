package organizations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UploadIdPMetadataCmd = &cobra.Command{
	Use: "upload-id-p-metadata",

	Short: "Upload IdP metadata",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		_, err := api.UploadIdPMetadata(client.NewContext(apiKey, appKey, site), *datadogV2.NewUploadIdPMetadataOptionalParameters())
		cmdutil.HandleError(err, "failed to upload-id-p-metadata")

	},
}

func init() {
	Cmd.AddCommand(UploadIdPMetadataCmd)
}
