package agentless_scanning

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAwsScanOptionsCmd = &cobra.Command{
	Use: "create-aws-scan-options",

	Short: "Create AWS scan options",
	Long: `Create AWS scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#create-aws-scan-options`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsScanOptionsResponse
		var err error

		var body datadogV2.AwsScanOptionsCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateAwsScanOptions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-scan-options")

		fmt.Println(cmdutil.FormatJSON(res, "aws_scan_option"))
	},
}

func init() {

	CreateAwsScanOptionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAwsScanOptionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAwsScanOptionsCmd)
}
