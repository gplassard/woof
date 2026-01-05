package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAwsOnDemandTaskCmd = &cobra.Command{
	Use: "create-aws-on-demand-task",

	Short: "Create AWS on demand task",
	Long: `Create AWS on demand task
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#create-aws-on-demand-task`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AwsOnDemandResponse
		var err error

		var body datadogV2.AwsOnDemandCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.CreateAwsOnDemandTask(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-on-demand-task")

		cmd.Println(cmdutil.FormatJSON(res, "aws_resource"))
	},
}

func init() {

	CreateAwsOnDemandTaskCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAwsOnDemandTaskCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAwsOnDemandTaskCmd)
}
