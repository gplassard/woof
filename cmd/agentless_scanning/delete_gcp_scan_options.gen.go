package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteGcpScanOptionsCmd = &cobra.Command{
	Use: "delete-gcp-scan-options [project_id]",

	Short: "Delete GCP scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err = api.DeleteGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-gcp-scan-options")

	},
}

func init() {
	Cmd.AddCommand(DeleteGcpScanOptionsCmd)
}
