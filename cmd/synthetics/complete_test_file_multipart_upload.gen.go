package synthetics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CompleteTestFileMultipartUploadCmd = &cobra.Command{
	Use: "complete-test-file-multipart-upload [public_id]",

	Short: "Complete a multipart upload of a test file",
	Long: `Complete a multipart upload of a test file
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#complete-test-file-multipart-upload`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SyntheticsTestFileCompleteMultipartUploadRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.CompleteTestFileMultipartUpload(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to complete-test-file-multipart-upload")

	},
}

func init() {

	CompleteTestFileMultipartUploadCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CompleteTestFileMultipartUploadCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CompleteTestFileMultipartUploadCmd)
}
