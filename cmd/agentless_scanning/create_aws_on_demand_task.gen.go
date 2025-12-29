package agentless_scanning

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAwsOnDemandTaskCmd = &cobra.Command{
	Use:   "create-aws-on-demand-task",
	
	Short: "Create AWS on demand task",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAwsOnDemandTask(client.NewContext(apiKey, appKey, site), datadogV2.AwsOnDemandCreateRequest{})
		cmdutil.HandleError(err, "failed to create-aws-on-demand-task")

		cmd.Println(cmdutil.FormatJSON(res, "aws_resource"))
	},
}

func init() {
	Cmd.AddCommand(CreateAwsOnDemandTaskCmd)
}
