package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateAwsScanOptionsCmd = &cobra.Command{
	Use: "create-aws-scan-options [payload]",

	Short: "Create AWS scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsScanOptionsResponse
		var err error

		var body datadogV2.AwsScanOptionsCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.CreateAwsScanOptions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "aws_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(CreateAwsScanOptionsCmd)
}
