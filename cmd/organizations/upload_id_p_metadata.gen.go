package organizations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UploadIdPMetadataCmd = &cobra.Command{
	Use: "upload-id-p-metadata [payload]",

	Short: "Upload IdP metadata",
	Long: `Upload IdP metadata
Documentation: https://docs.datadoghq.com/api/latest/organizations/#upload-id-p-metadata`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.UploadIdPMetadataOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		_, err = api.UploadIdPMetadata(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to upload-id-p-metadata")

	},
}

func init() {
	Cmd.AddCommand(UploadIdPMetadataCmd)
}
