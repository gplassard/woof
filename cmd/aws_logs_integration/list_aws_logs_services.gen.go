package aws_logs_integration

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSLogsServicesCmd = &cobra.Command{
	Use: "list-aws-logs-services",

	Short: "Get list of AWS log ready services",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSLogsIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListAWSLogsServices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-logs-services")

		cmd.Println(cmdutil.FormatJSON(res, "logs_services"))
	},
}

func init() {
	Cmd.AddCommand(ListAWSLogsServicesCmd)
}
