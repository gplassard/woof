package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTestFileDownloadUrlCmd = &cobra.Command{
	Use: "get-test-file-download-url [public_id]",

	Short: "Get a presigned URL for downloading a test file",
	Long: `Get a presigned URL for downloading a test file
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-test-file-download-url`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsTestFileDownloadResponse
		var err error

		var body datadogV2.SyntheticsTestFileDownloadRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTestFileDownloadUrl(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to get-test-file-download-url")

		fmt.Println(cmdutil.FormatJSON(res, "synthetics"))
	},
}

func init() {

	GetTestFileDownloadUrlCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetTestFileDownloadUrlCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetTestFileDownloadUrlCmd)
}
