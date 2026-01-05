package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateAzureScanOptionsCmd = &cobra.Command{
	Use: "create-azure-scan-options [payload]",

	Short: "Create Azure scan options",
	Long: `Create Azure scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#create-azure-scan-options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureScanOptions
		var err error

		var body datadogV2.AzureScanOptions
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.CreateAzureScanOptions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(CreateAzureScanOptionsCmd)
}
