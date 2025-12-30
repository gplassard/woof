package agentless_scanning

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAwsScanOptionsCmd = &cobra.Command{
	Use: "create-aws-scan-options",

	Short: "Create AWS scan options",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAwsScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.AwsScanOptionsCreateRequest{})
		cmdutil.HandleError(err, "failed to create-aws-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "aws_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(CreateAwsScanOptionsCmd)
}
