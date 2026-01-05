package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAzureScanOptionsCmd = &cobra.Command{
	Use: "create-azure-scan-options",

	Short: "Create Azure scan options",
	Long: `Create Azure scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#create-azure-scan-options`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureScanOptions
		var err error

		var body datadogV2.AzureScanOptions
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.CreateAzureScanOptions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {

	CreateAzureScanOptionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAzureScanOptionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAzureScanOptionsCmd)
}
