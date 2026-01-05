package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAwsOnDemandTasksCmd = &cobra.Command{
	Use: "list-aws-on-demand-tasks",

	Short: "List AWS on demand tasks",
	Long: `List AWS on demand tasks
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#list-aws-on-demand-tasks`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsOnDemandListResponse
		var err error

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.ListAwsOnDemandTasks(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-on-demand-tasks")

		cmd.Println(cmdutil.FormatJSON(res, "aws_resource"))
	},
}

func init() {

	Cmd.AddCommand(ListAwsOnDemandTasksCmd)
}
