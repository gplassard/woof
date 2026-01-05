package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAwsScanOptionsCmd = &cobra.Command{
	Use: "update-aws-scan-options [account_id]",

	Short: "Update AWS scan options",
	Long: `Update AWS scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#update-aws-scan-options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.AwsScanOptionsUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err = api.UpdateAwsScanOptions(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-aws-scan-options")

	},
}

func init() {

	UpdateAwsScanOptionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAwsScanOptionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAwsScanOptionsCmd)
}
