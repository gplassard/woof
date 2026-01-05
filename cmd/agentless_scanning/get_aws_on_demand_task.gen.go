package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAwsOnDemandTaskCmd = &cobra.Command{
	Use: "get-aws-on-demand-task [task_id]",

	Short: "Get AWS on demand task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.GetAwsOnDemandTask(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-aws-on-demand-task")

		cmd.Println(cmdutil.FormatJSON(res, "aws_resource"))
	},
}

func init() {
	Cmd.AddCommand(GetAwsOnDemandTaskCmd)
}
