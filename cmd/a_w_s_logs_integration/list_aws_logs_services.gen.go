package a_w_s_logs_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSLogsServicesCmd = &cobra.Command{
	Use: "list-aws-logs-services",

	Short: "Get list of AWS log ready services",
	Long: `Get list of AWS log ready services
Documentation: https://docs.datadoghq.com/api/latest/a-w-s-logs-integration/#list-aws-logs-services`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSLogsServicesResponse
		var err error

		api := datadogV2.NewAWSLogsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAWSLogsServices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-logs-services")

		fmt.Println(cmdutil.FormatJSON(res, "a_w_s_logs_service"))
	},
}

func init() {

	Cmd.AddCommand(ListAWSLogsServicesCmd)
}
