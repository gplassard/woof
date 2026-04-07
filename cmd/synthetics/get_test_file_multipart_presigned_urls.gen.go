package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTestFileMultipartPresignedUrlsCmd = &cobra.Command{
	Use: "get-test-file-multipart-presigned-urls [public_id]",

	Short: "Get presigned URLs for uploading a test file",
	Long: `Get presigned URLs for uploading a test file
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-test-file-multipart-presigned-urls`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsTestFileMultipartPresignedUrlsResponse
		var err error

		var body datadogV2.SyntheticsTestFileMultipartPresignedUrlsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTestFileMultipartPresignedUrls(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to get-test-file-multipart-presigned-urls")

		fmt.Println(cmdutil.FormatJSON(res, "synthetics"))
	},
}

func init() {

	GetTestFileMultipartPresignedUrlsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetTestFileMultipartPresignedUrlsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetTestFileMultipartPresignedUrlsCmd)
}
