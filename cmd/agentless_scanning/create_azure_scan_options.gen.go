package agentless_scanning

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAzureScanOptionsCmd = &cobra.Command{
	Use: "create-azure-scan-options",

	Short: "Create Azure scan options",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAzureScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.AzureScanOptions{})
		cmdutil.HandleError(err, "failed to create-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(CreateAzureScanOptionsCmd)
}
