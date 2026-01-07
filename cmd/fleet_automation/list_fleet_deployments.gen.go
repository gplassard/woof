package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetDeploymentsCmd = &cobra.Command{
	Use: "list-fleet-deployments",

	Short: "List all deployments",
	Long: `List all deployments
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#list-fleet-deployments`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentsResponse
		var err error

		optionalParams := datadogV2.NewListFleetDeploymentsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFleetDeployments(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-fleet-deployments")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {

	ListFleetDeploymentsCmd.Flags().Int64("page-size", 0, "Number of deployments to return per page. Maximum value is 100.")

	ListFleetDeploymentsCmd.Flags().Int64("page-offset", 0, "Index of the first deployment to return. Use this with 'page_size' to paginate through results.")

	Cmd.AddCommand(ListFleetDeploymentsCmd)
}
