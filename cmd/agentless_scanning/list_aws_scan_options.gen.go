package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAwsScanOptionsCmd = &cobra.Command{
	Use: "list-aws-scan-options",

	Short: "List AWS scan options",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.ListAwsScanOptions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "aws_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(ListAwsScanOptionsCmd)
}
