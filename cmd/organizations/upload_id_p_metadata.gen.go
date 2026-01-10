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
	Long: `Upload IdP metadata
Documentation: https://docs.datadoghq.com/api/latest/organizations/#upload-id-p-metadata`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		optionalParams := datadogV2.NewUploadIdPMetadataOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UploadIdPMetadata(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to upload-id-p-metadata")

	},
}

func init() {

	UploadIdPMetadataCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UploadIdPMetadataCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UploadIdPMetadataCmd)
}
