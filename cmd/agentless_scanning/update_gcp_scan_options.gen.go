package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateGcpScanOptionsCmd = &cobra.Command{
	Use: "update-gcp-scan-options [project_id]",

	Short: "Update GCP scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.UpdateGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0], datadogV2.GcpScanOptionsInputUpdate{})
		cmdutil.HandleError(err, "failed to update-gcp-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(UpdateGcpScanOptionsCmd)
}
