package synthetics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AbortTestFileMultipartUploadCmd = &cobra.Command{
	Use: "abort-test-file-multipart-upload [public_id]",

	Short: "Abort a multipart upload of a test file",
	Long: `Abort a multipart upload of a test file
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#abort-test-file-multipart-upload`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SyntheticsTestFileAbortMultipartUploadRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.AbortTestFileMultipartUpload(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to abort-test-file-multipart-upload")

	},
}

func init() {

	AbortTestFileMultipartUploadCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AbortTestFileMultipartUploadCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AbortTestFileMultipartUploadCmd)
}
