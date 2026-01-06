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
	Long: `Update GCP scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#update-gcp-scan-options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GcpScanOptions
		var err error

		var body datadogV2.GcpScanOptionsInputUpdate
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-gcp-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_scan_options"))
	},
}

func init() {

	UpdateGcpScanOptionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateGcpScanOptionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateGcpScanOptionsCmd)
}
