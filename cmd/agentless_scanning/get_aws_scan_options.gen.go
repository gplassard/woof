package agentless_scanning

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAwsScanOptionsCmd = &cobra.Command{
	Use: "get-aws-scan-options [account_id]",

	Short: "Get AWS scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.GetAwsScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-aws-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "aws_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(GetAwsScanOptionsCmd)
}
