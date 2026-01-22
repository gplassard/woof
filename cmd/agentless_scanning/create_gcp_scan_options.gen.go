package agentless_scanning

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateGcpScanOptionsCmd = &cobra.Command{
	Use: "create-gcp-scan-options",

	Short: "Create GCP scan options",
	Long: `Create GCP scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#create-gcp-scan-options`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GcpScanOptions
		var err error

		var body datadogV2.GcpScanOptions
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateGcpScanOptions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-gcp-scan-options")

		fmt.Println(cmdutil.FormatJSON(res, "gcp_scan_option"))
	},
}

func init() {

	CreateGcpScanOptionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateGcpScanOptionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateGcpScanOptionsCmd)
}
